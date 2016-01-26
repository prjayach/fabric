// Code generated by protoc-gen-go.
// source: chaincode.proto
// DO NOT EDIT!

package protos

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "google/protobuf"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Confidentiality Levels
type ConfidentialityLevel int32

const (
	ConfidentialityLevel_PUBLIC       ConfidentialityLevel = 0
	ConfidentialityLevel_CONFIDENTIAL ConfidentialityLevel = 1
)

var ConfidentialityLevel_name = map[int32]string{
	0: "PUBLIC",
	1: "CONFIDENTIAL",
}
var ConfidentialityLevel_value = map[string]int32{
	"PUBLIC":       0,
	"CONFIDENTIAL": 1,
}

func (x ConfidentialityLevel) String() string {
	return proto.EnumName(ConfidentialityLevel_name, int32(x))
}

type ChaincodeSpec_Type int32

const (
	ChaincodeSpec_UNDEFINED ChaincodeSpec_Type = 0
	ChaincodeSpec_GOLANG    ChaincodeSpec_Type = 1
	ChaincodeSpec_NODE      ChaincodeSpec_Type = 2
)

var ChaincodeSpec_Type_name = map[int32]string{
	0: "UNDEFINED",
	1: "GOLANG",
	2: "NODE",
}
var ChaincodeSpec_Type_value = map[string]int32{
	"UNDEFINED": 0,
	"GOLANG":    1,
	"NODE":      2,
}

func (x ChaincodeSpec_Type) String() string {
	return proto.EnumName(ChaincodeSpec_Type_name, int32(x))
}

type ChaincodeMessage_Type int32

const (
	ChaincodeMessage_UNDEFINED         ChaincodeMessage_Type = 0
	ChaincodeMessage_REGISTER          ChaincodeMessage_Type = 1
	ChaincodeMessage_REGISTERED        ChaincodeMessage_Type = 2
	ChaincodeMessage_INIT              ChaincodeMessage_Type = 3
	ChaincodeMessage_READY             ChaincodeMessage_Type = 4
	ChaincodeMessage_TRANSACTION       ChaincodeMessage_Type = 5
	ChaincodeMessage_COMPLETED         ChaincodeMessage_Type = 6
	ChaincodeMessage_ERROR             ChaincodeMessage_Type = 7
	ChaincodeMessage_GET_STATE         ChaincodeMessage_Type = 8
	ChaincodeMessage_PUT_STATE         ChaincodeMessage_Type = 9
	ChaincodeMessage_DEL_STATE         ChaincodeMessage_Type = 10
	ChaincodeMessage_INVOKE_CHAINCODE  ChaincodeMessage_Type = 11
	ChaincodeMessage_INVOKE_QUERY      ChaincodeMessage_Type = 12
	ChaincodeMessage_RESPONSE          ChaincodeMessage_Type = 13
	ChaincodeMessage_QUERY             ChaincodeMessage_Type = 14
	ChaincodeMessage_QUERY_COMPLETED   ChaincodeMessage_Type = 15
	ChaincodeMessage_QUERY_ERROR       ChaincodeMessage_Type = 16
	ChaincodeMessage_RANGE_QUERY_STATE ChaincodeMessage_Type = 17
)

var ChaincodeMessage_Type_name = map[int32]string{
	0:  "UNDEFINED",
	1:  "REGISTER",
	2:  "REGISTERED",
	3:  "INIT",
	4:  "READY",
	5:  "TRANSACTION",
	6:  "COMPLETED",
	7:  "ERROR",
	8:  "GET_STATE",
	9:  "PUT_STATE",
	10: "DEL_STATE",
	11: "INVOKE_CHAINCODE",
	12: "INVOKE_QUERY",
	13: "RESPONSE",
	14: "QUERY",
	15: "QUERY_COMPLETED",
	16: "QUERY_ERROR",
	17: "RANGE_QUERY_STATE",
}
var ChaincodeMessage_Type_value = map[string]int32{
	"UNDEFINED":         0,
	"REGISTER":          1,
	"REGISTERED":        2,
	"INIT":              3,
	"READY":             4,
	"TRANSACTION":       5,
	"COMPLETED":         6,
	"ERROR":             7,
	"GET_STATE":         8,
	"PUT_STATE":         9,
	"DEL_STATE":         10,
	"INVOKE_CHAINCODE":  11,
	"INVOKE_QUERY":      12,
	"RESPONSE":          13,
	"QUERY":             14,
	"QUERY_COMPLETED":   15,
	"QUERY_ERROR":       16,
	"RANGE_QUERY_STATE": 17,
}

