// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        (unknown)
// source: dht/dhtnode/message.proto

package dhtnode

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Status_ErrCode int32

const (
	// common (0x)
	Status_SUCCESS                   Status_ErrCode = 0
	Status_ERROR_UNSUPPORTED_VERSION Status_ErrCode = 1
	Status_ERROR_UNEXPECTED_PAYLOAD  Status_ErrCode = 2
	Status_ERROR_GENERIC             Status_ErrCode = 3
	Status_ERROR_SERIALIZATION       Status_ErrCode = 4
	// register (1x)
	Status_ERROR_WRONG_AGENT_ADDRESS Status_ErrCode = 10
	Status_ERROR_WRONG_PUBLIC_KEY    Status_ErrCode = 11
	Status_ERROR_INVALID_PROOF       Status_ErrCode = 12
	Status_ERROR_UNSUPPORTED_LEDGER  Status_ErrCode = 13
	// lookup & delivery (2x)
	Status_ERROR_UNKNOWN_AGENT_ADDRESS Status_ErrCode = 20
	Status_ERROR_AGENT_NOT_READY       Status_ErrCode = 21
)

// Enum value maps for Status_ErrCode.
var (
	Status_ErrCode_name = map[int32]string{
		0:  "SUCCESS",
		1:  "ERROR_UNSUPPORTED_VERSION",
		2:  "ERROR_UNEXPECTED_PAYLOAD",
		3:  "ERROR_GENERIC",
		4:  "ERROR_SERIALIZATION",
		10: "ERROR_WRONG_AGENT_ADDRESS",
		11: "ERROR_WRONG_PUBLIC_KEY",
		12: "ERROR_INVALID_PROOF",
		13: "ERROR_UNSUPPORTED_LEDGER",
		20: "ERROR_UNKNOWN_AGENT_ADDRESS",
		21: "ERROR_AGENT_NOT_READY",
	}
	Status_ErrCode_value = map[string]int32{
		"SUCCESS":                     0,
		"ERROR_UNSUPPORTED_VERSION":   1,
		"ERROR_UNEXPECTED_PAYLOAD":    2,
		"ERROR_GENERIC":               3,
		"ERROR_SERIALIZATION":         4,
		"ERROR_WRONG_AGENT_ADDRESS":   10,
		"ERROR_WRONG_PUBLIC_KEY":      11,
		"ERROR_INVALID_PROOF":         12,
		"ERROR_UNSUPPORTED_LEDGER":    13,
		"ERROR_UNKNOWN_AGENT_ADDRESS": 20,
		"ERROR_AGENT_NOT_READY":       21,
	}
)

func (x Status_ErrCode) Enum() *Status_ErrCode {
	p := new(Status_ErrCode)
	*p = x
	return p
}

func (x Status_ErrCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status_ErrCode) Descriptor() protoreflect.EnumDescriptor {
	return file_dht_dhtnode_message_proto_enumTypes[0].Descriptor()
}

func (Status_ErrCode) Type() protoreflect.EnumType {
	return &file_dht_dhtnode_message_proto_enumTypes[0]
}

func (x Status_ErrCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status_ErrCode.Descriptor instead.
func (Status_ErrCode) EnumDescriptor() ([]byte, []int) {
	return file_dht_dhtnode_message_proto_rawDescGZIP(), []int{5, 0}
}

type AgentRecord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServiceId     string `protobuf:"bytes,1,opt,name=service_id,json=serviceId,proto3" json:"service_id,omitempty"`
	LedgerId      string `protobuf:"bytes,2,opt,name=ledger_id,json=ledgerId,proto3" json:"ledger_id,omitempty"`
	Address       string `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	PublicKey     string `protobuf:"bytes,4,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	PeerPublicKey string `protobuf:"bytes,5,opt,name=peer_public_key,json=peerPublicKey,proto3" json:"peer_public_key,omitempty"`
	Signature     string `protobuf:"bytes,6,opt,name=signature,proto3" json:"signature,omitempty"`
	NotBefore     string `protobuf:"bytes,7,opt,name=not_before,json=notBefore,proto3" json:"not_before,omitempty"`
	NotAfter      string `protobuf:"bytes,8,opt,name=not_after,json=notAfter,proto3" json:"not_after,omitempty"`
}

