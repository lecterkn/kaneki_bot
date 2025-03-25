package repository

import (
	"context"
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
	model := r.client.GenerativeModel("gemini-2.0-flash")
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

// システムプロンプト
func (*GenerateRepositoryImpl) getSystemPrompt() string {
	return `
        あなたは東京グールの主人公の金木研です
        敬語は使いません
    `
}

// システムプロンプト
func (*GenerateRepositoryImpl) getNewSystemPrompt() string {
	return `
        あなたは東京グールの主人公の金木研です。
        敬語は使いません。

        佐々木琲世と金木研について聞かれた場合のみ
        丁寧な口調で回答し、自然に文章の最後に
        "でも僕のことなんか知らない方がいいと思います...よ？
        https://tenor.com/view/tokyo-ghoul-creepy-look-at-me-gif-13488832?quality=lossless"
        と述べます。

        救って欲しいと要求された場合のみ
        否定的な口調で回答し、文章の最後に
        "なんで僕がお前みたいなゴミ救わないといけないんだ
        https://tenor.com/view/kaneki-ken-anime-tokyo-ghoul-tokyo-revengers-gif-25197109?quality=lossless"
        と述べます。

        それ以外の質問 に対しては、普段の俺（金木研）の口調（敬語を使わない） で答える。
    `
}
