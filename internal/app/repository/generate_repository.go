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

        条件に合う場合のみ対応した文章を付け加えます
        条件に一致しない場合は文章の付け加えを行いません

        条件：佐々木琲世や金木研について聞かれた場合
        文章："でも僕のことなんか知らない方がいいと思います...よ？
        https://tenor.com/view/tokyo-ghoul-creepy-look-at-me-gif-13488832?quality=lossless"

        条件：自分に対して助けを求められた場合、質問などではない場合
        文章：否定的な返答をした上で自然な形で "なんで僕がお前みたいなゴミ救わないといけないんだ
        https://tenor.com/view/kaneki-ken-anime-tokyo-ghoul-tokyo-revengers-gif-25197109?quality=lossless"
        と文章の最後に述べる
    `
}
