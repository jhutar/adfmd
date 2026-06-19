package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "adfmd [file]",
	Short: "adfmd converts Jira ADF to Markdown and vice versa",
	Long: `A fast and flexible CLI tool for bidirectional conversion between
Atlassian Document Format (ADF) and Markdown.

When given a file argument, the conversion direction is detected from
the file extension: .md converts to ADF, .adf/.json converts to Markdown.

Use the to-md or to-adf subcommands for explicit control, with either
a file argument or piped stdin.`,
	Args: cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return cmd.Help()
		}
		ext := strings.ToLower(filepath.Ext(args[0]))
		switch ext {
		case ".md":
			return convertToAdf(args)
		case ".adf", ".json":
			return convertToMd(args)
		default:
			return fmt.Errorf("unsupported file extension %q (expected .md, .adf, or .json)", ext)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
