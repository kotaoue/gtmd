package main

import (
	"flag"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"

	"golang.org/x/net/html"
)

var (
	source string
	mode   string
)

func main() {
	f := flag.String("url", "", "source url")
	m := flag.String("mode", "", "mode")
	flag.Parse()

	if *f != "" {
		source = *f
	} else if len(flag.Args()) > 0 {
		source = flag.Args()[0]
	} else {
		source = "https://pkg.go.dev/"
	}
	mode = *m

	if err := Main(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func Main() error {
	n, err := node(source)
	if err != nil {
		return err
	}

	t := pageTitle(n)
	if mode == "bookmeter" {
		t = extractBookmeterTitle(t)
	}

	if err := touch(source, t); err != nil {
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
	defer r.Body.Close()

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

func extractBookmeterTitle(title string) string {
	re := regexp.MustCompile(`『(.*)』`)
	matches := re.FindStringSubmatch(title)
	if len(matches) > 1 {
		return strings.TrimSpace(matches[1])
	}
	return strings.TrimSpace(title)
}

func touch(url, title string) error {
	fp, err := os.Create(fmt.Sprintf("%s.md", title))
	if err != nil {
		return err
	}
	defer fp.Close()

	_, err = fp.WriteString(fmt.Sprintf("# [%s](%s)", title, url))
	if err != nil {
		return err
	}

	return nil
}
