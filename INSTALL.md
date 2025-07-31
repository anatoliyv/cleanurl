# Installation Guide for CleanURL

This guide will help you install and run the CleanURL tool on your local machine.

## Prerequisites

- **Go 1.21 or later** - Download from [golang.org](https://golang.org/dl/)
- **Git** (for cloning the repository)

## Installation Steps

### Step 1: Clone the Repository

```bash
git clone https://github.com/anatoliyv/cleanurl.git
cd cleanurl
```

### Step 2: Install Dependencies

```bash
go mod tidy
```

### Step 3: Build the Tool

```bash
go build -o cleanurl
```

### Step 4: Test the Installation

```bash
# Test the help command
./cleanurl --help

# Test with a simple example
echo '"https://example.com/"' | ./cleanurl
```

### Step 5: Install Globally (Optional)

#### On macOS/Linux:
```bash
sudo cp cleanurl /usr/local/bin/
```

#### On Windows (run as administrator):
```cmd
copy cleanurl.exe C:\Windows\System32\
```

## Alternative Installation Methods

### Using Go Install (Recommended for Users)

If the repository is published to GitHub, you can install directly:

```bash
go install github.com/anatoliyv/cleanurl@latest
```

### Using Make (if available)

```bash
# Build the tool
make build

# Run tests
make test

# Install globally
make install
```

## Verification

After installation, verify that the tool works correctly:

```bash
# Test basic functionality
echo -e '"https://example.com/"\nhttp://example.com\nhttps://example.com' | cleanurl

# Expected output:
# https://example.com
```

## Troubleshooting

### Common Issues

1. **"command not found" error**
   - Make sure the binary is in your PATH
   - Try using the full path: `./cleanurl`

2. **Permission denied**
   - Make sure the binary is executable: `chmod +x cleanurl`

3. **Go version too old**
   - Update Go to version 1.21 or later

4. **Dependency issues**
   - Run `go mod tidy` to fix dependency issues

### Getting Help

If you encounter any issues:

1. Check the [README.md](README.md) for usage examples
2. Run `./cleanurl --help` for command-line options
3. Check the [Issues](https://github.com/anatoliyv/cleanurl/issues) page
4. Create a new issue with details about your problem

## Development Setup

For developers who want to contribute:

```bash
# Clone the repository
git clone https://github.com/anatoliyv/cleanurl.git
cd cleanurl

# Install dependencies
go mod tidy

# Run tests
go test -v

# Run tests with coverage
go test -cover

# Build for multiple platforms
make build-all
```

## Next Steps

Once installed, you can:

1. Read the [README.md](README.md) for usage examples
2. Try the examples in the `examples/` directory
3. Explore the command-line options with `./cleanurl --help`
4. Use the tool in your own projects

## Uninstalling

To remove the tool:

```bash
# If installed globally
sudo rm /usr/local/bin/cleanurl

# On Windows
del C:\Windows\System32\cleanurl.exe
``` 