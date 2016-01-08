/*
Licensed to the Apache Software Foundation (ASF) under one
or more contributor license agreements.  See the NOTICE file
distributed with this work for additional information
regarding copyright ownership.  The ASF licenses this file
to you under the Apache License, Version 2.0 (the
"License"); you may not use this file except in compliance
with the License.  You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing,
software distributed under the License is distributed on an
"AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
KIND, either express or implied.  See the License for the
specific language governing permissions and limitations
under the License.
*/

package genesis

import (
	"fmt"
	"sync"

	"golang.org/x/net/context"

	"github.com/op/go-logging"
	"github.com/openblockchain/obc-peer/openchain"
	"github.com/openblockchain/obc-peer/openchain/chaincode"
	"github.com/openblockchain/obc-peer/openchain/container"
	"github.com/openblockchain/obc-peer/openchain/crypto"
	"github.com/openblockchain/obc-peer/openchain/ledger"
	"github.com/openblockchain/obc-peer/protos"
	"github.com/spf13/viper"
)

var genesisLogger = logging.MustGetLogger("genesis")

var makeGenesisError error
var once sync.Once

// MakeGenesis creates the genesis block based on configuration in openchain.yaml
// and adds it to the blockchain.
func MakeGenesis(secCxt crypto.Peer) error {
	once.Do(func() {
		ledger, err := ledger.GetLedger()
		if err != nil {
			makeGenesisError = err
			return
		}

		if ledger.GetBlockchainSize() > 0 {
			// genesis block already exists
			return
		}

		genesisLogger.Info("Creating genesis block.")

		ledger.BeginTxBatch(0)
		var genesisTransactions []*protos.Transaction

		genesis := viper.GetStringMap("ledger.blockchain.genesisBlock")

		if genesis == nil {
			genesisLogger.Info("No genesis block chaincodes defined.")
			ledger.CommitTxBatch(0, genesisTransactions, nil)
			return
		}

		chaincodes, chaincodesOK := genesis["chaincode"].([]interface{})
		if !chaincodesOK {
			genesisLogger.Info("No genesis block chaincodes defined.")
			ledger.CommitTxBatch(0, genesisTransactions, nil)
			return
		}

		genesisLogger.Debug("Genesis chaincodes are %s", chaincodes)

		ledger.BeginTxBatch(0)
		for i := 0; i < len(chaincodes); i++ {
			genesisLogger.Debug("Chaincode %d is %s", i, chaincodes[i])

			chaincodeMap, chaincodeMapOK := chaincodes[i].(map[interface{}]interface{})
			if !chaincodeMapOK {
				genesisLogger.Error("Invalid chaincode defined in genesis configuration:", chaincodes[i])
				makeGenesisError = fmt.Errorf("Invalid chaincode defined in genesis configuration: %s", chaincodes[i])
				return
			}

			chaincodeIDMap, chaincodIDOK := chaincodeMap["id"].(map[interface{}]interface{})
			if !chaincodIDOK {
				genesisLogger.Error("Invalid chaincode ID defined in genesis configuration:", chaincodeMap["id"])
				makeGenesisError = fmt.Errorf("Invalid chaincode ID defined in genesis configuration: %s", chaincodeMap["id"])
				return
			}

			url, urlOK := chaincodeIDMap["url"].(string)
			if !urlOK {
				genesisLogger.Error("Invalid chaincode URL defined in genesis configuration:", chaincodeIDMap["url"])
				makeGenesisError = fmt.Errorf("Invalid chaincode URL defined in genesis configuration: %s", chaincodeIDMap["url"])
				return
			}

			version, versionOK := chaincodeIDMap["version"].(string)
			if !versionOK {
				genesisLogger.Error("Invalid chaincode version defined in genesis configuration:", chaincodeIDMap["version"])
				makeGenesisError = fmt.Errorf("Invalid chaincode version defined in genesis configuration: %s", chaincodeIDMap["version"])
				return
			}

			chaincodeType, chaincodeTypeOK := chaincodeMap["type"].(string)
			if !chaincodeTypeOK {
				genesisLogger.Error("Invalid chaincode type defined in genesis configuration:", chaincodeMap["type"])
				makeGenesisError = fmt.Errorf("Invalid chaincode type defined in genesis configuration: %s", chaincodeMap["type"])
				return
			}

			chaincodeID := &protos.ChaincodeID{Url: url, Version: version}

			genesisLogger.Debug("Genesis chaincodeID %s", chaincodeID)
			genesisLogger.Debug("Genesis chaincode type %s", chaincodeType)

			constructorMap, constructorMapOK := chaincodeMap["constructor"].(map[interface{}]interface{})
			if !constructorMapOK {
				genesisLogger.Error("Invalid chaincode constructor defined in genesis configuration:", chaincodeMap["constructor"])
				makeGenesisError = fmt.Errorf("Invalid chaincode constructor defined in genesis configuration: %s", chaincodeMap["constructor"])
				return
			}

			var spec protos.ChaincodeSpec
			if constructorMap == nil {
				genesisLogger.Debug("Genesis chaincode has no constructor.")
				spec = protos.ChaincodeSpec{Type: protos.ChaincodeSpec_Type(protos.ChaincodeSpec_Type_value[chaincodeType]), ChaincodeID: chaincodeID}
			} else {

				ctorFunc, ctorFuncOK := constructorMap["func"].(string)
				if !ctorFuncOK {
					genesisLogger.Error("Invalid chaincode constructor function defined in genesis configuration:", constructorMap["func"])
					makeGenesisError = fmt.Errorf("Invalid chaincode constructor function args defined in genesis configuration: %s", constructorMap["func"])
					return
				}

				ctorArgs, ctorArgsOK := constructorMap["args"].([]interface{})
				if !ctorArgsOK {
					genesisLogger.Error("Invalid chaincode constructor args defined in genesis configuration:", constructorMap["args"])
					makeGenesisError = fmt.Errorf("Invalid chaincode constructor args defined in genesis configuration: %s", constructorMap["args"])
					return
				}

				genesisLogger.Debug("Genesis chaincode constructor func %s", ctorFunc)
				genesisLogger.Debug("Genesis chaincode constructor args %s", ctorArgs)
				var ctorArgsStringArray []string
				for j := 0; j < len(ctorArgs); j++ {
					ctorArgsStringArray = append(ctorArgsStringArray, ctorArgs[j].(string))
				}

				spec = protos.ChaincodeSpec{Type: protos.ChaincodeSpec_Type(protos.ChaincodeSpec_Type_value[chaincodeType]), ChaincodeID: chaincodeID, CtorMsg: &protos.ChaincodeInput{Function: ctorFunc, Args: ctorArgsStringArray}}
			}

			transaction, _, deployErr := DeployLocal(context.Background(), &spec, secCxt)
			if deployErr != nil {
				genesisLogger.Error("Error deploying chaincode for genesis block.", deployErr)
				makeGenesisError = deployErr
				return
			}

			genesisTransactions = append(genesisTransactions, transaction)

		}

		genesisLogger.Info("Adding %d system chaincodes to the genesis block.", len(genesisTransactions))
		ledger.CommitTxBatch(0, genesisTransactions, nil)

	})
	return makeGenesisError
}

