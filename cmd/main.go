package main

import (
	"fmt"
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"github.com/lecterkn/kaneki_bot/internal/app/di"
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
	// ハンドラー取得
	handlerSet := di.InitializeHandlers()
	// ハンドラー登録
	discord.AddHandler(handlerSet.MessageHandler.Mention)

	// 開始
	err = discord.Open()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("discord bot is running...")
	select {}
}
