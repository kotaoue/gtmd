package main

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func TestPageTitle(t *testing.T) {
	htmlString := `<html><head><title>Test Title</title></head><body></body></html>`
	doc, err := html.Parse(strings.NewReader(htmlString))
	if err != nil {
		t.Fatalf("failed to parse html: %v", err)
	}

	title := pageTitle(doc)
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

func TestTouch(t *testing.T) {
	title := "Test File"
	url := "http://example.com"
	filename := fmt.Sprintf("%s.md", title)

	// Clean up the file after the test
	defer os.Remove(filename)

	err := touch(url, title)
	if err != nil {
		t.Fatalf("touch() failed: %v", err)
	}

	// Check if file exists
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		t.Fatalf("touch() did not create file: %s", filename)
	}

	// Check file content
	content, err := os.ReadFile(filename)
	if err != nil {
		t.Fatalf("failed to read created file: %v", err)
	}

	expectedContent := fmt.Sprintf("# [%s](%s)", title, url)
	if string(content) != expectedContent {
		t.Errorf("expected content %q, but got %q", expectedContent, string(content))
	}
}