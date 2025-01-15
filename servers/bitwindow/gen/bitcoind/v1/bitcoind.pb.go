// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
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
	state         protoimpl.MessageState `protogen:"open.v1"`
	Count         int64                  `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
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

type Block struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	BlockTime     *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=block_time,json=blockTime,proto3" json:"block_time,omitempty"`
	BlockHeight   uint32                 `protobuf:"varint,2,opt,name=block_height,json=blockHeight,proto3" json:"block_height,omitempty"`
	Hash          string                 `protobuf:"bytes,3,opt,name=hash,proto3" json:"hash,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Block) Reset() {
	*x = Block{}
	mi := &file_bitcoind_v1_bitcoind_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Block) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Block) ProtoMessage() {}

func (x *Block) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Block.ProtoReflect.Descriptor instead.
func (*Block) Descriptor() ([]byte, []int) {
	return file_bitcoind_v1_bitcoind_proto_rawDescGZIP(), []int{1}
}

func (x *Block) GetBlockTime() *timestamppb.Timestamp {
	if x != nil {
		return x.BlockTime
	}
	return nil
}

func (x *Block) GetBlockHeight() uint32 {
	if x != nil {
		return x.BlockHeight
	}
	return 0
}

func (x *Block) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

type ListRecentBlocksResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	RecentBlocks  []*Block               `protobuf:"bytes,4,rep,name=recent_blocks,json=recentBlocks,proto3" json:"recent_blocks,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListRecentBlocksResponse) Reset() {
	*x = ListRecentBlocksResponse{}
	mi := &file_bitcoind_v1_bitcoind_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListRecentBlocksResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRecentBlocksResponse) ProtoMessage() {}

func (x *ListRecentBlocksResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use ListRecentBlocksResponse.ProtoReflect.Descriptor instead.
func (*ListRecentBlocksResponse) Descriptor() ([]byte, []int) {
	return file_bitcoind_v1_bitcoind_proto_rawDescGZIP(), []int{2}
}

func (x *ListRecentBlocksResponse) GetRecentBlocks() []*Block {
	if x != nil {
		return x.RecentBlocks
	}
	return nil
}

type ListRecentTransactionsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Count         int64                  `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListRecentTransactionsRequest) Reset() {
	*x = ListRecentTransactionsRequest{}
	mi := &file_bitcoind_v1_bitcoind_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListRecentTransactionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRecentTransactionsRequest) ProtoMessage() {}

func (x *ListRecentTransactionsRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use ListRecentTransactionsRequest.ProtoReflect.Descriptor instead.
func (*ListRecentTransactionsRequest) Descriptor() ([]byte, []int) {
	return file_bitcoind_v1_bitcoind_proto_rawDescGZIP(), []int{3}
}

func (x *ListRecentTransactionsRequest) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

type ListRecentTransactionsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Transactions  []*RecentTransaction   `protobuf:"bytes,1,rep,name=transactions,proto3" json:"transactions,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListRecentTransactionsResponse) Reset() {
	*x = ListRecentTransactionsResponse{}
	mi := &file_bitcoind_v1_bitcoind_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListRecentTransactionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRecentTransactionsResponse) ProtoMessage() {}

