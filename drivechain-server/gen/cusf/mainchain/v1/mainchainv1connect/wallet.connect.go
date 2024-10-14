// CUSF mainchain wallet service

// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: cusf/mainchain/v1/wallet.proto

package mainchainv1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/LayerTwo-Labs/sidesail/drivechain-server/gen/cusf/mainchain/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// WalletServiceName is the fully-qualified name of the WalletService service.
	WalletServiceName = "cusf.mainchain.v1.WalletService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// WalletServiceBroadcastWithdrawalBundleProcedure is the fully-qualified name of the
	// WalletService's BroadcastWithdrawalBundle RPC.
	WalletServiceBroadcastWithdrawalBundleProcedure = "/cusf.mainchain.v1.WalletService/BroadcastWithdrawalBundle"
	// WalletServiceCreateBmmCriticalDataTransactionProcedure is the fully-qualified name of the
	// WalletService's CreateBmmCriticalDataTransaction RPC.
	WalletServiceCreateBmmCriticalDataTransactionProcedure = "/cusf.mainchain.v1.WalletService/CreateBmmCriticalDataTransaction"
	// WalletServiceCreateDepositTransactionProcedure is the fully-qualified name of the WalletService's
	// CreateDepositTransaction RPC.
	WalletServiceCreateDepositTransactionProcedure = "/cusf.mainchain.v1.WalletService/CreateDepositTransaction"
	// WalletServiceCreateNewAddressProcedure is the fully-qualified name of the WalletService's
	// CreateNewAddress RPC.
	WalletServiceCreateNewAddressProcedure = "/cusf.mainchain.v1.WalletService/CreateNewAddress"
	// WalletServiceGenerateBlocksProcedure is the fully-qualified name of the WalletService's
	// GenerateBlocks RPC.
	WalletServiceGenerateBlocksProcedure = "/cusf.mainchain.v1.WalletService/GenerateBlocks"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	walletServiceServiceDescriptor                                = v1.File_cusf_mainchain_v1_wallet_proto.Services().ByName("WalletService")
	walletServiceBroadcastWithdrawalBundleMethodDescriptor        = walletServiceServiceDescriptor.Methods().ByName("BroadcastWithdrawalBundle")
	walletServiceCreateBmmCriticalDataTransactionMethodDescriptor = walletServiceServiceDescriptor.Methods().ByName("CreateBmmCriticalDataTransaction")
	walletServiceCreateDepositTransactionMethodDescriptor         = walletServiceServiceDescriptor.Methods().ByName("CreateDepositTransaction")
	walletServiceCreateNewAddressMethodDescriptor                 = walletServiceServiceDescriptor.Methods().ByName("CreateNewAddress")
	walletServiceGenerateBlocksMethodDescriptor                   = walletServiceServiceDescriptor.Methods().ByName("GenerateBlocks")
)

// WalletServiceClient is a client for the cusf.mainchain.v1.WalletService service.
type WalletServiceClient interface {
	BroadcastWithdrawalBundle(context.Context, *connect.Request[v1.BroadcastWithdrawalBundleRequest]) (*connect.Response[v1.BroadcastWithdrawalBundleResponse], error)
	CreateBmmCriticalDataTransaction(context.Context, *connect.Request[v1.CreateBmmCriticalDataTransactionRequest]) (*connect.Response[v1.CreateBmmCriticalDataTransactionResponse], error)
	CreateDepositTransaction(context.Context, *connect.Request[v1.CreateDepositTransactionRequest]) (*connect.Response[v1.CreateDepositTransactionResponse], error)
	CreateNewAddress(context.Context, *connect.Request[v1.CreateNewAddressRequest]) (*connect.Response[v1.CreateNewAddressResponse], error)
	// Regtest only
	GenerateBlocks(context.Context, *connect.Request[v1.GenerateBlocksRequest]) (*connect.Response[v1.GenerateBlocksResponse], error)
}

