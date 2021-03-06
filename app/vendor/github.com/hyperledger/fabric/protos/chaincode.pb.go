// Code generated by protoc-gen-go.
// source: chaincode.proto
// DO NOT EDIT!

package protos

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

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
func (ConfidentialityLevel) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

type ChaincodeSpec_Type int32

const (
	ChaincodeSpec_UNDEFINED ChaincodeSpec_Type = 0
	ChaincodeSpec_GOLANG    ChaincodeSpec_Type = 1
	ChaincodeSpec_NODE      ChaincodeSpec_Type = 2
	ChaincodeSpec_CAR       ChaincodeSpec_Type = 3
	ChaincodeSpec_JAVA      ChaincodeSpec_Type = 4
)

var ChaincodeSpec_Type_name = map[int32]string{
	0: "UNDEFINED",
	1: "GOLANG",
	2: "NODE",
	3: "CAR",
	4: "JAVA",
}
var ChaincodeSpec_Type_value = map[string]int32{
	"UNDEFINED": 0,
	"GOLANG":    1,
	"NODE":      2,
	"CAR":       3,
	"JAVA":      4,
}

func (x ChaincodeSpec_Type) String() string {
	return proto.EnumName(ChaincodeSpec_Type_name, int32(x))
}
func (ChaincodeSpec_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{2, 0} }

type ChaincodeDeploymentSpec_ExecutionEnvironment int32

const (
	ChaincodeDeploymentSpec_DOCKER ChaincodeDeploymentSpec_ExecutionEnvironment = 0
	ChaincodeDeploymentSpec_SYSTEM ChaincodeDeploymentSpec_ExecutionEnvironment = 1
)

var ChaincodeDeploymentSpec_ExecutionEnvironment_name = map[int32]string{
	0: "DOCKER",
	1: "SYSTEM",
}
var ChaincodeDeploymentSpec_ExecutionEnvironment_value = map[string]int32{
	"DOCKER": 0,
	"SYSTEM": 1,
}

func (x ChaincodeDeploymentSpec_ExecutionEnvironment) String() string {
	return proto.EnumName(ChaincodeDeploymentSpec_ExecutionEnvironment_name, int32(x))
}
func (ChaincodeDeploymentSpec_ExecutionEnvironment) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor1, []int{3, 0}
}

type ChaincodeMessage_Type int32

const (
	ChaincodeMessage_UNDEFINED               ChaincodeMessage_Type = 0
	ChaincodeMessage_REGISTER                ChaincodeMessage_Type = 1
	ChaincodeMessage_REGISTERED              ChaincodeMessage_Type = 2
	ChaincodeMessage_INIT                    ChaincodeMessage_Type = 3
	ChaincodeMessage_READY                   ChaincodeMessage_Type = 4
	ChaincodeMessage_TRANSACTION             ChaincodeMessage_Type = 5
	ChaincodeMessage_COMPLETED               ChaincodeMessage_Type = 6
	ChaincodeMessage_ERROR                   ChaincodeMessage_Type = 7
	ChaincodeMessage_GET_STATE               ChaincodeMessage_Type = 8
	ChaincodeMessage_PUT_STATE               ChaincodeMessage_Type = 9
	ChaincodeMessage_DEL_STATE               ChaincodeMessage_Type = 10
	ChaincodeMessage_INVOKE_CHAINCODE        ChaincodeMessage_Type = 11
	ChaincodeMessage_INVOKE_QUERY            ChaincodeMessage_Type = 12
	ChaincodeMessage_RESPONSE                ChaincodeMessage_Type = 13
	ChaincodeMessage_QUERY                   ChaincodeMessage_Type = 14
	ChaincodeMessage_QUERY_COMPLETED         ChaincodeMessage_Type = 15
	ChaincodeMessage_QUERY_ERROR             ChaincodeMessage_Type = 16
	ChaincodeMessage_RANGE_QUERY_STATE       ChaincodeMessage_Type = 17
	ChaincodeMessage_RANGE_QUERY_STATE_NEXT  ChaincodeMessage_Type = 18
	ChaincodeMessage_RANGE_QUERY_STATE_CLOSE ChaincodeMessage_Type = 19
	ChaincodeMessage_KEEPALIVE               ChaincodeMessage_Type = 20
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
	18: "RANGE_QUERY_STATE_NEXT",
	19: "RANGE_QUERY_STATE_CLOSE",
	20: "KEEPALIVE",
}
var ChaincodeMessage_Type_value = map[string]int32{
	"UNDEFINED":               0,
	"REGISTER":                1,
	"REGISTERED":              2,
	"INIT":                    3,
	"READY":                   4,
	"TRANSACTION":             5,
	"COMPLETED":               6,
	"ERROR":                   7,
	"GET_STATE":               8,
	"PUT_STATE":               9,
	"DEL_STATE":               10,
	"INVOKE_CHAINCODE":        11,
	"INVOKE_QUERY":            12,
	"RESPONSE":                13,
	"QUERY":                   14,
	"QUERY_COMPLETED":         15,
	"QUERY_ERROR":             16,
	"RANGE_QUERY_STATE":       17,
	"RANGE_QUERY_STATE_NEXT":  18,
	"RANGE_QUERY_STATE_CLOSE": 19,
	"KEEPALIVE":               20,
}