func (x ChaincodeMessage_Type) String() string {
	return proto.EnumName(ChaincodeMessage_Type_name, int32(x))
}

// ChaincodeID contains the path as specified by the deploy transaction
// that created it as well as the hashCode that is generated by the
// system for the path. From the user level (ie, CLI, REST API and so on)
// deploy transaction is expected to provide the path and other requests
// are expected to provide the hashCode. The other value will be ignored.
// Internally, the structure could contain both values. For instance, the
// hashCode will be set when first generated using the path
type ChaincodeID struct {
	// deploy transaction will use the path
	Path string `protobuf:"bytes,1,opt,name=path" json:"path,omitempty"`
	// all other requests will use the name (really a hashcode) generated by
	// the deploy transaction
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
}

func (m *ChaincodeID) Reset()         { *m = ChaincodeID{} }
func (m *ChaincodeID) String() string { return proto.CompactTextString(m) }
func (*ChaincodeID) ProtoMessage()    {}

// Carries the chaincode function and its arguments.
type ChaincodeInput struct {
	Function string   `protobuf:"bytes,1,opt,name=function" json:"function,omitempty"`
	Args     []string `protobuf:"bytes,2,rep,name=args" json:"args,omitempty"`
}

func (m *ChaincodeInput) Reset()         { *m = ChaincodeInput{} }
func (m *ChaincodeInput) String() string { return proto.CompactTextString(m) }
func (*ChaincodeInput) ProtoMessage()    {}

// Carries the chaincode specification. This is the actual metadata required for
// defining a chaincode.
type ChaincodeSpec struct {
	Type                 ChaincodeSpec_Type   `protobuf:"varint,1,opt,name=type,enum=protos.ChaincodeSpec_Type" json:"type,omitempty"`
	ChaincodeID          *ChaincodeID         `protobuf:"bytes,2,opt,name=chaincodeID" json:"chaincodeID,omitempty"`
	CtorMsg              *ChaincodeInput      `protobuf:"bytes,3,opt,name=ctorMsg" json:"ctorMsg,omitempty"`
	Timeout              int32                `protobuf:"varint,4,opt,name=timeout" json:"timeout,omitempty"`
	SecureContext        string               `protobuf:"bytes,5,opt,name=secureContext" json:"secureContext,omitempty"`
	ConfidentialityLevel ConfidentialityLevel `protobuf:"varint,6,opt,name=confidentialityLevel,enum=protos.ConfidentialityLevel" json:"confidentialityLevel,omitempty"`
	Metadata             []byte               `protobuf:"bytes,7,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (m *ChaincodeSpec) Reset()         { *m = ChaincodeSpec{} }
func (m *ChaincodeSpec) String() string { return proto.CompactTextString(m) }
func (*ChaincodeSpec) ProtoMessage()    {}

func (m *ChaincodeSpec) GetChaincodeID() *ChaincodeID {
	if m != nil {
		return m.ChaincodeID
	}
	return nil
}

func (m *ChaincodeSpec) GetCtorMsg() *ChaincodeInput {
	if m != nil {
		return m.CtorMsg
	}
	return nil
}

// Specify the deployment of a chaincode.
// TODO: Define `codePackage`.
type ChaincodeDeploymentSpec struct {
	ChaincodeSpec *ChaincodeSpec `protobuf:"bytes,1,opt,name=chaincodeSpec" json:"chaincodeSpec,omitempty"`
	// Controls when the chaincode becomes executable.
	EffectiveDate *google_protobuf.Timestamp `protobuf:"bytes,2,opt,name=effectiveDate" json:"effectiveDate,omitempty"`
	CodePackage   []byte                     `protobuf:"bytes,3,opt,name=codePackage,proto3" json:"codePackage,omitempty"`
}

func (m *ChaincodeDeploymentSpec) Reset()         { *m = ChaincodeDeploymentSpec{} }
func (m *ChaincodeDeploymentSpec) String() string { return proto.CompactTextString(m) }
func (*ChaincodeDeploymentSpec) ProtoMessage()    {}

func (m *ChaincodeDeploymentSpec) GetChaincodeSpec() *ChaincodeSpec {
	if m != nil {
		return m.ChaincodeSpec
	}
	return nil
}

func (m *ChaincodeDeploymentSpec) GetEffectiveDate() *google_protobuf.Timestamp {
	if m != nil {
		return m.EffectiveDate
	}
	return nil
}

// Carries the chaincode function and its arguments.
type ChaincodeInvocationSpec struct {
	ChaincodeSpec *ChaincodeSpec `protobuf:"bytes,1,opt,name=chaincodeSpec" json:"chaincodeSpec,omitempty"`
}

func (m *ChaincodeInvocationSpec) Reset()         { *m = ChaincodeInvocationSpec{} }
func (m *ChaincodeInvocationSpec) String() string { return proto.CompactTextString(m) }
func (*ChaincodeInvocationSpec) ProtoMessage()    {}

func (m *ChaincodeInvocationSpec) GetChaincodeSpec() *ChaincodeSpec {
	if m != nil {
		return m.ChaincodeSpec
	}
	return nil
}

// TODO: Merge this with ChaincodeID.
type ChaincodeIdentifier struct {
	// URL for accessing the Chaincode, eg. https://github.com/user/SampleContract
	Url string `protobuf:"bytes,1,opt,name=Url" json:"Url,omitempty"`
}

func (m *ChaincodeIdentifier) Reset()         { *m = ChaincodeIdentifier{} }
func (m *ChaincodeIdentifier) String() string { return proto.CompactTextString(m) }
func (*ChaincodeIdentifier) ProtoMessage()    {}

// Used by the peer to identify the requesting chaincode and allows for proper
// access to state.
type ChaincodeRequestContext struct {
	Id *ChaincodeIdentifier `protobuf:"bytes,1,opt,name=Id" json:"Id,omitempty"`
}

func (m *ChaincodeRequestContext) Reset()         { *m = ChaincodeRequestContext{} }
func (m *ChaincodeRequestContext) String() string { return proto.CompactTextString(m) }
func (*ChaincodeRequestContext) ProtoMessage()    {}

func (m *ChaincodeRequestContext) GetId() *ChaincodeIdentifier {
	if m != nil {
		return m.Id
	}
	return nil
}

// Provided by the peer to the chaincode to identify the requesting chaincode
// and allow for proper access to state.
type ChaincodeExecutionContext struct {
	ChaincodeId *ChaincodeIdentifier       `protobuf:"bytes,1,opt,name=ChaincodeId" json:"ChaincodeId,omitempty"`
	Timestamp   *google_protobuf.Timestamp `protobuf:"bytes,2,opt,name=Timestamp" json:"Timestamp,omitempty"`
}

func (m *ChaincodeExecutionContext) Reset()         { *m = ChaincodeExecutionContext{} }
func (m *ChaincodeExecutionContext) String() string { return proto.CompactTextString(m) }
func (*ChaincodeExecutionContext) ProtoMessage()    {}

func (m *ChaincodeExecutionContext) GetChaincodeId() *ChaincodeIdentifier {
	if m != nil {
		return m.ChaincodeId
	}
	return nil
}

func (m *ChaincodeExecutionContext) GetTimestamp() *google_protobuf.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

type ChaincodeMessage struct {
	Type      ChaincodeMessage_Type      `protobuf:"varint,1,opt,name=type,enum=protos.ChaincodeMessage_Type" json:"type,omitempty"`
	Timestamp *google_protobuf.Timestamp `protobuf:"bytes,2,opt,name=timestamp" json:"timestamp,omitempty"`
	Payload   []byte                     `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
	Uuid      string                     `protobuf:"bytes,4,opt,name=uuid" json:"uuid,omitempty"`
}

