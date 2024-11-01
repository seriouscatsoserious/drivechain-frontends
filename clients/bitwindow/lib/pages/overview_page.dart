import 'package:auto_route/auto_route.dart';
import 'package:bitwindow/gen/bitcoind/v1/bitcoind.pbgrpc.dart';
import 'package:bitwindow/providers/balance_provider.dart';
import 'package:bitwindow/providers/blockchain_provider.dart';
import 'package:bitwindow/widgets/error_container.dart';
import 'package:bitwindow/widgets/qt_icon_button.dart';
import 'package:flutter/foundation.dart';
import 'package:flutter/material.dart';
import 'package:get_it/get_it.dart';
import 'package:money2/money2.dart';
import 'package:sail_ui/sail_ui.dart';
import 'package:sail_ui/widgets/containers/qt_page.dart';
import 'package:stacked/stacked.dart';

@RoutePage()
class OverviewPage extends StatelessWidget {
  const OverviewPage({super.key});

  @override
  Widget build(BuildContext context) {
    return QtPage(
      child: SingleChildScrollView(
        child: SailColumn(
          spacing: SailStyleValues.padding16,
          crossAxisAlignment: CrossAxisAlignment.start,
          mainAxisSize: MainAxisSize.min,
          children: [
            ExperimentalBanner(),
            SailColumn(
              spacing: SailStyleValues.padding16,
              children: [
                CoinNewsView(),
                TransactionsView(),
              ],
            ),
          ],
        ),
      ),
    );
  }
}

class ExperimentalBanner extends StatelessWidget {
  const ExperimentalBanner({super.key});

  @override
  Widget build(BuildContext context) {
    return Container(
      decoration: BoxDecoration(
        borderRadius: BorderRadius.circular(6.0),
        gradient: LinearGradient(
          begin: Alignment.bottomCenter,
          end: Alignment.topCenter,
          colors: [
            SailTheme.of(context).colors.orange.withOpacity(0.25),
            SailTheme.of(context).colors.orangeLight.withOpacity(0.25),
          ],
        ),
        border: Border.all(
          color: SailTheme.of(context).colors.orange,
          width: 1.0,
        ),
      ),
      width: double.infinity,
      padding: const EdgeInsets.all(16.0),
      child: Row(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          SailSVG.fromAsset(SailSVGAsset.iconWarning),
          const SailSpacing(SailStyleValues.padding08),
          Expanded(
            child: Column(
              crossAxisAlignment: CrossAxisAlignment.start,
              mainAxisAlignment: MainAxisAlignment.start,
              children: [
                SailText.primary15('Warning', bold: true),
                SailText.primary15(
                  'This is experimental sidechain software. Use at your own risk!',
                ),
              ],
            ),
          ),
        ],
      ),
    );
  }
}

class TransactionsView extends StatelessWidget {
  const TransactionsView({super.key});

  @override
  Widget build(BuildContext context) {
    return ViewModelBuilder<BalancesViewModel>.reactive(
      viewModelBuilder: () => BalancesViewModel(),
      builder: (context, model, child) {
        if (model.hasErrorForKey('blockchain')) {
          return ErrorContainer(
            error: model.error('blockchain').toString(),
          );
        }
        return Row(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            Flexible(
              child: Column(
                crossAxisAlignment: CrossAxisAlignment.start,
                children: [
                  SailRawCard(
                    bottomPadding: false,
                    title: 'Latest Transactions',
                    child: SizedBox(
                      height: 300,
                      child: LatestTransactionTable(
                        entries: model.unconfirmedTransactions,
                      ),
                    ),
                  ),
                ],
              ),
            ),
            const SizedBox(width: SailStyleValues.padding16), // Add some space between the two tables
            Flexible(
              child: Column(
                crossAxisAlignment: CrossAxisAlignment.start,
                children: [
                  SailRawCard(
                    title: 'Latest blocks',
                    bottomPadding: false,
                    child: SizedBox(
                      height: 300,
                      child: LatestBlocksTable(
                        blocks: model.recentBlocks,
                      ),
                    ),
                  ),
                ],
              ),
            ),
          ],
        );
      },
    );
  }
}

