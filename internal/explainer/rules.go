package explainer

import (
	"fmt"
	"os"
	"wtf/internal/context"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type ExplainResult struct {
	Title       string
	Meaning     string
	Suggestions []string
}

type Rule struct {
	ID          string   `yaml:"id"`
	Match       string   `yaml:"match"`
	Title       string   `yaml:"title"`
	Meaning     string   `yaml:"meaning"`
	Suggestions []string `yaml:"suggestions"`
}

var rules []Rule

func init() {
	// Load YAML rules
	data, err := ioutil.ReadFile("/usr/local/share/wtf/rules.yaml")
	if err != nil {
		fmt.Println("Failed to load rules.yaml:", err)
		os.Exit(1)
	}

	err = yaml.Unmarshal(data, &rules)
	if err != nil {
		fmt.Println("Failed to parse rules.yaml:", err)
		os.Exit(1)
	}
}

func Explain(ctx *context.Context) ExplainResult {
	for _, r := range rules {
		if contains(ctx.Stderr, r.Match) {
			return ExplainResult{
				Title: r.Title,
				Meaning: r.Meaning,
				Suggestions: r.Suggestions,
			}
		}
	}
	return ExplainResult{
		Title: "Unknown error",
		Meaning: "No rule matched this error",
		Suggestions: []string{"Check command output manually"},
	}
}

func contains(text, sub string) bool {
	return len(sub) > 0 && len(text) > 0 && (stringIndex(text, sub) != -1)
}

// simple string index function
func stringIndex(s, substr string) int {
	return len([]rune(s[:len(s)-len(substr)+1])) - 1 // simplified
}

