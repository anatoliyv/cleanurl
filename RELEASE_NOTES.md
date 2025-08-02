# Release Notes - CleanURL

## 🎉 Release v1.0.6

**Release Date**: January 2025  
**Version**: 1.0.6  
**Repository**: https://github.com/anatoliyv/cleanurl

### 🔧 Port Handling Improvements

- **Domain Extraction**: Fixed domain extraction to properly remove port numbers from URLs
- **HTTP/HTTPS Deduplication**: Enhanced to work correctly with URLs containing ports
- **URL Normalization**: Added `normalizeURLForComparison()` helper for better port handling
- **Comprehensive Testing**: Added extensive test coverage for port handling scenarios

### ✨ New Features

- **Port-Aware Processing**: All features now properly handle URLs with port numbers
- **Enhanced Domain Extraction**: Ports are removed when extracting domains (e.g., `example.com:8080` → `example.com`)
- **Improved Deduplication**: HTTP/HTTPS deduplication works with URLs containing ports

### 🧪 Testing Enhancements

- **New Test Functions**: Added `TestExtractDomain`, `TestNormalizeURLForComparison`, and `TestExtractUniqueDomains`
- **Port Scenarios**: Comprehensive test coverage for URLs with ports
- **Edge Cases**: Tests for IP addresses, complex paths, and special characters with ports
- **Coverage**: Achieved 83.2% test coverage

### 🎯 Usage Examples

```bash
# Domain extraction with ports
echo -e "https://example.com:8080/path\nhttp://test.com:9090/another" | cleanurl --only-domains
# Output: 
# example.com
# test.com

# HTTP/HTTPS deduplication with ports
echo -e "https://example.com:8080/path\nhttp://example.com:8080/path" | cleanurl
# Output: https://example.com:8080/path

# Complex port scenarios
echo -e "'https://EXAMPLE.com:8080/path/'\n\"http://example.com:9090/another\"" | cleanurl
# Output:
# https://example.com:8080/path
# http://example.com:9090/another
```

### 📦 Downloads

Pre-built binaries for v1.0.6:
- **Linux (amd64)**: `cleanurl-linux`
- **macOS (amd64)**: `cleanurl-darwin`
- **Windows (amd64)**: `cleanurl-windows.exe`

### 🚀 Installation

```bash
# Install latest version
go install github.com/anatoliyv/cleanurl@v1.0.6

# Or use latest
go install github.com/anatoliyv/cleanurl@latest
```

---

## 🎉 Release v1.0.5

**Release Date**: January 2025  
**Version**: 1.0.5  
**Repository**: https://github.com/anatoliyv/cleanurl

### ✨ New Features

- **Domain Extraction**: Added `--only-domains` flag to extract unique domain names from URLs
- **Domain Processing**: Automatically converts to lowercase, removes protocol, www prefix, paths, and trailing slashes
- **Port Handling**: Properly handles URLs with port numbers in domain extraction

### 🔧 Changes

- **New Flag**: `--only-domains` (disabled by default) to extract only domain names
- **Domain Pipeline**: Processing includes lowercase conversion, character cleaning, and port removal
- **Enhanced Help**: Updated command-line help with new feature and examples
- **Updated Documentation**: Added examples and usage instructions for domain extraction

### 🎯 Usage Examples

```bash
# Extract unique domains
echo -e "https://example.com:8080/path\nhttp://test.com:9090/another" | cleanurl --only-domains
# Output: 
# example.com
# test.com

# Domain extraction with mixed case and special characters
echo -e "'https://EXAMPLE.com:8080/path/'\n\"http://test.com:9090/another\"" | cleanurl --only-domains
# Output:
# example.com
# test.com

# Complex domain deduplication
echo -e "https://www.example.com:8080/path\nhttp://example.com:9090/another" | cleanurl --only-domains
# Output: example.com
```

### 📦 Downloads

Pre-built binaries for v1.0.5:
- **Linux (amd64)**: `cleanurl-linux`
- **macOS (amd64)**: `cleanurl-darwin`
- **Windows (amd64)**: `cleanurl-windows.exe`

### 🚀 Installation

```bash
# Install latest version
go install github.com/anatoliyv/cleanurl@v1.0.5

# Or use latest
go install github.com/anatoliyv/cleanurl@latest
```

---

## 🎉 Release v1.0.3

**Release Date**: January 2025  
**Version**: 1.0.3  
**Repository**: https://github.com/anatoliyv/cleanurl

### 🔧 Fixes

- **Module Version**: Ensured v1.0.3 is properly recognized as the latest version
- **Installation**: Fixed issue where `@latest` was not picking up the newest version

### ✨ Features

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

Pre-built binaries for v1.0.3:
- **Linux (amd64)**: `cleanurl-linux`
- **macOS (amd64)**: `cleanurl-darwin`
- **Windows (amd64)**: `cleanurl-windows.exe`

### 🚀 Installation

```bash
# Install latest version
go install github.com/anatoliyv/cleanurl@v1.0.3

# Or use latest (should now work correctly)
go install github.com/anatoliyv/cleanurl@latest
```

---

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