package main

import (
	"fmt"
	"os"

	"github.com/atotto/clipboard"
	"github.com/spf13/cobra"
)

var (
	urlFlag        string
	formatFlag     string
	sourceTypeFlag string
)

var rootCmd = &cobra.Command{
	Use:   "gtmd [URL]",
	Short: "Get the title tag and make markdown",
	Long:  `gtmd fetches a web page, extracts its title, and either creates a markdown file or outputs a markdown link.`,
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		source := resolveSource(urlFlag, args)
		return Main(source, formatFlag, sourceTypeFlag)
	},
}

func init() {
	rootCmd.Flags().StringVarP(&urlFlag, "url", "u", "", "source url")
	rootCmd.Flags().StringVarP(&formatFlag, "format", "f", "", "output format (link, clipboard). Creates a markdown file if not set.")
	rootCmd.Flags().StringVarP(&sourceTypeFlag, "source", "s", "", "source type (bookmeter, manual). Auto-detected from URL when not set.")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func Main(source, format, sourceType string) error {
	if sourceType == "" {
		sourceType = detectMode(source)
	}

	n, err := fetchPage(source)
	if err != nil {
		return err
	}

	return output(source, extractTitle(n), format, sourceType)
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

func output(source, title, format, sourceType string) error {
	if sourceType == "bookmeter" {
		title = extractBookmeterTitle(title)
	}

	switch format {
	case "link":
		fmt.Printf("[%s](%s)\n", title, source)
		return nil
	case "clipboard":
		link := fmt.Sprintf("[%s](%s)", title, source)
		err := clipboard.WriteAll(link)
		if err != nil {
			return fmt.Errorf("failed to copy to clipboard: %v", err)
		}
		fmt.Printf("Copied to clipboard: %s\n", link)
		return nil
	default:
		return createMarkdownFile(source, title)
	}
}