func (x *ListRecentTransactionsResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use ListRecentTransactionsResponse.ProtoReflect.Descriptor instead.
func (*ListRecentTransactionsResponse) Descriptor() ([]byte, []int) {
	return file_bitcoind_v1_bitcoind_proto_rawDescGZIP(), []int{4}
}

func (x *ListRecentTransactionsResponse) GetTransactions() []*RecentTransaction {
	if x != nil {
		return x.Transactions
	}
	return nil
}

type RecentTransaction struct {
	state            protoimpl.MessageState `protogen:"open.v1"`
	VirtualSize      uint32                 `protobuf:"varint,1,opt,name=virtual_size,json=virtualSize,proto3" json:"virtual_size,omitempty"`
	Time             *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=time,proto3" json:"time,omitempty"`
	Txid             string                 `protobuf:"bytes,3,opt,name=txid,proto3" json:"txid,omitempty"`
	FeeSats          uint64                 `protobuf:"varint,4,opt,name=fee_sats,json=feeSats,proto3" json:"fee_sats,omitempty"`
	ConfirmedInBlock *Block                 `protobuf:"bytes,5,opt,name=confirmed_in_block,json=confirmedInBlock,proto3,oneof" json:"confirmed_in_block,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *RecentTransaction) Reset() {
	*x = RecentTransaction{}
	mi := &file_bitcoind_v1_bitcoind_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RecentTransaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecentTransaction) ProtoMessage() {}

func (x *RecentTransaction) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use RecentTransaction.ProtoReflect.Descriptor instead.
func (*RecentTransaction) Descriptor() ([]byte, []int) {
	return file_bitcoind_v1_bitcoind_proto_rawDescGZIP(), []int{5}
}

func (x *RecentTransaction) GetVirtualSize() uint32 {
	if x != nil {
		return x.VirtualSize
	}
	return 0
}

func (x *RecentTransaction) GetTime() *timestamppb.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

func (x *RecentTransaction) GetTxid() string {
	if x != nil {
		return x.Txid
	}
	return ""
}

func (x *RecentTransaction) GetFeeSats() uint64 {
	if x != nil {
		return x.FeeSats
	}
	return 0
}

func (x *RecentTransaction) GetConfirmedInBlock() *Block {
	if x != nil {
		return x.ConfirmedInBlock
	}
	return nil
}

type GetBlockchainInfoResponse struct {
	state                protoimpl.MessageState `protogen:"open.v1"`
	Chain                string                 `protobuf:"bytes,1,opt,name=chain,proto3" json:"chain,omitempty"`
	Blocks               uint32                 `protobuf:"varint,2,opt,name=blocks,proto3" json:"blocks,omitempty"`
	Headers              uint32                 `protobuf:"varint,3,opt,name=headers,proto3" json:"headers,omitempty"`
	BestBlockHash        string                 `protobuf:"bytes,4,opt,name=best_block_hash,json=bestBlockHash,proto3" json:"best_block_hash,omitempty"`
	InitialBlockDownload bool                   `protobuf:"varint,8,opt,name=initial_block_download,json=initialBlockDownload,proto3" json:"initial_block_download,omitempty"`
	unknownFields        protoimpl.UnknownFields
	sizeCache            protoimpl.SizeCache
}

func (x *GetBlockchainInfoResponse) Reset() {
	*x = GetBlockchainInfoResponse{}
	mi := &file_bitcoind_v1_bitcoind_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetBlockchainInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBlockchainInfoResponse) ProtoMessage() {}

func (x *GetBlockchainInfoResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GetBlockchainInfoResponse.ProtoReflect.Descriptor instead.
func (*GetBlockchainInfoResponse) Descriptor() ([]byte, []int) {
	return file_bitcoind_v1_bitcoind_proto_rawDescGZIP(), []int{6}
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
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Addr          string                 `protobuf:"bytes,2,opt,name=addr,proto3" json:"addr,omitempty"`
	SyncedBlocks  int32                  `protobuf:"varint,3,opt,name=synced_blocks,json=syncedBlocks,proto3" json:"synced_blocks,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Peer) Reset() {
	*x = Peer{}
	mi := &file_bitcoind_v1_bitcoind_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Peer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Peer) ProtoMessage() {}

func (x *Peer) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Peer.ProtoReflect.Descriptor instead.
func (*Peer) Descriptor() ([]byte, []int) {
	return file_bitcoind_v1_bitcoind_proto_rawDescGZIP(), []int{7}
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
	state         protoimpl.MessageState `protogen:"open.v1"`
	Peers         []*Peer                `protobuf:"bytes,1,rep,name=peers,proto3" json:"peers,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListPeersResponse) Reset() {
	*x = ListPeersResponse{}
	mi := &file_bitcoind_v1_bitcoind_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListPeersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPeersResponse) ProtoMessage() {}

func (x *ListPeersResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use ListPeersResponse.ProtoReflect.Descriptor instead.
func (*ListPeersResponse) Descriptor() ([]byte, []int) {
	return file_bitcoind_v1_bitcoind_proto_rawDescGZIP(), []int{8}
}

func (x *ListPeersResponse) GetPeers() []*Peer {
	if x != nil {
		return x.Peers
	}
	return nil
}

type EstimateSmartFeeRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ConfTarget    int64                  `protobuf:"varint,1,opt,name=conf_target,json=confTarget,proto3" json:"conf_target,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *EstimateSmartFeeRequest) Reset() {
	*x = EstimateSmartFeeRequest{}
	mi := &file_bitcoind_v1_bitcoind_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EstimateSmartFeeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EstimateSmartFeeRequest) ProtoMessage() {}

func (x *EstimateSmartFeeRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use EstimateSmartFeeRequest.ProtoReflect.Descriptor instead.
func (*EstimateSmartFeeRequest) Descriptor() ([]byte, []int) {
	return file_bitcoind_v1_bitcoind_proto_rawDescGZIP(), []int{9}
}

func (x *EstimateSmartFeeRequest) GetConfTarget() int64 {
	if x != nil {
		return x.ConfTarget
	}
	return 0
}

type EstimateSmartFeeResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	FeeRate       float64                `protobuf:"fixed64,1,opt,name=fee_rate,json=feeRate,proto3" json:"fee_rate,omitempty"`
	Errors        []string               `protobuf:"bytes,2,rep,name=errors,proto3" json:"errors,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *EstimateSmartFeeResponse) Reset() {
	*x = EstimateSmartFeeResponse{}
	mi := &file_bitcoind_v1_bitcoind_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EstimateSmartFeeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EstimateSmartFeeResponse) ProtoMessage() {}

func (x *EstimateSmartFeeResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use EstimateSmartFeeResponse.ProtoReflect.Descriptor instead.
func (*EstimateSmartFeeResponse) Descriptor() ([]byte, []int) {
	return file_bitcoind_v1_bitcoind_proto_rawDescGZIP(), []int{10}
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
	0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x79, 0x0a, 0x05, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x12, 0x39, 0x0a, 0x0a, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68,
	0x61, 0x73, 0x68, 0x22, 0x53, 0x0a, 0x18, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x63, 0x65, 0x6e,
	0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x37, 0x0a, 0x0d, 0x72, 0x65, 0x63, 0x65, 0x6e, 0x74, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x62, 0x69, 0x74, 0x63, 0x6f, 0x69, 0x6e,
	0x64, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x0c, 0x72, 0x65, 0x63, 0x65,
	0x6e, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x22, 0x35, 0x0a, 0x1d, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x63, 0x65, 0x6e, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22,
	0x64, 0x0a, 0x1e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x63, 0x65, 0x6e, 0x74, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x42, 0x0a, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x62, 0x69, 0x74, 0x63, 0x6f, 0x69,
	0x6e, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x63, 0x65, 0x6e, 0x74, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0xf3, 0x01, 0x0a, 0x11, 0x52, 0x65, 0x63, 0x65, 0x6e, 0x74,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x76,
	0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0b, 0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x2e,
	0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x78, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x78,
	0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x66, 0x65, 0x65, 0x5f, 0x73, 0x61, 0x74, 0x73, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x66, 0x65, 0x65, 0x53, 0x61, 0x74, 0x73, 0x12, 0x45, 0x0a,
	0x12, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x65, 0x64, 0x5f, 0x69, 0x6e, 0x5f, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x62, 0x69, 0x74, 0x63,
	0x6f, 0x69, 0x6e, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x00, 0x52,
	0x10, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x65, 0x64, 0x49, 0x6e, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x88, 0x01, 0x01, 0x42, 0x15, 0x0a, 0x13, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d,
	0x65, 0x64, 0x5f, 0x69, 0x6e, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x22, 0xc1, 0x01, 0x0a, 0x19,
	0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x68, 0x61,
	0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x12,
	0x16, 0x0a, 0x06, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x06, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x73, 0x12, 0x26, 0x0a, 0x0f, 0x62, 0x65, 0x73, 0x74, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f,
	0x68, 0x61, 0x73, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x62, 0x65, 0x73, 0x74,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x12, 0x34, 0x0a, 0x16, 0x69, 0x6e, 0x69,
	0x74, 0x69, 0x61, 0x6c, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x64, 0x6f, 0x77, 0x6e, 0x6c,
	0x6f, 0x61, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x14, 0x69, 0x6e, 0x69, 0x74, 0x69,
	0x61, 0x6c, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x22,
	0x4f, 0x0a, 0x04, 0x50, 0x65, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x64, 0x64, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x64, 0x64, 0x72, 0x12, 0x23, 0x0a, 0x0d, 0x73,
	0x79, 0x6e, 0x63, 0x65, 0x64, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0c, 0x73, 0x79, 0x6e, 0x63, 0x65, 0x64, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73,
	0x22, 0x3c, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x65, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x05, 0x70, 0x65, 0x65, 0x72, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x62, 0x69, 0x74, 0x63, 0x6f, 0x69, 0x6e, 0x64, 0x2e,
	0x76, 0x31, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x52, 0x05, 0x70, 0x65, 0x65, 0x72, 0x73, 0x22, 0x3a,
	0x0a, 0x17, 0x45, 0x73, 0x74, 0x69, 0x6d, 0x61, 0x74, 0x65, 0x53, 0x6d, 0x61, 0x72, 0x74, 0x46,
	0x65, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6f, 0x6e,
	0x66, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a,
	0x63, 0x6f, 0x6e, 0x66, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x22, 0x4d, 0x0a, 0x18, 0x45, 0x73,
	0x74, 0x69, 0x6d, 0x61, 0x74, 0x65, 0x53, 0x6d, 0x61, 0x72, 0x74, 0x46, 0x65, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x66, 0x65, 0x65, 0x5f, 0x72, 0x61,
	0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x07, 0x66, 0x65, 0x65, 0x52, 0x61, 0x74,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x32, 0xe0, 0x03, 0x0a, 0x0f, 0x42, 0x69,
	0x74, 0x63, 0x6f, 0x69, 0x6e, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x71, 0x0a,
	0x16, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x63, 0x65, 0x6e, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x2a, 0x2e, 0x62, 0x69, 0x74, 0x63, 0x6f, 0x69,
	0x6e, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x63, 0x65, 0x6e, 0x74,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x62, 0x69, 0x74, 0x63, 0x6f, 0x69, 0x6e, 0x64, 0x2e, 0x76,
	0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x63, 0x65, 0x6e, 0x74, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x5f, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x63, 0x65, 0x6e, 0x74, 0x42, 0x6c,
	0x6f, 0x63, 0x6b, 0x73, 0x12, 0x24, 0x2e, 0x62, 0x69, 0x74, 0x63, 0x6f, 0x69, 0x6e, 0x64, 0x2e,
	0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x63, 0x65, 0x6e, 0x74, 0x42, 0x6c, 0x6f,
	0x63, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x62, 0x69, 0x74,
	0x63, 0x6f, 0x69, 0x6e, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x63,
	0x65, 0x6e, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x53, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61,
	0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x26,
	0x2e, 0x62, 0x69, 0x74, 0x63, 0x6f, 0x69, 0x6e, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x65,
	0x65, 0x72, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1e, 0x2e, 0x62, 0x69,
	0x74, 0x63, 0x6f, 0x69, 0x6e, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x65,
	0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5f, 0x0a, 0x10, 0x45,
	0x73, 0x74, 0x69, 0x6d, 0x61, 0x74, 0x65, 0x53, 0x6d, 0x61, 0x72, 0x74, 0x46, 0x65, 0x65, 0x12,
	0x24, 0x2e, 0x62, 0x69, 0x74, 0x63, 0x6f, 0x69, 0x6e, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x73,
	0x74, 0x69, 0x6d, 0x61, 0x74, 0x65, 0x53, 0x6d, 0x61, 0x72, 0x74, 0x46, 0x65, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x62, 0x69, 0x74, 0x63, 0x6f, 0x69, 0x6e, 0x64,
	0x2e, 0x76, 0x31, 0x2e, 0x45, 0x73, 0x74, 0x69, 0x6d, 0x61, 0x74, 0x65, 0x53, 0x6d, 0x61, 0x72,
	0x74, 0x46, 0x65, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0xbd, 0x01, 0x0a,
	0x0f, 0x63, 0x6f, 0x6d, 0x2e, 0x62, 0x69, 0x74, 0x63, 0x6f, 0x69, 0x6e, 0x64, 0x2e, 0x76, 0x31,
	0x42, 0x0d, 0x42, 0x69, 0x74, 0x63, 0x6f, 0x69, 0x6e, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x4e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4c, 0x61,
	0x79, 0x65, 0x72, 0x54, 0x77, 0x6f, 0x2d, 0x4c, 0x61, 0x62, 0x73, 0x2f, 0x73, 0x69, 0x64, 0x65,
	0x73, 0x61, 0x69, 0x6c, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x2f, 0x62, 0x69, 0x74,
	0x77, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x69, 0x74, 0x63, 0x6f,
	0x69, 0x6e, 0x64, 0x2f, 0x76, 0x31, 0x3b, 0x62, 0x69, 0x74, 0x63, 0x6f, 0x69, 0x6e, 0x64, 0x76,
	0x31, 0xa2, 0x02, 0x03, 0x42, 0x58, 0x58, 0xaa, 0x02, 0x0b, 0x42, 0x69, 0x74, 0x63, 0x6f, 0x69,
	0x6e, 0x64, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0b, 0x42, 0x69, 0x74, 0x63, 0x6f, 0x69, 0x6e, 0x64,
	0x5c, 0x56, 0x31, 0xe2, 0x02, 0x17, 0x42, 0x69, 0x74, 0x63, 0x6f, 0x69, 0x6e, 0x64, 0x5c, 0x56,
	0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0c,
	0x42, 0x69, 0x74, 0x63, 0x6f, 0x69, 0x6e, 0x64, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
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
	(*ListRecentBlocksRequest)(nil),        // 0: bitcoind.v1.ListRecentBlocksRequest
	(*Block)(nil),                          // 1: bitcoind.v1.Block
	(*ListRecentBlocksResponse)(nil),       // 2: bitcoind.v1.ListRecentBlocksResponse
	(*ListRecentTransactionsRequest)(nil),  // 3: bitcoind.v1.ListRecentTransactionsRequest
	(*ListRecentTransactionsResponse)(nil), // 4: bitcoind.v1.ListRecentTransactionsResponse
	(*RecentTransaction)(nil),              // 5: bitcoind.v1.RecentTransaction
	(*GetBlockchainInfoResponse)(nil),      // 6: bitcoind.v1.GetBlockchainInfoResponse
	(*Peer)(nil),                           // 7: bitcoind.v1.Peer
	(*ListPeersResponse)(nil),              // 8: bitcoind.v1.ListPeersResponse
	(*EstimateSmartFeeRequest)(nil),        // 9: bitcoind.v1.EstimateSmartFeeRequest
	(*EstimateSmartFeeResponse)(nil),       // 10: bitcoind.v1.EstimateSmartFeeResponse
	(*timestamppb.Timestamp)(nil),          // 11: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),                  // 12: google.protobuf.Empty
}
var file_bitcoind_v1_bitcoind_proto_depIdxs = []int32{
	11, // 0: bitcoind.v1.Block.block_time:type_name -> google.protobuf.Timestamp
	1,  // 1: bitcoind.v1.ListRecentBlocksResponse.recent_blocks:type_name -> bitcoind.v1.Block
	5,  // 2: bitcoind.v1.ListRecentTransactionsResponse.transactions:type_name -> bitcoind.v1.RecentTransaction
	11, // 3: bitcoind.v1.RecentTransaction.time:type_name -> google.protobuf.Timestamp
	1,  // 4: bitcoind.v1.RecentTransaction.confirmed_in_block:type_name -> bitcoind.v1.Block
	7,  // 5: bitcoind.v1.ListPeersResponse.peers:type_name -> bitcoind.v1.Peer
	3,  // 6: bitcoind.v1.BitcoindService.ListRecentTransactions:input_type -> bitcoind.v1.ListRecentTransactionsRequest
	0,  // 7: bitcoind.v1.BitcoindService.ListRecentBlocks:input_type -> bitcoind.v1.ListRecentBlocksRequest
	12, // 8: bitcoind.v1.BitcoindService.GetBlockchainInfo:input_type -> google.protobuf.Empty
	12, // 9: bitcoind.v1.BitcoindService.ListPeers:input_type -> google.protobuf.Empty
	9,  // 10: bitcoind.v1.BitcoindService.EstimateSmartFee:input_type -> bitcoind.v1.EstimateSmartFeeRequest
	4,  // 11: bitcoind.v1.BitcoindService.ListRecentTransactions:output_type -> bitcoind.v1.ListRecentTransactionsResponse
	2,  // 12: bitcoind.v1.BitcoindService.ListRecentBlocks:output_type -> bitcoind.v1.ListRecentBlocksResponse
	6,  // 13: bitcoind.v1.BitcoindService.GetBlockchainInfo:output_type -> bitcoind.v1.GetBlockchainInfoResponse
	8,  // 14: bitcoind.v1.BitcoindService.ListPeers:output_type -> bitcoind.v1.ListPeersResponse
	10, // 15: bitcoind.v1.BitcoindService.EstimateSmartFee:output_type -> bitcoind.v1.EstimateSmartFeeResponse
	11, // [11:16] is the sub-list for method output_type
	6,  // [6:11] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_bitcoind_v1_bitcoind_proto_init() }
func file_bitcoind_v1_bitcoind_proto_init() {
	if File_bitcoind_v1_bitcoind_proto != nil {
		return
	}
	file_bitcoind_v1_bitcoind_proto_msgTypes[5].OneofWrappers = []any{}
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
