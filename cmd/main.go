package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"github.com/lecterkn/kaneki_bot/internal/app/client"
	"github.com/lecterkn/kaneki_bot/internal/app/handler"
	"github.com/lecterkn/kaneki_bot/internal/app/port"
	"github.com/lecterkn/kaneki_bot/internal/app/repository"
	"github.com/lecterkn/kaneki_bot/internal/app/usecase"
)

func main() {
	// 環境変数読み込み
	err := godotenv.Load()
	if err != nil {
		panic("\".env\" was not found")
	}
	// トークン
	discordApiToken, ok := os.LookupEnv("DISCORD_BOT_TOKEN")
	if !ok {
		panic("\"DISCORD_BOT_TOKEN\" is not set")
	}
	// discord起動
	discord, err := discordgo.New("Bot " + discordApiToken)
	if err != nil {
		panic("failed to initialize discord client")
	}
	// DI
	client := client.GetGeminiClient()
	var generateRepository port.GenerateRepository
	if strings.ToLower(os.Getenv("DISCORD_BOT_JAILBREAK")) == "true" {
		generateRepository = repository.NewGenerateJailbreakRepositoryImpl(client)
	} else {
		generateRepository = repository.NewGenerateCommonRepositoryImpl(client)
	}
	generateUsecase := usecase.NewGenerateUsecase(generateRepository)
	messageHandler := handler.NewMessageHandler(generateUsecase)
	// ハンドラー登録
	discord.AddHandler(messageHandler.Mention)

	// 開始
	err = discord.Open()
	if err != nil {
		log.Fatal(err)
	}
	model, ok := os.LookupEnv("DISCORD_BOT_GEMINI_MODEL")
	if !ok {
		model = fmt.Sprintf("default(%s)", repository.GEMINI_MODEL_FALLBACK)
	}
	fmt.Printf("discord bot is started\n running on %s\n", model)
	select {}
}
