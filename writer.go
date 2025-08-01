package main

import (
	"fmt"
	"os"
)

func createMarkdownFile(url, title string) error {
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