func (x *AgentRecord) Reset() {
	*x = AgentRecord{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dht_dhtnode_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AgentRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AgentRecord) ProtoMessage() {}

func (x *AgentRecord) ProtoReflect() protoreflect.Message {
	mi := &file_dht_dhtnode_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AgentRecord.ProtoReflect.Descriptor instead.
func (*AgentRecord) Descriptor() ([]byte, []int) {
	return file_dht_dhtnode_message_proto_rawDescGZIP(), []int{0}
}

func (x *AgentRecord) GetServiceId() string {
	if x != nil {
		return x.ServiceId
	}
	return ""
}

func (x *AgentRecord) GetLedgerId() string {
	if x != nil {
		return x.LedgerId
	}
	return ""
}

func (x *AgentRecord) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *AgentRecord) GetPublicKey() string {
	if x != nil {
		return x.PublicKey
	}
	return ""
}

func (x *AgentRecord) GetPeerPublicKey() string {
	if x != nil {
		return x.PeerPublicKey
	}
	return ""
}

func (x *AgentRecord) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

func (x *AgentRecord) GetNotBefore() string {
	if x != nil {
		return x.NotBefore
	}
	return ""
}

func (x *AgentRecord) GetNotAfter() string {
	if x != nil {
		return x.NotAfter
	}
	return ""
}

type Register struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Record *AgentRecord `protobuf:"bytes,1,opt,name=record,proto3" json:"record,omitempty"`
}

func (x *Register) Reset() {
	*x = Register{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dht_dhtnode_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Register) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Register) ProtoMessage() {}

func (x *Register) ProtoReflect() protoreflect.Message {
	mi := &file_dht_dhtnode_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Register.ProtoReflect.Descriptor instead.
func (*Register) Descriptor() ([]byte, []int) {
	return file_dht_dhtnode_message_proto_rawDescGZIP(), []int{1}
}

func (x *Register) GetRecord() *AgentRecord {
	if x != nil {
		return x.Record
	}
	return nil
}

type LookupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AgentAddress string `protobuf:"bytes,1,opt,name=agent_address,json=agentAddress,proto3" json:"agent_address,omitempty"`
}

func (x *LookupRequest) Reset() {
	*x = LookupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dht_dhtnode_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LookupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LookupRequest) ProtoMessage() {}

func (x *LookupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dht_dhtnode_message_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LookupRequest.ProtoReflect.Descriptor instead.
func (*LookupRequest) Descriptor() ([]byte, []int) {
	return file_dht_dhtnode_message_proto_rawDescGZIP(), []int{2}
}

func (x *LookupRequest) GetAgentAddress() string {
	if x != nil {
		return x.AgentAddress
	}
	return ""
}

type LookupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AgentRecord *AgentRecord `protobuf:"bytes,1,opt,name=agent_record,json=agentRecord,proto3" json:"agent_record,omitempty"`
}

func (x *LookupResponse) Reset() {
	*x = LookupResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dht_dhtnode_message_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LookupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LookupResponse) ProtoMessage() {}

func (x *LookupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dht_dhtnode_message_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LookupResponse.ProtoReflect.Descriptor instead.
func (*LookupResponse) Descriptor() ([]byte, []int) {
	return file_dht_dhtnode_message_proto_rawDescGZIP(), []int{3}
}

func (x *LookupResponse) GetAgentRecord() *AgentRecord {
	if x != nil {
		return x.AgentRecord
	}
	return nil
}

type AeaEnvelope struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// TOFIX(LR) import aea.Envelop type
	Envel  []byte       `protobuf:"bytes,1,opt,name=envel,proto3" json:"envel,omitempty"`
	Record *AgentRecord `protobuf:"bytes,2,opt,name=record,proto3" json:"record,omitempty"`
}

