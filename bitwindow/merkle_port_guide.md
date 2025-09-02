# Merkle Tree Calculator Port Guide / Discovery Questions

This guide provides comprehensive questions to assess a target codebase for porting the merkle calculator from the Drivechain/Bitcoin implementation.

## Flutter/Dart Implementation Status

✅ **COMPLETED**: A full Flutter/Dart implementation has been created in this repository with the following features:

### Implementation Files
- `lib/utils/merkle_tree.dart` - Core merkle tree computation logic
- `lib/widgets/merkle_tree_dialog.dart` - Flutter UI dialog with tree visualization

### Features Implemented
- ✅ Double SHA-256 hashing (`sha256d` function)
- ✅ Bitcoin's "duplicate-last-when-odd" behavior
- ✅ RCB (Reversed Byte Order) support with toggle
- ✅ Visual tree representation with node connections
- ✅ Duplicate node indicators (red border with "DUP" label)
- ✅ Interactive node selection
- ✅ Copy merkle root to clipboard
- ✅ Example transaction set
- ✅ Paste from clipboard functionality

### Known Issues
- ⚠️ Minor pixel overflow on duplicate nodes with long hashes (cosmetic issue)
- Can be fixed by conditionally hiding hash display for duplicate nodes

### Integration Points
- Added to Tools menu in `lib/pages/root_page.dart`
- Registered as SubWindowType in window provider system

## 1. Core Dependencies & Infrastructure

### Hash Implementation
- Does the target codebase have `CHash256()` or equivalent double SHA-256 implementation?
- What's the include path for hash functions? (`hash.h`, `crypto/sha256.h`, other?)
- Does it use the same `uint256` type or something different for hash values?

### Data Structures
- How are blocks represented? (Look for `CBlock` equivalent)
- How are transactions stored in blocks? (`vtx` vector or different?)
- Is there a `GetHash()` method on transactions?
- Is there SegWit support with `GetWitnessHash()`?

## 2. Existing Merkle Implementation

### Current State
- Is there already a `consensus/merkle.h` or similar?
- Search for: `ComputeMerkleRoot`, `BlockMerkleRoot`, `MerkleComputation`
- Are there existing merkle-related test files?
- Any existing merkle UI components?

## 3. UI Framework Questions (if porting the dialog)

### Qt Integration
- Is Qt used for UI? Which version?
- Where are dialog classes located? (`src/qt/`, other?)
- Look for `.ui` form files - what's the naming convention?
- How are new dialogs registered/accessed from main GUI?

### UI Patterns
- Find an existing dialog (search for `*dialog.cpp`)
- How do they handle:
  - Form setup (`ui->setupUi(this)` pattern?)
  - Slots and signals?
  - Text display widgets?

## 4. Build System Integration

### Source Files
- Where should merkle consensus code go?
- Where should UI dialog code go?
- How are new files added to build? (`Makefile.am`, `CMakeLists.txt`, `.pro` file?)

### UI Files
- How are `.ui` files processed?
- Where do generated ui headers go?
- Check for `forms/` directory structure

## 5. Code Style & Conventions

### Naming
- Function naming: `ComputeMerkleRoot` vs `compute_merkle_root` vs other?
- Variable prefixes: `vLeaf` (Hungarian) vs `leaf` vs `leaves`?
- Member variables: `m_`, `_` suffix, or plain?

### String Handling
- `std::string` or custom string class?
- How is uint256 converted to string? (`.ToString()`, `.GetHex()`, other?)
- Qt string conversion: `QString::fromStdString()` or different?

## 6. Special Considerations

### Bitcoin Quirks to Check
- Does target maintain Bitcoin's "duplicate-last-when-odd" behavior?
- Is CVE-2012-2459 protection needed/wanted?
- Should RCB (reversed byte) mode be included?

### Performance
- Is space-efficient implementation needed or can we use simpler full-tree version?
- Are there memory constraints?

## 7. Testing Infrastructure

### Test Framework
- What test framework? (Boost.Test, Google Test, other?)
- Where are tests located?
- How to add new test files?
- Look for existing merkle tests to extend/replace

## 8. Menu Integration (for UI)

