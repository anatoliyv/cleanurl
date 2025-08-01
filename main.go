package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var (
	// Flags
	characters bool
	cleanHTTP  bool
	backslash  bool
	lower      bool
	onlyDomains bool
)

var rootCmd = &cobra.Command{
	Use:   "cleanurl",
	Short: "Clean and deduplicate URLs from stdin",
	Long: `CleanURL is a command-line tool that processes URLs from stdin and applies various cleaning operations.

Features:
- Convert URLs to lowercase for consistent processing
- Remove unnecessary characters (quotes and exclamation marks) from URLs
- Remove HTTP duplicates when HTTPS version exists
- Remove trailing slashes to deduplicate URLs
- Extract unique domain names from URLs (--only-domains)
- Output cleaned URLs to stdout

Examples:
  echo "https://example.com/" | cleanurl
  cat urls.txt | cleanurl --no-characters
  echo "http://example.com" | cleanurl --no-clean-http
  echo "https://example.com/path" | cleanurl --only-domains`,
	Run: runCleanURL,
}

func init() {
	// Set default values to true (on by default)
	rootCmd.Flags().BoolVar(&characters, "characters", true, "Remove unnecessary characters (quotes and exclamation marks) from URLs")
	rootCmd.Flags().BoolVar(&cleanHTTP, "clean-http", true, "Remove HTTP duplicates when HTTPS version exists")
	rootCmd.Flags().BoolVar(&backslash, "backslash", true, "Remove trailing slashes to deduplicate URLs")
	rootCmd.Flags().BoolVar(&lower, "lower", true, "Convert URLs to lowercase")
	rootCmd.Flags().BoolVar(&onlyDomains, "only-domains", false, "Extract only unique domain names from URLs")
	
	// Add negative flags for convenience
	rootCmd.Flags().Bool("no-characters", false, "Disable character cleaning")
	rootCmd.Flags().Bool("no-clean-http", false, "Disable HTTP cleaning")
	rootCmd.Flags().Bool("no-backslash", false, "Disable backslash cleaning")
	rootCmd.Flags().Bool("no-lower", false, "Disable lowercase conversion")
}

func runCleanURL(cmd *cobra.Command, args []string) {
	// Handle negative flags
	if cmd.Flag("no-characters").Changed {
		characters = false
	}
	if cmd.Flag("no-clean-http").Changed {
		cleanHTTP = false
	}
	if cmd.Flag("no-backslash").Changed {
		backslash = false
	}
	if cmd.Flag("no-lower").Changed {
		lower = false
	}

	// Read URLs from stdin
	urls := readURLsFromStdin()
	
	// Apply cleaning operations
	var cleanedURLs []string
	if onlyDomains {
		cleanedURLs = extractUniqueDomains(urls)
	} else {
		cleanedURLs = cleanURLs(urls)
	}
	
	// Output results
	for _, url := range cleanedURLs {
		fmt.Println(url)
	}
}

func readURLsFromStdin() []string {
	var urls []string
	scanner := bufio.NewScanner(os.Stdin)
	
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			urls = append(urls, line)
		}
	}
	
	if len(urls) == 0 {
		return []string{}
	}
	return urls
}

func cleanURLs(urls []string) []string {
	if len(urls) == 0 {
		return []string{}
	}

	// Step 1: Convert to lowercase
	if lower {
		urls = convertToLowercase(urls)
	}

	// Step 2: Remove unnecessary characters
	if characters {
		urls = removeUnnecessaryCharacters(urls)
	}

	// Step 2: Create maps for tracking
	urlMap := make(map[string]bool)
	httpsMap := make(map[string]bool)
	var result []string

	// First pass: collect HTTPS URLs
	for _, url := range urls {
		// Track HTTPS URLs by their normalized form
		if strings.HasPrefix(url, "https://") {
			normalized := normalizeURLForComparison(url)
			httpsMap[normalized] = true
		}
	}
	
	// Second pass: process URLs
	for _, url := range urls {
		processedURL := url
		shouldAdd := true

		// Handle trailing slashes first
		if backslash {
			noSlash := strings.TrimSuffix(url, "/")
			if url != noSlash {
				processedURL = noSlash // Always remove trailing slash
			}
		}

		// Handle HTTP/HTTPS duplicates (check the processed URL)
		if cleanHTTP && strings.HasPrefix(processedURL, "http://") {
			normalized := normalizeURLForComparison(processedURL)
			if httpsMap[normalized] {
				shouldAdd = false // Skip HTTP if HTTPS exists
			}
		}

		// Add to result if not already processed and should be added
		if shouldAdd && !urlMap[processedURL] {
			urlMap[processedURL] = true
			result = append(result, processedURL)
		}
	}

	return result
}

func convertToLowercase(urls []string) []string {
	if len(urls) == 0 {
		return []string{}
	}
	var result []string
	for _, url := range urls {
		result = append(result, strings.ToLower(url))
	}
	return result
}

func removeUnnecessaryCharacters(urls []string) []string {
	if len(urls) == 0 {
		return []string{}
	}
	var result []string
	for _, url := range urls {
		cleaned := strings.Trim(url, `"'!`)
		result = append(result, cleaned)
	}
	return result
}

// normalizeURLForComparison removes protocol and trailing slash for HTTP/HTTPS comparison
func normalizeURLForComparison(url string) string {
	// Remove protocol
	if strings.HasPrefix(url, "https://") {
		url = strings.TrimPrefix(url, "https://")
	} else if strings.HasPrefix(url, "http://") {
		url = strings.TrimPrefix(url, "http://")
	}
	
	// Remove trailing slash
	url = strings.TrimSuffix(url, "/")
	
	return url
}

func extractUniqueDomains(urls []string) []string {
	if len(urls) == 0 {
		return []string{}
	}
	
	domainMap := make(map[string]bool)
	var result []string
	
	for _, url := range urls {
		// Convert to lowercase first
		url = strings.ToLower(url)
		
		// Remove unnecessary characters
		url = strings.Trim(url, `"'!`)
		
		// Extract domain
		domain := extractDomain(url)
		if domain != "" && !domainMap[domain] {
			domainMap[domain] = true
			result = append(result, domain)
		}
	}
	
	return result
}

func extractDomain(url string) string {
	// Remove protocol
	if strings.HasPrefix(url, "http://") {
		url = strings.TrimPrefix(url, "http://")
	} else if strings.HasPrefix(url, "https://") {
		url = strings.TrimPrefix(url, "https://")
	}
	
	// Remove www. prefix if present
	if strings.HasPrefix(url, "www.") {
		url = strings.TrimPrefix(url, "www.")
	}
	
	// Get the domain part (before the first slash or path)
	if idx := strings.Index(url, "/"); idx != -1 {
		url = url[:idx]
	}
	
	// Remove trailing slash if present
	url = strings.TrimSuffix(url, "/")
	
	// Remove port if present (everything after the last colon)
	if idx := strings.LastIndex(url, ":"); idx != -1 {
		url = url[:idx]
	}
	
	return url
}


func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
} 