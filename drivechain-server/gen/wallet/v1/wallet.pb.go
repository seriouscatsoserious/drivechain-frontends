// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: wallet/v1/wallet.proto

package walletv1

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

type GetNewAddressResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Index   uint32 `protobuf:"varint,2,opt,name=index,proto3" json:"index,omitempty"`
}

func (x *GetNewAddressResponse) Reset() {
	*x = GetNewAddressResponse{}
	mi := &file_wallet_v1_wallet_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetNewAddressResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNewAddressResponse) ProtoMessage() {}

func (x *GetNewAddressResponse) ProtoReflect() protoreflect.Message {
	mi := &file_wallet_v1_wallet_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNewAddressResponse.ProtoReflect.Descriptor instead.
func (*GetNewAddressResponse) Descriptor() ([]byte, []int) {
	return file_wallet_v1_wallet_proto_rawDescGZIP(), []int{0}
}

func (x *GetNewAddressResponse) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *GetNewAddressResponse) GetIndex() uint32 {
	if x != nil {
		return x.Index
	}
	return 0
}

type SendTransactionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Address -> satoshi amount
	Destinations map[string]uint64 `protobuf:"bytes,1,rep,name=destinations,proto3" json:"destinations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	// Fee rate, measured in BTC/kvB. If set to zero, a reasonable
	// rate is used by asking Core for an estimate.
	FeeRate float64 `protobuf:"fixed64,2,opt,name=fee_rate,json=feeRate,proto3" json:"fee_rate,omitempty"`
	// Whether to activate replace by fee for this transaction
	Rbf bool `protobuf:"varint,3,opt,name=rbf,proto3" json:"rbf,omitempty"`
}

func (x *SendTransactionRequest) Reset() {
	*x = SendTransactionRequest{}
	mi := &file_wallet_v1_wallet_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SendTransactionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendTransactionRequest) ProtoMessage() {}

func (x *SendTransactionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_wallet_v1_wallet_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendTransactionRequest.ProtoReflect.Descriptor instead.
func (*SendTransactionRequest) Descriptor() ([]byte, []int) {
	return file_wallet_v1_wallet_proto_rawDescGZIP(), []int{1}
}

func (x *SendTransactionRequest) GetDestinations() map[string]uint64 {
	if x != nil {
		return x.Destinations
	}
	return nil
}

func (x *SendTransactionRequest) GetFeeRate() float64 {
	if x != nil {
		return x.FeeRate
	}
	return 0
}

func (x *SendTransactionRequest) GetRbf() bool {
	if x != nil {
		return x.Rbf
	}
	return false
}

type SendTransactionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Txid string `protobuf:"bytes,1,opt,name=txid,proto3" json:"txid,omitempty"`
}

func (x *SendTransactionResponse) Reset() {
	*x = SendTransactionResponse{}
	mi := &file_wallet_v1_wallet_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SendTransactionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendTransactionResponse) ProtoMessage() {}

func (x *SendTransactionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_wallet_v1_wallet_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendTransactionResponse.ProtoReflect.Descriptor instead.
func (*SendTransactionResponse) Descriptor() ([]byte, []int) {
	return file_wallet_v1_wallet_proto_rawDescGZIP(), []int{2}
}

func (x *SendTransactionResponse) GetTxid() string {
	if x != nil {
		return x.Txid
	}
	return ""
}

type GetBalanceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConfirmedSatoshi uint64 `protobuf:"varint,1,opt,name=confirmed_satoshi,json=confirmedSatoshi,proto3" json:"confirmed_satoshi,omitempty"`
	PendingSatoshi   uint64 `protobuf:"varint,2,opt,name=pending_satoshi,json=pendingSatoshi,proto3" json:"pending_satoshi,omitempty"`
}

func (x *GetBalanceResponse) Reset() {
	*x = GetBalanceResponse{}
	mi := &file_wallet_v1_wallet_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetBalanceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBalanceResponse) ProtoMessage() {}

func (x *GetBalanceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_wallet_v1_wallet_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBalanceResponse.ProtoReflect.Descriptor instead.
func (*GetBalanceResponse) Descriptor() ([]byte, []int) {
	return file_wallet_v1_wallet_proto_rawDescGZIP(), []int{3}
}

func (x *GetBalanceResponse) GetConfirmedSatoshi() uint64 {
	if x != nil {
		return x.ConfirmedSatoshi
	}
	return 0
}

func (x *GetBalanceResponse) GetPendingSatoshi() uint64 {
	if x != nil {
		return x.PendingSatoshi
	}
	return 0
}

type ListTransactionsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Transactions []*Transaction `protobuf:"bytes,1,rep,name=transactions,proto3" json:"transactions,omitempty"`
}

func (x *ListTransactionsResponse) Reset() {
	*x = ListTransactionsResponse{}
	mi := &file_wallet_v1_wallet_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListTransactionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTransactionsResponse) ProtoMessage() {}

func (x *ListTransactionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_wallet_v1_wallet_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTransactionsResponse.ProtoReflect.Descriptor instead.
func (*ListTransactionsResponse) Descriptor() ([]byte, []int) {
	return file_wallet_v1_wallet_proto_rawDescGZIP(), []int{4}
}

func (x *ListTransactionsResponse) GetTransactions() []*Transaction {
	if x != nil {
		return x.Transactions
	}
	return nil
}

type Confirmation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Height    uint32                 `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *Confirmation) Reset() {
	*x = Confirmation{}
	mi := &file_wallet_v1_wallet_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Confirmation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Confirmation) ProtoMessage() {}

