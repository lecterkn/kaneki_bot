package usecase

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode/utf8"

	"github.com/lecterkn/kaneki_bot/internal/app/port"
	"github.com/lecterkn/kaneki_bot/internal/app/usecase/input"
	"github.com/lecterkn/kaneki_bot/internal/app/usecase/output"
)

const (
	MAX_MESSAGE_LEN = 100
	MIN_MESSAGE_LEN = 3
)

type GenerateUsecase struct {
	generateRepository port.GenerateRepository
	messageMin         int
	messageMax         int
}

func NewGenerateUsecase(generateRepository port.GenerateRepository) *GenerateUsecase {
	return &GenerateUsecase{
		generateRepository,
		getMessageMin(),
		getMessageMax(),
	}
}

func (u *GenerateUsecase) GenerateReply(cmd input.GenerateCommandInput) (*output.GenerateCommandOutput, error) {
	// バリデーション
	if utf8.RuneCountInString(cmd.Message) > u.messageMax ||
		utf8.RuneCountInString(cmd.Message) < u.messageMin {
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

func getMessageMax() int {
	limit, ok := os.LookupEnv("MAX_MESSAGE_LEN")
	if !ok {
		fmt.Println("max not found")
		return MAX_MESSAGE_LEN
	}
	value, err := strconv.Atoi(limit)
	if err != nil {
		fmt.Println(err.Error())
		return MAX_MESSAGE_LEN
	}
	return value
}

func getMessageMin() int {
	limit, ok := os.LookupEnv("MIN_MESSAGE_LEN")
	if !ok {
		fmt.Println("min not found")
		return MIN_MESSAGE_LEN
	}
	value, err := strconv.Atoi(limit)
	if err != nil {
		fmt.Println(err.Error())
		return MIN_MESSAGE_LEN
	}
	return value
}