func (x ChaincodeMessage_Type) String() string {
	return proto.EnumName(ChaincodeMessage_Type_name, int32(x))
}
func (ChaincodeMessage_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{6, 0} }

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

func (m *ChaincodeID) Reset()                    { *m = ChaincodeID{} }
func (m *ChaincodeID) String() string            { return proto.CompactTextString(m) }
func (*ChaincodeID) ProtoMessage()               {}
func (*ChaincodeID) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

// Carries the chaincode function and its arguments.
// UnmarshalJSON in transaction.go converts the string-based REST/JSON input to
// the []byte-based current ChaincodeInput structure.
type ChaincodeInput struct {
	Args [][]byte `protobuf:"bytes,1,rep,name=args,proto3" json:"args,omitempty"`
}

func (m *ChaincodeInput) Reset()                    { *m = ChaincodeInput{} }
func (m *ChaincodeInput) String() string            { return proto.CompactTextString(m) }
func (*ChaincodeInput) ProtoMessage()               {}
func (*ChaincodeInput) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

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
	Attributes           []string             `protobuf:"bytes,8,rep,name=attributes" json:"attributes,omitempty"`
}

func (m *ChaincodeSpec) Reset()                    { *m = ChaincodeSpec{} }
func (m *ChaincodeSpec) String() string            { return proto.CompactTextString(m) }
func (*ChaincodeSpec) ProtoMessage()               {}
func (*ChaincodeSpec) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

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
	EffectiveDate *google_protobuf.Timestamp                   `protobuf:"bytes,2,opt,name=effectiveDate" json:"effectiveDate,omitempty"`
	CodePackage   []byte                                       `protobuf:"bytes,3,opt,name=codePackage,proto3" json:"codePackage,omitempty"`
	ExecEnv       ChaincodeDeploymentSpec_ExecutionEnvironment `protobuf:"varint,4,opt,name=execEnv,enum=protos.ChaincodeDeploymentSpec_ExecutionEnvironment" json:"execEnv,omitempty"`
}

func (m *ChaincodeDeploymentSpec) Reset()                    { *m = ChaincodeDeploymentSpec{} }
func (m *ChaincodeDeploymentSpec) String() string            { return proto.CompactTextString(m) }
func (*ChaincodeDeploymentSpec) ProtoMessage()               {}
func (*ChaincodeDeploymentSpec) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

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
	// This field can contain a user-specified ID generation algorithm
	// If supplied, this will be used to generate a ID
	// If not supplied (left empty), sha256base64 will be used
	// The algorithm consists of two parts:
	//  1, a hash function
	//  2, a decoding used to decode user (string) input to bytes
	// Currently, SHA256 with BASE64 is supported (e.g. idGenerationAlg='sha256base64')
	IdGenerationAlg string `protobuf:"bytes,2,opt,name=idGenerationAlg" json:"idGenerationAlg,omitempty"`
}

func (m *ChaincodeInvocationSpec) Reset()                    { *m = ChaincodeInvocationSpec{} }
func (m *ChaincodeInvocationSpec) String() string            { return proto.CompactTextString(m) }
func (*ChaincodeInvocationSpec) ProtoMessage()               {}
func (*ChaincodeInvocationSpec) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func (m *ChaincodeInvocationSpec) GetChaincodeSpec() *ChaincodeSpec {
	if m != nil {
		return m.ChaincodeSpec
	}
	return nil
}

