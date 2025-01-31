import 'dart:async';
import 'dart:io';

import 'package:auto_route/auto_route.dart';
import 'package:bitwindow/providers/balance_provider.dart';
import 'package:bitwindow/routing/router.dart';
import 'package:flutter/material.dart';
import 'package:get_it/get_it.dart';
import 'package:logger/logger.dart';
import 'package:sail_ui/config/binaries.dart';
import 'package:sail_ui/gen/bitcoind/v1/bitcoind.pbgrpc.dart';
import 'package:sail_ui/rpcs/bitwindow_api.dart';
import 'package:sail_ui/rpcs/enforcer_rpc.dart';
import 'package:sail_ui/rpcs/mainchain_rpc.dart';
import 'package:sail_ui/sail_ui.dart';
import 'package:sail_ui/widgets/nav/top_nav.dart';
import 'package:stacked/stacked.dart';

@RoutePage()
class RootPage extends StatelessWidget {
  const RootPage({super.key});

  @override
  Widget build(BuildContext context) {
    return AutoTabsRouter.tabBar(
      animatePageTransition: false,
      routes: const [
        OverviewRoute(),
        WalletRoute(),
        SidechainsRoute(),
      ],
      builder: (context, child, controller) {
        final theme = SailTheme.of(context);

        return Scaffold(
          backgroundColor: theme.colors.background,
          appBar: PreferredSize(
            preferredSize: const Size.fromHeight(80),
            child: DecoratedBox(
              decoration: BoxDecoration(
                color: theme.colors.background,
              ),
              child: Builder(
                builder: (context) {
                  final tabsRouter = AutoTabsRouter.of(context);
                  return Row(
                    children: [
                      QtTab(
                        icon: SailSVGAsset.iconHome,
                        label: 'Overview',
                        active: tabsRouter.activeIndex == 0,
                        onTap: () => tabsRouter.setActiveIndex(0),
                      ),
                      QtTab(
                        icon: SailSVGAsset.iconSend,
                        label: 'Send/Receive',
                        active: tabsRouter.activeIndex == 1,
                        onTap: () => tabsRouter.setActiveIndex(1),
                      ),
                      QtTab(
                        icon: SailSVGAsset.iconSidechains,
                        label: 'Sidechains',
                        active: tabsRouter.activeIndex == 2,
                        onTap: () => tabsRouter.setActiveIndex(2),
                        end: true,
                      ),
                      Expanded(child: Container()),
                      const ToggleThemeButton(),
                    ],
                  );
                },
              ),
            ),
          ),
          body: Column(
            children: [
              const Divider(
                height: 1,
                thickness: 1,
                color: Colors.grey,
              ),
              Expanded(child: child),
              const StatusBar(),
            ],
          ),
        );
      },
    );
  }
}

class StatusBar extends StatefulWidget {
  const StatusBar({super.key});

  @override
  State<StatusBar> createState() => _StatusBarState();
}

class _StatusBarState extends State<StatusBar> {
  BlockchainProvider get blockchainProvider => GetIt.I.get<BlockchainProvider>();
  BalanceProvider get balanceProvider => GetIt.I.get<BalanceProvider>();
  late Timer _timer;

  @override
  void initState() {
    super.initState();
    _timer = Timer.periodic(const Duration(seconds: 1), (_) => setState(() {}));
    balanceProvider.addListener(setstate);
  }

  void setstate() {
    setState(() {});
  }

  String _getTimeSinceLastBlock() {
    if (blockchainProvider.lastBlockAt == null) {
      return 'Unknown';
    }

    final now = DateTime.now();
    final lastBlockTime = blockchainProvider.lastBlockAt!.toDateTime().toLocal();
    final difference = now.difference(lastBlockTime);

    if (difference.inDays > 0) {
      return '${formatTimeDifference(difference.inDays, 'day')} ago';
    } else if (difference.inHours > 0) {
      return '${formatTimeDifference(difference.inHours, 'hour')} ago';
    } else if (difference.inMinutes > 0) {
      return '${formatTimeDifference(difference.inMinutes, 'minute')} ago';
    } else {
      return '${formatTimeDifference(difference.inSeconds, 'second')} ago';
    }
  }

