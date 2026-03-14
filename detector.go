package main

import (
	"net/url"
	"strings"
)

// detectMode returns a mode string inferred from the URL domain.
// It returns an empty string when no specific mode is detected.
func detectMode(source string) string {
	u, err := url.Parse(source)
	if err != nil {
		return ""
	}

	switch {
	case strings.Contains(u.Host, "bookmeter.com"):
		return "bookmeter"
	default:
		return ""
	}
}
