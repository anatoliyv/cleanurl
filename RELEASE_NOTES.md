# Release Notes - CleanURL v1.0.0

## 🎉 Initial Release

This is the first official release of CleanURL, a powerful command-line tool for cleaning and deduplicating URLs.

## ✨ Features

### Core Functionality
- **Character Cleaning**: Remove unnecessary quotes (`'` and `"`) from URLs
- **HTTP/HTTPS Deduplication**: Remove HTTP duplicates when HTTPS version exists
- **Trailing Slash Removal**: Remove trailing slashes to deduplicate URLs
- **Stream Processing**: Process URLs from stdin and output to stdout

### Command Line Interface
- **Configurable Options**: Enable/disable individual cleaning features
- **User-friendly Flags**: Both positive and negative flag options
- **Comprehensive Help**: Detailed help with examples
- **Cross-platform**: Works on Windows, macOS, and Linux

## 🚀 Installation

### Quick Install
```bash
go install github.com/anatoliyv/cleanurl@latest
```

### Manual Installation
```bash
git clone https://github.com/anatoliyv/cleanurl.git
cd cleanurl
go build -o cleanurl
```

## 📦 Downloads

Pre-built binaries are available for:
- **Linux (amd64)**: `cleanurl-linux`
- **macOS (amd64)**: `cleanurl-darwin`
- **Windows (amd64)**: `cleanurl-windows.exe`

## 🎯 Usage Examples

### Basic Usage
```bash
echo '"https://example.com/"' | cleanurl
# Output: https://example.com
```

### Process File
```bash
cat urls.txt | cleanurl
```

### Disable Specific Features
```bash
echo "http://example.com" | cleanurl --no-clean-http
```

### Complex Example
```bash
echo -e '"https://example.com/"\nhttp://example.com\nhttps://example.com' | cleanurl
# Output: https://example.com
```

## 🧪 Testing

The release includes comprehensive tests:
- Unit tests for all functions
- Integration tests for the complete pipeline
- Command-line flag tests
- Edge case handling

Run tests with:
```bash
go test -v
```

## 📚 Documentation

- **README.md**: Complete documentation with examples
- **INSTALL.md**: Detailed installation guide
- **LICENSE**: MIT License
- **Examples**: Sample files for testing

## 🔧 Technical Details

- **Language**: Go 1.21+
- **Dependencies**: 
  - `github.com/spf13/cobra` - Command-line interface
  - `github.com/stretchr/testify` - Testing utilities
- **Architecture**: Modular design with clear separation of concerns
- **Performance**: Optimized for large URL lists

## 🐛 Bug Reports & Contributions

If you find any issues or want to contribute:
1. Check existing [Issues](https://github.com/anatoliyv/cleanurl/issues)
2. Create a new issue with detailed description
3. Fork the repository and submit a pull request

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- Built with [Cobra](https://github.com/spf13/cobra) for the command-line interface
- Tested with [Testify](https://github.com/stretchr/testify) for comprehensive testing
- Inspired by the need for efficient URL processing tools

---

**Release Date**: January 2025  
**Version**: 1.0.0  
**Repository**: https://github.com/anatoliyv/cleanurl 