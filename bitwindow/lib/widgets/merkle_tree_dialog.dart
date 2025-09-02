import 'package:bitwindow/utils/merkle_tree.dart';
import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:sail_ui/sail_ui.dart';

class MerkleTreeDialog extends StatefulWidget {
  final List<String>? initialTxids;
  final String? blockHash;
  final int? blockHeight;

  const MerkleTreeDialog({
    super.key,
    this.initialTxids,
    this.blockHash,
    this.blockHeight,
  });

  @override
  State<MerkleTreeDialog> createState() => _MerkleTreeDialogState();
}

class _MerkleTreeDialogState extends State<MerkleTreeDialog> {
  final TextEditingController _inputController = TextEditingController();
  
  MerkleTreeResult? _result;
  bool _reverseHashes = true;
  String _error = '';
  int? _selectedNodeLevel;
  int? _selectedNodeIndex;
  
  @override
  void initState() {
    super.initState();
    if (widget.initialTxids != null) {
      _inputController.text = widget.initialTxids!.join('\n');
      _calculateMerkleTree();
    }
  }

  void _calculateMerkleTree() {
    final input = _inputController.text.trim();
    if (input.isEmpty) {
      setState(() {
        _result = null;
        _error = '';
      });
      return;
    }

    try {
      final lines = input.split('\n').where((line) => line.trim().isNotEmpty).toList();
      
      for (String line in lines) {
        final trimmed = line.trim();
        if (!RegExp(r'^[0-9a-fA-F]{64}$').hasMatch(trimmed)) {
          setState(() {
            _error = 'Invalid transaction hash: $trimmed';
            _result = null;
          });
          return;
        }
      }

      final result = MerkleTree.computeMerkleTree(lines, reverseHashes: _reverseHashes);
      
      setState(() {
        _result = result;
        _error = '';
      });
    } catch (e) {
      setState(() {
        _error = 'Error: $e';
        _result = null;
      });
    }
  }

  @override
  Widget build(BuildContext context) {
    final theme = SailTheme.of(context);

    return SailCard(
      title: 'Merkle Tree Calculator',
      subtitle: widget.blockHeight != null 
          ? 'Block #${widget.blockHeight}' 
          : 'Calculate merkle root from transaction hashes',
      padding: true,
      widgetHeaderEnd: SailButton(
        label: '?',
        onPressed: () async {
          _showHelpDialog(context);
        },
        variant: ButtonVariant.ghost,
      ),
      withCloseButton: false,
      child: Padding(
        padding: const EdgeInsets.all(SailStyleValues.padding16),
        child: SailColumn(
          spacing: SailStyleValues.padding12,
          crossAxisAlignment: CrossAxisAlignment.stretch,
          children: [
            SailCard(
              padding: true,
              secondary: true,
              child: SailColumn(
                spacing: SailStyleValues.padding12,
                crossAxisAlignment: CrossAxisAlignment.stretch,
                children: [
                  SailRow(
                    spacing: SailStyleValues.padding08,
                    mainAxisAlignment: MainAxisAlignment.start,
                    children: [
                      SailText.primary13('Transaction Hashes'),
                      SailCheckbox(
                        value: _reverseHashes,
                        onChanged: (value) {
                          setState(() {
                            _reverseHashes = value;
                            if (_inputController.text.isNotEmpty) {
                              _calculateMerkleTree();
                            }
                          });
                        },
                        label: 'RCB (Reversed Byte Order)',
                      ),
                      const Spacer(),
                      SailButton(
                        label: 'Clear',
                        onPressed: () async {
                          _inputController.clear();
                          setState(() {
                            _result = null;
                            _error = '';
                          });
                        },
                        variant: ButtonVariant.secondary,
                      ),
                      SailButton(
                        label: 'Paste',
                        onPressed: () async {
                          final data = await Clipboard.getData(Clipboard.kTextPlain);
                          if (data?.text != null) {
                            _inputController.text = data!.text!;
                            _calculateMerkleTree();
                          }
                        },
                        variant: ButtonVariant.secondary,
                      ),
                      SailButton(
                        label: 'Example',
                        onPressed: () async {
                          _inputController.text = [
                            '51b78168dcceb0b14f4e9f05a49b60158ba0afea7fb72323c1b8fb5c9a1ea310',
                            '84f3c6e12e8d3f3d3e86e4e75f9e948c3c4d3e8f3e5c3e2d3e1c3e0d3e9c3e8d',
                            '0a3b4c5d6e7f8091a2b3c4d5e6f7081928374658492837465920183746582019',
                          ].join('\n');
                          _calculateMerkleTree();
                        },
                        variant: ButtonVariant.secondary,
                      ),
                    ],
                  ),
                  SailTextField(
                    controller: _inputController,
                    hintText: 'Enter transaction hashes (one per line)',
                    textFieldType: TextFieldType.text,
                    size: TextFieldSize.small,
                    maxLines: 5,
                  ),
                  if (_error.isNotEmpty)
                    SailText.primary13(
                      _error,
                      color: theme.colors.error,
                    ),
                ],
              ),
            ),
            
            if (_result != null) ...[
              SailCard(
                padding: true,
                secondary: true,
                child: SailColumn(
                  spacing: SailStyleValues.padding08,
                  crossAxisAlignment: CrossAxisAlignment.stretch,
                  children: [
                    SailText.primary13('Merkle Root', color: theme.colors.text),
                    SailRow(
                      spacing: SailStyleValues.padding08,
                      children: [
                        Expanded(
                          child: SelectableText(
                            _result!.root,
                            style: TextStyle(
                              fontFamily: 'monospace',
                              fontSize: 13,
                              color: theme.colors.success,
                            ),
                          ),
                        ),
                        SailButton(
                          label: 'Copy',
                          onPressed: () async {
                            await Clipboard.setData(ClipboardData(text: _result!.root));
                          },
                          variant: ButtonVariant.secondary,
                        ),
                      ],
                    ),
                    SailText.primary13(
                      'Levels: ${_result!.levels} | Transactions: ${_result!.tree.isNotEmpty ? _result!.tree[0].length : 0}',
                      color: theme.colors.textTertiary,
                    ),
                  ],
                ),
              ),
              
              Expanded(
                child: SailCard(
                  padding: true,
                  secondary: true,
                  child: Column(
                    crossAxisAlignment: CrossAxisAlignment.stretch,
                    children: [
                      SailText.primary13('Merkle Tree Visualization', color: theme.colors.text),
                      const SailSpacing(SailStyleValues.padding08),
                      Expanded(
                        child: SingleChildScrollView(
                          scrollDirection: Axis.horizontal,
                          child: SingleChildScrollView(
                            scrollDirection: Axis.vertical,
                            child: _buildTreeVisualization(theme),
                          ),
                        ),
                      ),
                    ],
                  ),
                ),
              ),
            ],
          ],
        ),
      ),
    );
  }

