package main

import (
	"regexp"
	"strings"
)

var (
	bookmeterTitleRe = regexp.MustCompile(`『(.*)』`)
)

func extractBookmeterTitle(title string) string {
	matches := bookmeterTitleRe.FindStringSubmatch(title)
	if len(matches) > 1 {
		return strings.TrimSpace(matches[1])
	}
	return strings.TrimSpace(title)
}
