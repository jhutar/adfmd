package cmd

import (
	"fmt"

	adfmarkdown "github.com/ajbeck/adf-to-markdown"
	"github.com/spf13/cobra"
)

var toMdCmd = &cobra.Command{
	Use:   "to-md [file]",
	Short: "Convert ADF JSON to Markdown",
	Long:  "Convert ADF JSON to Markdown. Reads from a file argument or stdin.",
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return convertToMd(args)
	},
}

func convertToMd(args []string) error {
	inputData, err := readInput(args)
	if err != nil {
		return err
	}

	mdOutput, err := adfmarkdown.UnmarshalADF(inputData)
	if err != nil {
		return fmt.Errorf("failed to convert ADF to Markdown: %w", err)
	}

	fmt.Print(string(mdOutput))
	return nil
}

func init() {
	rootCmd.AddCommand(toMdCmd)
}