func (m *ChaincodeMessage) Reset()         { *m = ChaincodeMessage{} }
func (m *ChaincodeMessage) String() string { return proto.CompactTextString(m) }
func (*ChaincodeMessage) ProtoMessage()    {}

func (m *ChaincodeMessage) GetTimestamp() *google_protobuf.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

type PutStateInfo struct {
	Key   string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Value []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *PutStateInfo) Reset()         { *m = PutStateInfo{} }
func (m *PutStateInfo) String() string { return proto.CompactTextString(m) }
func (*PutStateInfo) ProtoMessage()    {}

type RangeQueryStateInfo struct {
	StartKey string `protobuf:"bytes,1,opt,name=startKey" json:"startKey,omitempty"`
	EndKey   string `protobuf:"bytes,2,opt,name=endKey" json:"endKey,omitempty"`
	Limit    uint32 `protobuf:"varint,3,opt,name=limit" json:"limit,omitempty"`
}

func (m *RangeQueryStateInfo) Reset()         { *m = RangeQueryStateInfo{} }
func (m *RangeQueryStateInfo) String() string { return proto.CompactTextString(m) }
func (*RangeQueryStateInfo) ProtoMessage()    {}

type RangeQueryStateKeyValue struct {
	Key   string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Value []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *RangeQueryStateKeyValue) Reset()         { *m = RangeQueryStateKeyValue{} }
func (m *RangeQueryStateKeyValue) String() string { return proto.CompactTextString(m) }
func (*RangeQueryStateKeyValue) ProtoMessage()    {}

