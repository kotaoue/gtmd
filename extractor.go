package main

import (
	"fmt"
	"net/http"
	"net/url"

	"golang.org/x/net/html"
)

func fetchPage(source string) (*html.Node, error) {
	u, err := url.Parse(source)
	if err != nil {
		return nil, fmt.Errorf("can't parse url")
	}

	r, err := http.Get(u.String())
	if err != nil {
		return nil, fmt.Errorf("can't get page")
	}
	defer r.Body.Close()

	page, err := html.Parse(r.Body)
	if err != nil {
		return nil, fmt.Errorf("can't parse page")
	}
	return page, nil
}

func extractTitle(n *html.Node) (title string) {
	if n == nil {
		return ""
	}
	if n.Type == html.ElementNode && n.Data == "title" {
		if n.FirstChild != nil {
			return n.FirstChild.Data
		}
		return ""
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		title = extractTitle(c)
		if title != "" {
			break
		}
	}
	return title
}