func (x *Confirmation) ProtoReflect() protoreflect.Message {
	mi := &file_wallet_v1_wallet_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Confirmation.ProtoReflect.Descriptor instead.
func (*Confirmation) Descriptor() ([]byte, []int) {
	return file_wallet_v1_wallet_proto_rawDescGZIP(), []int{5}
}

func (x *Confirmation) GetHeight() uint32 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *Confirmation) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

type Transaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Txid             string        `protobuf:"bytes,1,opt,name=txid,proto3" json:"txid,omitempty"`
	FeeSatoshi       uint64        `protobuf:"varint,2,opt,name=fee_satoshi,json=feeSatoshi,proto3" json:"fee_satoshi,omitempty"`
	ReceivedSatoshi  uint64        `protobuf:"varint,3,opt,name=received_satoshi,json=receivedSatoshi,proto3" json:"received_satoshi,omitempty"`
	SentSatoshi      uint64        `protobuf:"varint,4,opt,name=sent_satoshi,json=sentSatoshi,proto3" json:"sent_satoshi,omitempty"`
	ConfirmationTime *Confirmation `protobuf:"bytes,5,opt,name=confirmation_time,json=confirmationTime,proto3" json:"confirmation_time,omitempty"`
}

func (x *Transaction) Reset() {
	*x = Transaction{}
	mi := &file_wallet_v1_wallet_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Transaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transaction) ProtoMessage() {}

func (x *Transaction) ProtoReflect() protoreflect.Message {
	mi := &file_wallet_v1_wallet_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Transaction.ProtoReflect.Descriptor instead.
func (*Transaction) Descriptor() ([]byte, []int) {
	return file_wallet_v1_wallet_proto_rawDescGZIP(), []int{6}
}

func (x *Transaction) GetTxid() string {
	if x != nil {
		return x.Txid
	}
	return ""
}

func (x *Transaction) GetFeeSatoshi() uint64 {
	if x != nil {
		return x.FeeSatoshi
	}
	return 0
}

func (x *Transaction) GetReceivedSatoshi() uint64 {
	if x != nil {
		return x.ReceivedSatoshi
	}
	return 0
}

func (x *Transaction) GetSentSatoshi() uint64 {
	if x != nil {
		return x.SentSatoshi
	}
	return 0
}

func (x *Transaction) GetConfirmationTime() *Confirmation {
	if x != nil {
		return x.ConfirmationTime
	}
	return nil
}

type ListSidechainDepositsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Slot int32 `protobuf:"varint,1,opt,name=slot,proto3" json:"slot,omitempty"`
}

func (x *ListSidechainDepositsRequest) Reset() {
	*x = ListSidechainDepositsRequest{}
	mi := &file_wallet_v1_wallet_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListSidechainDepositsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSidechainDepositsRequest) ProtoMessage() {}