// This structure contain transaction data that we send to the chaincode
// container shim and allow the chaincode to access through the shim interface.
// TODO: Consider remove this message and just pass the transaction object
// to the shim and/or allow the chaincode to query transactions.
type ChaincodeSecurityContext struct {
	CallerCert     []byte                     `protobuf:"bytes,1,opt,name=callerCert,proto3" json:"callerCert,omitempty"`
	CallerSign     []byte                     `protobuf:"bytes,2,opt,name=callerSign,proto3" json:"callerSign,omitempty"`
	Payload        []byte                     `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
	Binding        []byte                     `protobuf:"bytes,4,opt,name=binding,proto3" json:"binding,omitempty"`
	Metadata       []byte                     `protobuf:"bytes,5,opt,name=metadata,proto3" json:"metadata,omitempty"`
	ParentMetadata []byte                     `protobuf:"bytes,6,opt,name=parentMetadata,proto3" json:"parentMetadata,omitempty"`
	TxTimestamp    *google_protobuf.Timestamp `protobuf:"bytes,7,opt,name=txTimestamp" json:"txTimestamp,omitempty"`
}

func (m *ChaincodeSecurityContext) Reset()                    { *m = ChaincodeSecurityContext{} }
func (m *ChaincodeSecurityContext) String() string            { return proto.CompactTextString(m) }
func (*ChaincodeSecurityContext) ProtoMessage()               {}
func (*ChaincodeSecurityContext) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{5} }

func (m *ChaincodeSecurityContext) GetTxTimestamp() *google_protobuf.Timestamp {
	if m != nil {
		return m.TxTimestamp
	}
	return nil
}

type ChaincodeMessage struct {
	Type            ChaincodeMessage_Type      `protobuf:"varint,1,opt,name=type,enum=protos.ChaincodeMessage_Type" json:"type,omitempty"`
	Timestamp       *google_protobuf.Timestamp `protobuf:"bytes,2,opt,name=timestamp" json:"timestamp,omitempty"`
	Payload         []byte                     `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
	Txid            string                     `protobuf:"bytes,4,opt,name=txid" json:"txid,omitempty"`
	SecurityContext *ChaincodeSecurityContext  `protobuf:"bytes,5,opt,name=securityContext" json:"securityContext,omitempty"`
	// event emmited by chaincode. Used only with Init or Invoke.
	// This event is then stored (currently)
	// with Block.NonHashData.TransactionResult
	ChaincodeEvent *ChaincodeEvent `protobuf:"bytes,6,opt,name=chaincodeEvent" json:"chaincodeEvent,omitempty"`
}

func (m *ChaincodeMessage) Reset()                    { *m = ChaincodeMessage{} }
func (m *ChaincodeMessage) String() string            { return proto.CompactTextString(m) }
func (*ChaincodeMessage) ProtoMessage()               {}
func (*ChaincodeMessage) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{6} }

func (m *ChaincodeMessage) GetTimestamp() *google_protobuf.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *ChaincodeMessage) GetSecurityContext() *ChaincodeSecurityContext {
	if m != nil {
		return m.SecurityContext
	}
	return nil
}

func (m *ChaincodeMessage) GetChaincodeEvent() *ChaincodeEvent {
	if m != nil {
		return m.ChaincodeEvent
	}
	return nil
}

type PutStateInfo struct {
	Key   string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Value []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *PutStateInfo) Reset()                    { *m = PutStateInfo{} }
func (m *PutStateInfo) String() string            { return proto.CompactTextString(m) }
func (*PutStateInfo) ProtoMessage()               {}
func (*PutStateInfo) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{7} }

type RangeQueryState struct {
	StartKey string `protobuf:"bytes,1,opt,name=startKey" json:"startKey,omitempty"`
	EndKey   string `protobuf:"bytes,2,opt,name=endKey" json:"endKey,omitempty"`
}

func (m *RangeQueryState) Reset()                    { *m = RangeQueryState{} }
func (m *RangeQueryState) String() string            { return proto.CompactTextString(m) }
func (*RangeQueryState) ProtoMessage()               {}
func (*RangeQueryState) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{8} }

