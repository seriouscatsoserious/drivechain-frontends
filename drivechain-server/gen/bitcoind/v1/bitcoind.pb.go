// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: bitcoind/v1/bitcoind.proto

package bitcoindv1

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

type ListRecentBlocksRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count int64 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *ListRecentBlocksRequest) Reset() {
	*x = ListRecentBlocksRequest{}
	mi := &file_bitcoind_v1_bitcoind_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListRecentBlocksRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRecentBlocksRequest) ProtoMessage() {}

func (x *ListRecentBlocksRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bitcoind_v1_bitcoind_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRecentBlocksRequest.ProtoReflect.Descriptor instead.
func (*ListRecentBlocksRequest) Descriptor() ([]byte, []int) {
	return file_bitcoind_v1_bitcoind_proto_rawDescGZIP(), []int{0}
}

func (x *ListRecentBlocksRequest) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

type ListRecentBlocksResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RecentBlocks []*ListRecentBlocksResponse_RecentBlock `protobuf:"bytes,4,rep,name=recent_blocks,json=recentBlocks,proto3" json:"recent_blocks,omitempty"`
}

func (x *ListRecentBlocksResponse) Reset() {
	*x = ListRecentBlocksResponse{}
	mi := &file_bitcoind_v1_bitcoind_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListRecentBlocksResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRecentBlocksResponse) ProtoMessage() {}

func (x *ListRecentBlocksResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bitcoind_v1_bitcoind_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRecentBlocksResponse.ProtoReflect.Descriptor instead.
func (*ListRecentBlocksResponse) Descriptor() ([]byte, []int) {
	return file_bitcoind_v1_bitcoind_proto_rawDescGZIP(), []int{1}
}

func (x *ListRecentBlocksResponse) GetRecentBlocks() []*ListRecentBlocksResponse_RecentBlock {
	if x != nil {
		return x.RecentBlocks
	}
	return nil
}

type ListUnconfirmedTransactionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count int64 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *ListUnconfirmedTransactionsRequest) Reset() {
	*x = ListUnconfirmedTransactionsRequest{}
	mi := &file_bitcoind_v1_bitcoind_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListUnconfirmedTransactionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListUnconfirmedTransactionsRequest) ProtoMessage() {}

func (x *ListUnconfirmedTransactionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bitcoind_v1_bitcoind_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListUnconfirmedTransactionsRequest.ProtoReflect.Descriptor instead.
func (*ListUnconfirmedTransactionsRequest) Descriptor() ([]byte, []int) {
	return file_bitcoind_v1_bitcoind_proto_rawDescGZIP(), []int{2}
}

func (x *ListUnconfirmedTransactionsRequest) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

type ListUnconfirmedTransactionsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UnconfirmedTransactions []*UnconfirmedTransaction `protobuf:"bytes,1,rep,name=unconfirmed_transactions,json=unconfirmedTransactions,proto3" json:"unconfirmed_transactions,omitempty"`
}

func (x *ListUnconfirmedTransactionsResponse) Reset() {
	*x = ListUnconfirmedTransactionsResponse{}
	mi := &file_bitcoind_v1_bitcoind_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListUnconfirmedTransactionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListUnconfirmedTransactionsResponse) ProtoMessage() {}

func (x *ListUnconfirmedTransactionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bitcoind_v1_bitcoind_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListUnconfirmedTransactionsResponse.ProtoReflect.Descriptor instead.
func (*ListUnconfirmedTransactionsResponse) Descriptor() ([]byte, []int) {
	return file_bitcoind_v1_bitcoind_proto_rawDescGZIP(), []int{3}
}

func (x *ListUnconfirmedTransactionsResponse) GetUnconfirmedTransactions() []*UnconfirmedTransaction {
	if x != nil {
		return x.UnconfirmedTransactions
	}
	return nil
}

type UnconfirmedTransaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VirtualSize uint32                 `protobuf:"varint,1,opt,name=virtual_size,json=virtualSize,proto3" json:"virtual_size,omitempty"`
	Weight      uint32                 `protobuf:"varint,2,opt,name=weight,proto3" json:"weight,omitempty"`
	Time        *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=time,proto3" json:"time,omitempty"`
	Txid        string                 `protobuf:"bytes,4,opt,name=txid,proto3" json:"txid,omitempty"`
	FeeSatoshi  uint64                 `protobuf:"varint,5,opt,name=fee_satoshi,json=feeSatoshi,proto3" json:"fee_satoshi,omitempty"`
}