  Widget _buildTreeVisualization(SailThemeData theme) {
    if (_result == null) return const SizedBox.shrink();

    const double nodeWidth = 140.0;
    const double nodeHeight = 50.0;
    const double horizontalSpacing = 20.0;
    const double verticalSpacing = 80.0;

    double maxWidth = 0;
    for (var level in _result!.tree) {
      double levelWidth = level.length * (nodeWidth + horizontalSpacing);
      if (levelWidth > maxWidth) {
        maxWidth = levelWidth;
      }
    }

    double totalHeight = _result!.levels * (nodeHeight + verticalSpacing);

    return SizedBox(
      width: maxWidth + 100,
      height: totalHeight + 100,
      child: CustomPaint(
        painter: _MerkleTreePainter(
          result: _result!,
          nodeWidth: nodeWidth,
          nodeHeight: nodeHeight,
          horizontalSpacing: horizontalSpacing,
          verticalSpacing: verticalSpacing,
          theme: theme,
        ),
        child: Stack(
          children: [
            for (int levelIndex = _result!.levels - 1; levelIndex >= 0; levelIndex--)
              ..._buildLevelNodes(
                levelIndex,
                nodeWidth,
                nodeHeight,
                horizontalSpacing,
                verticalSpacing,
                maxWidth,
                theme,
              ),
          ],
        ),
      ),
    );
  }

  List<Widget> _buildLevelNodes(
    int levelIndex,
    double nodeWidth,
    double nodeHeight,
    double horizontalSpacing,
    double verticalSpacing,
    double maxWidth,
    SailThemeData theme,
  ) {
    final level = _result!.tree[levelIndex];
    final y = (_result!.levels - 1 - levelIndex) * (nodeHeight + verticalSpacing) + 20;
    
    final levelWidth = level.length * (nodeWidth + horizontalSpacing) - horizontalSpacing;
    final startX = (maxWidth - levelWidth) / 2;

    return level.asMap().entries.map((entry) {
      final index = entry.key;
      final node = entry.value;
      final x = startX + index * (nodeWidth + horizontalSpacing);
      
      final isSelected = _selectedNodeLevel == levelIndex && _selectedNodeIndex == index;
      final isRoot = levelIndex == _result!.levels - 1;
      
      return Positioned(
        left: x,
        top: y,
        width: nodeWidth,
        height: nodeHeight,
        child: GestureDetector(
          onTap: () {
            setState(() {
              _selectedNodeLevel = levelIndex;
              _selectedNodeIndex = index;
            });
          },
          child: Container(
            decoration: BoxDecoration(
              color: isRoot 
                  ? theme.colors.success.withValues(alpha: 0.2)
                  : isSelected 
                      ? theme.colors.primary.withValues(alpha: 0.2)
                      : theme.colors.backgroundSecondary,
              border: Border.all(
                color: isRoot
                    ? theme.colors.success
                    : isSelected
                        ? theme.colors.primary
                        : node.isDuplicate
                            ? theme.colors.error.withValues(alpha: 0.8)
                            : theme.colors.divider,
                width: isRoot || isSelected ? 2 : 1,
              ),
              borderRadius: BorderRadius.circular(4),
            ),
            padding: const EdgeInsets.all(4),
            child: Column(
              mainAxisAlignment: MainAxisAlignment.center,
              children: [
                if (isRoot)
                  SailText.primary10('ROOT', color: theme.colors.success),
                if (node.isDuplicate)
                  SailText.primary10('DUP', color: theme.colors.error.withValues(alpha: 0.8)),
                SailText.primary10(
                  '${node.displayHash.substring(0, 8)}...',
                  color: theme.colors.text,
                ),
                SailText.primary10(
                  'L${node.level} #${node.index}',
                  color: theme.colors.textTertiary,
                ),
              ],
            ),
          ),
        ),
      );
    }).toList();
  }