type RangeQueryStateNext struct {
	ID string `protobuf:"bytes,1,opt,name=ID" json:"ID,omitempty"`
}

func (m *RangeQueryStateNext) Reset()                    { *m = RangeQueryStateNext{} }
func (m *RangeQueryStateNext) String() string            { return proto.CompactTextString(m) }
func (*RangeQueryStateNext) ProtoMessage()               {}
func (*RangeQueryStateNext) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{9} }

type RangeQueryStateClose struct {
	ID string `protobuf:"bytes,1,opt,name=ID" json:"ID,omitempty"`
}

func (m *RangeQueryStateClose) Reset()                    { *m = RangeQueryStateClose{} }
func (m *RangeQueryStateClose) String() string            { return proto.CompactTextString(m) }
func (*RangeQueryStateClose) ProtoMessage()               {}
func (*RangeQueryStateClose) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{10} }

type RangeQueryStateKeyValue struct {
	Key   string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Value []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *RangeQueryStateKeyValue) Reset()                    { *m = RangeQueryStateKeyValue{} }
func (m *RangeQueryStateKeyValue) String() string            { return proto.CompactTextString(m) }
func (*RangeQueryStateKeyValue) ProtoMessage()               {}
func (*RangeQueryStateKeyValue) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{11} }

type RangeQueryStateResponse struct {
	KeysAndValues []*RangeQueryStateKeyValue `protobuf:"bytes,1,rep,name=keysAndValues" json:"keysAndValues,omitempty"`
	HasMore       bool                       `protobuf:"varint,2,opt,name=hasMore" json:"hasMore,omitempty"`
	ID            string                     `protobuf:"bytes,3,opt,name=ID" json:"ID,omitempty"`
}

func (m *RangeQueryStateResponse) Reset()                    { *m = RangeQueryStateResponse{} }
func (m *RangeQueryStateResponse) String() string            { return proto.CompactTextString(m) }
func (*RangeQueryStateResponse) ProtoMessage()               {}
func (*RangeQueryStateResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{12} }

func (m *RangeQueryStateResponse) GetKeysAndValues() []*RangeQueryStateKeyValue {
	if m != nil {
		return m.KeysAndValues
	}
	return nil
}