func (x *UnconfirmedTransaction) Reset() {
	*x = UnconfirmedTransaction{}
	mi := &file_bitcoind_v1_bitcoind_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UnconfirmedTransaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnconfirmedTransaction) ProtoMessage() {}

func (x *UnconfirmedTransaction) ProtoReflect() protoreflect.Message {
	mi := &file_bitcoind_v1_bitcoind_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnconfirmedTransaction.ProtoReflect.Descriptor instead.
func (*UnconfirmedTransaction) Descriptor() ([]byte, []int) {
	return file_bitcoind_v1_bitcoind_proto_rawDescGZIP(), []int{4}
}

func (x *UnconfirmedTransaction) GetVirtualSize() uint32 {
	if x != nil {
		return x.VirtualSize
	}
	return 0
}

func (x *UnconfirmedTransaction) GetWeight() uint32 {
	if x != nil {
		return x.Weight
	}
	return 0
}

func (x *UnconfirmedTransaction) GetTime() *timestamppb.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

func (x *UnconfirmedTransaction) GetTxid() string {
	if x != nil {
		return x.Txid
	}
	return ""
}

func (x *UnconfirmedTransaction) GetFeeSatoshi() uint64 {
	if x != nil {
		return x.FeeSatoshi
	}
	return 0
}

type GetBlockchainInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Chain                string `protobuf:"bytes,1,opt,name=chain,proto3" json:"chain,omitempty"`
	Blocks               uint32 `protobuf:"varint,2,opt,name=blocks,proto3" json:"blocks,omitempty"`
	Headers              uint32 `protobuf:"varint,3,opt,name=headers,proto3" json:"headers,omitempty"`
	BestBlockHash        string `protobuf:"bytes,4,opt,name=best_block_hash,json=bestBlockHash,proto3" json:"best_block_hash,omitempty"`
	InitialBlockDownload bool   `protobuf:"varint,8,opt,name=initial_block_download,json=initialBlockDownload,proto3" json:"initial_block_download,omitempty"`
}

func (x *GetBlockchainInfoResponse) Reset() {
	*x = GetBlockchainInfoResponse{}
	mi := &file_bitcoind_v1_bitcoind_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetBlockchainInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBlockchainInfoResponse) ProtoMessage() {}

func (x *GetBlockchainInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bitcoind_v1_bitcoind_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBlockchainInfoResponse.ProtoReflect.Descriptor instead.
func (*GetBlockchainInfoResponse) Descriptor() ([]byte, []int) {
	return file_bitcoind_v1_bitcoind_proto_rawDescGZIP(), []int{5}
}

func (x *GetBlockchainInfoResponse) GetChain() string {
	if x != nil {
		return x.Chain
	}
	return ""
}

func (x *GetBlockchainInfoResponse) GetBlocks() uint32 {
	if x != nil {
		return x.Blocks
	}
	return 0
}

func (x *GetBlockchainInfoResponse) GetHeaders() uint32 {
	if x != nil {
		return x.Headers
	}
	return 0
}

func (x *GetBlockchainInfoResponse) GetBestBlockHash() string {
	if x != nil {
		return x.BestBlockHash
	}
	return ""
}

func (x *GetBlockchainInfoResponse) GetInitialBlockDownload() bool {
	if x != nil {
		return x.InitialBlockDownload
	}
	return false
}

type Peer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Addr         string `protobuf:"bytes,2,opt,name=addr,proto3" json:"addr,omitempty"`
	SyncedBlocks int32  `protobuf:"varint,3,opt,name=synced_blocks,json=syncedBlocks,proto3" json:"synced_blocks,omitempty"`
}

func (x *Peer) Reset() {
	*x = Peer{}
	mi := &file_bitcoind_v1_bitcoind_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Peer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Peer) ProtoMessage() {}

func (x *Peer) ProtoReflect() protoreflect.Message {
	mi := &file_bitcoind_v1_bitcoind_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Peer.ProtoReflect.Descriptor instead.
func (*Peer) Descriptor() ([]byte, []int) {
	return file_bitcoind_v1_bitcoind_proto_rawDescGZIP(), []int{6}
}

func (x *Peer) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Peer) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *Peer) GetSyncedBlocks() int32 {
	if x != nil {
		return x.SyncedBlocks
	}
	return 0
}

type ListPeersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Peers []*Peer `protobuf:"bytes,1,rep,name=peers,proto3" json:"peers,omitempty"`
}

