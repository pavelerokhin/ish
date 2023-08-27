package openai

import (
	"context"
	"errors"
	"github.com/sashabaranov/go-openai"
	"github.com/spf13/viper"
)

var (
	client *openai.Client
)

func Init() {
	key := viper.GetString("open-ai-api-key")
	if key != "" {
		client = openai.NewClient(key)
	}
}

func MakeCompletion(prompt string) (string, error) {
	ctx := context.Background()

	var request openai.CompletionRequest
	request.Model = viper.GetString("open-ai-model")
	request.Prompt = prompt

	var response openai.CompletionResponse
	var err error

	if client == nil {
		return "", errors.New("OpenAI client is not initialized")
	}
	response, err = client.CreateCompletion(ctx, request)

	if err != nil {
		return "", err
	}

	return response.Choices[0].Text, nil
}
