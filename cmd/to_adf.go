package cmd

import (
	"fmt"

	"github.com/ajbeck/goldmark-adf"
	"github.com/spf13/cobra"
)

var toAdfCmd = &cobra.Command{
	Use:   "to-adf [file]",
	Short: "Convert Markdown to ADF JSON",
	Long:  "Convert Markdown to ADF JSON. Reads from a file argument or stdin.",
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return convertToAdf(args)
	},
}

func convertToAdf(args []string) error {
	inputData, err := readInput(args)
	if err != nil {
		return err
	}

	adfOutput, err := adf.ConvertWithGFM(inputData)
	if err != nil {
		return fmt.Errorf("failed to convert Markdown to ADF: %w", err)
	}

	fmt.Println(string(adfOutput))
	return nil
}

func init() {
	rootCmd.AddCommand(toAdfCmd)
}