  @override
  Widget build(BuildContext context) {
    return ViewModelBuilder.reactive(
      viewModelBuilder: () => BottomNavViewModel(),
      fireOnViewModelReadyOnce: true,
      onViewModelReady: (model) {
        if (!model.mainchainConnected || !model.enforcerConnected || !model.serverConnected) {
          WidgetsBinding.instance.addPostFrameCallback((_) {
            displayConnectionStatusDialog(context);
          });
        }
      },
      builder: ((context, model, child) {
        return SizedBox(
          height: 36,
          child: DecoratedBox(
            decoration: const BoxDecoration(
              border: Border(
                top: BorderSide(
                  color: Colors.grey,
                ),
              ),
            ),
            child: SailRow(
              mainAxisAlignment: MainAxisAlignment.end,
              spacing: SailStyleValues.padding04,
              leadingSpacing: true,
              trailingSpacing: true,
              children: [
                SailSpacing(SailStyleValues.padding04),
                Tooltip(
                  message: 'Confirmed balance',
                  child: SailRow(
                    spacing: SailStyleValues.padding08,
                    children: [
                      SailSVG.icon(
                        SailSVGAsset.iconCoins,
                        color: SailColorScheme.green,
                        width: SailStyleValues.iconSizeSecondary,
                        height: SailStyleValues.iconSizeSecondary,
                      ),
                      SailText.secondary12(formatBitcoin(satoshiToBTC(balanceProvider.balance), symbol: 'BTC')),
                    ],
                  ),
                ),
                const DividerDot(),
                Tooltip(
                  message: 'Unconfirmed balance',
                  child: SailRow(
                    spacing: SailStyleValues.padding08,
                    children: [
                      SailSVG.icon(
                        SailSVGAsset.iconCoins,
                        width: SailStyleValues.iconSizeSecondary,
                        height: SailStyleValues.iconSizeSecondary,
                      ),
                      SailText.secondary12(formatBitcoin(satoshiToBTC(balanceProvider.pendingBalance), symbol: 'BTC')),
                    ],
                  ),
                ),
                Expanded(child: Container()),
                SailRawButton(
                  onPressed: () => displayConnectionStatusDialog(context),
                  disabled: false,
                  loading: false,
                  child: Tooltip(
                    message: model.connectionStatus,
                    child: SailSVG.fromAsset(
                      SailSVGAsset.iconConnectionStatus,
                      color: model.connectionColor,
                    ),
                  ),
                ),
                Tooltip(
                  message: blockchainProvider.recentBlocks.firstOrNull?.toPretty() ?? '',
                  child: SailText.primary12('Last block: ${_getTimeSinceLastBlock()}'),
                ),
                const DividerDot(),
                if (blockchainProvider.blockchainInfo.initialBlockDownload &&
                    blockchainProvider.blockchainInfo.blocks != blockchainProvider.blockchainInfo.headers)
                  Tooltip(
                    message:
                        'Current height: ${blockchainProvider.blockchainInfo.blocks}\nHeader height: ${blockchainProvider.blockchainInfo.headers}',
                    child: SailText.primary12(
                      'Downloading blocks (${blockchainProvider.verificationProgress}%)',
                    ),
                  ),
                if (blockchainProvider.blockchainInfo.initialBlockDownload &&
                    blockchainProvider.blockchainInfo.blocks != blockchainProvider.blockchainInfo.headers)
                  const DividerDot(),
                SailText.primary12(
                  '${formatWithThousandSpacers(blockchainProvider.blockchainInfo.blocks)} blocks',
                ),
                const DividerDot(),
                Tooltip(
                  message: blockchainProvider.peers.map((e) => 'Peer id=${e.id} addr=${e.addr}').join('\n'),
                  child: SailText.primary12(
                    formatTimeDifference(blockchainProvider.peers.length, 'peer'),
                  ),
                ),
                SailSpacing(SailStyleValues.padding04),
              ],
            ),
          ),
        );
      }),
    );
  }

  Future<void> displayConnectionStatusDialog(
    BuildContext context,
  ) async {
    await widgetDialog(
      context: context,
      title: 'Daemon Status',
      subtitle:
          "You can use BitWindow without the enforcer, but it's not that interesting because the wallet does not work.",
      child: ViewModelBuilder.reactive(
        viewModelBuilder: () => BottomNavViewModel(),
        builder: ((context, model, child) {
          return SailColumn(
            spacing: SailStyleValues.padding20,
            mainAxisAlignment: MainAxisAlignment.start,
            crossAxisAlignment: CrossAxisAlignment.start,
            mainAxisSize: MainAxisSize.min,
            children: [
              const SailSpacing(SailStyleValues.padding08),
              SailColumn(
                spacing: SailStyleValues.padding12,
                crossAxisAlignment: CrossAxisAlignment.start,
                children: [
                  DaemonConnectionCard(
                    connection: model.mainchain,
                    restartDaemon: () => model.initMainchainBinary(context),
                    infoMessage: null,
                    navigateToLogs: model.navigateToLogs,
                  ),
                  DaemonConnectionCard(
                    connection: model.enforcer,
                    infoMessage: model.mainchainInitializing
                        ? 'Waiting for mainchain to finish init'
                        : model.inIBD
                            ? 'Waiting for L1 initial block download to complete...'
                            : null,
                    restartDaemon: () => model.initEnforcerBinary(context),
                    navigateToLogs: model.navigateToLogs,
                  ),
                  DaemonConnectionCard(
                    connection: model.server,
                    infoMessage:
                        !model.serverConnected && model.enforcerInitializing ? 'Waiting for enforcer to start' : null,
                    restartDaemon: () => model.initServerBinary(context),
                    navigateToLogs: model.navigateToLogs,
                  ),
                ],
              ),
            ],
          );
        }),
      ),
    );
  }

  @override
  void dispose() {
    _timer.cancel();
    balanceProvider.removeListener(setstate);
    super.dispose();
  }
}

String formatTimeDifference(int value, String unit) {
  return '$value $unit${value == 1 ? '' : 's'}';
}

