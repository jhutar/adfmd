package cmd

import (
	"fmt"
	"os"

	"github.com/ajbeck/goldmark-adf"
	"github.com/spf13/cobra"
)

var toAdfCmd = &cobra.Command{
	Use:   "to-adf [input.md]",
	Short: "Convert Markdown to ADF JSON",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		inputFile := args[0]

		inputData, err := os.ReadFile(inputFile)
		if err != nil {
			return fmt.Errorf("failed to read input file: %w", err)
		}

		adfOutput, err := adf.ConvertWithGFM(inputData)
		if err != nil {
			return fmt.Errorf("failed to convert Markdown to ADF: %w", err)
		}

		fmt.Println(string(adfOutput))
		return nil
	},
}

func init() {
	rootCmd.AddCommand(toAdfCmd)
}