//BuildLocal builds a given chaincode code
func BuildLocal(context context.Context, spec *protos.ChaincodeSpec) (*protos.ChaincodeDeploymentSpec, error) {
	genesisLogger.Debug("Received build request for chaincode spec: %v", spec)
	mode := viper.GetString("chaincode.chaincoderunmode")
	var codePackageBytes []byte
	if mode != chaincode.DevModeUserRunsChaincode {
		if err := openchain.CheckSpec(spec); err != nil {
			genesisLogger.Debug("check spec failed: %s", err)
			return nil, err
		}
		// Get new VM and as for building of container image
		vm, err := container.NewVM()
		if err != nil {
			genesisLogger.Error(fmt.Sprintf("Error getting VM: %s", err))
			return nil, err
		}
		// Build the spec
		codePackageBytes, err = vm.BuildChaincodeContainer(spec)
		if err != nil {
			genesisLogger.Error(fmt.Sprintf("Error getting VM: %s", err))
			return nil, err
		}
	}
	chaincodeDeploymentSpec := &protos.ChaincodeDeploymentSpec{ChaincodeSpec: spec, CodePackage: codePackageBytes}
	return chaincodeDeploymentSpec, nil
}

// DeployLocal deploys the supplied chaincode image to the local peer
func DeployLocal(ctx context.Context, spec *protos.ChaincodeSpec, secCxt crypto.Peer) (*protos.Transaction, []byte, error) {
	// First build and get the deployment spec
	chaincodeDeploymentSpec, err := BuildLocal(ctx, spec)

	if err != nil {
		genesisLogger.Error(fmt.Sprintf("Error deploying chaincode spec: %v\n\n error: %s", spec, err))
		return nil, nil, err
	}

	// Now create the Transactions message and send to Peer.
	// The UUID must not be random so it will match on all peers.
	uuid := "genesis_" + spec.GetChaincodeID().Url + "_" + spec.GetChaincodeID().Version
	transaction, err := protos.NewChaincodeDeployTransaction(chaincodeDeploymentSpec, uuid)
	if err != nil {
		return nil, nil, fmt.Errorf("Error deploying chaincode: %s ", err)
	}
	//chaincode.NewChaincodeSupport(chaincode.DefaultChain, peer.GetPeerEndpoint, false, 120000)
	result, err := chaincode.Execute(ctx, chaincode.GetChain(chaincode.DefaultChain), transaction, secCxt)
	return transaction, result, err
}
