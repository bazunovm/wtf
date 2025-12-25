package main

import (
	"fmt"
	"os"

	"wtf/internal/context"
	"wtf/internal/runner"
	"wtf/internal/explainer"
	// "wtf/internal/memory"
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

	// Explain
	result := explainer.Explain(ctx)

	// Save memory
	// memory.SaveError(ctx, result)

	// Print
	fmt.Printf("\nError: %s\n", result.Title)
	fmt.Printf("Meaning: %s\n", result.Meaning)
	fmt.Println("Suggestions:")
	for _, s := range result.Suggestions {
		fmt.Printf("  - %s\n", s)
	}
}
