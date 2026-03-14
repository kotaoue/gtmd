package main

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func TestExtractTitle(t *testing.T) {
	htmlString := `<html><head><title>Test Title</title></head><body></body></html>`
	doc, err := html.Parse(strings.NewReader(htmlString))
	if err != nil {
		t.Fatalf("failed to parse html: %v", err)
	}

	title := extractTitle(doc)
	expected := "Test Title"
	if title != expected {
		t.Errorf("expected title %q, but got %q", expected, title)
	}
}

func TestExtractBookmeterTitle(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "Bookmeter title",
			input: "『The Go Programming Language』の感想 - 読書メーター",
			want:  "The Go Programming Language",
		},
		{
			name:  "Bookmeter title with space",
			input: "『   Another Book   』の感想 - 読書メーター",
			want:  "Another Book",
		},
		{
			name:  "No brackets",
			input: "Just a regular title",
			want:  "Just a regular title",
		},
		{
			name:  "Empty input",
			input: "",
			want:  "",
		},
		{
			name:  "Only brackets",
			input: "『』",
			want:  "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := extractBookmeterTitle(tt.input)
			if got != tt.want {
				t.Errorf("extractBookmeterTitle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDetectMode(t *testing.T) {
	tests := []struct {
		name   string
		source string
		want   string
	}{
		{
			name:   "Bookmeter URL",
			source: "https://bookmeter.com/books/556977",
			want:   "bookmeter",
		},
		{
			name:   "Bookmeter URL with subdomain",
			source: "https://www.bookmeter.com/books/556977",
			want:   "bookmeter",
		},
		{
			name:   "Non-bookmeter URL",
			source: "https://example.com/page",
			want:   "",
		},
		{
			name:   "pkg.go.dev URL",
			source: "https://pkg.go.dev/",
			want:   "",
		},
		{
			name:   "Invalid URL",
			source: "://invalid",
			want:   "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := detectMode(tt.source)
			if got != tt.want {
				t.Errorf("detectMode(%q) = %q, want %q", tt.source, got, tt.want)
			}
		})
	}
}

func TestOutputManualMode(t *testing.T) {
	title := "『Manual Mode Book』の感想 - 読書メーター"
	source := "https://bookmeter.com/books/123"
	filename := fmt.Sprintf("%s.md", title)
	defer os.Remove(filename)

	err := output(source, title, "", "manual")
	if err != nil {
		t.Errorf("output() error = %v, want nil", err)
	}

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		t.Errorf("output() with manual mode did not create file %s", filename)
	}

	content, err := os.ReadFile(filename)
	if err != nil {
		t.Fatalf("failed to read created file: %v", err)
	}

	expectedContent := fmt.Sprintf("# [%s](%s)", title, source)
	if string(content) != expectedContent {
		t.Errorf("expected content %q, but got %q", expectedContent, string(content))
	}
}

func TestCreateMarkdownFile(t *testing.T) {
	title := "Test File"
	url := "http://example.com"
	filename := fmt.Sprintf("%s.md", title)

	defer os.Remove(filename)

	err := createMarkdownFile(url, title)
	if err != nil {
		t.Fatalf("createMarkdownFile() failed: %v", err)
	}

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		t.Fatalf("createMarkdownFile() did not create file: %s", filename)
	}

	content, err := os.ReadFile(filename)
	if err != nil {
		t.Fatalf("failed to read created file: %v", err)
	}

	expectedContent := fmt.Sprintf("# [%s](%s)", title, url)
	if string(content) != expectedContent {
		t.Errorf("expected content %q, but got %q", expectedContent, string(content))
	}
}

func TestResolveSource(t *testing.T) {
	tests := []struct {
		name     string
		urlFlag  string
		args     []string
		expected string
	}{
		{
			name:     "URL flag provided",
			urlFlag:  "https://example.com",
			args:     []string{},
			expected: "https://example.com",
		},
		{
			name:     "Args provided",
			urlFlag:  "",
			args:     []string{"https://golang.org"},
			expected: "https://golang.org",
		},
		{
			name:     "Default URL when empty",
			urlFlag:  "",
			args:     []string{},
			expected: "https://pkg.go.dev/",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := resolveSource(tt.urlFlag, tt.args)
			if got != tt.expected {
				t.Errorf("resolveSource() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestOutputDefaultMode(t *testing.T) {
	title := "Test Output Default"
	source := "https://example.com"
	filename := fmt.Sprintf("%s.md", title)
	defer os.Remove(filename)

	err := output(source, title, "", "")
	if err != nil {
		t.Errorf("output() error = %v, want nil", err)
	}

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		t.Errorf("output() with default mode did not create file %s", filename)
	}
}

func TestOutputBookmeterMode(t *testing.T) {
	title := "『Test Book』の感想 - 読書メーター"
	source := "https://bookmeter.com/books/123"
	filename := fmt.Sprintf("%s.md", extractBookmeterTitle(title))
	defer os.Remove(filename)

	err := output(source, title, "", "bookmeter")
	if err != nil {
		t.Errorf("output() error = %v, want nil", err)
	}

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		t.Errorf("output() with bookmeter mode did not create file %s", filename)
	}
}

func TestOutput(t *testing.T) {
	tests := []struct {
		name       string
		source     string
		title      string
		format     string
		sourceType string
		want       error
	}{
		{
			name:       "Link format",
			source:     "https://example.com",
			title:      "Example Title",
			format:     "link",
			sourceType: "",
			want:       nil,
		},
		{
			name:       "Bookmeter source with link format",
			source:     "https://bookmeter.com/books/123",
			title:      "『Test Book』の感想",
			format:     "link",
			sourceType: "bookmeter",
			want:       nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.format == "link" {
				err := output(tt.source, tt.title, tt.format, tt.sourceType)
				if err != tt.want {
					t.Errorf("output() error = %v, want %v", err, tt.want)
				}
			}
		})
	}
}
