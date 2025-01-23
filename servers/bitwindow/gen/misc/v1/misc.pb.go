// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.3
// 	protoc        (unknown)
// source: misc/v1/misc.proto

package miscv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ListOPReturnResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	OpReturns     []*OPReturn            `protobuf:"bytes,1,rep,name=op_returns,json=opReturns,proto3" json:"op_returns,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListOPReturnResponse) Reset() {
	*x = ListOPReturnResponse{}
	mi := &file_misc_v1_misc_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListOPReturnResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListOPReturnResponse) ProtoMessage() {}

func (x *ListOPReturnResponse) ProtoReflect() protoreflect.Message {
	mi := &file_misc_v1_misc_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListOPReturnResponse.ProtoReflect.Descriptor instead.
func (*ListOPReturnResponse) Descriptor() ([]byte, []int) {
	return file_misc_v1_misc_proto_rawDescGZIP(), []int{0}
}

func (x *ListOPReturnResponse) GetOpReturns() []*OPReturn {
	if x != nil {
		return x.OpReturns
	}
	return nil
}

type OPReturn struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Message       string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Txid          string                 `protobuf:"bytes,3,opt,name=txid,proto3" json:"txid,omitempty"`
	Vout          int32                  `protobuf:"varint,4,opt,name=vout,proto3" json:"vout,omitempty"`
	Height        *int32                 `protobuf:"varint,5,opt,name=height,proto3,oneof" json:"height,omitempty"`
	FeeSats       int64                  `protobuf:"varint,6,opt,name=fee_sats,json=feeSats,proto3" json:"fee_sats,omitempty"`
	CreateTime    *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OPReturn) Reset() {
	*x = OPReturn{}
	mi := &file_misc_v1_misc_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OPReturn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OPReturn) ProtoMessage() {}

func (x *OPReturn) ProtoReflect() protoreflect.Message {
	mi := &file_misc_v1_misc_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OPReturn.ProtoReflect.Descriptor instead.
func (*OPReturn) Descriptor() ([]byte, []int) {
	return file_misc_v1_misc_proto_rawDescGZIP(), []int{1}
}

func (x *OPReturn) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *OPReturn) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *OPReturn) GetTxid() string {
	if x != nil {
		return x.Txid
	}
	return ""
}

func (x *OPReturn) GetVout() int32 {
	if x != nil {
		return x.Vout
	}
	return 0
}

func (x *OPReturn) GetHeight() int32 {
	if x != nil && x.Height != nil {
		return *x.Height
	}
	return 0
}

func (x *OPReturn) GetFeeSats() int64 {
	if x != nil {
		return x.FeeSats
	}
	return 0
}

func (x *OPReturn) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

type BroadcastNewsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Topic         string                 `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	Headline      string                 `protobuf:"bytes,2,opt,name=headline,proto3" json:"headline,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BroadcastNewsRequest) Reset() {
	*x = BroadcastNewsRequest{}
	mi := &file_misc_v1_misc_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BroadcastNewsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BroadcastNewsRequest) ProtoMessage() {}

func (x *BroadcastNewsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_misc_v1_misc_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BroadcastNewsRequest.ProtoReflect.Descriptor instead.
func (*BroadcastNewsRequest) Descriptor() ([]byte, []int) {
	return file_misc_v1_misc_proto_rawDescGZIP(), []int{2}
}

func (x *BroadcastNewsRequest) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *BroadcastNewsRequest) GetHeadline() string {
	if x != nil {
		return x.Headline
	}
	return ""
}

type BroadcastNewsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Txid          string                 `protobuf:"bytes,1,opt,name=txid,proto3" json:"txid,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BroadcastNewsResponse) Reset() {
	*x = BroadcastNewsResponse{}
	mi := &file_misc_v1_misc_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BroadcastNewsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BroadcastNewsResponse) ProtoMessage() {}