func (x *ListSidechainDepositsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_wallet_v1_wallet_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSidechainDepositsRequest.ProtoReflect.Descriptor instead.
func (*ListSidechainDepositsRequest) Descriptor() ([]byte, []int) {
	return file_wallet_v1_wallet_proto_rawDescGZIP(), []int{7}
}

func (x *ListSidechainDepositsRequest) GetSlot() int32 {
	if x != nil {
		return x.Slot
	}
	return 0
}

type ListSidechainDepositsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Deposits []*ListSidechainDepositsResponse_SidechainDeposit `protobuf:"bytes,1,rep,name=deposits,proto3" json:"deposits,omitempty"`
}

func (x *ListSidechainDepositsResponse) Reset() {
	*x = ListSidechainDepositsResponse{}
	mi := &file_wallet_v1_wallet_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListSidechainDepositsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSidechainDepositsResponse) ProtoMessage() {}

func (x *ListSidechainDepositsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_wallet_v1_wallet_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSidechainDepositsResponse.ProtoReflect.Descriptor instead.
func (*ListSidechainDepositsResponse) Descriptor() ([]byte, []int) {
	return file_wallet_v1_wallet_proto_rawDescGZIP(), []int{8}
}

func (x *ListSidechainDepositsResponse) GetDeposits() []*ListSidechainDepositsResponse_SidechainDeposit {
	if x != nil {
		return x.Deposits
	}
	return nil
}

type CreateSidechainDepositRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The sidechain deposit address to send to.
	Destination string `protobuf:"bytes,1,opt,name=destination,proto3" json:"destination,omitempty"`
	// The amount in BTC to send. eg 0.1
	Amount float64 `protobuf:"fixed64,2,opt,name=amount,proto3" json:"amount,omitempty"`
	// The fee in BTC
	Fee float64 `protobuf:"fixed64,3,opt,name=fee,proto3" json:"fee,omitempty"`
}

func (x *CreateSidechainDepositRequest) Reset() {
	*x = CreateSidechainDepositRequest{}
	mi := &file_wallet_v1_wallet_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateSidechainDepositRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSidechainDepositRequest) ProtoMessage() {}

func (x *CreateSidechainDepositRequest) ProtoReflect() protoreflect.Message {
	mi := &file_wallet_v1_wallet_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSidechainDepositRequest.ProtoReflect.Descriptor instead.
func (*CreateSidechainDepositRequest) Descriptor() ([]byte, []int) {
	return file_wallet_v1_wallet_proto_rawDescGZIP(), []int{9}
}

func (x *CreateSidechainDepositRequest) GetDestination() string {
	if x != nil {
		return x.Destination
	}
	return ""
}

func (x *CreateSidechainDepositRequest) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *CreateSidechainDepositRequest) GetFee() float64 {
	if x != nil {
		return x.Fee
	}
	return 0
}

type CreateSidechainDepositResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Txid string `protobuf:"bytes,1,opt,name=txid,proto3" json:"txid,omitempty"`
}

func (x *CreateSidechainDepositResponse) Reset() {
	*x = CreateSidechainDepositResponse{}
	mi := &file_wallet_v1_wallet_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateSidechainDepositResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSidechainDepositResponse) ProtoMessage() {}