func (x *ListPeersResponse) Reset() {
	*x = ListPeersResponse{}
	mi := &file_bitcoind_v1_bitcoind_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListPeersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPeersResponse) ProtoMessage() {}

func (x *ListPeersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bitcoind_v1_bitcoind_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPeersResponse.ProtoReflect.Descriptor instead.
func (*ListPeersResponse) Descriptor() ([]byte, []int) {
	return file_bitcoind_v1_bitcoind_proto_rawDescGZIP(), []int{7}
}

func (x *ListPeersResponse) GetPeers() []*Peer {
	if x != nil {
		return x.Peers
	}
	return nil
}

type EstimateSmartFeeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConfTarget int64 `protobuf:"varint,1,opt,name=conf_target,json=confTarget,proto3" json:"conf_target,omitempty"`
}

func (x *EstimateSmartFeeRequest) Reset() {
	*x = EstimateSmartFeeRequest{}
	mi := &file_bitcoind_v1_bitcoind_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EstimateSmartFeeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EstimateSmartFeeRequest) ProtoMessage() {}

func (x *EstimateSmartFeeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bitcoind_v1_bitcoind_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EstimateSmartFeeRequest.ProtoReflect.Descriptor instead.
func (*EstimateSmartFeeRequest) Descriptor() ([]byte, []int) {
	return file_bitcoind_v1_bitcoind_proto_rawDescGZIP(), []int{8}
}

func (x *EstimateSmartFeeRequest) GetConfTarget() int64 {
	if x != nil {
		return x.ConfTarget
	}
	return 0
}

type EstimateSmartFeeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FeeRate float64  `protobuf:"fixed64,1,opt,name=fee_rate,json=feeRate,proto3" json:"fee_rate,omitempty"`
	Errors  []string `protobuf:"bytes,2,rep,name=errors,proto3" json:"errors,omitempty"`
}

func (x *EstimateSmartFeeResponse) Reset() {
	*x = EstimateSmartFeeResponse{}
	mi := &file_bitcoind_v1_bitcoind_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EstimateSmartFeeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EstimateSmartFeeResponse) ProtoMessage() {}

