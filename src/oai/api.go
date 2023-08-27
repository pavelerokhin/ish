package oai

import (
	"context"
	"errors"
	"github.com/sashabaranov/go-openai"
	"os"
)

const (
	MODEL = openai.GPT3Dot5Turbo
)

var (
	client *openai.Client
)

func Init() {
	if os.Getenv("OPENAI_API_KEY") == "" {
		panic("OPENAI_API_KEY is not set")
	}
	client = openai.NewClient(os.Getenv("OPENAI_API_KEY"))
}

func MakeCompletion(prompt string) (string, error) {
	ctx := context.Background()

	request := openai.ChatCompletionRequest{
		Model: MODEL,
		//MaxTokens: 5,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    "assistant",
				Content: "You help by making text clearer",
			},
			{
				Role:    "user",
				Content: prompt,
			},
		},
	}

	var response openai.ChatCompletionResponse
	var err error

	if client == nil {
		return "", errors.New("OpenAI client is not initialized")
	}
	response, err = client.CreateChatCompletion(ctx, request)

	if err != nil {
		return "", err
	}

	return response.Choices[0].Message.Content, nil
}