  void _showHelpDialog(BuildContext context) {
    final theme = SailTheme.of(context);
    showDialog(
      context: context,
      builder: (context) => Dialog(
        backgroundColor: theme.colors.backgroundSecondary,
        child: IntrinsicWidth(
          child: SailCard(
            title: 'Merkle Tree Calculator Help',
            subtitle: '',
            padding: true,
            withCloseButton: true,
            child: SailColumn(
              spacing: SailStyleValues.padding08,
              mainAxisSize: MainAxisSize.min,
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                SailText.primary13(
                  'The Merkle Tree Calculator computes the merkle root from a list of transaction hashes.',
                ),
                SailText.primary13(
                  'Transaction Hashes: Enter one 64-character hex transaction ID per line.',
                ),
                SailText.primary13(
                  'RCB Mode: Reversed Byte Order - Bitcoin uses little-endian format for hashes.',
                ),
                SailText.primary13(
                  'Merkle Root: The final hash at the top of the tree that represents all transactions.',
                ),
                SailText.primary13(
                  'Tree Visualization: Shows how transaction hashes are paired and hashed to form the tree.',
                ),
                SailText.primary13(
                  'DUP: Indicates a duplicate hash (when odd number of elements at a level).',
                ),
                SailText.primary13(
                  'The merkle root allows verification of transaction inclusion without storing all transactions.',
                ),
              ],
            ),
          ),
        ),
      ),
    );
  }

  @override
  void dispose() {
    _inputController.removeListener(_calculateMerkleTree);
    _inputController.dispose();
    super.dispose();
  }
}

class _MerkleTreePainter extends CustomPainter {
  final MerkleTreeResult result;
  final double nodeWidth;
  final double nodeHeight;
  final double horizontalSpacing;
  final double verticalSpacing;
  final SailThemeData theme;

  _MerkleTreePainter({
    required this.result,
    required this.nodeWidth,
    required this.nodeHeight,
    required this.horizontalSpacing,
    required this.verticalSpacing,
    required this.theme,
  });

  @override
  void paint(Canvas canvas, Size size) {
    final paint = Paint()
      ..color = theme.colors.divider
      ..strokeWidth = 1
      ..style = PaintingStyle.stroke;

    double maxWidth = 0;
    for (var level in result.tree) {
      double levelWidth = level.length * (nodeWidth + horizontalSpacing);
      if (levelWidth > maxWidth) {
        maxWidth = levelWidth;
      }
    }

    for (int levelIndex = 0; levelIndex < result.levels - 1; levelIndex++) {
      final level = result.tree[levelIndex];
      final nextLevel = result.tree[levelIndex + 1];
      
      final y1 = (result.levels - 1 - levelIndex) * (nodeHeight + verticalSpacing) + 20 + nodeHeight;
      final y2 = (result.levels - 2 - levelIndex) * (nodeHeight + verticalSpacing) + 20;
      
      final levelWidth = level.length * (nodeWidth + horizontalSpacing) - horizontalSpacing;
      final startX = (maxWidth - levelWidth) / 2;
      
      final nextLevelWidth = nextLevel.length * (nodeWidth + horizontalSpacing) - horizontalSpacing;
      final nextStartX = (maxWidth - nextLevelWidth) / 2;
      
      for (int i = 0; i < level.length; i++) {
        final x1 = startX + i * (nodeWidth + horizontalSpacing) + nodeWidth / 2;
        final parentIndex = i ~/ 2;
        final x2 = nextStartX + parentIndex * (nodeWidth + horizontalSpacing) + nodeWidth / 2;
        
        canvas.drawLine(
          Offset(x1, y1),
          Offset(x2, y2),
          paint,
        );
      }
    }
  }

  @override
  bool shouldRepaint(_MerkleTreePainter oldDelegate) => true;
}