// NewWalletServiceClient constructs a client for the cusf.mainchain.v1.WalletService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewWalletServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) WalletServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &walletServiceClient{
		broadcastWithdrawalBundle: connect.NewClient[v1.BroadcastWithdrawalBundleRequest, v1.BroadcastWithdrawalBundleResponse](
			httpClient,
			baseURL+WalletServiceBroadcastWithdrawalBundleProcedure,
			connect.WithSchema(walletServiceBroadcastWithdrawalBundleMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		createBmmCriticalDataTransaction: connect.NewClient[v1.CreateBmmCriticalDataTransactionRequest, v1.CreateBmmCriticalDataTransactionResponse](
			httpClient,
			baseURL+WalletServiceCreateBmmCriticalDataTransactionProcedure,
			connect.WithSchema(walletServiceCreateBmmCriticalDataTransactionMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		createDepositTransaction: connect.NewClient[v1.CreateDepositTransactionRequest, v1.CreateDepositTransactionResponse](
			httpClient,
			baseURL+WalletServiceCreateDepositTransactionProcedure,
			connect.WithSchema(walletServiceCreateDepositTransactionMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		createNewAddress: connect.NewClient[v1.CreateNewAddressRequest, v1.CreateNewAddressResponse](
			httpClient,
			baseURL+WalletServiceCreateNewAddressProcedure,
			connect.WithSchema(walletServiceCreateNewAddressMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		generateBlocks: connect.NewClient[v1.GenerateBlocksRequest, v1.GenerateBlocksResponse](
			httpClient,
			baseURL+WalletServiceGenerateBlocksProcedure,
			connect.WithSchema(walletServiceGenerateBlocksMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// walletServiceClient implements WalletServiceClient.
type walletServiceClient struct {
	broadcastWithdrawalBundle        *connect.Client[v1.BroadcastWithdrawalBundleRequest, v1.BroadcastWithdrawalBundleResponse]
	createBmmCriticalDataTransaction *connect.Client[v1.CreateBmmCriticalDataTransactionRequest, v1.CreateBmmCriticalDataTransactionResponse]
	createDepositTransaction         *connect.Client[v1.CreateDepositTransactionRequest, v1.CreateDepositTransactionResponse]
	createNewAddress                 *connect.Client[v1.CreateNewAddressRequest, v1.CreateNewAddressResponse]
	generateBlocks                   *connect.Client[v1.GenerateBlocksRequest, v1.GenerateBlocksResponse]
}

// BroadcastWithdrawalBundle calls cusf.mainchain.v1.WalletService.BroadcastWithdrawalBundle.
func (c *walletServiceClient) BroadcastWithdrawalBundle(ctx context.Context, req *connect.Request[v1.BroadcastWithdrawalBundleRequest]) (*connect.Response[v1.BroadcastWithdrawalBundleResponse], error) {
	return c.broadcastWithdrawalBundle.CallUnary(ctx, req)
}

// CreateBmmCriticalDataTransaction calls
// cusf.mainchain.v1.WalletService.CreateBmmCriticalDataTransaction.
func (c *walletServiceClient) CreateBmmCriticalDataTransaction(ctx context.Context, req *connect.Request[v1.CreateBmmCriticalDataTransactionRequest]) (*connect.Response[v1.CreateBmmCriticalDataTransactionResponse], error) {
	return c.createBmmCriticalDataTransaction.CallUnary(ctx, req)
}

// CreateDepositTransaction calls cusf.mainchain.v1.WalletService.CreateDepositTransaction.
func (c *walletServiceClient) CreateDepositTransaction(ctx context.Context, req *connect.Request[v1.CreateDepositTransactionRequest]) (*connect.Response[v1.CreateDepositTransactionResponse], error) {
	return c.createDepositTransaction.CallUnary(ctx, req)
}

// CreateNewAddress calls cusf.mainchain.v1.WalletService.CreateNewAddress.
func (c *walletServiceClient) CreateNewAddress(ctx context.Context, req *connect.Request[v1.CreateNewAddressRequest]) (*connect.Response[v1.CreateNewAddressResponse], error) {
	return c.createNewAddress.CallUnary(ctx, req)
}

// GenerateBlocks calls cusf.mainchain.v1.WalletService.GenerateBlocks.
func (c *walletServiceClient) GenerateBlocks(ctx context.Context, req *connect.Request[v1.GenerateBlocksRequest]) (*connect.Response[v1.GenerateBlocksResponse], error) {
	return c.generateBlocks.CallUnary(ctx, req)
}

// WalletServiceHandler is an implementation of the cusf.mainchain.v1.WalletService service.
type WalletServiceHandler interface {
	BroadcastWithdrawalBundle(context.Context, *connect.Request[v1.BroadcastWithdrawalBundleRequest]) (*connect.Response[v1.BroadcastWithdrawalBundleResponse], error)
	CreateBmmCriticalDataTransaction(context.Context, *connect.Request[v1.CreateBmmCriticalDataTransactionRequest]) (*connect.Response[v1.CreateBmmCriticalDataTransactionResponse], error)
	CreateDepositTransaction(context.Context, *connect.Request[v1.CreateDepositTransactionRequest]) (*connect.Response[v1.CreateDepositTransactionResponse], error)
	CreateNewAddress(context.Context, *connect.Request[v1.CreateNewAddressRequest]) (*connect.Response[v1.CreateNewAddressResponse], error)
	// Regtest only
	GenerateBlocks(context.Context, *connect.Request[v1.GenerateBlocksRequest]) (*connect.Response[v1.GenerateBlocksResponse], error)
}

// NewWalletServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewWalletServiceHandler(svc WalletServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	walletServiceBroadcastWithdrawalBundleHandler := connect.NewUnaryHandler(
		WalletServiceBroadcastWithdrawalBundleProcedure,
		svc.BroadcastWithdrawalBundle,
		connect.WithSchema(walletServiceBroadcastWithdrawalBundleMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	walletServiceCreateBmmCriticalDataTransactionHandler := connect.NewUnaryHandler(
		WalletServiceCreateBmmCriticalDataTransactionProcedure,
		svc.CreateBmmCriticalDataTransaction,
		connect.WithSchema(walletServiceCreateBmmCriticalDataTransactionMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	walletServiceCreateDepositTransactionHandler := connect.NewUnaryHandler(
		WalletServiceCreateDepositTransactionProcedure,
		svc.CreateDepositTransaction,
		connect.WithSchema(walletServiceCreateDepositTransactionMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	walletServiceCreateNewAddressHandler := connect.NewUnaryHandler(
		WalletServiceCreateNewAddressProcedure,
		svc.CreateNewAddress,
		connect.WithSchema(walletServiceCreateNewAddressMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	walletServiceGenerateBlocksHandler := connect.NewUnaryHandler(
		WalletServiceGenerateBlocksProcedure,
		svc.GenerateBlocks,
		connect.WithSchema(walletServiceGenerateBlocksMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/cusf.mainchain.v1.WalletService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case WalletServiceBroadcastWithdrawalBundleProcedure:
			walletServiceBroadcastWithdrawalBundleHandler.ServeHTTP(w, r)
		case WalletServiceCreateBmmCriticalDataTransactionProcedure:
			walletServiceCreateBmmCriticalDataTransactionHandler.ServeHTTP(w, r)
		case WalletServiceCreateDepositTransactionProcedure:
			walletServiceCreateDepositTransactionHandler.ServeHTTP(w, r)
		case WalletServiceCreateNewAddressProcedure:
			walletServiceCreateNewAddressHandler.ServeHTTP(w, r)
		case WalletServiceGenerateBlocksProcedure:
			walletServiceGenerateBlocksHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedWalletServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedWalletServiceHandler struct{}

func (UnimplementedWalletServiceHandler) BroadcastWithdrawalBundle(context.Context, *connect.Request[v1.BroadcastWithdrawalBundleRequest]) (*connect.Response[v1.BroadcastWithdrawalBundleResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("cusf.mainchain.v1.WalletService.BroadcastWithdrawalBundle is not implemented"))
}

func (UnimplementedWalletServiceHandler) CreateBmmCriticalDataTransaction(context.Context, *connect.Request[v1.CreateBmmCriticalDataTransactionRequest]) (*connect.Response[v1.CreateBmmCriticalDataTransactionResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("cusf.mainchain.v1.WalletService.CreateBmmCriticalDataTransaction is not implemented"))
}

func (UnimplementedWalletServiceHandler) CreateDepositTransaction(context.Context, *connect.Request[v1.CreateDepositTransactionRequest]) (*connect.Response[v1.CreateDepositTransactionResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("cusf.mainchain.v1.WalletService.CreateDepositTransaction is not implemented"))
}

func (UnimplementedWalletServiceHandler) CreateNewAddress(context.Context, *connect.Request[v1.CreateNewAddressRequest]) (*connect.Response[v1.CreateNewAddressResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("cusf.mainchain.v1.WalletService.CreateNewAddress is not implemented"))
}

func (UnimplementedWalletServiceHandler) GenerateBlocks(context.Context, *connect.Request[v1.GenerateBlocksRequest]) (*connect.Response[v1.GenerateBlocksResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("cusf.mainchain.v1.WalletService.GenerateBlocks is not implemented"))
}
