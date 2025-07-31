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
)

var rootCmd = &cobra.Command{
	Use:   "cleanurl",
	Short: "Clean and deduplicate URLs from stdin",
	Long: `CleanURL is a command-line tool that processes URLs from stdin and applies various cleaning operations.

Features:
- Remove unnecessary characters (quotes and exclamation marks) from URLs
- Remove HTTP duplicates when HTTPS version exists
- Remove trailing slashes to deduplicate URLs
- Output cleaned URLs to stdout

Examples:
  echo "https://example.com/" | cleanurl
  cat urls.txt | cleanurl --no-characters
  echo "http://example.com" | cleanurl --no-clean-http`,
	Run: runCleanURL,
}

func init() {
	// Set default values to true (on by default)
	rootCmd.Flags().BoolVar(&characters, "characters", true, "Remove unnecessary characters (quotes and exclamation marks) from URLs")
	rootCmd.Flags().BoolVar(&cleanHTTP, "clean-http", true, "Remove HTTP duplicates when HTTPS version exists")
	rootCmd.Flags().BoolVar(&backslash, "backslash", true, "Remove trailing slashes to deduplicate URLs")
	
	// Add negative flags for convenience
	rootCmd.Flags().Bool("no-characters", false, "Disable character cleaning")
	rootCmd.Flags().Bool("no-clean-http", false, "Disable HTTP cleaning")
	rootCmd.Flags().Bool("no-backslash", false, "Disable backslash cleaning")
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

	// Read URLs from stdin
	urls := readURLsFromStdin()
	
	// Apply cleaning operations
	cleanedURLs := cleanURLs(urls)
	
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

	// Step 1: Remove unnecessary characters
	if characters {
		urls = removeUnnecessaryCharacters(urls)
	}

	// Step 2: Create maps for tracking
	urlMap := make(map[string]bool)
	httpsMap := make(map[string]bool)
	noSlashMap := make(map[string]bool)
	var result []string

	// First pass: collect HTTPS URLs and URLs without trailing slashes
	for _, url := range urls {
		// Track HTTPS URLs
		if strings.HasPrefix(url, "https://") {
			httpsMap[strings.Replace(url, "https://", "http://", 1)] = true
		}
		
		// Track URLs without trailing slashes
		noSlash := strings.TrimSuffix(url, "/")
		if url != noSlash {
			noSlashMap[noSlash] = true
		}
	}

	// Second pass: process URLs
	for _, url := range urls {
		processedURL := url
		shouldAdd := true

		// Handle HTTP/HTTPS duplicates
		if cleanHTTP && strings.HasPrefix(url, "http://") {
			if httpsMap[url] {
				shouldAdd = false // Skip HTTP if HTTPS exists
			}
		}

		// Handle trailing slashes
		if backslash && shouldAdd {
			noSlash := strings.TrimSuffix(url, "/")
			if url != noSlash && noSlashMap[noSlash] {
				shouldAdd = false // Skip URL with trailing slash if version without exists
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



func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
} 