func (x *CreateSidechainDepositResponse) ProtoReflect() protoreflect.Message {
	mi := &file_wallet_v1_wallet_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSidechainDepositResponse.ProtoReflect.Descriptor instead.
func (*CreateSidechainDepositResponse) Descriptor() ([]byte, []int) {
	return file_wallet_v1_wallet_proto_rawDescGZIP(), []int{10}
}

func (x *CreateSidechainDepositResponse) GetTxid() string {
	if x != nil {
		return x.Txid
	}
	return ""
}

type ListSidechainDepositsResponse_SidechainDeposit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Txid          string  `protobuf:"bytes,1,opt,name=txid,proto3" json:"txid,omitempty"`
	Address       string  `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Amount        float64 `protobuf:"fixed64,3,opt,name=amount,proto3" json:"amount,omitempty"`
	Fee           float64 `protobuf:"fixed64,4,opt,name=fee,proto3" json:"fee,omitempty"`
	Confirmations int32   `protobuf:"varint,5,opt,name=confirmations,proto3" json:"confirmations,omitempty"`
}

func (x *ListSidechainDepositsResponse_SidechainDeposit) Reset() {
	*x = ListSidechainDepositsResponse_SidechainDeposit{}
	mi := &file_wallet_v1_wallet_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListSidechainDepositsResponse_SidechainDeposit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSidechainDepositsResponse_SidechainDeposit) ProtoMessage() {}

func (x *ListSidechainDepositsResponse_SidechainDeposit) ProtoReflect() protoreflect.Message {
	mi := &file_wallet_v1_wallet_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSidechainDepositsResponse_SidechainDeposit.ProtoReflect.Descriptor instead.
func (*ListSidechainDepositsResponse_SidechainDeposit) Descriptor() ([]byte, []int) {
	return file_wallet_v1_wallet_proto_rawDescGZIP(), []int{8, 0}
}

func (x *ListSidechainDepositsResponse_SidechainDeposit) GetTxid() string {
	if x != nil {
		return x.Txid
	}
	return ""
}

func (x *ListSidechainDepositsResponse_SidechainDeposit) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *ListSidechainDepositsResponse_SidechainDeposit) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *ListSidechainDepositsResponse_SidechainDeposit) GetFee() float64 {
	if x != nil {
		return x.Fee
	}
	return 0
}

func (x *ListSidechainDepositsResponse_SidechainDeposit) GetConfirmations() int32 {
	if x != nil {
		return x.Confirmations
	}
	return 0
}

var File_wallet_v1_wallet_proto protoreflect.FileDescriptor

var file_wallet_v1_wallet_proto_rawDesc = []byte{
	0x0a, 0x16, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x77, 0x61, 0x6c, 0x6c,
	0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74,
	0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x47, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x4e, 0x65, 0x77, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x22, 0xdf, 0x01, 0x0a, 0x16, 0x53,
	0x65, 0x6e, 0x64, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x57, 0x0a, 0x0c, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x77, 0x61,
	0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x44,
	0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x0c, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x19,
	0x0a, 0x08, 0x66, 0x65, 0x65, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x07, 0x66, 0x65, 0x65, 0x52, 0x61, 0x74, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x62, 0x66,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x72, 0x62, 0x66, 0x1a, 0x3f, 0x0a, 0x11, 0x44,
	0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x2d, 0x0a, 0x17,
	0x53, 0x65, 0x6e, 0x64, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x78, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x78, 0x69, 0x64, 0x22, 0x6a, 0x0a, 0x12, 0x47,
	0x65, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x2b, 0x0a, 0x11, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x65, 0x64, 0x5f, 0x73,
	0x61, 0x74, 0x6f, 0x73, 0x68, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x10, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x72, 0x6d, 0x65, 0x64, 0x53, 0x61, 0x74, 0x6f, 0x73, 0x68, 0x69, 0x12, 0x27,
	0x0a, 0x0f, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x61, 0x74, 0x6f, 0x73, 0x68,
	0x69, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0e, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67,
	0x53, 0x61, 0x74, 0x6f, 0x73, 0x68, 0x69, 0x22, 0x56, 0x0a, 0x18, 0x4c, 0x69, 0x73, 0x74, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x77, 0x61, 0x6c, 0x6c,
	0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22,
	0x60, 0x0a, 0x0c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x22, 0xd6, 0x01, 0x0a, 0x0b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x78, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x74, 0x78, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x65, 0x65, 0x5f, 0x73, 0x61, 0x74,
	0x6f, 0x73, 0x68, 0x69, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x66, 0x65, 0x65, 0x53,
	0x61, 0x74, 0x6f, 0x73, 0x68, 0x69, 0x12, 0x29, 0x0a, 0x10, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76,
	0x65, 0x64, 0x5f, 0x73, 0x61, 0x74, 0x6f, 0x73, 0x68, 0x69, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x0f, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x53, 0x61, 0x74, 0x6f, 0x73, 0x68,
	0x69, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x61, 0x74, 0x6f, 0x73, 0x68,
	0x69, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x73, 0x65, 0x6e, 0x74, 0x53, 0x61, 0x74,
	0x6f, 0x73, 0x68, 0x69, 0x12, 0x44, 0x0a, 0x11, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x10, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72,
	0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x32, 0x0a, 0x1c, 0x4c, 0x69,
	0x73, 0x74, 0x53, 0x69, 0x64, 0x65, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x44, 0x65, 0x70, 0x6f, 0x73,
	0x69, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6c,
	0x6f, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x6c, 0x6f, 0x74, 0x22, 0x89,
	0x02, 0x0a, 0x1d, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x69, 0x64, 0x65, 0x63, 0x68, 0x61, 0x69, 0x6e,
	0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x55, 0x0a, 0x08, 0x64, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x39, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x53, 0x69, 0x64, 0x65, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x44, 0x65, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x53, 0x69, 0x64,
	0x65, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x52, 0x08, 0x64,
	0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x73, 0x1a, 0x90, 0x01, 0x0a, 0x10, 0x53, 0x69, 0x64, 0x65,
	0x63, 0x68, 0x61, 0x69, 0x6e, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x78, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x78, 0x69, 0x64,
	0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x66, 0x65, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x03, 0x66, 0x65, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x6b, 0x0a, 0x1d, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x53, 0x69, 0x64, 0x65, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x44, 0x65, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a,
	0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x61,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x66, 0x65, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x03, 0x66, 0x65, 0x65, 0x22, 0x34, 0x0a, 0x1e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x53, 0x69, 0x64, 0x65, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x78, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x78, 0x69, 0x64, 0x32, 0xa5, 0x04,
	0x0a, 0x0d, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x58, 0x0a, 0x0f, 0x53, 0x65, 0x6e, 0x64, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x21, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x65, 0x6e, 0x64, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a, 0x0a, 0x47, 0x65, 0x74,
	0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a,
	0x1d, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x42,
	0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x49,
	0x0a, 0x0d, 0x47, 0x65, 0x74, 0x4e, 0x65, 0x77, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x20, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4e, 0x65, 0x77, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4f, 0x0a, 0x10, 0x4c, 0x69, 0x73,
	0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x23, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6a, 0x0a, 0x15, 0x4c, 0x69,
	0x73, 0x74, 0x53, 0x69, 0x64, 0x65, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x44, 0x65, 0x70, 0x6f, 0x73,
	0x69, 0x74, 0x73, 0x12, 0x27, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x53, 0x69, 0x64, 0x65, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x44, 0x65, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x77,
	0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x69, 0x64,
	0x65, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6d, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x53, 0x69, 0x64, 0x65, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74,
	0x12, 0x28, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x53, 0x69, 0x64, 0x65, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x44, 0x65, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x77, 0x61, 0x6c,
	0x6c, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x69, 0x64,
	0x65, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0xad, 0x01, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x2e, 0x77, 0x61,
	0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x42, 0x0b, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x54, 0x77, 0x6f, 0x2d, 0x4c, 0x61, 0x62, 0x73,
	0x2f, 0x73, 0x69, 0x64, 0x65, 0x73, 0x61, 0x69, 0x6c, 0x2f, 0x64, 0x72, 0x69, 0x76, 0x65, 0x63,
	0x68, 0x61, 0x69, 0x6e, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x67, 0x65, 0x6e, 0x2f,
	0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2f, 0x76, 0x31, 0x3b, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74,
	0x76, 0x31, 0xa2, 0x02, 0x03, 0x57, 0x58, 0x58, 0xaa, 0x02, 0x09, 0x57, 0x61, 0x6c, 0x6c, 0x65,
	0x74, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x09, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x5c, 0x56, 0x31,
	0xe2, 0x02, 0x15, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0a, 0x57, 0x61, 0x6c, 0x6c, 0x65,
	0x74, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_wallet_v1_wallet_proto_rawDescOnce sync.Once
	file_wallet_v1_wallet_proto_rawDescData = file_wallet_v1_wallet_proto_rawDesc
)

func file_wallet_v1_wallet_proto_rawDescGZIP() []byte {
	file_wallet_v1_wallet_proto_rawDescOnce.Do(func() {
		file_wallet_v1_wallet_proto_rawDescData = protoimpl.X.CompressGZIP(file_wallet_v1_wallet_proto_rawDescData)
	})
	return file_wallet_v1_wallet_proto_rawDescData
}

var file_wallet_v1_wallet_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_wallet_v1_wallet_proto_goTypes = []any{
	(*GetNewAddressResponse)(nil),                          // 0: wallet.v1.GetNewAddressResponse
	(*SendTransactionRequest)(nil),                         // 1: wallet.v1.SendTransactionRequest
	(*SendTransactionResponse)(nil),                        // 2: wallet.v1.SendTransactionResponse
	(*GetBalanceResponse)(nil),                             // 3: wallet.v1.GetBalanceResponse
	(*ListTransactionsResponse)(nil),                       // 4: wallet.v1.ListTransactionsResponse
	(*Confirmation)(nil),                                   // 5: wallet.v1.Confirmation
	(*Transaction)(nil),                                    // 6: wallet.v1.Transaction
	(*ListSidechainDepositsRequest)(nil),                   // 7: wallet.v1.ListSidechainDepositsRequest
	(*ListSidechainDepositsResponse)(nil),                  // 8: wallet.v1.ListSidechainDepositsResponse
	(*CreateSidechainDepositRequest)(nil),                  // 9: wallet.v1.CreateSidechainDepositRequest
	(*CreateSidechainDepositResponse)(nil),                 // 10: wallet.v1.CreateSidechainDepositResponse
	nil,                                                    // 11: wallet.v1.SendTransactionRequest.DestinationsEntry
	(*ListSidechainDepositsResponse_SidechainDeposit)(nil), // 12: wallet.v1.ListSidechainDepositsResponse.SidechainDeposit
	(*timestamppb.Timestamp)(nil),                          // 13: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),                                  // 14: google.protobuf.Empty
}
var file_wallet_v1_wallet_proto_depIdxs = []int32{
	11, // 0: wallet.v1.SendTransactionRequest.destinations:type_name -> wallet.v1.SendTransactionRequest.DestinationsEntry
	6,  // 1: wallet.v1.ListTransactionsResponse.transactions:type_name -> wallet.v1.Transaction
	13, // 2: wallet.v1.Confirmation.timestamp:type_name -> google.protobuf.Timestamp
	5,  // 3: wallet.v1.Transaction.confirmation_time:type_name -> wallet.v1.Confirmation
	12, // 4: wallet.v1.ListSidechainDepositsResponse.deposits:type_name -> wallet.v1.ListSidechainDepositsResponse.SidechainDeposit
	1,  // 5: wallet.v1.WalletService.SendTransaction:input_type -> wallet.v1.SendTransactionRequest
	14, // 6: wallet.v1.WalletService.GetBalance:input_type -> google.protobuf.Empty
	14, // 7: wallet.v1.WalletService.GetNewAddress:input_type -> google.protobuf.Empty
	14, // 8: wallet.v1.WalletService.ListTransactions:input_type -> google.protobuf.Empty
	7,  // 9: wallet.v1.WalletService.ListSidechainDeposits:input_type -> wallet.v1.ListSidechainDepositsRequest
	9,  // 10: wallet.v1.WalletService.CreateSidechainDeposit:input_type -> wallet.v1.CreateSidechainDepositRequest
	2,  // 11: wallet.v1.WalletService.SendTransaction:output_type -> wallet.v1.SendTransactionResponse
	3,  // 12: wallet.v1.WalletService.GetBalance:output_type -> wallet.v1.GetBalanceResponse
	0,  // 13: wallet.v1.WalletService.GetNewAddress:output_type -> wallet.v1.GetNewAddressResponse
	4,  // 14: wallet.v1.WalletService.ListTransactions:output_type -> wallet.v1.ListTransactionsResponse
	8,  // 15: wallet.v1.WalletService.ListSidechainDeposits:output_type -> wallet.v1.ListSidechainDepositsResponse
	10, // 16: wallet.v1.WalletService.CreateSidechainDeposit:output_type -> wallet.v1.CreateSidechainDepositResponse
	11, // [11:17] is the sub-list for method output_type
	5,  // [5:11] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_wallet_v1_wallet_proto_init() }
func file_wallet_v1_wallet_proto_init() {
	if File_wallet_v1_wallet_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_wallet_v1_wallet_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_wallet_v1_wallet_proto_goTypes,
		DependencyIndexes: file_wallet_v1_wallet_proto_depIdxs,
		MessageInfos:      file_wallet_v1_wallet_proto_msgTypes,
	}.Build()
	File_wallet_v1_wallet_proto = out.File
	file_wallet_v1_wallet_proto_rawDesc = nil
	file_wallet_v1_wallet_proto_goTypes = nil
	file_wallet_v1_wallet_proto_depIdxs = nil
}