class BalancesViewModel extends BaseViewModel {
  final BlockchainProvider blockchainProvider = GetIt.I.get<BlockchainProvider>();
  final BalanceProvider balanceProvider = GetIt.I.get<BalanceProvider>();

  BalancesViewModel() {
    balanceProvider.addListener(notifyListeners);
    blockchainProvider.addListener(notifyListeners);
  }

  void errorListener() {
    if (balanceProvider.error != null) {
      setErrorForObject('balance', balanceProvider.error);
    }
    if (blockchainProvider.error != null) {
      setErrorForObject('blockchain', blockchainProvider.error);
    }
  }

  int get confirmedBalance => balanceProvider.balance;
  int get pendingBalance => balanceProvider.pendingBalance;
  int get totalBalance => balanceProvider.balance + balanceProvider.pendingBalance;
  double get totalBalanceUSD => (satoshiToBTC(totalBalance) * 50000);

  List<ListRecentBlocksResponse_RecentBlock> get recentBlocks => blockchainProvider.recentBlocks;
  List<UnconfirmedTransaction> get unconfirmedTransactions => blockchainProvider.unconfirmedTXs;
}

class QtSeparator extends StatelessWidget {
  final double width;

  const QtSeparator({
    super.key,
    this.width = double.infinity,
  });

  @override
  Widget build(BuildContext context) {
    return ConstrainedBox(
      constraints: BoxConstraints(
        maxWidth: width,
      ),
      child: Container(
        height: 1.5,
        decoration: BoxDecoration(
          gradient: LinearGradient(
            begin: Alignment.topCenter,
            end: Alignment.bottomCenter,
            stops: const [0.0, 0.35, 0.36, 1.0],
            colors: [
              Colors.grey,
              Colors.grey.withOpacity(0.3),
              Colors.white,
              Colors.white,
            ],
          ),
        ),
      ),
    );
  }
}

class BitcoinPrice extends StatelessWidget {
  final Money money;

  const BitcoinPrice({super.key, required this.money});

  @override
  Widget build(BuildContext context) {
    return Row(
      children: [
        SailText.primary13(
          '${money.format('S###,###.##')}/BTC',
        ),
        const SizedBox(width: 8.0),
        QtIconButton(
          onPressed: () {
            showSnackBar(context, 'Not implemented');
          },
          icon: Icon(Icons.settings),
        ),
      ],
    );
  }
}

class LatestTransactionTable extends StatefulWidget {
  final List<UnconfirmedTransaction> entries;

  const LatestTransactionTable({
    super.key,
    required this.entries,
  });

  @override
  State<LatestTransactionTable> createState() => _LatestTransactionTableState();
}

class _LatestTransactionTableState extends State<LatestTransactionTable> {
  String sortColumn = 'time';
  bool sortAscending = true;
  List<UnconfirmedTransaction> entries = [];

  @override
  void initState() {
    super.initState();
    entries = widget.entries;
  }

  @override
  void didChangeDependencies() {
    super.didChangeDependencies();
    if (!listEquals(entries, widget.entries)) {
      entries = widget.entries;
      onSort(sortColumn);
    }
  }

  void onSort(String column) {
    if (sortColumn == column) {
      sortAscending = !sortAscending;
    } else {
      sortColumn = column;
      sortAscending = true;
    }
    sortEntries();
    setState(() {});
  }

