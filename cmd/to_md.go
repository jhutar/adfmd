package cmd

import (
	"fmt"
	"os"

	adfmarkdown "github.com/ajbeck/adf-to-markdown"
	"github.com/spf13/cobra"
)

var toMdCmd = &cobra.Command{
	Use:   "to-md [input.json]",
	Short: "Convert ADF JSON to Markdown",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		inputFile := args[0]

		inputData, err := os.ReadFile(inputFile)
		if err != nil {
			return fmt.Errorf("failed to read input file: %w", err)
		}

		mdOutput, err := adfmarkdown.UnmarshalADF(inputData)
		if err != nil {
			return fmt.Errorf("failed to convert ADF to Markdown: %w", err)
		}

		fmt.Print(string(mdOutput))
		return nil
	},
}

func init() {
	rootCmd.AddCommand(toMdCmd)
}
