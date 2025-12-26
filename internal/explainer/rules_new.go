package explainer

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
	"wtf/internal/context"
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
	paths := []string{
		"./rules.yaml",
		"/usr/local/share/wtf/rules.yaml",
		filepath.Join(os.Getenv("HOME"), ".config/wtf/rules.yaml"),
	}

	var data []byte
	var err error

	for _, p := range paths {
		data, err = os.ReadFile(p)
		if err == nil {
			break
		}
	}

	if err != nil {
		fmt.Println("Failed to load rules.yaml")
		os.Exit(1)
	}

	if err := yaml.Unmarshal(data, &rules); err != nil {
		fmt.Println("Failed to parse rules.yaml:", err)
		os.Exit(1)
	}
}

func Explain(ctx *context.Context) ExplainResult {
	for _, r := range rules {
		if r.Match == "" {
			continue
		}
		if strings.Contains(ctx.Stderr, r.Match) ||
		strings.Contains(ctx.Stdout, r.Match) {
			return ExplainResult{
				Title:       r.Title,
				Meaning:     r.Meaning,
				Suggestions: r.Suggestions,
			}
		}
	}

	return ExplainResult{
		Title:   "",
		Meaning: "",
	}
}

