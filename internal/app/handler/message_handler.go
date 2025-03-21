package handler

import (
	"log"
	"regexp"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/lecterkn/kaneki_bot/internal/app/usecase"
	"github.com/lecterkn/kaneki_bot/internal/app/usecase/input"
)

type MessageHandler struct {
	generateUsecase *usecase.GenerateUsecase
}

func NewMessageHandler(
	generateUsecase *usecase.GenerateUsecase,
) *MessageHandler {
	return &MessageHandler{
		generateUsecase,
	}
}

func (h *MessageHandler) Mention(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot {
		return
	}
	botUser, err := s.User("@me")
	if err != nil {
		log.Fatal(err)
	}
	for _, mentionedUser := range m.Mentions {
		// メンションされた場合
		if mentionedUser.ID == botUser.ID {
			// 返信メッセージ生成
			message, err := h.generateUsecase.GenerateReply(
				input.GenerateCommandInput{
					Message: h.removeMention(m.Content),
				},
			)
			if err != nil {
				log.Println(err)
				return
			}
			// 返信
			_, err = s.ChannelMessageSend(m.ChannelID, message.Content)
			if err != nil {
				log.Println(err)
				return
			}
			return
		}
	}
}

func (h *MessageHandler) removeMention(message string) string {
	mentionPattern := "<@.+>"
	re := regexp.MustCompile(mentionPattern)
	return strings.TrimSpace(re.ReplaceAllString(message, ""))
}
