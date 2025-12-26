package ai

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"time"

	"wtf/internal/context"
	"wtf/internal/explainer"
)

const geminiURL = "https://generativelanguage.googleapis.com/v1/models/gemini-1.5-flash:generateContent"

type geminiRequest struct {
	Contents []geminiContent `json:"contents"`
}

type geminiContent struct {
	Role  string        `json:"role"`
	Parts []geminiPart  `json:"parts"`
}

type geminiPart struct {
	Text string `json:"text"`
}

type geminiResponse struct {
	Candidates []struct {
		Content struct {
			Parts []struct {
				Text string `json:"text"`
			} `json:"parts"`
		} `json:"content"`
	} `json:"candidates"`
}

func Explain(ctx *context.Context) (explainer.ExplainResult, error) {
	apiKey := os.Getenv("WTF_AI_API_KEY")
	if apiKey == "" {
		return explainer.ExplainResult{}, errors.New("AI API key not set")
	}

	prompt := BuildPrompt(ctx)

	reqBody := geminiRequest{
		Contents: []geminiContent{
			{
				Role: "user",
				Parts: []geminiPart{
					{Text: prompt},
				},
			},
		},
	}

	body, _ := json.Marshal(reqBody)

	req, err := http.NewRequest(
		"POST",
		geminiURL+"?key="+apiKey,
		bytes.NewBuffer(body),
	)
	if err != nil {
		return explainer.ExplainResult{}, err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return explainer.ExplainResult{}, err
	}
	defer resp.Body.Close()

	var gResp geminiResponse
	if err := json.NewDecoder(resp.Body).Decode(&gResp); err != nil {
		return explainer.ExplainResult{}, err
	}

	if len(gResp.Candidates) == 0 ||
		len(gResp.Candidates[0].Content.Parts) == 0 {
		return explainer.ExplainResult{}, errors.New("empty Gemini response")
	}

	raw := gResp.Candidates[0].Content.Parts[0].Text

	var result explainer.ExplainResult
	if err := json.Unmarshal([]byte(raw), &result); err != nil {
		return explainer.ExplainResult{}, err
	}

	return result, nil
}

