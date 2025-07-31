package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertToLowercase(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected []string
	}{
		{
			name:     "Convert mixed case URLs",
			input:    []string{"HTTPS://EXAMPLE.COM", "Http://Test.Com", "https://lowercase.com"},
			expected: []string{"https://example.com", "http://test.com", "https://lowercase.com"},
		},
		{
			name:     "Convert uppercase URLs",
			input:    []string{"HTTPS://EXAMPLE.COM", "HTTP://TEST.COM"},
			expected: []string{"https://example.com", "http://test.com"},
		},
		{
			name:     "Already lowercase URLs",
			input:    []string{"https://example.com", "http://test.com"},
			expected: []string{"https://example.com", "http://test.com"},
		},
		{
			name:     "Empty input",
			input:    []string{},
			expected: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := convertToLowercase(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestRemoveUnnecessaryCharacters(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected []string
	}{
		{
			name:     "Remove single quotes",
			input:    []string{"'https://example.com'", "https://test.com"},
			expected: []string{"https://example.com", "https://test.com"},
		},
		{
			name:     "Remove double quotes",
			input:    []string{`"https://example.com"`, "https://test.com"},
			expected: []string{"https://example.com", "https://test.com"},
		},
		{
			name:     "Remove mixed quotes",
			input:    []string{`"https://example.com"`, "'https://test.com'"},
			expected: []string{"https://example.com", "https://test.com"},
		},
		{
			name:     "Remove exclamation marks",
			input:    []string{"!https://example.com!", "https://test.com"},
			expected: []string{"https://example.com", "https://test.com"},
		},
		{
			name:     "Remove mixed characters",
			input:    []string{`"https://example.com"`, "'https://test.com'", "!https://another.com!"},
			expected: []string{"https://example.com", "https://test.com", "https://another.com"},
		},
		{
			name:     "No quotes to remove",
			input:    []string{"https://example.com", "https://test.com"},
			expected: []string{"https://example.com", "https://test.com"},
		},
		{
			name:     "Empty input",
			input:    []string{},
			expected: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := removeUnnecessaryCharacters(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}



func TestCleanURLs(t *testing.T) {
	tests := []struct {
		name        string
		input       []string
		characters  bool
		cleanHTTP   bool
		backslash   bool
		lower       bool
		expected    []string
	}{
		{
			name:        "All features enabled",
			input:       []string{`"https://example.com/"`, "http://example.com", "https://example.com"},
			characters:  true,
			cleanHTTP:   true,
			backslash:   true,
			lower:       true,
			expected:    []string{"https://example.com"},
		},
		{
			name:        "Characters disabled",
			input:       []string{`"https://example.com"`, "https://test.com"},
			characters:  false,
			cleanHTTP:   true,
			backslash:   true,
			lower:       true,
			expected:    []string{`"https://example.com"`, "https://test.com"},
		},
		{
			name:        "HTTP cleaning disabled",
			input:       []string{"http://example.com", "https://example.com"},
			characters:  true,
			cleanHTTP:   false,
			backslash:   true,
			lower:       true,
			expected:    []string{"http://example.com", "https://example.com"},
		},
		{
			name:        "Backslash cleaning disabled",
			input:       []string{"https://example.com/", "https://example.com"},
			characters:  true,
			cleanHTTP:   true,
			backslash:   false,
			lower:       true,
			expected:    []string{"https://example.com/", "https://example.com"},
		},
		{
			name:        "Empty input",
			input:       []string{},
			characters:  true,
			cleanHTTP:   true,
			backslash:   true,
			lower:       true,
			expected:    []string{},
		},
		{
			name:        "Complex deduplication",
			input:       []string{`"http://example.com/"`, "https://example.com", "'https://test.com/'", "https://test.com"},
			characters:  true,
			cleanHTTP:   true,
			backslash:   true,
			lower:       true,
			expected:    []string{"https://example.com", "https://test.com"},
		},
		{
			name:        "Lowercase conversion with mixed case",
			input:       []string{"HTTPS://EXAMPLE.COM", "Http://Test.Com", "!https://UPPERCASE.com!"},
			characters:  true,
			cleanHTTP:   true,
			backslash:   true,
			lower:       true,
			expected:    []string{"https://example.com", "http://test.com", "https://uppercase.com"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Set global flags
			characters = tt.characters
			cleanHTTP = tt.cleanHTTP
			backslash = tt.backslash
			lower = tt.lower

			result := cleanURLs(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestReadURLsFromStdin(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:     "Single URL",
			input:    "https://example.com\n",
			expected: []string{"https://example.com"},
		},
		{
			name:     "Multiple URLs",
			input:    "https://example.com\nhttps://test.com\n",
			expected: []string{"https://example.com", "https://test.com"},
		},
		{
			name:     "URLs with whitespace",
			input:    "  https://example.com  \n  https://test.com  \n",
			expected: []string{"https://example.com", "https://test.com"},
		},
		{
			name:     "Empty lines",
			input:    "https://example.com\n\nhttps://test.com\n",
			expected: []string{"https://example.com", "https://test.com"},
		},
		{
			name:     "Empty input",
			input:    "",
			expected: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a temporary file with test input
			tmpfile, err := os.CreateTemp("", "test")
			if err != nil {
				t.Fatal(err)
			}
			defer os.Remove(tmpfile.Name())

			if _, err := tmpfile.WriteString(tt.input); err != nil {
				t.Fatal(err)
			}
			if err := tmpfile.Close(); err != nil {
				t.Fatal(err)
			}

			// Redirect stdin to the temporary file
			oldStdin := os.Stdin
			defer func() { os.Stdin = oldStdin }()

			file, err := os.Open(tmpfile.Name())
			if err != nil {
				t.Fatal(err)
			}
			os.Stdin = file
			defer file.Close()

			result := readURLsFromStdin()
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestIntegration(t *testing.T) {
	// Test the complete pipeline with real input
	input := `"https://example.com/"
http://example.com
https://example.com
'https://test.com/'
https://test.com
https://unique.com`

	expected := []string{
		"https://example.com",
		"https://test.com",
		"https://unique.com",
	}

	// Create a temporary file with test input
	tmpfile, err := os.CreateTemp("", "test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	if _, err := tmpfile.WriteString(input); err != nil {
		t.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		t.Fatal(err)
	}

	// Redirect stdin to the temporary file
	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }()

	file, err := os.Open(tmpfile.Name())
	if err != nil {
		t.Fatal(err)
	}
	os.Stdin = file
	defer file.Close()

	// Set all features enabled
	characters = true
	cleanHTTP = true
	backslash = true
	lower = true

	urls := readURLsFromStdin()
	result := cleanURLs(urls)

	assert.Equal(t, expected, result)
}

func TestCommandLineFlags(t *testing.T) {
	// Test that the command can be created without errors
	assert.NotNil(t, rootCmd)
	
	// Test that our custom flags exist
	charactersFlag := rootCmd.Flags().Lookup("characters")
	assert.NotNil(t, charactersFlag)
	
	cleanHTTPFlag := rootCmd.Flags().Lookup("clean-http")
	assert.NotNil(t, cleanHTTPFlag)
	
	backslashFlag := rootCmd.Flags().Lookup("backslash")
	assert.NotNil(t, backslashFlag)
	
	lowerFlag := rootCmd.Flags().Lookup("lower")
	assert.NotNil(t, lowerFlag)
	
	// Test that negative flags exist
	noCharactersFlag := rootCmd.Flags().Lookup("no-characters")
	assert.NotNil(t, noCharactersFlag)
	
	noCleanHTTPFlag := rootCmd.Flags().Lookup("no-clean-http")
	assert.NotNil(t, noCleanHTTPFlag)
	
	noBackslashFlag := rootCmd.Flags().Lookup("no-backslash")
	assert.NotNil(t, noBackslashFlag)
	
	noLowerFlag := rootCmd.Flags().Lookup("no-lower")
	assert.NotNil(t, noLowerFlag)
} 