  void sortEntries() {
    entries.sort((a, b) {
      dynamic aValue = '';
      dynamic bValue = '';

      switch (sortColumn) {
        case 'time':
          aValue = a.time.toDateTime().millisecondsSinceEpoch;
          bValue = b.time.toDateTime().millisecondsSinceEpoch;
          break;
        case 'fee':
          aValue = a.feeSatoshi;
          bValue = b.feeSatoshi;
          break;
        case 'txid':
          aValue = a.txid;
          bValue = b.txid;
          break;
        case 'size':
          aValue = a.virtualSize;
          bValue = b.virtualSize;
          break;
      }

      return sortAscending ? aValue.compareTo(bValue) : bValue.compareTo(aValue);
    });
  }

  @override
  Widget build(BuildContext context) {
    return SailTable(
      getRowId: (index) => entries[index].txid,
      headerBuilder: (context) => [
        SailTableHeaderCell(
          name: 'Time',
          onSort: () => onSort('time'),
        ),
        SailTableHeaderCell(
          name: 'sat/vB',
          onSort: () => onSort('fee'),
        ),
        SailTableHeaderCell(
          name: 'TxID',
          onSort: () => onSort('txid'),
        ),
        SailTableHeaderCell(
          name: 'Size',
          onSort: () => onSort('size'),
        ),
      ],
      rowBuilder: (context, row, selected) {
        final entry = entries[row];
        return [
          SailTableCell(child: SailText.primary12(entry.time.toDateTime().format())),
          SailTableCell(child: SailText.primary12(entry.feeSatoshi.toString())),
          SailTableCell(child: SailText.primary12(entry.txid)),
          SailTableCell(child: SailText.primary12(entry.virtualSize.toString())),
        ];
      },
      rowCount: entries.length,
      columnWidths: const [150, 100, 200, 100],
      drawGrid: true,
      sortColumnIndex: ['time', 'fee', 'txid', 'size'].indexOf(sortColumn),
      sortAscending: sortAscending,
      onSort: (columnIndex, ascending) {
        onSort(['time', 'fee', 'txid', 'size'][columnIndex]);
      },
    );
  }
}

class LatestBlocksTable extends StatefulWidget {
  final List<ListRecentBlocksResponse_RecentBlock> blocks;

  const LatestBlocksTable({
    super.key,
    required this.blocks,
  });

  @override
  State<LatestBlocksTable> createState() => _LatestBlocksTableState();
}

class _LatestBlocksTableState extends State<LatestBlocksTable> {
  String sortColumn = 'time';
  bool sortAscending = true;
  List<ListRecentBlocksResponse_RecentBlock> blocks = [];

  @override
  void initState() {
    super.initState();
    blocks = widget.blocks;
  }

  @override
  void didChangeDependencies() {
    super.didChangeDependencies();
    if (!listEquals(blocks, widget.blocks)) {
      blocks = widget.blocks;
      onSort(sortColumn);
    }
  }

  void onSort(String column) {
    if (sortColumn == column) {
      sortAscending = !sortAscending;
    } else {
      sortColumn = column;
      sortAscending = true;
    }
    sortEntries();
    setState(() {});
  }

  void sortEntries() {
    blocks.sort((a, b) {
      dynamic aValue = '';
      dynamic bValue = '';

      switch (sortColumn) {
        case 'time':
          aValue = a.blockTime.toDateTime().millisecondsSinceEpoch;
          bValue = b.blockTime.toDateTime().millisecondsSinceEpoch;
          break;
        case 'height':
          aValue = a.blockHeight;
          bValue = b.blockHeight;
          break;
        case 'hash':
          aValue = a.hash;
          bValue = b.hash;
          break;
      }

      return sortAscending ? aValue.compareTo(bValue) : bValue.compareTo(aValue);
    });
  }

