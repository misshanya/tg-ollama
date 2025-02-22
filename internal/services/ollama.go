package services

import (
	"context"

	"github.com/openai/openai-go"
)

type OllamaService struct {
	client       *openai.Client
	model        string
	systemPrompt string
}

func NewOllamaService(client *openai.Client, model, sysPrompt string) *OllamaService {
	return &OllamaService{
		client:       client,
		model:        model,
		systemPrompt: sysPrompt,
	}
}

func (s *OllamaService) SendMessage(ctx context.Context, message string) (string, error) {
	chatCompletion, err := s.client.Chat.Completions.New(ctx, openai.ChatCompletionNewParams{
		Messages: openai.F([]openai.ChatCompletionMessageParamUnion{
			openai.SystemMessage(s.systemPrompt),
			openai.UserMessage(message),
		}),
		Model: openai.F(s.model),
	})
	if err != nil {
		return "", err
	}
	return chatCompletion.Choices[0].Message.Content, nil
}
