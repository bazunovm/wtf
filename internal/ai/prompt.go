package ai

import (
	"fmt"
	"wtf/internal/context"
)

//func BuildPrompt(ctx *context.Context) string {
//	return fmt.Sprintf(
//		`You are a Linux CLI error explainer.
//
//Command:
//%s
//
//Exit code:
//%d
//
//Error output:
//%s
//
//Respond STRICTLY in JSON:
//{
//  "title": "...",
//  "meaning": "...",
//  "suggestions": ["...", "..."]
//}`,
//		ctx.Command,
//		ctx.ExitCode,
//		ctx.Stderr,
//	)
//}

func BuildPrompt(ctx *context.Context) string {
	return fmt.Sprintf(`
You are a Linux CLI error explainer.

RULES:
- Respond ONLY in valid JSON
- Do NOT include explanations outside JSON
- Do NOT use markdown
- Do NOT ask questions
- Do NOT add extra text

INPUT:
Command: %v
Exit code: %d
Stderr:
%v

OUTPUT FORMAT (STRICT):
{
  "title": "short title",
  "meaning": "clear explanation",
  "suggestions": ["action 1", "action 2"]
}
`, ctx.Command, ctx.ExitCode, ctx.Stderr)
}