extension on Block {
  String toPretty() {
    return 'Block $blockHeight\nBlockTime=${blockTime.toDateTime().format()}\nHash=$hash';
  }
}

class BottomNavViewModel extends BaseViewModel {
  final log = Logger(level: Level.debug);
  AppRouter get router => GetIt.I.get<AppRouter>();

  MainchainRPC get mainchain => GetIt.I.get<MainchainRPC>();
  EnforcerRPC get enforcer => GetIt.I.get<EnforcerRPC>();
  BitwindowRPC get server => GetIt.I.get<BitwindowRPC>();
  BalanceProvider get _balanceProvider => GetIt.I.get<BalanceProvider>();

  int get balance => _balanceProvider.balance;
  int get pendingBalance => _balanceProvider.pendingBalance;

  bool get mainchainConnected => mainchain.connected;
  bool get enforcerConnected => enforcer.connected;
  bool get serverConnected => server.connected;

  bool get mainchainInitializing => mainchain.initializingBinary;
  bool get enforcerInitializing => enforcer.initializingBinary;
  bool get serverInitializing => server.initializingBinary;
  bool get inIBD => mainchain.inIBD;

  int get blockCount => mainchain.blockCount;

  String? get mainchainError => mainchain.connectionError;
  String? get enforcerError => enforcer.connectionError;
  String? get serverError => server.connectionError;

  BottomNavViewModel() {
    mainchain.addListener(notifyListeners);
    enforcer.addListener(notifyListeners);
    server.addListener(notifyListeners);
  }

  Future<void> initMainchainBinary(BuildContext context) async {
    return mainchain.initBinary(
      context,
    );
  }

  Future<void> deleteMainchainBlocks(BuildContext context) async {
    try {
      final datadir = mainchain.chain.datadir();
      final signetDir = filePath([datadir, 'signet']);

      // List of files/directories to delete
      final toDelete = [
        'blocks',
        'chainstate',
        'indexes',
        'banlist.json',
        'bitcoind.pid',
        'debug.log',
        'fee_estimates.dat',
        'mempool.dat',
        'peers.dat',
        'settings.json',
      ];

      for (final item in toDelete) {
        final path = filePath([signetDir, item]);
        final entity = File(path);
        if (await entity.exists()) {
          await entity.delete(recursive: true);
        } else {
          final dir = Directory(path);
          if (await dir.exists()) {
            await dir.delete(recursive: true);
          }
        }
      }

      if (context.mounted) {
        ScaffoldMessenger.of(context).showSnackBar(
          const SnackBar(content: Text('Successfully deleted blockchain data')),
        );
      }
    } catch (e) {
      if (context.mounted) {
        ScaffoldMessenger.of(context).showSnackBar(
          SnackBar(content: Text('Failed to delete blockchain data: $e')),
        );
      }
    }
  }

  Future<void> initEnforcerBinary(
    BuildContext context,
  ) async {
    await enforcer.initBinary(context);
  }

  Future<void> initServerBinary(
    BuildContext context,
  ) async {
    await server.initBinary(context);
  }

  Color get connectionColor {
    if (mainchainConnected && enforcerConnected && serverConnected) {
      return SailColorScheme.green;
    }
    if (!mainchainConnected || !serverConnected) {
      return SailColorScheme.red;
    }

    // only the enforcer is not running
    return SailColorScheme.orange;
  }

  String get connectionStatus {
    if (mainchainConnected && enforcerConnected && serverConnected) {
      return 'All binaries connected';
    }

    List<String> errors = [];
    if (!mainchainConnected && mainchainError != null) {
      errors.add('Mainchain: ${mainchainError!}');
    }
    if (!enforcerConnected && enforcerError != null) {
      errors.add('Enforcer: ${enforcerError!}');
    }
    if (!serverConnected && serverError != null) {
      errors.add('Server: ${serverError!}');
    }

    if (errors.length == 1) {
      return errors.first;
    } else {
      return errors.join('\n');
    }
  }

  void navigateToLogs(String name, String logPath) {
    router.push(
      SailLogRoute(
        name: name,
        logPath: logPath,
      ),
    );
  }

  @override
  void dispose() {
    super.dispose();
    mainchain.removeListener(notifyListeners);
    enforcer.removeListener(notifyListeners);
    server.removeListener(notifyListeners);
  }
}

class Separator extends StatelessWidget {
  final Widget child;
  final bool right;

  const Separator({
    super.key,
    required this.child,
    this.right = false,
  });

  @override
  Widget build(BuildContext context) {
    return Container(
      padding: const EdgeInsets.symmetric(
        horizontal: 4.0,
        vertical: 2.0,
      ),
      decoration: BoxDecoration(
        border: Border(
          left: right ? BorderSide.none : BorderSide(color: Colors.grey),
          right: right ? BorderSide(color: Colors.grey) : BorderSide.none,
        ),
      ),
      child: child,
    );
  }
}

class DividerDot extends StatelessWidget {
  const DividerDot({super.key});

  @override
  Widget build(BuildContext context) {
    return SailSVG.fromAsset(SailSVGAsset.dividerDot);
  }
}
