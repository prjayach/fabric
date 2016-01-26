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

package crypto

import (
	obcca "github.com/openblockchain/obc-peer/obc-ca/protos"

	"errors"
	"github.com/openblockchain/obc-peer/openchain/crypto/utils"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func (node *nodeImpl) retrieveTCACertsChain(userID string) error {
	// Retrieve TCA certificate and verify it
	tcaCertRaw, err := node.getTCACertificate()
	if err != nil {
		node.log.Error("Failed getting TCA certificate [%s].", err.Error())

		return err
	}
	node.log.Debug("TCA certificate [%s]", utils.EncodeBase64(tcaCertRaw))

	// TODO: Test TCA cert againt root CA
	_, err = utils.DERToX509Certificate(tcaCertRaw)
	if err != nil {
		node.log.Error("Failed parsing TCA certificate [%s].", err.Error())

		return err
	}

	// Store TCA cert
	node.log.Debug("Storing TCA certificate for [%s]...", userID)

	if err := node.ks.storeCert(node.conf.getTCACertsChainFilename(), tcaCertRaw); err != nil {
		node.log.Error("Failed storing tca certificate [%s].", err.Error())
		return err
	}

	return nil
}

func (node *nodeImpl) loadTCACertsChain() error {
	// Load TCA certs chain
	node.log.Debug("Loading TCA certificates chain...")

	cert, err := node.ks.loadCert(node.conf.getTCACertsChainFilename())
	if err != nil {
		node.log.Error("Failed loading TCA certificates chain [%s].", err.Error())

		return err
	}

	ok := node.rootsCertPool.AppendCertsFromPEM(cert)
	if !ok {
		node.log.Error("Failed appending TCA certificates chain.")

		return errors.New("Failed appending TCA certificates chain.")
	}

	return nil
}

func (node *nodeImpl) getTCAClient() (*grpc.ClientConn, obcca.TCAPClient, error) {
	socket, err := grpc.Dial(node.conf.getTCAPAddr(), grpc.WithInsecure())
	if err != nil {
		node.log.Error("Failed dailing in [%s].", err.Error())

		return nil, nil, err
	}
	tcaPClient := obcca.NewTCAPClient(socket)

	return socket, tcaPClient, nil
}

func (node *nodeImpl) callTCAReadCACertificate(ctx context.Context, opts ...grpc.CallOption) (*obcca.Cert, error) {
	// Get a TCA Client
	sock, tcaP, err := node.getTCAClient()
	defer sock.Close()

	// Issue the request
	cert, err := tcaP.ReadCACertificate(ctx, &obcca.Empty{}, opts...)
	if err != nil {
		node.log.Error("Failed requesting tca read certificate [%s].", err.Error())

		return nil, err
	}

	return cert, nil
}

func (node *nodeImpl) getTCACertificate() ([]byte, error) {
	pbCert, err := node.callTCAReadCACertificate(context.Background())
	if err != nil {
		node.log.Error("Failed requesting tca certificate [%s].", err.Error())

		return nil, err
	}

	// TODO Verify pbCert.Cert

	return pbCert.Cert, nil
}
