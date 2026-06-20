package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/ajbeck/goldmark-adf"
	"github.com/google/uuid"
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

	fixed, err := fixAdf(adfOutput)
	if err != nil {
		return fmt.Errorf("failed to post-process ADF: %w", err)
	}

	fmt.Println(string(fixed))
	return nil
}

func fixAdf(data []byte) ([]byte, error) {
	var doc any
	if err := json.Unmarshal(data, &doc); err != nil {
		return nil, err
	}
	walkAndFix(doc)
	return json.Marshal(doc)
}

func walkAndFix(v any) {
	obj, ok := v.(map[string]any)
	if !ok {
		return
	}

	nodeType, _ := obj["type"].(string)

	fillEmptyLocalId(obj, nodeType)
	stripCheckboxPrefix(obj, nodeType)
	dropMarksConflictingWithCode(obj, nodeType)

	if content, ok := obj["content"].([]any); ok {
		for _, child := range content {
			walkAndFix(child)
		}
	}
}

func fillEmptyLocalId(obj map[string]any, nodeType string) {
	switch nodeType {
	case "taskList", "taskItem", "decisionList", "decisionItem":
	default:
		return
	}
	if attrs, ok := obj["attrs"].(map[string]any); ok {
		if id, _ := attrs["localId"].(string); id == "" {
			attrs["localId"] = uuid.New().String()
		}
	}
}

func stripCheckboxPrefix(obj map[string]any, nodeType string) {
	if nodeType != "taskItem" {
		return
	}
	content, ok := obj["content"].([]any)
	if !ok {
		return
	}
	for _, child := range content {
		para, ok := child.(map[string]any)
		if !ok {
			continue
		}
		if t, _ := para["type"].(string); t != "paragraph" {
			continue
		}
		texts, ok := para["content"].([]any)
		if !ok || len(texts) == 0 {
			continue
		}
		first, ok := texts[0].(map[string]any)
		if !ok {
			continue
		}
		if t, _ := first["type"].(string); t != "text" {
			continue
		}
		text, _ := first["text"].(string)
		if text == "[ ] " || text == "[x] " {
			para["content"] = texts[1:]
		}
	}
}

func dropMarksConflictingWithCode(obj map[string]any, nodeType string) {
	if nodeType != "text" {
		return
	}
	marks, ok := obj["marks"].([]any)
	if !ok || len(marks) <= 1 {
		return
	}
	hasCode := false
	for _, m := range marks {
		if mark, ok := m.(map[string]any); ok {
			if t, _ := mark["type"].(string); t == "code" {
				hasCode = true
				break
			}
		}
	}
	if !hasCode {
		return
	}
	kept := make([]any, 0, 1)
	for _, m := range marks {
		if mark, ok := m.(map[string]any); ok {
			if t, _ := mark["type"].(string); t != "code" {
				kept = append(kept, m)
			}
		}
	}
	obj["marks"] = kept
}

func init() {
	rootCmd.AddCommand(toAdfCmd)
}