type RangeQueryStateResponse struct {
	KeysAndValues []*RangeQueryStateKeyValue `protobuf:"bytes,1,rep,name=keysAndValues" json:"keysAndValues,omitempty"`
	HasMore       bool                       `protobuf:"varint,2,opt,name=hasMore" json:"hasMore,omitempty"`
}

func (m *RangeQueryStateResponse) Reset()         { *m = RangeQueryStateResponse{} }
func (m *RangeQueryStateResponse) String() string { return proto.CompactTextString(m) }
func (*RangeQueryStateResponse) ProtoMessage()    {}

func (m *RangeQueryStateResponse) GetKeysAndValues() []*RangeQueryStateKeyValue {
	if m != nil {
		return m.KeysAndValues
	}
	return nil
}

func init() {
	proto.RegisterEnum("protos.ConfidentialityLevel", ConfidentialityLevel_name, ConfidentialityLevel_value)
	proto.RegisterEnum("protos.ChaincodeSpec_Type", ChaincodeSpec_Type_name, ChaincodeSpec_Type_value)
	proto.RegisterEnum("protos.ChaincodeMessage_Type", ChaincodeMessage_Type_name, ChaincodeMessage_Type_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for ChaincodeSupport service

type ChaincodeSupportClient interface {
	// Return the datetime.
	GetExecutionContext(ctx context.Context, in *ChaincodeRequestContext, opts ...grpc.CallOption) (*ChaincodeExecutionContext, error)
	Register(ctx context.Context, opts ...grpc.CallOption) (ChaincodeSupport_RegisterClient, error)
}

type chaincodeSupportClient struct {
	cc *grpc.ClientConn
}

func NewChaincodeSupportClient(cc *grpc.ClientConn) ChaincodeSupportClient {
	return &chaincodeSupportClient{cc}
}

func (c *chaincodeSupportClient) GetExecutionContext(ctx context.Context, in *ChaincodeRequestContext, opts ...grpc.CallOption) (*ChaincodeExecutionContext, error) {
	out := new(ChaincodeExecutionContext)
	err := grpc.Invoke(ctx, "/protos.ChaincodeSupport/GetExecutionContext", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chaincodeSupportClient) Register(ctx context.Context, opts ...grpc.CallOption) (ChaincodeSupport_RegisterClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_ChaincodeSupport_serviceDesc.Streams[0], c.cc, "/protos.ChaincodeSupport/Register", opts...)
	if err != nil {
		return nil, err
	}
	x := &chaincodeSupportRegisterClient{stream}
	return x, nil
}

type ChaincodeSupport_RegisterClient interface {
	Send(*ChaincodeMessage) error
	Recv() (*ChaincodeMessage, error)
	grpc.ClientStream
}

type chaincodeSupportRegisterClient struct {
	grpc.ClientStream
}

func (x *chaincodeSupportRegisterClient) Send(m *ChaincodeMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *chaincodeSupportRegisterClient) Recv() (*ChaincodeMessage, error) {
	m := new(ChaincodeMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for ChaincodeSupport service

type ChaincodeSupportServer interface {
	// Return the datetime.
	GetExecutionContext(context.Context, *ChaincodeRequestContext) (*ChaincodeExecutionContext, error)
	Register(ChaincodeSupport_RegisterServer) error
}

func RegisterChaincodeSupportServer(s *grpc.Server, srv ChaincodeSupportServer) {
	s.RegisterService(&_ChaincodeSupport_serviceDesc, srv)
}

func _ChaincodeSupport_GetExecutionContext_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(ChaincodeRequestContext)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(ChaincodeSupportServer).GetExecutionContext(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _ChaincodeSupport_Register_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ChaincodeSupportServer).Register(&chaincodeSupportRegisterServer{stream})
}

type ChaincodeSupport_RegisterServer interface {
	Send(*ChaincodeMessage) error
	Recv() (*ChaincodeMessage, error)
	grpc.ServerStream
}

type chaincodeSupportRegisterServer struct {
	grpc.ServerStream
}

func (x *chaincodeSupportRegisterServer) Send(m *ChaincodeMessage) error {
	return x.ServerStream.SendMsg(m)
}

func (x *chaincodeSupportRegisterServer) Recv() (*ChaincodeMessage, error) {
	m := new(ChaincodeMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _ChaincodeSupport_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protos.ChaincodeSupport",
	HandlerType: (*ChaincodeSupportServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetExecutionContext",
			Handler:    _ChaincodeSupport_GetExecutionContext_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Register",
			Handler:       _ChaincodeSupport_Register_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
}