func (x *BroadcastNewsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_misc_v1_misc_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BroadcastNewsResponse.ProtoReflect.Descriptor instead.
func (*BroadcastNewsResponse) Descriptor() ([]byte, []int) {
	return file_misc_v1_misc_proto_rawDescGZIP(), []int{3}
}

func (x *BroadcastNewsResponse) GetTxid() string {
	if x != nil {
		return x.Txid
	}
	return ""
}

type CreateTopicRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Topic         string                 `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateTopicRequest) Reset() {
	*x = CreateTopicRequest{}
	mi := &file_misc_v1_misc_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateTopicRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTopicRequest) ProtoMessage() {}

func (x *CreateTopicRequest) ProtoReflect() protoreflect.Message {
	mi := &file_misc_v1_misc_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTopicRequest.ProtoReflect.Descriptor instead.
func (*CreateTopicRequest) Descriptor() ([]byte, []int) {
	return file_misc_v1_misc_proto_rawDescGZIP(), []int{4}
}

func (x *CreateTopicRequest) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *CreateTopicRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CreateTopicResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Txid          string                 `protobuf:"bytes,1,opt,name=txid,proto3" json:"txid,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateTopicResponse) Reset() {
	*x = CreateTopicResponse{}
	mi := &file_misc_v1_misc_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateTopicResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTopicResponse) ProtoMessage() {}

func (x *CreateTopicResponse) ProtoReflect() protoreflect.Message {
	mi := &file_misc_v1_misc_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTopicResponse.ProtoReflect.Descriptor instead.
func (*CreateTopicResponse) Descriptor() ([]byte, []int) {
	return file_misc_v1_misc_proto_rawDescGZIP(), []int{5}
}

func (x *CreateTopicResponse) GetTxid() string {
	if x != nil {
		return x.Txid
	}
	return ""
}

type Topic struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Topic         string                 `protobuf:"bytes,2,opt,name=topic,proto3" json:"topic,omitempty"`
	Name          string                 `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	CreateTime    *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Topic) Reset() {
	*x = Topic{}
	mi := &file_misc_v1_misc_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Topic) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Topic) ProtoMessage() {}

func (x *Topic) ProtoReflect() protoreflect.Message {
	mi := &file_misc_v1_misc_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Topic.ProtoReflect.Descriptor instead.
func (*Topic) Descriptor() ([]byte, []int) {
	return file_misc_v1_misc_proto_rawDescGZIP(), []int{6}
}

func (x *Topic) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Topic) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *Topic) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Topic) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

type ListTopicsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Topics        []*Topic               `protobuf:"bytes,1,rep,name=topics,proto3" json:"topics,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListTopicsResponse) Reset() {
	*x = ListTopicsResponse{}
	mi := &file_misc_v1_misc_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListTopicsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTopicsResponse) ProtoMessage() {}

func (x *ListTopicsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_misc_v1_misc_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTopicsResponse.ProtoReflect.Descriptor instead.
func (*ListTopicsResponse) Descriptor() ([]byte, []int) {
	return file_misc_v1_misc_proto_rawDescGZIP(), []int{7}
}

func (x *ListTopicsResponse) GetTopics() []*Topic {
	if x != nil {
		return x.Topics
	}
	return nil
}

type ListCoinNewsRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// if set, only return news for this topic
	Topic         *string `protobuf:"bytes,1,opt,name=topic,proto3,oneof" json:"topic,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListCoinNewsRequest) Reset() {
	*x = ListCoinNewsRequest{}
	mi := &file_misc_v1_misc_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListCoinNewsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCoinNewsRequest) ProtoMessage() {}

func (x *ListCoinNewsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_misc_v1_misc_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCoinNewsRequest.ProtoReflect.Descriptor instead.
func (*ListCoinNewsRequest) Descriptor() ([]byte, []int) {
	return file_misc_v1_misc_proto_rawDescGZIP(), []int{8}
}

func (x *ListCoinNewsRequest) GetTopic() string {
	if x != nil && x.Topic != nil {
		return *x.Topic
	}
	return ""
}

type CoinNews struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Topic         string                 `protobuf:"bytes,2,opt,name=topic,proto3" json:"topic,omitempty"`
	Headline      string                 `protobuf:"bytes,3,opt,name=headline,proto3" json:"headline,omitempty"`
	Content       string                 `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
	FeeSats       int64                  `protobuf:"varint,5,opt,name=fee_sats,json=feeSats,proto3" json:"fee_sats,omitempty"`
	CreateTime    *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CoinNews) Reset() {
	*x = CoinNews{}
	mi := &file_misc_v1_misc_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CoinNews) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CoinNews) ProtoMessage() {}

