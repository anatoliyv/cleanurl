# Release Notes - CleanURL

## 🎉 Release v1.0.2

**Release Date**: January 2025  
**Version**: 1.0.2  
**Repository**: https://github.com/anatoliyv/cleanurl

### ✨ New Features

- **Lowercase Conversion**: Added automatic conversion of all URLs to lowercase for consistent processing
- **Enhanced Pipeline**: Lowercase conversion is now the first step in the cleaning pipeline
- **Improved Deduplication**: Better deduplication through case-insensitive URL processing

### 🔧 Changes

- **New Flag**: `--lower` (enabled by default) to control lowercase conversion
- **New Flag**: `--no-lower` to disable lowercase conversion
- **Updated Pipeline**: Processing order is now: lowercase → character cleaning → HTTP deduplication → trailing slash removal
- **Enhanced Tests**: Added comprehensive tests for lowercase conversion functionality
- **Updated Examples**: Example files now include mixed-case URLs

### 🎯 Usage Examples

```bash
# Convert to lowercase
echo 'HTTPS://EXAMPLE.COM' | cleanurl
# Output: https://example.com

# Convert mixed case with character cleaning
echo '!HTTPS://UPPERCASE.COM!' | cleanurl
# Output: https://uppercase.com

# Disable lowercase conversion
echo 'HTTPS://EXAMPLE.COM' | cleanurl --no-lower
# Output: HTTPS://EXAMPLE.COM
```

### 📦 Downloads

Pre-built binaries for v1.0.2:
- **Linux (amd64)**: `cleanurl-linux`
- **macOS (amd64)**: `cleanurl-darwin`
- **Windows (amd64)**: `cleanurl-windows.exe`

### 🚀 Installation

```bash
# Update to latest version
go install github.com/anatoliyv/cleanurl@latest
```

---

## 🎉 Release v1.0.1

**Release Date**: January 2025  
**Version**: 1.0.1  
**Repository**: https://github.com/anatoliyv/cleanurl

### ✨ New Features

- **Enhanced Character Cleaning**: Added exclamation mark (`!`) to the list of characters that are automatically removed from URLs
- **Improved Documentation**: Updated all documentation to reflect the new character cleaning capabilities

### 🔧 Changes

- **Character Cleaning**: Now removes `'`, `"`, and `!` characters from URLs
- **Updated Help Text**: Command-line help now shows the complete list of characters being cleaned
- **Enhanced Tests**: Added comprehensive tests for exclamation mark removal
- **Updated Examples**: Example files now include URLs with exclamation marks

### 🎯 Usage Examples

```bash
# Remove exclamation marks
echo '!https://example.com!' | cleanurl
# Output: https://example.com

# Remove mixed characters
echo -e '"https://example.com"\n!https://test.com!\n'\''https://another.com'\'' | cleanurl
# Output: 
# https://example.com
# https://test.com
# https://another.com
```

### 📦 Downloads

Pre-built binaries for v1.0.1:
- **Linux (amd64)**: `cleanurl-linux`
- **macOS (amd64)**: `cleanurl-darwin`
- **Windows (amd64)**: `cleanurl-windows.exe`

### 🚀 Installation

```bash
# Update to latest version
go install github.com/anatoliyv/cleanurl@latest
```

---

## 🎉 Release v1.0.0

**Release Date**: January 2025  
**Version**: 1.0.0  
**Repository**: https://github.com/anatoliyv/cleanurl

### ✨ Features

#### Core Functionality
- **Character Cleaning**: Remove unnecessary quotes (`'` and `"`) from URLs
- **HTTP/HTTPS Deduplication**: Remove HTTP duplicates when HTTPS version exists
- **Trailing Slash Removal**: Remove trailing slashes to deduplicate URLs
- **Stream Processing**: Process URLs from stdin and output to stdout

#### Command Line Interface
- **Configurable Options**: Enable/disable individual cleaning features
- **User-friendly Flags**: Both positive and negative flag options
- **Comprehensive Help**: Detailed help with examples
- **Cross-platform**: Works on Windows, macOS, and Linux

### 🚀 Installation

#### Quick Install
```bash
go install github.com/anatoliyv/cleanurl@latest
```

#### Manual Installation
```bash
git clone https://github.com/anatoliyv/cleanurl.git
cd cleanurl
go build -o cleanurl
```

### 📦 Downloads

Pre-built binaries are available for:
- **Linux (amd64)**: `cleanurl-linux`
- **macOS (amd64)**: `cleanurl-darwin`
- **Windows (amd64)**: `cleanurl-windows.exe`

### 🎯 Usage Examples

#### Basic Usage
```bash
echo '"https://example.com/"' | cleanurl
# Output: https://example.com
```

#### Process File
```bash
cat urls.txt | cleanurl
```

#### Disable Specific Features
```bash
echo "http://example.com" | cleanurl --no-clean-http
```

#### Complex Example
```bash
echo -e '"https://example.com/"\nhttp://example.com\nhttps://example.com' | cleanurl
# Output: https://example.com
```

### 🧪 Testing

The release includes comprehensive tests:
- Unit tests for all functions
- Integration tests for the complete pipeline
- Command-line flag tests
- Edge case handling

Run tests with:
```bash
go test -v
```

### 📚 Documentation

- **README.md**: Complete documentation with examples
- **INSTALL.md**: Detailed installation guide
- **LICENSE**: MIT License
- **Examples**: Sample files for testing

### 🔧 Technical Details

- **Language**: Go 1.21+
- **Dependencies**: 
  - `github.com/spf13/cobra` - Command-line interface
  - `github.com/stretchr/testify` - Testing utilities
- **Architecture**: Modular design with clear separation of concerns
- **Performance**: Optimized for large URL lists

### 🐛 Bug Reports & Contributions

If you find any issues or want to contribute:
1. Check existing [Issues](https://github.com/anatoliyv/cleanurl/issues)
2. Create a new issue with detailed description
3. Fork the repository and submit a pull request

### 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

### 🙏 Acknowledgments

- Built with [Cobra](https://github.com/spf13/cobra) for the command-line interface
- Tested with [Testify](https://github.com/stretchr/testify) for comprehensive testing
- Inspired by the need for efficient URL processing tools 