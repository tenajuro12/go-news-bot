package summary

import (
	"context"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"log"
	"strings"
	"sync"
)

type OpenAiSummarizer struct {
	client  *openai.Client
	prompt  string
	enabled bool
	mu      sync.Mutex
}

func NewOpenAiSummarizer(apiKey string, prompt string) *OpenAiSummarizer {
	s := &OpenAiSummarizer{client: openai.NewClient(apiKey), prompt: prompt}
	log.Printf("openai summarizer enabled", apiKey != "")

	if apiKey != "" {
		s.enabled = true
	}
	return s
}

func (s *OpenAiSummarizer) Summarize(ctx context.Context, text string) (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if !s.enabled {
		return "", nil
	}

	req := openai.ChatCompletionRequest{
		Model: "o3-mini",
		Messages: []openai.ChatCompletionMessage{{
			Role:    openai.ChatMessageRoleSystem,
			Content: fmt.Sprintf("%s%s", text, s.prompt),
		}},
		MaxTokens:   256,
		Temperature: 0.7,
		TopP:        1}

	resp, err := s.client.CreateChatCompletion(ctx, req)
	if err != nil {
		return "", err
	}
	rawSummary := strings.TrimSpace(resp.Choices[0].Message.Content)

	if strings.HasSuffix(rawSummary, ".") {
		return rawSummary, nil
	}
	sentences := strings.Split(rawSummary, ".")

	return strings.Join(sentences[:len(sentences)-1], ".") + ".", nil
}
