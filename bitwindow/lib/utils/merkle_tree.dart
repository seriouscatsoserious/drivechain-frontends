import 'dart:typed_data';
import 'package:crypto/crypto.dart';

class MerkleTree {
  static String sha256d(Uint8List data) {
    final firstHash = sha256.convert(data);
    final secondHash = sha256.convert(firstHash.bytes);
    return secondHash.toString();
  }

  static Uint8List sha256dBytes(Uint8List data) {
    final firstHash = sha256.convert(data);
    final secondHash = sha256.convert(firstHash.bytes);
    return Uint8List.fromList(secondHash.bytes);
  }

  static Uint8List hexToBytes(String hex) {
    if (hex.length % 2 != 0) {
      throw ArgumentError('Invalid hex string length');
    }
    
    final bytes = <int>[];
    for (var i = 0; i < hex.length; i += 2) {
      bytes.add(int.parse(hex.substring(i, i + 2), radix: 16));
    }
    return Uint8List.fromList(bytes);
  }

  static String bytesToHex(Uint8List bytes) {
    return bytes.map((b) => b.toRadixString(16).padLeft(2, '0')).join();
  }

  static Uint8List reverseBytes(Uint8List bytes) {
    return Uint8List.fromList(bytes.reversed.toList());
  }

  static MerkleTreeResult computeMerkleTree(
    List<String> txids, {
    bool reverseHashes = true,
  }) {
    if (txids.isEmpty) {
      return MerkleTreeResult(
        root: '0' * 64,
        tree: [],
        levels: 0,
      );
    }

    final List<List<MerkleNode>> tree = [];
    
    List<MerkleNode> currentLevel = [];
    for (int i = 0; i < txids.length; i++) {
      String txid = txids[i];
      
      Uint8List hashBytes;
      if (reverseHashes) {
        hashBytes = reverseBytes(hexToBytes(txid));
      } else {
        hashBytes = hexToBytes(txid);
      }
      
      currentLevel.add(MerkleNode(
        hash: bytesToHex(hashBytes),
        displayHash: txid,
        index: i,
        level: 0,
        isLeft: i % 2 == 0,
        isDuplicate: false,
      ));
    }
    tree.add(currentLevel);

    int level = 0;
    while (currentLevel.length > 1) {
      level++;
      List<MerkleNode> nextLevel = [];
      
      for (int i = 0; i < currentLevel.length; i += 2) {
        MerkleNode left = currentLevel[i];
        MerkleNode right;
        bool isDuplicate = false;
        
        if (i + 1 < currentLevel.length) {
          right = currentLevel[i + 1];
        } else {
          right = currentLevel[i];
          isDuplicate = true;
        }
        
        final leftBytes = hexToBytes(left.hash);
        final rightBytes = hexToBytes(right.hash);
        final combined = Uint8List.fromList([...leftBytes, ...rightBytes]);
        final parentHash = sha256dBytes(combined);
        
        nextLevel.add(MerkleNode(
          hash: bytesToHex(parentHash),
          displayHash: reverseHashes ? bytesToHex(reverseBytes(parentHash)) : bytesToHex(parentHash),
          index: i ~/ 2,
          level: level,
          isLeft: (i ~/ 2) % 2 == 0,
          isDuplicate: isDuplicate,
          leftChild: left,
          rightChild: right,
        ));
      }
      
      currentLevel = nextLevel;
      tree.add(currentLevel);
    }
    
    final root = currentLevel[0];
    return MerkleTreeResult(
      root: root.displayHash,
      tree: tree,
      levels: tree.length,
    );
  }

  static List<String> getMerkleProof(
    List<String> txids,
    int txIndex, {
    bool reverseHashes = true,
  }) {
    if (txids.isEmpty || txIndex < 0 || txIndex >= txids.length) {
      return [];
    }

    final result = computeMerkleTree(txids, reverseHashes: reverseHashes);
    final proof = <String>[];
    
    int currentIndex = txIndex;
    for (int level = 0; level < result.levels - 1; level++) {
      final levelNodes = result.tree[level];
      final isLeft = currentIndex % 2 == 0;
      final siblingIndex = isLeft ? currentIndex + 1 : currentIndex - 1;
      
      if (siblingIndex < levelNodes.length) {
        proof.add(levelNodes[siblingIndex].displayHash);
      } else {
        proof.add(levelNodes[currentIndex].displayHash);
      }
      
      currentIndex = currentIndex ~/ 2;
    }
    
    return proof;
  }

  static bool verifyMerkleProof(
    String txid,
    List<String> proof,
    String merkleRoot, {
    bool reverseHashes = true,
  }) {
    if (proof.isEmpty) {
      return false;
    }

    Uint8List currentHash;
    if (reverseHashes) {
      currentHash = reverseBytes(hexToBytes(txid));
    } else {
      currentHash = hexToBytes(txid);
    }

    for (String proofElement in proof) {
      Uint8List proofHash;
      if (reverseHashes) {
        proofHash = reverseBytes(hexToBytes(proofElement));
      } else {
        proofHash = hexToBytes(proofElement);
      }

      final combined = Uint8List.fromList([...currentHash, ...proofHash]);
      currentHash = sha256dBytes(combined);
    }

    final computedRoot = reverseHashes 
        ? bytesToHex(reverseBytes(currentHash)) 
        : bytesToHex(currentHash);
    
    return computedRoot == merkleRoot;
  }
}

class MerkleTreeResult {
  final String root;
  final List<List<MerkleNode>> tree;
  final int levels;

  MerkleTreeResult({
    required this.root,
    required this.tree,
    required this.levels,
  });
}

class MerkleNode {
  final String hash;
  final String displayHash;
  final int index;
  final int level;
  final bool isLeft;
  final bool isDuplicate;
  final MerkleNode? leftChild;
  final MerkleNode? rightChild;

  MerkleNode({
    required this.hash,
    required this.displayHash,
    required this.index,
    required this.level,
    required this.isLeft,
    required this.isDuplicate,
    this.leftChild,
    this.rightChild,
  });
}