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
	
	"github.com/openblockchain/obc-peer/openchain/crypto/utils"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"github.com/golang/protobuf/proto"
	"google/protobuf"
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/x509"
	"time"
	"github.com/openblockchain/obc-peer/openchain/util"
)

func (node *nodeImpl) getTLSCertificateFromTLSCA(id, affiliation string) (interface{}, []byte, error) {
	node.log.Info("getTLSCertificate...")
	
	priv, err := utils.NewECDSAKey()

	if err != nil {
		node.log.Error("Failed generating key: %s", err)

		return nil, nil, err
	}
	
	uuid, err := util.GenerateUUID()
	if err != nil {
		node.log.Error("Failed generating uuid: %s", err)

		return nil, nil, err
	}

	// Prepare the request
	pubraw, _ := x509.MarshalPKIXPublicKey(&priv.PublicKey)
	now := time.Now()
	timestamp := google_protobuf.Timestamp{int64(now.Second()), int32(now.Nanosecond())}
	
	req := &obcca.TLSCertCreateReq{
		&timestamp,
		&obcca.Identity{Id: id + "-" + uuid},
		&obcca.Password{Pw: ""},
		&obcca.PublicKey{
			Type: obcca.CryptoType_ECDSA,
			Key:  pubraw,
		}, nil}
	rawreq, _ := proto.Marshal(req)
	r, s, err := ecdsa.Sign(rand.Reader, priv, utils.Hash(rawreq))
	if err != nil {
		panic(err)
	}
	R, _ := r.MarshalText()
	S, _ := s.MarshalText()
	req.Sig = &obcca.Signature{obcca.CryptoType_ECDSA, R, S}

	pbCert, err := node.callTLSCACreateCertificate(context.Background(), req)
	if err != nil {
		node.log.Error("Failed requesting tls certificate: %s", err)

		return nil, nil, err
	}

	node.log.Info("Verifing tls certificate...")

	tlsCert, err := utils.DERToX509Certificate(pbCert.Cert)
	certPK := tlsCert.PublicKey.(*ecdsa.PublicKey)
	utils.VerifySignCapability(priv, certPK)

	node.log.Info("Verifing tls certificate...done!")

	return priv, pbCert.Cert, nil
}

func (node *nodeImpl) callTLSCACreateCertificate(ctx context.Context, in *obcca.TLSCertCreateReq, opts ...grpc.CallOption) (*obcca.Cert, error) {
	sockP, err := grpc.Dial(node.conf.getTLSCAPAddr(), grpc.WithInsecure())
	if err != nil {
		node.log.Error("Failed dialing in: %s", err)

		return nil, err
	}
	defer sockP.Close()

	tlscaP := obcca.NewTLSCAPClient(sockP)

	cert, err := tlscaP.CreateCertificate(context.Background(), in)
	if err != nil {
		node.log.Error("Failed requesting tls certificate: %s", err)

		return nil, err
	}

	return cert, nil
}