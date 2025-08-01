package main

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestCreateMarkdownFileFunction(t *testing.T) {
	tests := []struct {
		name     string
		url      string
		title    string
		wantErr  bool
	}{
		{
			name:    "Valid input",
			url:     "https://example.com",
			title:   "Example Title",
			wantErr: false,
		},
		{
			name:    "Title with spaces",
			url:     "https://golang.org",
			title:   "Go Programming Language",
			wantErr: false,
		},
		{
			name:    "Special characters in title",
			url:     "https://test.com",
			title:   "Test: Special & Characters",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filename := fmt.Sprintf("%s.md", tt.title)
			defer os.Remove(filename)

			err := createMarkdownFile(tt.url, tt.title)
			if (err != nil) != tt.wantErr {
				t.Errorf("createMarkdownFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				if _, err := os.Stat(filename); os.IsNotExist(err) {
					t.Errorf("createMarkdownFile() did not create file: %s", filename)
					return
				}

				content, err := os.ReadFile(filename)
				if err != nil {
					t.Errorf("failed to read created file: %v", err)
					return
				}

				expectedContent := fmt.Sprintf("# [%s](%s)", tt.title, tt.url)
				if string(content) != expectedContent {
					t.Errorf("expected content %q, but got %q", expectedContent, string(content))
				}
			}
		})
	}
}

func TestCreateMarkdownFileInvalidPath(t *testing.T) {
	invalidPath := filepath.Join("nonexistent", "directory", "file.md")
	
	err := createMarkdownFile("https://example.com", invalidPath)
	if err == nil {
		t.Error("createMarkdownFile() should fail with invalid path")
		os.Remove(fmt.Sprintf("%s.md", invalidPath))
	}
}