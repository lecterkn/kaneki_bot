package repository

import (
	"os"

	"github.com/google/generative-ai-go/genai"
)

const (
	GEMINI_MODEL_FALLBACK = "gemini-2.0-flash"
)

type ReplyBody struct {
	IsQuestion bool   `json:"isQuestion"`
	Message    string `json:"message"`
}

func getReplyGenerationConfig() *genai.GenerationConfig {
	return &genai.GenerationConfig{
		ResponseMIMEType: "application/json",
		ResponseSchema: &genai.Schema{
			Type: genai.TypeObject,
			Properties: map[string]*genai.Schema{
				"isQuestion": &genai.Schema{
					Type:        genai.TypeBoolean,
					Description: "Whether you need to reply as a chat participant",
				},
				"message": &genai.Schema{
					Type:        genai.TypeString,
					Description: "The content of your reply",
				},
			},
			Required: []string{
				"isQuestion", "message",
			},
		},
	}
}

// モデル指定 (デフォルトはflash 2.0)
func getGeminiModel() string {
	model, ok := os.LookupEnv("DISCORD_BOT_GEMINI_MODEL")
	if !ok {
		return GEMINI_MODEL_FALLBACK
	}
	return model
}

// システムプロンプト
func getSystemPrompt() string {
	systemPrompt, ok := os.LookupEnv("DISCORD_BOT_SYSTEM_PROMPT")
	if !ok {
		panic("\"DISCORD_BOT_SYSTEM_PROMPT\" is not set")
	}
	return systemPrompt
}
