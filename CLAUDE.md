# Drivechain Frontends Project Guide

## Project Overview
This is a monorepo containing multiple Drivechain-related Flutter/Dart applications:
- **bitassets** - Asset management application
- **bitnames** - Names management application  
- **bitwindow** - Main frontend application
- **faucet** - Faucet service
- **sail_ui** - Shared UI components
- **thunder** - Thunder application
- **zside** - ZSide application

## Tech Stack
- **Language**: Dart/Flutter
- **Platform**: Cross-platform (Linux, macOS, Windows, Web)
- **Build System**: Flutter SDK

## Development Commands

### Flutter Commands
```bash
# Run the application
flutter run

# Build for release
flutter build linux
flutter build macos
flutter build windows
flutter build web

# Run tests
flutter test

# Analyze code
flutter analyze

# Format code
flutter format .
```

### Project Scripts
```bash
# Using Justfile (if available in subdirectories)
just build
just run
just test
```

## Project Structure
```
/
├── bitassets/     # Asset management app
├── bitnames/      # Names management app
├── bitwindow/     # Main frontend app
│   ├── lib/       # Dart source code
│   ├── server/    # Server components
│   ├── assets/    # Static assets
│   └── test/      # Tests
├── faucet/        # Faucet service
├── sail_ui/       # Shared UI components
├── thunder/       # Thunder app
└── zside/         # ZSide app
```

## Code Style Guidelines
- Follow Dart/Flutter conventions
- Use `flutter format` before committing
- Run `flutter analyze` to check for issues
- Avoid adding comments unless necessary
- Match existing code patterns in each application

## Testing
- Write tests in the `test/` directory of each application
- Run tests with `flutter test`
- Ensure all tests pass before committing changes

## Important Notes
- This is a multi-application repository
- Each application may have its own pubspec.yaml and dependencies
- Shared components should go in sail_ui when appropriate
- Check individual application README files for specific details