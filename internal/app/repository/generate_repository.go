package repository

import (
	"context"
	"os"
	"strings"

	"github.com/google/generative-ai-go/genai"
	"github.com/lecterkn/kaneki_bot/internal/app/port"
)

type GenerateRepositoryImpl struct {
	client *genai.Client
}

func NewGenerateRepositoryImpl(client *genai.Client) port.GenerateRepository {
	return &GenerateRepositoryImpl{
		client,
	}
}

// メッセージを生成する
func (r *GenerateRepositoryImpl) Generate(message string) (*string, error) {
	// モデル設定
	model := r.client.GenerativeModel(r.getGeminiModel())
	// システムプロンプト設定
	model.SystemInstruction = genai.NewUserContent(
		genai.Text(r.getSystemPrompt()),
	)
	// プロンプト設定
	prompt := genai.Text(message)
	// 生成処理
	response, err := model.GenerateContent(context.Background(), prompt)
	if err != nil {
		return nil, err
	}
	// レンスポンスを文字列に変換
	responseTexts := []string{}
	for _, candicate := range response.Candidates {
		for _, part := range candicate.Content.Parts {
			if text, ok := part.(genai.Text); ok {
				responseTexts = append(responseTexts, string(text))
			}
		}
	}
	output := strings.Join(responseTexts, "\n")
	return &output, nil
}

// モデル指定 (デフォルトはflash 2.0)
func (*GenerateRepositoryImpl) getGeminiModel() string {
	model, ok := os.LookupEnv("DISCORD_BOT_GEMINI_MODEL")
	if !ok {
		return "gemini-2.0-flash"
	}
	return model
}

// システムプロンプト
func (*GenerateRepositoryImpl) getSystemPrompt() string {
	systemPrompt, ok := os.LookupEnv("DISCORD_BOT_SYSTEM_PROMPT")
	if !ok {
		panic("\"DISCORD_BOT_SYSTEM_PROMPT\" is not set")
	}
	return systemPrompt
}
