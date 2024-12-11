package services

import (
	"errors"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var bot *tgbotapi.BotAPI

func InitTelegramBot() {
	token := os.Getenv("TELEGRAM_BOT_TOKEN")
	if token == "" {
		log.Fatal("TELEGRAM_BOT_TOKEN is not set")
	}

	var err error
	bot, err = tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatalf("Failed to initialize Telegram bot: %v", err)
	}
}

func PublishToTelegram(channelID, message string) error {
	if bot == nil {
		return errors.New("Telegram bot is not initialized")
	}

	msg := tgbotapi.NewMessageToChannel(channelID, message)
	_, err := bot.Send(msg)
	return err
}
