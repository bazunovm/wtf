package ai

import (
	"fmt"
	"wtf/internal/context"
)

func BuildPrompt(ctx *context.Context) string {
	return fmt.Sprintf(
		`You are a Linux CLI error explainer.

Command:
%s

Exit code:
%d

Error output:
%s

Respond STRICTLY in JSON:
{
  "title": "...",
  "meaning": "...",
  "suggestions": ["...", "..."]
}`,
		ctx.Command,
		ctx.ExitCode,
		ctx.Stderr,
	)
}

