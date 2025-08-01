package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func TestExtractTitleFromNode(t *testing.T) {
	tests := []struct {
		name     string
		htmlStr  string
		expected string
	}{
		{
			name:     "Simple title",
			htmlStr:  `<html><head><title>Simple Title</title></head></html>`,
			expected: "Simple Title",
		},
		{
			name:     "Nested structure",
			htmlStr:  `<html><head><meta charset="utf-8"><title>Nested Title</title></head><body></body></html>`,
			expected: "Nested Title",
		},
		{
			name:     "No title tag",
			htmlStr:  `<html><head></head><body></body></html>`,
			expected: "",
		},
		{
			name:     "Empty title",
			htmlStr:  `<html><head><title></title></head></html>`,
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			doc, err := html.Parse(strings.NewReader(tt.htmlStr))
			if err != nil {
				t.Fatalf("failed to parse HTML: %v", err)
			}

			got := extractTitle(doc)
			if got != tt.expected {
				t.Errorf("extractTitle() = %q, want %q", got, tt.expected)
			}
		})
	}
}

func TestFetchPage(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`<html><head><title>Test Server</title></head><body></body></html>`))
	}))
	defer server.Close()

	node, err := fetchPage(server.URL)
	if err != nil {
		t.Fatalf("fetchPage() failed: %v", err)
	}

	if node == nil {
		t.Fatal("fetchPage() returned nil node")
	}

	title := extractTitle(node)
	if title != "Test Server" {
		t.Errorf("expected title 'Test Server', got %q", title)
	}

	_, err = fetchPage("invalid-url")
	if err == nil {
		t.Error("fetchPage() should fail with invalid URL")
	}
}
