package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	urlFlag  string
	modeFlag string
)

var rootCmd = &cobra.Command{
	Use:   "gtmd [URL]",
	Short: "Get the title tag and make markdown",
	Long:  `gtmd fetches a web page, extracts its title, and either creates a markdown file or outputs a markdown link.`,
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		source := resolveSource(urlFlag, args)
		return Main(source, modeFlag)
	},
}

func init() {
	rootCmd.Flags().StringVarP(&urlFlag, "url", "u", "", "source url")
	rootCmd.Flags().StringVarP(&modeFlag, "mode", "m", "", "mode (link, bookmeter)")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func Main(source, mode string) error {
	n, err := fetchPage(source)
	if err != nil {
		return err
	}

	return output(source, extractTitle(n), mode)
}

func resolveSource(urlFlag string, args []string) string {
	switch {
	case urlFlag != "":
		return urlFlag
	case len(args) > 0:
		return args[0]
	default:
		return "https://pkg.go.dev/"
	}
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
