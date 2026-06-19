package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "adfmd",
	Short: "adfmd converts Jira ADF to Markdown and vice versa",
	Long:  `A fast and flexible CLI tool for bidirectional conversion between Atlassian Document Format (ADF) and Markdown.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