func (x *EstimateSmartFeeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bitcoind_v1_bitcoind_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EstimateSmartFeeResponse.ProtoReflect.Descriptor instead.
func (*EstimateSmartFeeResponse) Descriptor() ([]byte, []int) {
	return file_bitcoind_v1_bitcoind_proto_rawDescGZIP(), []int{9}
}

func (x *EstimateSmartFeeResponse) GetFeeRate() float64 {
	if x != nil {
		return x.FeeRate
	}
	return 0
}

func (x *EstimateSmartFeeResponse) GetErrors() []string {
	if x != nil {
		return x.Errors
	}
	return nil
}

type ListRecentBlocksResponse_RecentBlock struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockTime   *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=block_time,json=blockTime,proto3" json:"block_time,omitempty"`
	BlockHeight uint32                 `protobuf:"varint,2,opt,name=block_height,json=blockHeight,proto3" json:"block_height,omitempty"`
	Hash        string                 `protobuf:"bytes,3,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (x *ListRecentBlocksResponse_RecentBlock) Reset() {
	*x = ListRecentBlocksResponse_RecentBlock{}
	mi := &file_bitcoind_v1_bitcoind_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListRecentBlocksResponse_RecentBlock) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRecentBlocksResponse_RecentBlock) ProtoMessage() {}

func (x *ListRecentBlocksResponse_RecentBlock) ProtoReflect() protoreflect.Message {
	mi := &file_bitcoind_v1_bitcoind_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRecentBlocksResponse_RecentBlock.ProtoReflect.Descriptor instead.
func (*ListRecentBlocksResponse_RecentBlock) Descriptor() ([]byte, []int) {
	return file_bitcoind_v1_bitcoind_proto_rawDescGZIP(), []int{1, 0}
}

func (x *ListRecentBlocksResponse_RecentBlock) GetBlockTime() *timestamppb.Timestamp {
	if x != nil {
		return x.BlockTime
	}
	return nil
}

func (x *ListRecentBlocksResponse_RecentBlock) GetBlockHeight() uint32 {
	if x != nil {
		return x.BlockHeight
	}
	return 0
}

func (x *ListRecentBlocksResponse_RecentBlock) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

var File_bitcoind_v1_bitcoind_proto protoreflect.FileDescriptor

var file_bitcoind_v1_bitcoind_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x62, 0x69, 0x74, 0x63, 0x6f, 0x69, 0x6e, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x69,
	0x74, 0x63, 0x6f, 0x69, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x62, 0x69,
	0x74, 0x63, 0x6f, 0x69, 0x6e, 0x64, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2f, 0x0a, 0x17, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x63, 0x65, 0x6e, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0xf3, 0x01, 0x0a, 0x18, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x63, 0x65, 0x6e, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x56, 0x0a, 0x0d, 0x72, 0x65, 0x63, 0x65, 0x6e, 0x74, 0x5f,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x62,
	0x69, 0x74, 0x63, 0x6f, 0x69, 0x6e, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x63, 0x65, 0x6e, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x63, 0x65, 0x6e, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52,
	0x0c, 0x72, 0x65, 0x63, 0x65, 0x6e, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x1a, 0x7f, 0x0a,
	0x0b, 0x52, 0x65, 0x63, 0x65, 0x6e, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x39, 0x0a, 0x0a,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x62, 0x6c, 0x6f, 0x63, 0x6b,
	0x5f, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61,
	0x73, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x22, 0x3a,
	0x0a, 0x22, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x6e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x65,
	0x64, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x85, 0x01, 0x0a, 0x23, 0x4c,
	0x69, 0x73, 0x74, 0x55, 0x6e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x65, 0x64, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x5e, 0x0a, 0x18, 0x75, 0x6e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x65,
	0x64, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x62, 0x69, 0x74, 0x63, 0x6f, 0x69, 0x6e, 0x64, 0x2e,
	0x76, 0x31, 0x2e, 0x55, 0x6e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x65, 0x64, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x17, 0x75, 0x6e, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x72, 0x6d, 0x65, 0x64, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x22, 0xb8, 0x01, 0x0a, 0x16, 0x55, 0x6e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d,
	0x65, 0x64, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a,
	0x0c, 0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0b, 0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x53, 0x69, 0x7a, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x2e, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x78, 0x69, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x78, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b,
	0x66, 0x65, 0x65, 0x5f, 0x73, 0x61, 0x74, 0x6f, 0x73, 0x68, 0x69, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x0a, 0x66, 0x65, 0x65, 0x53, 0x61, 0x74, 0x6f, 0x73, 0x68, 0x69, 0x22, 0xc1, 0x01,
	0x0a, 0x19, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63,
	0x68, 0x61, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x68, 0x61, 0x69,
	0x6e, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x06, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x68, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x68, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x62, 0x65, 0x73, 0x74, 0x5f, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x62, 0x65,
	0x73, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x12, 0x34, 0x0a, 0x16, 0x69,
	0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x64, 0x6f, 0x77,
	0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x14, 0x69, 0x6e, 0x69,
	0x74, 0x69, 0x61, 0x6c, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61,
	0x64, 0x22, 0x4f, 0x0a, 0x04, 0x50, 0x65, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x64, 0x64,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x64, 0x64, 0x72, 0x12, 0x23, 0x0a,
	0x0d, 0x73, 0x79, 0x6e, 0x63, 0x65, 0x64, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x73, 0x79, 0x6e, 0x63, 0x65, 0x64, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x73, 0x22, 0x3c, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x65, 0x65, 0x72, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x05, 0x70, 0x65, 0x65, 0x72, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x62, 0x69, 0x74, 0x63, 0x6f, 0x69, 0x6e,
	0x64, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x52, 0x05, 0x70, 0x65, 0x65, 0x72, 0x73,
	0x22, 0x3a, 0x0a, 0x17, 0x45, 0x73, 0x74, 0x69, 0x6d, 0x61, 0x74, 0x65, 0x53, 0x6d, 0x61, 0x72,
	0x74, 0x46, 0x65, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x63,
	0x6f, 0x6e, 0x66, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x66, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x22, 0x4d, 0x0a, 0x18,
	0x45, 0x73, 0x74, 0x69, 0x6d, 0x61, 0x74, 0x65, 0x53, 0x6d, 0x61, 0x72, 0x74, 0x46, 0x65, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x66, 0x65, 0x65, 0x5f,
	0x72, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x07, 0x66, 0x65, 0x65, 0x52,
	0x61, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x32, 0xf0, 0x03, 0x0a, 0x0f,
	0x42, 0x69, 0x74, 0x63, 0x6f, 0x69, 0x6e, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x80, 0x01, 0x0a, 0x1b, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x6e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72,
	0x6d, 0x65, 0x64, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12,
	0x2f, 0x2e, 0x62, 0x69, 0x74, 0x63, 0x6f, 0x69, 0x6e, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x55, 0x6e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x65, 0x64, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x30, 0x2e, 0x62, 0x69, 0x74, 0x63, 0x6f, 0x69, 0x6e, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x55, 0x6e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x65, 0x64, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x5f, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x63, 0x65, 0x6e, 0x74,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x12, 0x24, 0x2e, 0x62, 0x69, 0x74, 0x63, 0x6f, 0x69, 0x6e,
	0x64, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x63, 0x65, 0x6e, 0x74, 0x42,
	0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x62,
	0x69, 0x74, 0x63, 0x6f, 0x69, 0x6e, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x63, 0x65, 0x6e, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x53, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x63,
	0x68, 0x61, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x1a, 0x26, 0x2e, 0x62, 0x69, 0x74, 0x63, 0x6f, 0x69, 0x6e, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74,
	0x50, 0x65, 0x65, 0x72, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1e, 0x2e,
	0x62, 0x69, 0x74, 0x63, 0x6f, 0x69, 0x6e, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x50, 0x65, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5f, 0x0a,
	0x10, 0x45, 0x73, 0x74, 0x69, 0x6d, 0x61, 0x74, 0x65, 0x53, 0x6d, 0x61, 0x72, 0x74, 0x46, 0x65,
	0x65, 0x12, 0x24, 0x2e, 0x62, 0x69, 0x74, 0x63, 0x6f, 0x69, 0x6e, 0x64, 0x2e, 0x76, 0x31, 0x2e,
	0x45, 0x73, 0x74, 0x69, 0x6d, 0x61, 0x74, 0x65, 0x53, 0x6d, 0x61, 0x72, 0x74, 0x46, 0x65, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x62, 0x69, 0x74, 0x63, 0x6f, 0x69,
	0x6e, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x73, 0x74, 0x69, 0x6d, 0x61, 0x74, 0x65, 0x53, 0x6d,
	0x61, 0x72, 0x74, 0x46, 0x65, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0xbd,
	0x01, 0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x2e, 0x62, 0x69, 0x74, 0x63, 0x6f, 0x69, 0x6e, 0x64, 0x2e,
	0x76, 0x31, 0x42, 0x0d, 0x42, 0x69, 0x74, 0x63, 0x6f, 0x69, 0x6e, 0x64, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x4e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x4c, 0x61, 0x79, 0x65, 0x72, 0x54, 0x77, 0x6f, 0x2d, 0x4c, 0x61, 0x62, 0x73, 0x2f, 0x73, 0x69,
	0x64, 0x65, 0x73, 0x61, 0x69, 0x6c, 0x2f, 0x64, 0x72, 0x69, 0x76, 0x65, 0x63, 0x68, 0x61, 0x69,
	0x6e, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x69, 0x74,
	0x63, 0x6f, 0x69, 0x6e, 0x64, 0x2f, 0x76, 0x31, 0x3b, 0x62, 0x69, 0x74, 0x63, 0x6f, 0x69, 0x6e,
	0x64, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x42, 0x58, 0x58, 0xaa, 0x02, 0x0b, 0x42, 0x69, 0x74, 0x63,
	0x6f, 0x69, 0x6e, 0x64, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0b, 0x42, 0x69, 0x74, 0x63, 0x6f, 0x69,
	0x6e, 0x64, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x17, 0x42, 0x69, 0x74, 0x63, 0x6f, 0x69, 0x6e, 0x64,
	0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x0c, 0x42, 0x69, 0x74, 0x63, 0x6f, 0x69, 0x6e, 0x64, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bitcoind_v1_bitcoind_proto_rawDescOnce sync.Once
	file_bitcoind_v1_bitcoind_proto_rawDescData = file_bitcoind_v1_bitcoind_proto_rawDesc
)

func file_bitcoind_v1_bitcoind_proto_rawDescGZIP() []byte {
	file_bitcoind_v1_bitcoind_proto_rawDescOnce.Do(func() {
		file_bitcoind_v1_bitcoind_proto_rawDescData = protoimpl.X.CompressGZIP(file_bitcoind_v1_bitcoind_proto_rawDescData)
	})
	return file_bitcoind_v1_bitcoind_proto_rawDescData
}

var file_bitcoind_v1_bitcoind_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_bitcoind_v1_bitcoind_proto_goTypes = []any{
	(*ListRecentBlocksRequest)(nil),              // 0: bitcoind.v1.ListRecentBlocksRequest
	(*ListRecentBlocksResponse)(nil),             // 1: bitcoind.v1.ListRecentBlocksResponse
	(*ListUnconfirmedTransactionsRequest)(nil),   // 2: bitcoind.v1.ListUnconfirmedTransactionsRequest
	(*ListUnconfirmedTransactionsResponse)(nil),  // 3: bitcoind.v1.ListUnconfirmedTransactionsResponse
	(*UnconfirmedTransaction)(nil),               // 4: bitcoind.v1.UnconfirmedTransaction
	(*GetBlockchainInfoResponse)(nil),            // 5: bitcoind.v1.GetBlockchainInfoResponse
	(*Peer)(nil),                                 // 6: bitcoind.v1.Peer
	(*ListPeersResponse)(nil),                    // 7: bitcoind.v1.ListPeersResponse
	(*EstimateSmartFeeRequest)(nil),              // 8: bitcoind.v1.EstimateSmartFeeRequest
	(*EstimateSmartFeeResponse)(nil),             // 9: bitcoind.v1.EstimateSmartFeeResponse
	(*ListRecentBlocksResponse_RecentBlock)(nil), // 10: bitcoind.v1.ListRecentBlocksResponse.RecentBlock
	(*timestamppb.Timestamp)(nil),                // 11: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),                        // 12: google.protobuf.Empty
}
var file_bitcoind_v1_bitcoind_proto_depIdxs = []int32{
	10, // 0: bitcoind.v1.ListRecentBlocksResponse.recent_blocks:type_name -> bitcoind.v1.ListRecentBlocksResponse.RecentBlock
	4,  // 1: bitcoind.v1.ListUnconfirmedTransactionsResponse.unconfirmed_transactions:type_name -> bitcoind.v1.UnconfirmedTransaction
	11, // 2: bitcoind.v1.UnconfirmedTransaction.time:type_name -> google.protobuf.Timestamp
	6,  // 3: bitcoind.v1.ListPeersResponse.peers:type_name -> bitcoind.v1.Peer
	11, // 4: bitcoind.v1.ListRecentBlocksResponse.RecentBlock.block_time:type_name -> google.protobuf.Timestamp
	2,  // 5: bitcoind.v1.BitcoindService.ListUnconfirmedTransactions:input_type -> bitcoind.v1.ListUnconfirmedTransactionsRequest
	0,  // 6: bitcoind.v1.BitcoindService.ListRecentBlocks:input_type -> bitcoind.v1.ListRecentBlocksRequest
	12, // 7: bitcoind.v1.BitcoindService.GetBlockchainInfo:input_type -> google.protobuf.Empty
	12, // 8: bitcoind.v1.BitcoindService.ListPeers:input_type -> google.protobuf.Empty
	8,  // 9: bitcoind.v1.BitcoindService.EstimateSmartFee:input_type -> bitcoind.v1.EstimateSmartFeeRequest
	3,  // 10: bitcoind.v1.BitcoindService.ListUnconfirmedTransactions:output_type -> bitcoind.v1.ListUnconfirmedTransactionsResponse
	1,  // 11: bitcoind.v1.BitcoindService.ListRecentBlocks:output_type -> bitcoind.v1.ListRecentBlocksResponse
	5,  // 12: bitcoind.v1.BitcoindService.GetBlockchainInfo:output_type -> bitcoind.v1.GetBlockchainInfoResponse
	7,  // 13: bitcoind.v1.BitcoindService.ListPeers:output_type -> bitcoind.v1.ListPeersResponse
	9,  // 14: bitcoind.v1.BitcoindService.EstimateSmartFee:output_type -> bitcoind.v1.EstimateSmartFeeResponse
	10, // [10:15] is the sub-list for method output_type
	5,  // [5:10] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_bitcoind_v1_bitcoind_proto_init() }
func file_bitcoind_v1_bitcoind_proto_init() {
	if File_bitcoind_v1_bitcoind_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_bitcoind_v1_bitcoind_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_bitcoind_v1_bitcoind_proto_goTypes,
		DependencyIndexes: file_bitcoind_v1_bitcoind_proto_depIdxs,
		MessageInfos:      file_bitcoind_v1_bitcoind_proto_msgTypes,
	}.Build()
	File_bitcoind_v1_bitcoind_proto = out.File
	file_bitcoind_v1_bitcoind_proto_rawDesc = nil
	file_bitcoind_v1_bitcoind_proto_goTypes = nil
	file_bitcoind_v1_bitcoind_proto_depIdxs = nil
}