func (x *CoinNews) ProtoReflect() protoreflect.Message {
	mi := &file_misc_v1_misc_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CoinNews.ProtoReflect.Descriptor instead.
func (*CoinNews) Descriptor() ([]byte, []int) {
	return file_misc_v1_misc_proto_rawDescGZIP(), []int{9}
}

func (x *CoinNews) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CoinNews) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *CoinNews) GetHeadline() string {
	if x != nil {
		return x.Headline
	}
	return ""
}

func (x *CoinNews) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *CoinNews) GetFeeSats() int64 {
	if x != nil {
		return x.FeeSats
	}
	return 0
}

func (x *CoinNews) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

type ListCoinNewsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	CoinNews      []*CoinNews            `protobuf:"bytes,1,rep,name=coin_news,json=coinNews,proto3" json:"coin_news,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListCoinNewsResponse) Reset() {
	*x = ListCoinNewsResponse{}
	mi := &file_misc_v1_misc_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListCoinNewsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCoinNewsResponse) ProtoMessage() {}

func (x *ListCoinNewsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_misc_v1_misc_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCoinNewsResponse.ProtoReflect.Descriptor instead.
func (*ListCoinNewsResponse) Descriptor() ([]byte, []int) {
	return file_misc_v1_misc_proto_rawDescGZIP(), []int{10}
}

func (x *ListCoinNewsResponse) GetCoinNews() []*CoinNews {
	if x != nil {
		return x.CoinNews
	}
	return nil
}

var File_misc_v1_misc_proto protoreflect.FileDescriptor

var file_misc_v1_misc_proto_rawDesc = []byte{
	0x0a, 0x12, 0x6d, 0x69, 0x73, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x69, 0x73, 0x63, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x6d, 0x69, 0x73, 0x63, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65,
	0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x48, 0x0a, 0x14, 0x4c,
	0x69, 0x73, 0x74, 0x4f, 0x50, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x0a, 0x6f, 0x70, 0x5f, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x6d, 0x69, 0x73, 0x63, 0x2e, 0x76,
	0x31, 0x2e, 0x4f, 0x50, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x52, 0x09, 0x6f, 0x70, 0x52, 0x65,
	0x74, 0x75, 0x72, 0x6e, 0x73, 0x22, 0xdc, 0x01, 0x0a, 0x08, 0x4f, 0x50, 0x52, 0x65, 0x74, 0x75,
	0x72, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x78, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x78, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x76, 0x6f, 0x75, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x76, 0x6f, 0x75, 0x74, 0x12, 0x1b, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x88, 0x01,
	0x01, 0x12, 0x19, 0x0a, 0x08, 0x66, 0x65, 0x65, 0x5f, 0x73, 0x61, 0x74, 0x73, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x07, 0x66, 0x65, 0x65, 0x53, 0x61, 0x74, 0x73, 0x12, 0x3b, 0x0a, 0x0b,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x68, 0x65,
	0x69, 0x67, 0x68, 0x74, 0x22, 0x48, 0x0a, 0x14, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73,
	0x74, 0x4e, 0x65, 0x77, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x6f, 0x70, 0x69, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x70,
	0x69, 0x63, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x22, 0x2b,
	0x0a, 0x15, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x4e, 0x65, 0x77, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x78, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x78, 0x69, 0x64, 0x22, 0x3e, 0x0a, 0x12, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x29, 0x0a, 0x13, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x78, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x78, 0x69, 0x64, 0x22, 0x7e, 0x0a, 0x05, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x3c, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x6f,
	0x70, 0x69, 0x63, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x06,
	0x74, 0x6f, 0x70, 0x69, 0x63, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6d,
	0x69, 0x73, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x52, 0x06, 0x74, 0x6f,
	0x70, 0x69, 0x63, 0x73, 0x22, 0x3a, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x69, 0x6e,
	0x4e, 0x65, 0x77, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x05, 0x74,
	0x6f, 0x70, 0x69, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x74, 0x6f,
	0x70, 0x69, 0x63, 0x88, 0x01, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x74, 0x6f, 0x70, 0x69, 0x63,
	0x22, 0xbe, 0x01, 0x0a, 0x08, 0x43, 0x6f, 0x69, 0x6e, 0x4e, 0x65, 0x77, 0x73, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f,
	0x70, 0x69, 0x63, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x66, 0x65, 0x65,
	0x5f, 0x73, 0x61, 0x74, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x66, 0x65, 0x65,
	0x53, 0x61, 0x74, 0x73, 0x12, 0x3b, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d,
	0x65, 0x22, 0x46, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x69, 0x6e, 0x4e, 0x65, 0x77,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x09, 0x63, 0x6f, 0x69,
	0x6e, 0x5f, 0x6e, 0x65, 0x77, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x6d,
	0x69, 0x73, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x69, 0x6e, 0x4e, 0x65, 0x77, 0x73, 0x52,
	0x08, 0x63, 0x6f, 0x69, 0x6e, 0x4e, 0x65, 0x77, 0x73, 0x32, 0xfe, 0x02, 0x0a, 0x0b, 0x4d, 0x69,
	0x73, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x45, 0x0a, 0x0c, 0x4c, 0x69, 0x73,
	0x74, 0x4f, 0x50, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x1a, 0x1d, 0x2e, 0x6d, 0x69, 0x73, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x4f, 0x50, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x4e, 0x0a, 0x0d, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x4e, 0x65, 0x77,
	0x73, 0x12, 0x1d, 0x2e, 0x6d, 0x69, 0x73, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x72, 0x6f, 0x61,
	0x64, 0x63, 0x61, 0x73, 0x74, 0x4e, 0x65, 0x77, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1e, 0x2e, 0x6d, 0x69, 0x73, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x72, 0x6f, 0x61, 0x64,
	0x63, 0x61, 0x73, 0x74, 0x4e, 0x65, 0x77, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x48, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x12,
	0x1b, 0x2e, 0x6d, 0x69, 0x73, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x54, 0x6f, 0x70, 0x69, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x6d,
	0x69, 0x73, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x70,
	0x69, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x0a, 0x4c, 0x69,
	0x73, 0x74, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x1a, 0x1b, 0x2e, 0x6d, 0x69, 0x73, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54,
	0x6f, 0x70, 0x69, 0x63, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4b, 0x0a,
	0x0c, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x69, 0x6e, 0x4e, 0x65, 0x77, 0x73, 0x12, 0x1c, 0x2e,
	0x6d, 0x69, 0x73, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x69, 0x6e,
	0x4e, 0x65, 0x77, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x6d, 0x69,
	0x73, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x69, 0x6e, 0x4e, 0x65,
	0x77, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x9d, 0x01, 0x0a, 0x0b, 0x63,
	0x6f, 0x6d, 0x2e, 0x6d, 0x69, 0x73, 0x63, 0x2e, 0x76, 0x31, 0x42, 0x09, 0x4d, 0x69, 0x73, 0x63,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x46, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x54, 0x77, 0x6f, 0x2d, 0x4c, 0x61, 0x62,
	0x73, 0x2f, 0x73, 0x69, 0x64, 0x65, 0x73, 0x61, 0x69, 0x6c, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x73, 0x2f, 0x62, 0x69, 0x74, 0x77, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x2f, 0x67, 0x65, 0x6e,
	0x2f, 0x6d, 0x69, 0x73, 0x63, 0x2f, 0x76, 0x31, 0x3b, 0x6d, 0x69, 0x73, 0x63, 0x76, 0x31, 0xa2,
	0x02, 0x03, 0x4d, 0x58, 0x58, 0xaa, 0x02, 0x07, 0x4d, 0x69, 0x73, 0x63, 0x2e, 0x56, 0x31, 0xca,
	0x02, 0x07, 0x4d, 0x69, 0x73, 0x63, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x13, 0x4d, 0x69, 0x73, 0x63,
	0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x08, 0x4d, 0x69, 0x73, 0x63, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_misc_v1_misc_proto_rawDescOnce sync.Once
	file_misc_v1_misc_proto_rawDescData = file_misc_v1_misc_proto_rawDesc
)

func file_misc_v1_misc_proto_rawDescGZIP() []byte {
	file_misc_v1_misc_proto_rawDescOnce.Do(func() {
		file_misc_v1_misc_proto_rawDescData = protoimpl.X.CompressGZIP(file_misc_v1_misc_proto_rawDescData)
	})
	return file_misc_v1_misc_proto_rawDescData
}

var file_misc_v1_misc_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_misc_v1_misc_proto_goTypes = []any{
	(*ListOPReturnResponse)(nil),  // 0: misc.v1.ListOPReturnResponse
	(*OPReturn)(nil),              // 1: misc.v1.OPReturn
	(*BroadcastNewsRequest)(nil),  // 2: misc.v1.BroadcastNewsRequest
	(*BroadcastNewsResponse)(nil), // 3: misc.v1.BroadcastNewsResponse
	(*CreateTopicRequest)(nil),    // 4: misc.v1.CreateTopicRequest
	(*CreateTopicResponse)(nil),   // 5: misc.v1.CreateTopicResponse
	(*Topic)(nil),                 // 6: misc.v1.Topic
	(*ListTopicsResponse)(nil),    // 7: misc.v1.ListTopicsResponse
	(*ListCoinNewsRequest)(nil),   // 8: misc.v1.ListCoinNewsRequest
	(*CoinNews)(nil),              // 9: misc.v1.CoinNews
	(*ListCoinNewsResponse)(nil),  // 10: misc.v1.ListCoinNewsResponse
	(*timestamppb.Timestamp)(nil), // 11: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),         // 12: google.protobuf.Empty
}
var file_misc_v1_misc_proto_depIdxs = []int32{
	1,  // 0: misc.v1.ListOPReturnResponse.op_returns:type_name -> misc.v1.OPReturn
	11, // 1: misc.v1.OPReturn.create_time:type_name -> google.protobuf.Timestamp
	11, // 2: misc.v1.Topic.create_time:type_name -> google.protobuf.Timestamp
	6,  // 3: misc.v1.ListTopicsResponse.topics:type_name -> misc.v1.Topic
	11, // 4: misc.v1.CoinNews.create_time:type_name -> google.protobuf.Timestamp
	9,  // 5: misc.v1.ListCoinNewsResponse.coin_news:type_name -> misc.v1.CoinNews
	12, // 6: misc.v1.MiscService.ListOPReturn:input_type -> google.protobuf.Empty
	2,  // 7: misc.v1.MiscService.BroadcastNews:input_type -> misc.v1.BroadcastNewsRequest
	4,  // 8: misc.v1.MiscService.CreateTopic:input_type -> misc.v1.CreateTopicRequest
	12, // 9: misc.v1.MiscService.ListTopics:input_type -> google.protobuf.Empty
	8,  // 10: misc.v1.MiscService.ListCoinNews:input_type -> misc.v1.ListCoinNewsRequest
	0,  // 11: misc.v1.MiscService.ListOPReturn:output_type -> misc.v1.ListOPReturnResponse
	3,  // 12: misc.v1.MiscService.BroadcastNews:output_type -> misc.v1.BroadcastNewsResponse
	5,  // 13: misc.v1.MiscService.CreateTopic:output_type -> misc.v1.CreateTopicResponse
	7,  // 14: misc.v1.MiscService.ListTopics:output_type -> misc.v1.ListTopicsResponse
	10, // 15: misc.v1.MiscService.ListCoinNews:output_type -> misc.v1.ListCoinNewsResponse
	11, // [11:16] is the sub-list for method output_type
	6,  // [6:11] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_misc_v1_misc_proto_init() }
func file_misc_v1_misc_proto_init() {
	if File_misc_v1_misc_proto != nil {
		return
	}
	file_misc_v1_misc_proto_msgTypes[1].OneofWrappers = []any{}
	file_misc_v1_misc_proto_msgTypes[8].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_misc_v1_misc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_misc_v1_misc_proto_goTypes,
		DependencyIndexes: file_misc_v1_misc_proto_depIdxs,
		MessageInfos:      file_misc_v1_misc_proto_msgTypes,
	}.Build()
	File_misc_v1_misc_proto = out.File
	file_misc_v1_misc_proto_rawDesc = nil
	file_misc_v1_misc_proto_goTypes = nil
	file_misc_v1_misc_proto_depIdxs = nil
}
