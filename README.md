# CleanURL

A fast and efficient command-line tool for cleaning and deduplicating URLs from stdin. CleanURL processes URLs through a pipeline of cleaning operations and outputs clean, unique URLs to stdout.

## Features

- **Lowercase Conversion**: Convert all URLs to lowercase for consistent processing
- **Character Cleaning**: Remove unnecessary quotes (`'` and `"`) and exclamation marks (`!`) from URLs
- **HTTP/HTTPS Deduplication**: Remove HTTP duplicates when HTTPS version exists
- **Trailing Slash Removal**: Remove trailing slashes to deduplicate URLs
- **Stream Processing**: Process URLs from stdin and output to stdout
- **Configurable Options**: Enable/disable individual cleaning features
- **Cross-Platform**: Works on Windows, macOS, and Linux

## Installation

### Prerequisites

- Go 1.21 or later

### Local Installation

1. Clone the repository:
```bash
git clone https://github.com/anatoliyv/cleanurl.git
cd cleanurl
```

2. Build the binary:
```bash
go build -o cleanurl
```

3. Install globally (optional):
```bash
# On macOS/Linux
sudo cp cleanurl /usr/local/bin/

# On Windows (run as administrator)
copy cleanurl.exe C:\Windows\System32\
```

### Using Go Install

```bash
go install github.com/anatoliyv/cleanurl@latest
```

## Usage

### Basic Usage

Process URLs from stdin:
```bash
echo "https://example.com/" | cleanurl
```

Process URLs from a file:
```bash
cat urls.txt | cleanurl
```

### Command Line Options

All cleaning features are enabled by default. You can disable specific features using the `--no-*` flags:

```bash
# Disable character cleaning
echo "https://example.com" | cleanurl --no-characters

# Disable HTTP cleaning
echo "http://example.com" | cleanurl --no-clean-http

# Disable trailing slash removal
echo "https://example.com/" | cleanurl --no-backslash

# Disable multiple features
echo "https://example.com" | cleanurl --no-characters --no-clean-http
```

### Available Flags

| Flag | Description | Default |
|------|-------------|---------|
| `-h, --help` | Show help information | - |
| `--lower` | Convert URLs to lowercase | `true` |
| `--characters` | Remove unnecessary characters (quotes and exclamation marks) from URLs | `true` |
| `--clean-http` | Remove HTTP duplicates when HTTPS version exists | `true` |
| `--backslash` | Remove trailing slashes to deduplicate URLs | `true` |
| `--no-lower` | Disable lowercase conversion | - |
| `--no-characters` | Disable character cleaning | - |
| `--no-clean-http` | Disable HTTP cleaning | - |
| `--no-backslash` | Disable backslash cleaning | - |

## Examples

### Example 1: Basic URL Cleaning

**Input:**
```
"https://example.com/"
http://example.com
https://example.com
'https://test.com/'
https://test.com
```

**Command:**
```bash
cat input.txt | cleanurl
```

**Output:**
```
https://example.com
https://test.com
```

### Example 2: Disable HTTP Cleaning

**Input:**
```
http://example.com
https://example.com
```

**Command:**
```bash
echo -e "http://example.com\nhttps://example.com" | cleanurl --no-clean-http
```

**Output:**
```
http://example.com
https://example.com
```

### Example 3: Disable Character Cleaning

**Input:**
```
"https://example.com"
'https://test.com'
```

**Command:**
```bash
echo -e '"https://example.com"\n'\''https://test.com'\'' | cleanurl --no-characters
```

**Output:**
```
"https://example.com"
'https://test.com'
```

### Example 4: Complex Pipeline

**Input:**
```
"http://example.com/"
https://example.com
'https://test.com/'
https://test.com
https://unique.com/
```

**Command:**
```bash
cat complex_input.txt | cleanurl
```

**Output:**
```
https://example.com
https://test.com
https://unique.com
```

### Example 5: Lowercase Conversion and Character Cleaning

**Input:**
```
"HTTPS://EXAMPLE.COM/"
!https://UPPERCASE.com!
'https://MixedCase.Com/'
```

**Command:**
```bash
echo -e '"HTTPS://EXAMPLE.COM/"\n!https://UPPERCASE.com!\n'\''https://MixedCase.Com/'\'' | cleanurl
```

**Output:**
```
https://example.com
https://uppercase.com
https://mixedcase.com
```

### Example 6: Disable Multiple Features

**Input:**
```
"HTTPS://EXAMPLE.COM/"
!https://UPPERCASE.com!
'https://MixedCase.Com/'
```

**Command:**
```bash
echo -e '"HTTPS://EXAMPLE.COM/"\n!https://UPPERCASE.com!\n'\''https://MixedCase.Com/'\'' | cleanurl --no-lower --no-characters
```

**Output:**
```
"HTTPS://EXAMPLE.COM/"
!https://UPPERCASE.com!
'https://MixedCase.Com/'
```

## How It Works

CleanURL processes URLs through the following pipeline:

1. **Lowercase Conversion** (enabled by default)
   - Converts all URLs to lowercase for consistent processing
   - Example: `HTTPS://EXAMPLE.COM` → `https://example.com`

2. **Character Cleaning** (enabled by default)
   - Removes single quotes (`'`), double quotes (`"`), and exclamation marks (`!`) from URLs
   - Example: `"https://example.com"` → `https://example.com`
   - Example: `!https://example.com!` → `https://example.com`

3. **Trailing Slash Removal** (enabled by default)
   - Always removes trailing slashes for consistency
   - Example: `https://example.com/` → `https://example.com`

4. **HTTP/HTTPS Deduplication** (enabled by default)
   - If both HTTP and HTTPS versions of the same URL exist, keeps only HTTPS
   - Example: `http://example.com` + `https://example.com` → `https://example.com`

## Testing

Run the test suite:

```bash
go test
```

Run tests with verbose output:

```bash
go test -v
```

Run tests with coverage:

```bash
go test -cover
```

## Development

### Project Structure

```
cleanurl/
├── main.go          # Main application code
├── main_test.go     # Test suite
├── go.mod           # Go module file
├── go.sum           # Go module checksums
└── README.md        # This file
```

### Building

```bash
# Build for current platform
go build -o cleanurl

# Build for specific platforms
GOOS=linux GOARCH=amd64 go build -o cleanurl-linux
GOOS=darwin GOARCH=amd64 go build -o cleanurl-macos
GOOS=windows GOARCH=amd64 go build -o cleanurl-windows.exe
```

### Dependencies

- `github.com/spf13/cobra` - Command-line interface framework
- `github.com/stretchr/testify` - Testing utilities

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Support

If you encounter any issues or have questions, please:

1. Check the [Issues](https://github.com/anatoliyv/cleanurl/issues) page
2. Create a new issue with a detailed description of the problem
3. Include your operating system and Go version

## Changelog

### v1.0.4 (Latest)
- Fixed trailing slash removal logic
- Improved HTTP/HTTPS deduplication

### v1.0.3
- Added lowercase conversion feature
- Added `--lower` and `--no-lower` flags

### v1.0.2
- Added lowercase conversion feature
- Added `--lower` and `--no-lower` flags

### v1.0.1
- Added exclamation mark (`!`) to character cleaning
- Updated documentation and tests

### v1.0.0
- Initial release with character cleaning, HTTP/HTTPS deduplication, and trailing slash removal 