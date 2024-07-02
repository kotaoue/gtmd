package main

import (
	"flag"
	"fmt"
	"net/http"
	"net/url"
	"os"

	"golang.org/x/net/html"
)

var (
	target string
)

func init() {
	flag.StringVar(&target, "url", "https://pkg.go.dev/", "source url")
	flag.Parse()
}

func main() {
	if err := Main(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func Main() error {
	n, err := node(target)
	if err != nil {
		return err
	}

	t := pageTitle(n)

	if err := touch(target, t); err != nil {
		return err
	}

	return nil
}

func node(s string) (*html.Node, error) {
	u, err := url.Parse(s)
	if err != nil {
		return nil, fmt.Errorf("can't parse url")
	}

	r, err := http.Get(u.String())
	if err != nil {
		return nil, fmt.Errorf("can't get page")
	}

	page, err := html.Parse(r.Body)
	if err != nil {
		return nil, fmt.Errorf("can't parse page")
	}
	return page, nil
}

func pageTitle(n *html.Node) (title string) {
	if n.Type == html.ElementNode && n.Data == "title" {
		return n.FirstChild.Data
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		title = pageTitle(c)
		if title != "" {
			break
		}
	}
	return title
}

func touch(url, title string) error {
	fp, err := os.Create(fmt.Sprintf("%s.md", title))
	if err != nil {
		return err
	}
	defer fp.Close()

	fp.WriteString(fmt.Sprintf("# [%s](%s)", title, target))

	return nil
}
