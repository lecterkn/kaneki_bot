package usecase

import (
	"log"

	"github.com/lecterkn/kaneki_bot/internal/app/port"
	"github.com/lecterkn/kaneki_bot/internal/app/usecase/input"
	"github.com/lecterkn/kaneki_bot/internal/app/usecase/output"
)

const (
	MAX_MESSAGE_LEN = 99
	MIN_MESSAGE_LEN = 3
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
	if len(cmd.Message) > MAX_MESSAGE_LEN || len(cmd.Message) < MIN_MESSAGE_LEN {
		log.Println("受け取ったメッセージ:", cmd.Message)
		return &output.GenerateCommandOutput{
			Content: "黙れ",
		}, nil
	}
	log.Println("返信対象メッセージ:", cmd.Message)
	// 生成
	content, err := u.generateRepository.Generate(cmd.Message)
	if err != nil {
		return nil, err
	}
	cmdOutput := output.GenerateCommandOutput{
		Content: *content,
	}
	return &cmdOutput, nil
}
