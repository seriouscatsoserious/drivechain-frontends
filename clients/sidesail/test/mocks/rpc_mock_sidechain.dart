import 'package:sail_ui/sail_ui.dart';
import 'package:sidesail/rpc/rpc_sidechain.dart';
import 'package:sidesail/rpc/rpc_zcash.dart';

class MockSidechainRPC extends SidechainRPC {
  MockSidechainRPC()
      : super(
          conf: NodeConnectionSettings('mock town', 'mock mock', 1337, '', '', true),
          chain: TestSidechain(),
        );

  @override
  Future<(double, double)> getBalance() async {
    return (1.12345678, 2.24680);
  }

  @override
  Future callRAW(String method, [dynamic params]) async {
    return;
  }

  @override
  bool get connected => true;

  @override
  Future<List<CoreTransaction>> listTransactions() async {
    return List.empty();
  }

  @override
  List<String> binaryArgs(
    NodeConnectionSettings mainchainConf,
  ) {
    return List.empty();
  }

  @override
  Future<String> getDepositAddress() async {
    return formatDepositAddress('3CUZ683astRsmACdRKyx7eFb1y9yvMRzGi', 0);
  }

  @override
  Future<String> mainSend(String address, double amount, double sidechainFee, double mainchainFee) async {
    return 'txiddeadbeef';
  }

  @override
  Future<double> sideEstimateFee() async {
    return zcashFee;
  }

  @override
  Future<int> ping() async {
    return 69;
  }

  @override
  Future<String> sideSend(String address, double amount, bool subtractFeeFromAmount) async {
    return 'deadbeefdeadbeefdeadbeef';
  }

  @override
  Future<void> stop() async {
    return;
  }

  @override
  Future<String> getSideAddress() async {
    return 'taddress';
  }

  @override
  Future<BlockchainInfo> getBlockchainInfo() async {
    return BlockchainInfo(initialBlockDownload: false, blockHeight: 70);
  }
}