func init() {
	proto.RegisterType((*ChaincodeID)(nil), "protos.ChaincodeID")
	proto.RegisterType((*ChaincodeInput)(nil), "protos.ChaincodeInput")
	proto.RegisterType((*ChaincodeSpec)(nil), "protos.ChaincodeSpec")
	proto.RegisterType((*ChaincodeDeploymentSpec)(nil), "protos.ChaincodeDeploymentSpec")
	proto.RegisterType((*ChaincodeInvocationSpec)(nil), "protos.ChaincodeInvocationSpec")
	proto.RegisterType((*ChaincodeSecurityContext)(nil), "protos.ChaincodeSecurityContext")
	proto.RegisterType((*ChaincodeMessage)(nil), "protos.ChaincodeMessage")
	proto.RegisterType((*PutStateInfo)(nil), "protos.PutStateInfo")
	proto.RegisterType((*RangeQueryState)(nil), "protos.RangeQueryState")
	proto.RegisterType((*RangeQueryStateNext)(nil), "protos.RangeQueryStateNext")
	proto.RegisterType((*RangeQueryStateClose)(nil), "protos.RangeQueryStateClose")
	proto.RegisterType((*RangeQueryStateKeyValue)(nil), "protos.RangeQueryStateKeyValue")
	proto.RegisterType((*RangeQueryStateResponse)(nil), "protos.RangeQueryStateResponse")
	proto.RegisterEnum("protos.ConfidentialityLevel", ConfidentialityLevel_name, ConfidentialityLevel_value)
	proto.RegisterEnum("protos.ChaincodeSpec_Type", ChaincodeSpec_Type_name, ChaincodeSpec_Type_value)
	proto.RegisterEnum("protos.ChaincodeDeploymentSpec_ExecutionEnvironment", ChaincodeDeploymentSpec_ExecutionEnvironment_name, ChaincodeDeploymentSpec_ExecutionEnvironment_value)
	proto.RegisterEnum("protos.ChaincodeMessage_Type", ChaincodeMessage_Type_name, ChaincodeMessage_Type_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for ChaincodeSupport service

type ChaincodeSupportClient interface {
	Register(ctx context.Context, opts ...grpc.CallOption) (ChaincodeSupport_RegisterClient, error)
}

type chaincodeSupportClient struct {
	cc *grpc.ClientConn
}

func NewChaincodeSupportClient(cc *grpc.ClientConn) ChaincodeSupportClient {
	return &chaincodeSupportClient{cc}
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
	Register(ChaincodeSupport_RegisterServer) error
}

func RegisterChaincodeSupportServer(s *grpc.Server, srv ChaincodeSupportServer) {
	s.RegisterService(&_ChaincodeSupport_serviceDesc, srv)
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
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Register",
			Handler:       _ChaincodeSupport_Register_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: fileDescriptor1,
}

func init() { proto.RegisterFile("chaincode.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 1176 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xa4, 0x56, 0xcd, 0x6e, 0xdb, 0x46,
	0x10, 0x8e, 0x7e, 0x6c, 0x4b, 0xa3, 0xbf, 0xcd, 0x5a, 0x71, 0x08, 0xb5, 0x4d, 0x04, 0x22, 0x0d,
	0x84, 0x1e, 0x94, 0x54, 0x4d, 0x8a, 0x02, 0x2d, 0x82, 0x32, 0xe4, 0xc6, 0x65, 0x2c, 0x53, 0xca,
	0x8a, 0x36, 0x92, 0x93, 0x41, 0x53, 0x6b, 0x99, 0x88, 0x4c, 0x12, 0xe4, 0x4a, 0xb0, 0x6e, 0x3d,
	0xf7, 0xd4, 0x67, 0xe8, 0x43, 0xf4, 0xd0, 0x07, 0xea, 0x73, 0x14, 0xbb, 0x24, 0x65, 0xfd, 0x25,
	0x08, 0xd0, 0x93, 0x76, 0x66, 0xbe, 0x99, 0x9d, 0xfd, 0xf6, 0xdb, 0xa1, 0xa0, 0xe1, 0x5e, 0x3b,
	0x9e, 0xef, 0x06, 0x63, 0xd6, 0x0d, 0xa3, 0x80, 0x07, 0x78, 0x5f, 0xfe, 0xc4, 0xad, 0xe6, 0x32,
	0xc0, 0xe6, 0xcc, 0xe7, 0x49, 0xb4, 0xf5, 0x78, 0x12, 0x04, 0x93, 0x29, 0x7b, 0x26, 0xad, 0xcb,
	0xd9, 0xd5, 0x33, 0xee, 0xdd, 0xb0, 0x98, 0x3b, 0x37, 0x61, 0x02, 0x50, 0x5f, 0x42, 0x45, 0xcf,
	0x12, 0x4d, 0x03, 0x63, 0x28, 0x86, 0x0e, 0xbf, 0x56, 0x72, 0xed, 0x5c, 0xa7, 0x4c, 0xe5, 0x5a,
	0xf8, 0x7c, 0xe7, 0x86, 0x29, 0xf9, 0xc4, 0x27, 0xd6, 0xea, 0x13, 0xa8, 0xdf, 0xa5, 0xf9, 0xe1,
	0x8c, 0x0b, 0x94, 0x13, 0x4d, 0x62, 0x25, 0xd7, 0x2e, 0x74, 0xaa, 0x54, 0xae, 0xd5, 0xbf, 0x0b,
	0x50, 0x5b, 0xc2, 0x46, 0x21, 0x73, 0x71, 0x17, 0x8a, 0x7c, 0x11, 0x32, 0x59, 0xbf, 0xde, 0x6b,
	0x25, 0x4d, 0xc4, 0xdd, 0x35, 0x50, 0xd7, 0x5e, 0x84, 0x8c, 0x4a, 0x1c, 0x7e, 0x09, 0x15, 0xf7,
	0xae, 0x3d, 0xd9, 0x42, 0xa5, 0x77, 0xb8, 0x95, 0x66, 0x1a, 0x74, 0x15, 0x87, 0x9f, 0xc3, 0x81,
	0xcb, 0x83, 0xe8, 0x34, 0x9e, 0x28, 0x05, 0x99, 0x72, 0xb4, 0x9d, 0x22, 0xba, 0xa6, 0x19, 0x0c,
	0x2b, 0x70, 0x20, 0xa8, 0x09, 0x66, 0x5c, 0x29, 0xb6, 0x73, 0x9d, 0x3d, 0x9a, 0x99, 0xf8, 0x09,
	0xd4, 0x62, 0xe6, 0xce, 0x22, 0xa6, 0x07, 0x3e, 0x67, 0xb7, 0x5c, 0xd9, 0x93, 0x3c, 0xac, 0x3b,
	0xf1, 0x10, 0x9a, 0x6e, 0xe0, 0x5f, 0x79, 0x63, 0xe6, 0x73, 0xcf, 0x99, 0x7a, 0x7c, 0xd1, 0x67,
	0x73, 0x36, 0x55, 0xf6, 0xe5, 0x41, 0xbf, 0x5e, 0x6e, 0xbf, 0x03, 0x43, 0x77, 0x66, 0xe2, 0x16,
	0x94, 0x6e, 0x18, 0x77, 0xc6, 0x0e, 0x77, 0x94, 0x83, 0x76, 0xae, 0x53, 0xa5, 0x4b, 0x1b, 0x3f,
	0x02, 0x70, 0x38, 0x8f, 0xbc, 0xcb, 0x19, 0x67, 0xb1, 0x52, 0x6a, 0x17, 0x3a, 0x65, 0xba, 0xe2,
	0x51, 0x5f, 0x41, 0x51, 0x90, 0x88, 0x6b, 0x50, 0x3e, 0xb3, 0x0c, 0xf2, 0xc6, 0xb4, 0x88, 0x81,
	0xee, 0x61, 0x80, 0xfd, 0xe3, 0x41, 0x5f, 0xb3, 0x8e, 0x51, 0x0e, 0x97, 0xa0, 0x68, 0x0d, 0x0c,
	0x82, 0xf2, 0xf8, 0x00, 0x0a, 0xba, 0x46, 0x51, 0x41, 0xb8, 0xde, 0x6a, 0xe7, 0x1a, 0x2a, 0xaa,
	0xff, 0xe4, 0xe1, 0xe1, 0x92, 0x29, 0x83, 0x85, 0xd3, 0x60, 0x71, 0xc3, 0x7c, 0x2e, 0xaf, 0xf0,
	0x67, 0xa8, 0xb9, 0xab, 0xd7, 0x25, 0xef, 0xb2, 0xd2, 0x7b, 0xb0, 0xf3, 0x2e, 0xe9, 0x3a, 0x16,
	0xff, 0x0a, 0x35, 0x76, 0x75, 0xc5, 0x5c, 0xee, 0xcd, 0x99, 0xe1, 0x70, 0x96, 0xde, 0x68, 0xab,
	0x9b, 0xe8, 0xb4, 0x9b, 0xe9, 0xb4, 0x6b, 0x67, 0x3a, 0xa5, 0xeb, 0x09, 0xb8, 0x0d, 0x15, 0x51,
	0x6d, 0xe8, 0xb8, 0x1f, 0x9d, 0x09, 0x93, 0xd7, 0x5b, 0xa5, 0xab, 0x2e, 0x6c, 0xc1, 0x01, 0xbb,
	0x65, 0x2e, 0xf1, 0xe7, 0xf2, 0x2a, 0xeb, 0xbd, 0x17, 0x5b, 0xad, 0xad, 0x1f, 0xa9, 0x4b, 0x6e,
	0x99, 0x3b, 0xe3, 0x5e, 0xe0, 0x13, 0x7f, 0xee, 0x45, 0x81, 0x2f, 0x02, 0x34, 0x2b, 0xa2, 0x76,
	0xa1, 0xb9, 0x0b, 0x20, 0xd8, 0x34, 0x06, 0xfa, 0x09, 0xa1, 0x09, 0xb3, 0xa3, 0x0f, 0x23, 0x9b,
	0x9c, 0xa2, 0x9c, 0xfa, 0x7b, 0x6e, 0x85, 0x3c, 0xd3, 0x9f, 0x07, 0xae, 0x23, 0x52, 0xff, 0x3f,
	0x79, 0x1d, 0x68, 0x78, 0xe3, 0x63, 0xe6, 0xb3, 0x48, 0x16, 0xd4, 0xa6, 0x93, 0xf4, 0x4d, 0x6e,
	0xba, 0xd5, 0x3f, 0xf3, 0xa0, 0xdc, 0x95, 0x12, 0x42, 0xf5, 0xf8, 0x22, 0x93, 0xea, 0x23, 0x00,
	0xd7, 0x99, 0x4e, 0x59, 0xa4, 0xb3, 0x88, 0xcb, 0x06, 0xaa, 0x74, 0xc5, 0x73, 0x17, 0x1f, 0x79,
	0x13, 0x5f, 0xee, 0xb0, 0x8c, 0x0b, 0x8f, 0x78, 0x2a, 0xa1, 0xb3, 0x98, 0x06, 0xce, 0x38, 0x65,
	0x3f, 0x33, 0x45, 0xe4, 0xd2, 0xf3, 0xc7, 0x9e, 0x3f, 0x91, 0xcc, 0x57, 0x69, 0x66, 0xae, 0x89,
	0x79, 0x6f, 0x43, 0xcc, 0x4f, 0xa1, 0x1e, 0x3a, 0x11, 0xf3, 0xf9, 0x69, 0x86, 0xd8, 0x97, 0x88,
	0x0d, 0x2f, 0xfe, 0x05, 0x2a, 0xfc, 0x76, 0xa9, 0x0b, 0xf9, 0x26, 0x3e, 0xaf, 0x9c, 0x55, 0xb8,
	0xfa, 0xd7, 0x1e, 0xa0, 0x25, 0x25, 0xa7, 0x2c, 0x8e, 0x85, 0x54, 0xbe, 0x5f, 0x1b, 0x47, 0xdf,
	0x6c, 0xdd, 0x42, 0x8a, 0x5b, 0x9d, 0x48, 0x3f, 0x41, 0x79, 0x39, 0x43, 0xbf, 0x40, 0xbd, 0x77,
	0xe0, 0xcf, 0xf0, 0x86, 0xa1, 0xc8, 0x6f, 0xbd, 0xb1, 0x24, 0xad, 0x4c, 0xe5, 0x1a, 0xbf, 0x85,
	0x46, 0xbc, 0x7e, 0x71, 0x92, 0xb8, 0x4a, 0xaf, 0xbd, 0xad, 0x95, 0x75, 0x1c, 0xdd, 0x4c, 0xc4,
	0xaf, 0xa0, 0xbe, 0x54, 0x12, 0x11, 0x5f, 0x07, 0xc9, 0xf0, 0xae, 0xa9, 0x28, 0xa3, 0x74, 0x03,
	0xad, 0xfe, 0x9b, 0xdf, 0x3d, 0x4f, 0xaa, 0x50, 0xa2, 0xe4, 0xd8, 0x1c, 0xd9, 0x84, 0xa2, 0x1c,
	0xae, 0x03, 0x64, 0x16, 0x31, 0x50, 0x5e, 0x8c, 0x13, 0xd3, 0x32, 0x6d, 0x54, 0xc0, 0x65, 0xd8,
	0xa3, 0x44, 0x33, 0x3e, 0xa0, 0x22, 0x6e, 0x40, 0xc5, 0xa6, 0x9a, 0x35, 0xd2, 0x74, 0xdb, 0x1c,
	0x58, 0x68, 0x4f, 0x94, 0xd4, 0x07, 0xa7, 0xc3, 0x3e, 0xb1, 0x89, 0x81, 0xf6, 0x05, 0x94, 0x50,
	0x3a, 0xa0, 0xe8, 0x40, 0x44, 0x8e, 0x89, 0x7d, 0x31, 0xb2, 0x35, 0x9b, 0xa0, 0x92, 0x30, 0x87,
	0x67, 0x99, 0x59, 0x16, 0xa6, 0x41, 0xfa, 0xa9, 0x09, 0xb8, 0x09, 0xc8, 0xb4, 0xce, 0x07, 0x27,
	0xe4, 0x42, 0xff, 0x4d, 0x33, 0x2d, 0x5d, 0x8c, 0xb6, 0x0a, 0x46, 0x50, 0x4d, 0xbd, 0xef, 0xce,
	0x08, 0xfd, 0x80, 0xaa, 0x49, 0xcb, 0xa3, 0xe1, 0xc0, 0x1a, 0x11, 0x54, 0x13, 0xbb, 0x25, 0x81,
	0x3a, 0x3e, 0x84, 0x86, 0x5c, 0x5e, 0xdc, 0x75, 0xd3, 0x10, 0xdd, 0x26, 0xce, 0xa4, 0x27, 0x84,
	0x1f, 0xc0, 0x7d, 0xaa, 0x59, 0xc7, 0x69, 0xbd, 0x74, 0xf7, 0xfb, 0xb8, 0x05, 0x47, 0x5b, 0xee,
	0x0b, 0x8b, 0xbc, 0xb7, 0x11, 0xc6, 0x5f, 0xc1, 0xc3, 0xed, 0x98, 0xde, 0x1f, 0x8c, 0x08, 0x3a,
	0x14, 0xa7, 0x38, 0x21, 0x64, 0xa8, 0xf5, 0xcd, 0x73, 0x82, 0x9a, 0xea, 0x8f, 0x50, 0x1d, 0xce,
	0xf8, 0x88, 0x3b, 0x9c, 0x99, 0xfe, 0x55, 0x80, 0x11, 0x14, 0x3e, 0xb2, 0x45, 0xfa, 0x35, 0x16,
	0x4b, 0xdc, 0x84, 0xbd, 0xb9, 0x33, 0x9d, 0xb1, 0xf4, 0x5d, 0x26, 0x86, 0x4a, 0xa0, 0x41, 0x1d,
	0x7f, 0xc2, 0xde, 0xcd, 0x58, 0xb4, 0x90, 0xe9, 0xe2, 0xc5, 0xc5, 0xdc, 0x89, 0xf8, 0xc9, 0x32,
	0x7f, 0x69, 0xe3, 0x23, 0xd8, 0x67, 0xfe, 0x58, 0x44, 0x92, 0xf9, 0x91, 0x5a, 0xea, 0xb7, 0x70,
	0xb8, 0x51, 0xc6, 0x12, 0xf2, 0xa9, 0x43, 0xde, 0x34, 0xd2, 0x22, 0x79, 0xd3, 0x50, 0x9f, 0x42,
	0x73, 0x03, 0xa6, 0x4f, 0x83, 0x98, 0x6d, 0xe1, 0x34, 0x78, 0xb8, 0x81, 0x3b, 0x61, 0x8b, 0x73,
	0xd1, 0xf0, 0x17, 0x1f, 0xec, 0x8f, 0xdc, 0x56, 0x0d, 0xca, 0xe2, 0x30, 0xf0, 0x63, 0x86, 0x09,
	0xd4, 0x3e, 0xb2, 0x45, 0xac, 0xf9, 0x63, 0x59, 0x33, 0xf9, 0xeb, 0x51, 0xe9, 0x3d, 0xce, 0x44,
	0xfd, 0x89, 0xbd, 0xe9, 0x7a, 0x96, 0x78, 0x96, 0xd7, 0x4e, 0x7c, 0x1a, 0x44, 0xc9, 0xd6, 0x25,
	0x9a, 0x99, 0xe9, 0x79, 0x0a, 0xd9, 0x79, 0xbe, 0x7b, 0x01, 0xcd, 0x5d, 0xdf, 0x6f, 0x31, 0xfc,
	0x87, 0x67, 0xaf, 0xfb, 0xa6, 0x8e, 0xee, 0x09, 0xc5, 0xe9, 0x03, 0xeb, 0x8d, 0x69, 0x10, 0xcb,
	0x36, 0xb5, 0x3e, 0xca, 0xf5, 0xde, 0xaf, 0xcc, 0x9d, 0xd1, 0x2c, 0x0c, 0x83, 0x88, 0x63, 0x03,
	0x4a, 0x94, 0x4d, 0xbc, 0x98, 0xb3, 0x08, 0x2b, 0x9f, 0x9a, 0x3a, 0xad, 0x4f, 0x46, 0xd4, 0x7b,
	0x9d, 0xdc, 0xf3, 0xdc, 0x6b, 0x05, 0x8e, 0x82, 0x68, 0xd2, 0xbd, 0x5e, 0x84, 0x2c, 0x9a, 0xb2,
	0xf1, 0x84, 0x45, 0x69, 0xc2, 0x65, 0xf2, 0xa7, 0xf0, 0x87, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff,
	0xda, 0x5e, 0xbc, 0x46, 0x2e, 0x0a, 0x00, 0x00,
}