  @override
  Widget build(BuildContext context) {
    return SailTable(
      getRowId: (index) => blocks[index].hash,
      headerBuilder: (context) => [
        SailTableHeaderCell(name: 'Time'),
        SailTableHeaderCell(name: 'Height'),
        SailTableHeaderCell(name: 'Hash'),
      ],
      rowBuilder: (context, row, selected) {
        final entry = blocks[row];
        return [
          SailTableCell(child: SailText.primary12(entry.blockTime.toDateTime().format())),
          SailTableCell(child: SailText.primary12(entry.blockHeight.toString())),
          SailTableCell(child: SailText.primary12(entry.hash)),
        ];
      },
      rowCount: blocks.length,
      columnWidths: const [150, 100, 200],
      drawGrid: true,
      sortColumnIndex: ['time', 'height', 'hash'].indexOf(sortColumn),
      sortAscending: sortAscending,
      onSort: (columnIndex, ascending) {
        onSort(['time', 'height', 'hash'][columnIndex]);
      },
    );
  }
}

class CoinNewsViewModel extends ChangeNotifier {
  String leftRegion = 'US';
  String rightRegion = 'Japan';
  final List<CoinNewsEntry> leftEntries = [];
  final List<CoinNewsEntry> rightEntries = [];

  bool _sortAscending = true;
  String _sortColumn = 'date';

  void setLeftRegion(String region) {
    leftRegion = region;
    notifyListeners();
  }

  void setRightRegion(String region) {
    rightRegion = region;
    notifyListeners();
  }

  void sortEntries(String column) {
    if (_sortColumn == column) {
      _sortAscending = !_sortAscending;
    } else {
      _sortColumn = column;
      _sortAscending = true;
    }
    notifyListeners();
  }
}

class CoinNewsView extends StatelessWidget {
  const CoinNewsView({
    super.key,
  });

  @override
  Widget build(BuildContext context) {
    return ViewModelBuilder<CoinNewsViewModel>.reactive(
      viewModelBuilder: () => CoinNewsViewModel(),
      builder: (context, viewModel, child) {
        return SailRawCard(
          header: Padding(
            padding: EdgeInsets.only(bottom: SailStyleValues.padding16),
            child: SailRow(
              spacing: SailStyleValues.padding08,
              children: [
                SailText.primary15(
                  'Coin News',
                  bold: true,
                ),
                Expanded(child: Container()),
                QtButton(
                  label: 'Broadcast News',
                  onPressed: () => displayBroadcastNewsDialog(context),
                  size: ButtonSize.small,
                ),
              ],
            ),
          ),
          child: Row(
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              Flexible(
                child: SailColumn(
                  crossAxisAlignment: CrossAxisAlignment.start,
                  spacing: SailStyleValues.padding16,
                  mainAxisSize: MainAxisSize.min,
                  children: [
                    SailDropdownButton<String>(
                      items: [
                        SailDropdownItem(
                          value: 'US',
                          child: SailText.primary12('US Weekly'),
                        ),
                        SailDropdownItem(
                          value: 'Japan',
                          child: SailText.primary12('Japan Weekly'),
                        ),
                      ],
                      onChanged: viewModel.setLeftRegion,
                      value: viewModel.leftRegion,
                    ),
                    SizedBox(
                      height: 300,
                      child: CoinNewsTable(
                        entries: viewModel.leftEntries,
                        onSort: viewModel.sortEntries,
                      ),
                    ),
                  ],
                ),
              ),
              const SizedBox(width: SailStyleValues.padding16),
              Flexible(
                child: SailColumn(
                  spacing: SailStyleValues.padding16,
                  mainAxisSize: MainAxisSize.min,
                  children: [
                    SailDropdownButton<String>(
                      items: [
                        SailDropdownItem(
                          value: 'Japan',
                          child: SailText.primary12('Japan Weekly'),
                        ),
                        SailDropdownItem(
                          value: 'US',
                          child: SailText.primary12('US Weekly'),
                        ),
                      ],
                      onChanged: viewModel.setRightRegion,
                      value: viewModel.rightRegion,
                    ),
                    SizedBox(
                      height: 300,
                      child: CoinNewsTable(
                        entries: viewModel.rightEntries,
                        onSort: viewModel.sortEntries,
                      ),
                    ),
                  ],
                ),
              ),
            ],
          ),
        );
      },
    );
  }

  Future<void> displayBroadcastNewsDialog(BuildContext context) async {
    await widgetDialog(
      context: context,
      title: 'Broadcast News',
      child: BroadcastNewsView(),
    );
  }
}

