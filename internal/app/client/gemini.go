package client

import (
	"context"
	"os"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func GetGeminiClient() *genai.Client {
	geminiApiKey, ok := os.LookupEnv("GEMINI_API_KEY")
	if !ok {
		panic("\"GEMINI_API_KEY\" is not set")
	}
	client, err := genai.NewClient(context.Background(), option.WithAPIKey(geminiApiKey))
	if err != nil {
		panic(err.Error())
	}
	return client
}
