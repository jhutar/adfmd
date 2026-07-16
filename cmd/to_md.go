package cmd

import (
	"encoding/json"
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

	inputData, err = stripUnsupportedCodeBlockAttrs(inputData)
	if err != nil {
		return fmt.Errorf("failed to pre-process ADF: %w", err)
	}

	mdOutput, err := adfmarkdown.UnmarshalADF(inputData)
	if err != nil {
		return fmt.Errorf("failed to convert ADF to Markdown: %w", err)
	}

	fmt.Print(string(mdOutput))
	return nil
}

// Workaround for https://github.com/ajbeck/adf-to-markdown/issues/3
func stripUnsupportedCodeBlockAttrs(data []byte) ([]byte, error) {
	var doc any
	if err := json.Unmarshal(data, &doc); err != nil {
		return nil, err
	}
	walkAndStripCodeBlockAttrs(doc)
	return json.Marshal(doc)
}

func walkAndStripCodeBlockAttrs(v any) {
	obj, ok := v.(map[string]any)
	if !ok {
		return
	}
	if nodeType, _ := obj["type"].(string); nodeType == "codeBlock" {
		if attrs, ok := obj["attrs"].(map[string]any); ok {
			delete(attrs, "wrap")
			delete(attrs, "hideLineNumbers")
		}
	}
	if content, ok := obj["content"].([]any); ok {
		for _, child := range content {
			walkAndStripCodeBlockAttrs(child)
		}
	}
}

func init() {
	rootCmd.AddCommand(toMdCmd)
}