class BroadcastNewsView extends StatelessWidget {
  const BroadcastNewsView({
    super.key,
  });

  @override
  Widget build(BuildContext context) {
    return ViewModelBuilder<BroadcastNewsViewModel>.reactive(
      viewModelBuilder: () => BroadcastNewsViewModel(),
      builder: (context, viewModel, child) {
        return SailColumn(
          crossAxisAlignment: CrossAxisAlignment.start,
          spacing: SailStyleValues.padding16,
          mainAxisSize: MainAxisSize.min,
          leadingSpacing: true,
          children: [
            SailDropdownButton<String>(
              items: [
                SailDropdownItem(
                  value: 'US',
                  child: SailText.primary12('US Weekly'),
                ),
                SailDropdownItem(
                  value: 'Japan',
                  child: SailText.primary12('Japan Weekly'),
                ),
              ],
              onChanged: viewModel.setRegion,
              value: viewModel.region,
            ),
            SailTextField(
              label: 'Headline (max 64 characters)',
              controller: viewModel.headlineController,
              hintText: 'Enter a headline',
              size: TextFieldSize.small,
            ),
            QtButton(
              label: 'Broadcast',
              onPressed: () => viewModel.broadcastNews(context),
              size: ButtonSize.small,
              disabled: viewModel.headlineController.text.isEmpty,
            ),
          ],
        );
      },
    );
  }
}

class BroadcastNewsViewModel extends BaseViewModel {
  String _region = 'US';
  String get region => _region;

  final TextEditingController headlineController = TextEditingController();

  BroadcastNewsViewModel() {
    headlineController.addListener(notifyListeners);
  }

  void setRegion(String newRegion) {
    _region = newRegion;
    notifyListeners();
  }

  Future<void> broadcastNews(BuildContext context) async {
    if (headlineController.text.isEmpty) {
      return;
    }

    if (headlineController.text.length > 64) {
      showSnackBar(context, 'Headline must be 64 characters or less');
      return;
    }

    showSnackBar(context, 'TODO: Actually broadcast the news');
  }

  @override
  void dispose() {
    headlineController.removeListener(notifyListeners);
    headlineController.dispose();
    super.dispose();
  }
}

class CoinNewsTable extends StatelessWidget {
  final List<CoinNewsEntry> entries;
  final Function(String) onSort;

  const CoinNewsTable({
    super.key,
    required this.entries,
    required this.onSort,
  });

  @override
  Widget build(BuildContext context) {
    return SailTable(
      getRowId: (index) => entries[index].id,
      headerBuilder: (context) => [
        SailTableHeaderCell(name: 'Fe'),
        SailTableHeaderCell(name: 'Time'),
        SailTableHeaderCell(name: 'Headline'),
      ],
      rowBuilder: (context, row, selected) {
        final entry = entries[row];
        return [
          SailTableCell(child: SailText.primary12(entry.fe.toString())),
          SailTableCell(child: SailText.primary12(entry.time.format())),
          SailTableCell(child: SailText.primary12(entry.headline)),
        ];
      },
      rowCount: entries.length,
      columnWidths: const [150, 200, 150],
      drawGrid: true,
      onSort: (columnIndex, ascending) {
        onSort(['fe', 'time', 'headline'][columnIndex]);
      },
    );
  }
}

class CoinNewsEntry {
  final String id;
  final num fe;
  final DateTime time;
  final String headline;

  CoinNewsEntry({
    required this.id,
    required this.fe,
    required this.time,
    required this.headline,
  });
}