### Access Points
- Where should merkle calculator be accessible from?
- Look for menu bar setup (`createMenuBar()`, `setupUi()`)
- How are actions connected to dialogs?
- Search for patterns like `connect(action, SIGNAL(triggered()), ...)`

## 9. Block Explorer Integration

### Context
- Is there a block explorer/viewer?
- How to get transaction list from a specific block?
- Can we add "Calculate Merkle Tree" context menu or button?

## 10. Feature Flags/Configuration

### Optional Features
- Should merkle calculator be optional compile-time feature?
- Any `--with-merkle-ui` type flags needed?
- Check for `ENABLE_*` or `USE_*` preprocessor defines

---

## Quick Discovery Commands

Run these in the target repo to answer many questions quickly:

```bash
# Find hash implementation
grep -r "SHA256\|sha256" --include="*.h" --include="*.cpp"

# Find block/transaction structures  
grep -r "class.*Block\|struct.*Block" --include="*.h"

# Find existing merkle code
grep -r -i "merkle" --include="*.h" --include="*.cpp"

# Find UI dialog patterns
find . -name "*dialog.cpp" | head -5

# Find build system files
find . -name "Makefile.am" -o -name "CMakeLists.txt" -o -name "*.pro"

# Find Qt UI forms
find . -name "*.ui"

# Check for uint256 type
grep -r "uint256\|Hash256" --include="*.h"

# Find transaction class
grep -r "class.*Transaction\|struct.*Transaction" --include="*.h"

# Check Qt version and setup
grep -r "QT_VERSION\|QtCore" --include="*.h" --include="*.cpp" --include="*.pro"

# Find existing dialog registration patterns
grep -r "new.*Dialog\|show.*Dialog" --include="*.cpp"
```

## Response Template

After investigating the target codebase, provide answers in this format:

```
HASH IMPLEMENTATION: 
- Found at: src/crypto/sha256.h
- Usage: SHA256D() function, not CHash256
- uint256 type: Located in src/uint256.h

BLOCK STRUCTURE:
- Block class: CBlock in src/primitives/block.h
- Transaction vector: vtx (vector<CTransactionRef>)
- Hash method: GetHash() returns uint256

UI FRAMEWORK:
- Qt version: 5.x found
- Dialogs at: src/qt/dialogs/
- Pattern: Inherits from BaseDialog, not QDialog directly
- UI files: src/qt/forms/*.ui

BUILD SYSTEM:
- Type: CMake
- Add files in: src/CMakeLists.txt
- UI processing: automatic with qt5_wrap_ui()

CODE STYLE:
- Functions: CamelCase (ComputeMerkleRoot)
- Variables: camelCase with type prefix (vLeaves)
- Members: m_ prefix

EXISTING MERKLE:
- Found basic implementation in: src/consensus/merkle.cpp
- No UI component exists
- Tests in: src/test/merkle_tests.cpp

INTEGRATION POINTS:
- Menu bar: src/qt/bitcoingui.cpp::createMenuBar()
- Block explorer: src/qt/transactiontablemodel.cpp
- Can add action to: "Tools" menu

SPECIAL NOTES:
- [Any special considerations found]
```

## Files to Port

Based on the original implementation, these files need porting:

### Core Consensus Files
- `src/consensus/merkle.h` - Header with function declarations
- `src/consensus/merkle.cpp` - Core merkle computation implementation

### UI Files (optional)
- `src/qt/merkletreedialog.h` - Dialog header
- `src/qt/merkletreedialog.cpp` - Dialog implementation  
- `src/qt/forms/merkletreedialog.ui` - Qt Designer form file

### Test Files
- `src/test/merkle_tests.cpp` - Unit tests for merkle functions

### Integration Points
- Modify main GUI file to add menu action
- Update build system files to include new sources
- Add to translation files if internationalization is used

## Porting Checklist

- [ ] Identify hash function implementation
- [ ] Locate uint256 or equivalent type
- [ ] Find block and transaction structures
- [ ] Check for existing merkle code
- [ ] Understand UI framework patterns
- [ ] Determine build system integration
- [ ] Match code style conventions
- [ ] Plan test integration
- [ ] Identify menu/GUI integration points
- [ ] Handle any special Bitcoin consensus rules
- [ ] Consider optional feature flags
- [ ] Update documentation/help files