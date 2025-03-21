package usecase

import (
	"log"

	"github.com/lecterkn/kaneki_bot/internal/app/port"
	"github.com/lecterkn/kaneki_bot/internal/app/usecase/input"
	"github.com/lecterkn/kaneki_bot/internal/app/usecase/output"
)

type GenerateUsecase struct {
	generateRepository port.GenerateRepository
}

func NewGenerateUsecase(generateRepository port.GenerateRepository) *GenerateUsecase {
	return &GenerateUsecase{
		generateRepository,
	}
}

func (u *GenerateUsecase) GenerateReply(cmd input.GenerateCommandInput) (*output.GenerateCommandOutput, error) {
	// バリデーション
	if len(cmd.Message) > 63 || len(cmd.Message) < 3 {
		log.Println("受け取ったメッセージ:", cmd.Message)
		return &output.GenerateCommandOutput{
			Content: "黙れ",
		}, nil
	}
	// 生成
	log.Println("返信対象メッセージ:", cmd.Message)
	content, err := u.generateRepository.Generate(cmd.Message)
	if err != nil {
		return nil, err
	}
	cmdOutput := output.GenerateCommandOutput{
		Content: *content,
	}
	return &cmdOutput, nil
}
