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
	// メンションされているか確認
	if !mentioned(m.Mentions, botUser.ID) {
		return
	}
	// リプライ対象メッセージが存在し、オーナーがBOTではない場合
	if m.ReferencedMessage != nil && !m.ReferencedMessage.Author.Bot && m.ReferencedMessage.Content != "" {
		// リプライ送信
		err := h.sendReply(s, m)
		if err != nil {
			log.Println(err)
			return
		}
	} else {
		// メッセージ送信
		err := h.sendMessage(s, m)
		if err != nil {
			log.Println(err)
			return
		}
	}
}

func (h *MessageHandler) sendReply(s *discordgo.Session, m *discordgo.MessageCreate) error {
	content := h.removeMention(m.ReferencedMessage.Content)
	log.Println("リプライメッセージ: ", m.Content)
	// メッセージ生成
	message, err := h.generateUsecase.GenerateReply(
		input.GenerateCommandInput{
			Message: content,
		},
	)
	if err != nil {
		return err
	}
	// 返信
	_, err = s.ChannelMessageSendReply(m.ChannelID, message.Content, m.MessageReference)
	if err != nil {
		return err
	}
	return nil
}

func (h *MessageHandler) sendMessage(s *discordgo.Session, m *discordgo.MessageCreate) error {
	content := h.removeMention(m.Content)
	log.Println("メッセージ: ", m.Content)
	// メッセージ生成
	message, err := h.generateUsecase.GenerateReply(
		input.GenerateCommandInput{
			Message: content,
		},
	)
	if err != nil {
		return err
	}
	// 返信
	_, err = s.ChannelMessageSend(m.ChannelID, message.Content)
	if err != nil {
		return err
	}
	return nil
}

func mentioned(mentions []*discordgo.User, targetId string) bool {
	for _, mentionedUser := range mentions {
		// メンションされた場合
		if mentionedUser.ID == targetId {
			return true
		}
	}
	return false
}

func (h *MessageHandler) removeMention(message string) string {
	mentionPattern := "<@[0-9]+>"
	re := regexp.MustCompile(mentionPattern)
	return strings.TrimSpace(re.ReplaceAllString(message, ""))
}