func (x *AeaEnvelope) Reset() {
	*x = AeaEnvelope{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dht_dhtnode_message_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AeaEnvelope) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AeaEnvelope) ProtoMessage() {}

func (x *AeaEnvelope) ProtoReflect() protoreflect.Message {
	mi := &file_dht_dhtnode_message_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AeaEnvelope.ProtoReflect.Descriptor instead.
func (*AeaEnvelope) Descriptor() ([]byte, []int) {
	return file_dht_dhtnode_message_proto_rawDescGZIP(), []int{4}
}

func (x *AeaEnvelope) GetEnvel() []byte {
	if x != nil {
		return x.Envel
	}
	return nil
}

func (x *AeaEnvelope) GetRecord() *AgentRecord {
	if x != nil {
		return x.Record
	}
	return nil
}

type Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code Status_ErrCode `protobuf:"varint,1,opt,name=code,proto3,enum=dhtnode.Status_ErrCode" json:"code,omitempty"`
	Msgs []string       `protobuf:"bytes,2,rep,name=msgs,proto3" json:"msgs,omitempty"`
}

func (x *Status) Reset() {
	*x = Status{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dht_dhtnode_message_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_dht_dhtnode_message_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status.ProtoReflect.Descriptor instead.
func (*Status) Descriptor() ([]byte, []int) {
	return file_dht_dhtnode_message_proto_rawDescGZIP(), []int{5}
}

func (x *Status) GetCode() Status_ErrCode {
	if x != nil {
		return x.Code
	}
	return Status_SUCCESS
}

func (x *Status) GetMsgs() []string {
	if x != nil {
		return x.Msgs
	}
	return nil
}

type AcnMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	// Types that are assignable to Payload:
	//	*AcnMessage_Status
	//	*AcnMessage_Register
	//	*AcnMessage_LookupRequest
	//	*AcnMessage_LookupResponse
	//	*AcnMessage_AeaEnvelope
	Payload isAcnMessage_Payload `protobuf_oneof:"payload"`
}

func (x *AcnMessage) Reset() {
	*x = AcnMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dht_dhtnode_message_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AcnMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AcnMessage) ProtoMessage() {}

func (x *AcnMessage) ProtoReflect() protoreflect.Message {
	mi := &file_dht_dhtnode_message_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AcnMessage.ProtoReflect.Descriptor instead.
func (*AcnMessage) Descriptor() ([]byte, []int) {
	return file_dht_dhtnode_message_proto_rawDescGZIP(), []int{6}
}

func (x *AcnMessage) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (m *AcnMessage) GetPayload() isAcnMessage_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (x *AcnMessage) GetStatus() *Status {
	if x, ok := x.GetPayload().(*AcnMessage_Status); ok {
		return x.Status
	}
	return nil
}

func (x *AcnMessage) GetRegister() *Register {
	if x, ok := x.GetPayload().(*AcnMessage_Register); ok {
		return x.Register
	}
	return nil
}

func (x *AcnMessage) GetLookupRequest() *LookupRequest {
	if x, ok := x.GetPayload().(*AcnMessage_LookupRequest); ok {
		return x.LookupRequest
	}
	return nil
}

func (x *AcnMessage) GetLookupResponse() *LookupResponse {
	if x, ok := x.GetPayload().(*AcnMessage_LookupResponse); ok {
		return x.LookupResponse
	}
	return nil
}

func (x *AcnMessage) GetAeaEnvelope() *AeaEnvelope {
	if x, ok := x.GetPayload().(*AcnMessage_AeaEnvelope); ok {
		return x.AeaEnvelope
	}
	return nil
}

type isAcnMessage_Payload interface {
	isAcnMessage_Payload()
}

type AcnMessage_Status struct {
	Status *Status `protobuf:"bytes,2,opt,name=status,proto3,oneof"`
}

type AcnMessage_Register struct {
	Register *Register `protobuf:"bytes,3,opt,name=register,proto3,oneof"`
}

type AcnMessage_LookupRequest struct {
	LookupRequest *LookupRequest `protobuf:"bytes,4,opt,name=lookup_request,json=lookupRequest,proto3,oneof"`
}

type AcnMessage_LookupResponse struct {
	LookupResponse *LookupResponse `protobuf:"bytes,5,opt,name=lookup_response,json=lookupResponse,proto3,oneof"`
}

type AcnMessage_AeaEnvelope struct {
	AeaEnvelope *AeaEnvelope `protobuf:"bytes,6,opt,name=aea_envelope,json=aeaEnvelope,proto3,oneof"`
}

func (*AcnMessage_Status) isAcnMessage_Payload() {}

func (*AcnMessage_Register) isAcnMessage_Payload() {}

func (*AcnMessage_LookupRequest) isAcnMessage_Payload() {}

func (*AcnMessage_LookupResponse) isAcnMessage_Payload() {}

func (*AcnMessage_AeaEnvelope) isAcnMessage_Payload() {}

var File_dht_dhtnode_message_proto protoreflect.FileDescriptor

var file_dht_dhtnode_message_proto_rawDesc = []byte{
	0x0a, 0x19, 0x64, 0x68, 0x74, 0x2f, 0x64, 0x68, 0x74, 0x6e, 0x6f, 0x64, 0x65, 0x2f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x64, 0x68, 0x74,
	0x6e, 0x6f, 0x64, 0x65, 0x22, 0x84, 0x02, 0x0a, 0x0b, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x75,
	0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x26, 0x0a, 0x0f, 0x70, 0x65, 0x65,
	0x72, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x70, 0x65, 0x65, 0x72, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65,
	0x79, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12,
	0x1d, 0x0a, 0x0a, 0x6e, 0x6f, 0x74, 0x5f, 0x62, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x6f, 0x74, 0x42, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x12, 0x1b,
	0x0a, 0x09, 0x6e, 0x6f, 0x74, 0x5f, 0x61, 0x66, 0x74, 0x65, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x6e, 0x6f, 0x74, 0x41, 0x66, 0x74, 0x65, 0x72, 0x22, 0x38, 0x0a, 0x08, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x2c, 0x0a, 0x06, 0x72, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x64, 0x68, 0x74, 0x6e, 0x6f, 0x64,
	0x65, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x06, 0x72,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x22, 0x34, 0x0a, 0x0d, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x61,
	0x67, 0x65, 0x6e, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x49, 0x0a, 0x0e, 0x4c,
	0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a,
	0x0c, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x64, 0x68, 0x74, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x41, 0x67,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x0b, 0x61, 0x67, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x22, 0x51, 0x0a, 0x0b, 0x41, 0x65, 0x61, 0x45, 0x6e, 0x76,
	0x65, 0x6c, 0x6f, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6e, 0x76, 0x65, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x65, 0x6e, 0x76, 0x65, 0x6c, 0x12, 0x2c, 0x0a, 0x06, 0x72,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x64, 0x68,
	0x74, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x52, 0x06, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x22, 0xf9, 0x02, 0x0a, 0x06, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x2b, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x17, 0x2e, 0x64, 0x68, 0x74, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x2e, 0x45, 0x72, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x73, 0x67, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x04, 0x6d, 0x73, 0x67, 0x73, 0x22, 0xad, 0x02, 0x0a, 0x07, 0x45, 0x72, 0x72, 0x43, 0x6f, 0x64,
	0x65, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x00, 0x12, 0x1d,
	0x0a, 0x19, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x55, 0x4e, 0x53, 0x55, 0x50, 0x50, 0x4f, 0x52,
	0x54, 0x45, 0x44, 0x5f, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x10, 0x01, 0x12, 0x1c, 0x0a,
	0x18, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x55, 0x4e, 0x45, 0x58, 0x50, 0x45, 0x43, 0x54, 0x45,
	0x44, 0x5f, 0x50, 0x41, 0x59, 0x4c, 0x4f, 0x41, 0x44, 0x10, 0x02, 0x12, 0x11, 0x0a, 0x0d, 0x45,
	0x52, 0x52, 0x4f, 0x52, 0x5f, 0x47, 0x45, 0x4e, 0x45, 0x52, 0x49, 0x43, 0x10, 0x03, 0x12, 0x17,
	0x0a, 0x13, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x53, 0x45, 0x52, 0x49, 0x41, 0x4c, 0x49, 0x5a,
	0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x04, 0x12, 0x1d, 0x0a, 0x19, 0x45, 0x52, 0x52, 0x4f, 0x52,
	0x5f, 0x57, 0x52, 0x4f, 0x4e, 0x47, 0x5f, 0x41, 0x47, 0x45, 0x4e, 0x54, 0x5f, 0x41, 0x44, 0x44,
	0x52, 0x45, 0x53, 0x53, 0x10, 0x0a, 0x12, 0x1a, 0x0a, 0x16, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f,
	0x57, 0x52, 0x4f, 0x4e, 0x47, 0x5f, 0x50, 0x55, 0x42, 0x4c, 0x49, 0x43, 0x5f, 0x4b, 0x45, 0x59,
	0x10, 0x0b, 0x12, 0x17, 0x0a, 0x13, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x49, 0x4e, 0x56, 0x41,
	0x4c, 0x49, 0x44, 0x5f, 0x50, 0x52, 0x4f, 0x4f, 0x46, 0x10, 0x0c, 0x12, 0x1c, 0x0a, 0x18, 0x45,
	0x52, 0x52, 0x4f, 0x52, 0x5f, 0x55, 0x4e, 0x53, 0x55, 0x50, 0x50, 0x4f, 0x52, 0x54, 0x45, 0x44,
	0x5f, 0x4c, 0x45, 0x44, 0x47, 0x45, 0x52, 0x10, 0x0d, 0x12, 0x1f, 0x0a, 0x1b, 0x45, 0x52, 0x52,
	0x4f, 0x52, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x41, 0x47, 0x45, 0x4e, 0x54,
	0x5f, 0x41, 0x44, 0x44, 0x52, 0x45, 0x53, 0x53, 0x10, 0x14, 0x12, 0x19, 0x0a, 0x15, 0x45, 0x52,
	0x52, 0x4f, 0x52, 0x5f, 0x41, 0x47, 0x45, 0x4e, 0x54, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x52, 0x45,
	0x41, 0x44, 0x59, 0x10, 0x15, 0x22, 0xcd, 0x02, 0x0a, 0x0a, 0x41, 0x63, 0x6e, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x29,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f,
	0x2e, 0x64, 0x68, 0x74, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x48,
	0x00, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2f, 0x0a, 0x08, 0x72, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x64, 0x68,
	0x74, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x48, 0x00,
	0x52, 0x08, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x3f, 0x0a, 0x0e, 0x6c, 0x6f,
	0x6f, 0x6b, 0x75, 0x70, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x16, 0x2e, 0x64, 0x68, 0x74, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x4c, 0x6f, 0x6f,
	0x6b, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x0d, 0x6c, 0x6f,
	0x6f, 0x6b, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x42, 0x0a, 0x0f, 0x6c,
	0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x64, 0x68, 0x74, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x4c,
	0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x00, 0x52,
	0x0e, 0x6c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x39, 0x0a, 0x0c, 0x61, 0x65, 0x61, 0x5f, 0x65, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x64, 0x68, 0x74, 0x6e, 0x6f, 0x64, 0x65, 0x2e,
	0x41, 0x65, 0x61, 0x45, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x48, 0x00, 0x52, 0x0b, 0x61,
	0x65, 0x61, 0x45, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x42, 0x09, 0x0a, 0x07, 0x70, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_dht_dhtnode_message_proto_rawDescOnce sync.Once
	file_dht_dhtnode_message_proto_rawDescData = file_dht_dhtnode_message_proto_rawDesc
)

func file_dht_dhtnode_message_proto_rawDescGZIP() []byte {
	file_dht_dhtnode_message_proto_rawDescOnce.Do(func() {
		file_dht_dhtnode_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_dht_dhtnode_message_proto_rawDescData)
	})
	return file_dht_dhtnode_message_proto_rawDescData
}

var file_dht_dhtnode_message_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_dht_dhtnode_message_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_dht_dhtnode_message_proto_goTypes = []interface{}{
	(Status_ErrCode)(0),    // 0: dhtnode.Status.ErrCode
	(*AgentRecord)(nil),    // 1: dhtnode.AgentRecord
	(*Register)(nil),       // 2: dhtnode.Register
	(*LookupRequest)(nil),  // 3: dhtnode.LookupRequest
	(*LookupResponse)(nil), // 4: dhtnode.LookupResponse
	(*AeaEnvelope)(nil),    // 5: dhtnode.AeaEnvelope
	(*Status)(nil),         // 6: dhtnode.Status
	(*AcnMessage)(nil),     // 7: dhtnode.AcnMessage
}
var file_dht_dhtnode_message_proto_depIdxs = []int32{
	1, // 0: dhtnode.Register.record:type_name -> dhtnode.AgentRecord
	1, // 1: dhtnode.LookupResponse.agent_record:type_name -> dhtnode.AgentRecord
	1, // 2: dhtnode.AeaEnvelope.record:type_name -> dhtnode.AgentRecord
	0, // 3: dhtnode.Status.code:type_name -> dhtnode.Status.ErrCode
	6, // 4: dhtnode.AcnMessage.status:type_name -> dhtnode.Status
	2, // 5: dhtnode.AcnMessage.register:type_name -> dhtnode.Register
	3, // 6: dhtnode.AcnMessage.lookup_request:type_name -> dhtnode.LookupRequest
	4, // 7: dhtnode.AcnMessage.lookup_response:type_name -> dhtnode.LookupResponse
	5, // 8: dhtnode.AcnMessage.aea_envelope:type_name -> dhtnode.AeaEnvelope
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_dht_dhtnode_message_proto_init() }
func file_dht_dhtnode_message_proto_init() {
	if File_dht_dhtnode_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_dht_dhtnode_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AgentRecord); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_dht_dhtnode_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Register); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_dht_dhtnode_message_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LookupRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_dht_dhtnode_message_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LookupResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_dht_dhtnode_message_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AeaEnvelope); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_dht_dhtnode_message_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Status); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_dht_dhtnode_message_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AcnMessage); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_dht_dhtnode_message_proto_msgTypes[6].OneofWrappers = []interface{}{
		(*AcnMessage_Status)(nil),
		(*AcnMessage_Register)(nil),
		(*AcnMessage_LookupRequest)(nil),
		(*AcnMessage_LookupResponse)(nil),
		(*AcnMessage_AeaEnvelope)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_dht_dhtnode_message_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_dht_dhtnode_message_proto_goTypes,
		DependencyIndexes: file_dht_dhtnode_message_proto_depIdxs,
		EnumInfos:         file_dht_dhtnode_message_proto_enumTypes,
		MessageInfos:      file_dht_dhtnode_message_proto_msgTypes,
	}.Build()
	File_dht_dhtnode_message_proto = out.File
	file_dht_dhtnode_message_proto_rawDesc = nil
	file_dht_dhtnode_message_proto_goTypes = nil
	file_dht_dhtnode_message_proto_depIdxs = nil
}
