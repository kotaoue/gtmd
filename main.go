package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	source, mode := setup()

	if err := Main(source, mode); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func setup() (string, string) {
	f := flag.String("url", "", "source url")
	m := flag.String("mode", "", "mode")
	flag.Parse()

	return resolveSource(*f), *m
}

func resolveSource(urlFlag string) string {
	switch {
	case urlFlag != "":
		return urlFlag
	case len(flag.Args()) > 0:
		return flag.Args()[0]
	default:
		return "https://pkg.go.dev/"
	}
}

func Main(source, mode string) error {
	n, err := fetchPage(source)
	if err != nil {
		return err
	}

	return output(source, extractTitle(n), mode)
}


func output(source, title, mode string) error {
	if mode == "bookmeter" {
		title = extractBookmeterTitle(title)
	}

	switch mode {
	case "link":
		fmt.Printf("[%s](%s)\n", title, source)
		return nil
	default:
		return createMarkdownFile(source, title)
	}
}

