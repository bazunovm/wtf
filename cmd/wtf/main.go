package main

import (
	"fmt"
	"os"

	"wtf/internal/context"
	"wtf/internal/explainer"
	"wtf/internal/runner"
	"wtf/internal/ai"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: wtf <command>")
		return
	}

	cmd := os.Args[1:]

	// Build context
	ctx := context.BuildContext(cmd)

	// Run command
	stdout, stderr, exitCode := runner.RunCommand(cmd)
	ctx.Stdout = stdout
	ctx.Stderr = stderr
	ctx.ExitCode = exitCode

	// Try rules first
	result := explainer.Explain(ctx)

	// If no rule matched → use AI
	if result.Title == "" {
		aiResult, err := ai.Explain(ctx)
		if err != nil {
			result = explainer.ExplainResult{
				Title:   "Unknown error",
				Meaning: "No rule matched this error and AI failed",
				Suggestions: []string{
					"Check command output manually",
				},
			}
		} else {
			result = aiResult
		}
	}

	// Print result
	fmt.Printf("\nError: %s\n", result.Title)
	fmt.Printf("Meaning: %s\n", result.Meaning)
	fmt.Println("Suggestions:")
	for _, s := range result.Suggestions {
		fmt.Printf("  - %s\n", s)
	}
}

