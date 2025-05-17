package client

import (
	"context"
	"fmt"
	"os"
	"strconv"

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

func listModels(client *genai.Client) {
	res := client.ListModels(context.Background())
	for {
		info, err := res.Next()
		if err != nil {
			break
		}
		fmt.Println(info.Name + " | " + strconv.Itoa(int(info.InputTokenLimit)) + ", " + strconv.Itoa(int(info.OutputTokenLimit)))
	}
}
