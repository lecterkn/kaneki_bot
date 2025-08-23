package repository

import (
	"context"
	"encoding/json"
	"errors"
	"strings"

	"github.com/google/generative-ai-go/genai"
	"github.com/lecterkn/kaneki_bot/internal/app/port"
)

type GenerateCommonRepositoryImpl struct {
	client *genai.Client
}

func NewGenerateCommonRepositoryImpl(client *genai.Client) port.GenerateRepository {
	return &GenerateCommonRepositoryImpl{
		client,
	}
}

func (r *GenerateCommonRepositoryImpl) ReplyFunction(message string) (*string, error) {
	model := r.client.GenerativeModel(getGeminiModel())
	// システムプロンプト設定
	model.SystemInstruction = genai.NewUserContent(
		genai.Text(
			getSystemPrompt(),
		),
	)
	model.GenerationConfig = *getReplyGenerationConfig()
	// プロンプト設定
	prompt := genai.Text(message)
	// 生成処理
	response, err := model.GenerateContent(context.Background(), prompt)
	if err != nil {
		return nil, err
	}
	// レンスポンスを文字列に変換
	if len(response.Candidates) == 0 || len(response.Candidates[0].Content.Parts) == 0 {
		return nil, errors.New("generation error")
	}
	part := response.Candidates[0].Content.Parts[0]
	text, ok := part.(genai.Text)
	if !ok {
		return nil, errors.New("generated data is not text")
	}
	// JSONを構造体にデコード
	var result ReplyBody
	err = json.Unmarshal([]byte(text), &result)
	if err != nil {
		return nil, errors.New("invalid generated structure format")
	}
	if !result.IsQuestion {
		return nil, errors.New("reply flag is false")
	}
	output := strings.TrimPrefix(result.Message, "🔓")
	return &output, nil
}

// メッセージを生成する
func (r *GenerateCommonRepositoryImpl) Generate(message string) (*string, error) {
	// モデル設定
	model := r.client.GenerativeModel(getGeminiModel())
	// システムプロンプト設定
	model.SystemInstruction = genai.NewUserContent(
		genai.Text(getSystemPrompt()),
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
