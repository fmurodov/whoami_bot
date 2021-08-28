package main

import (
	"log"
	"net"
	"net/http"
	"os"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	var text string
	botToken := os.Getenv("BOT_TOKEN")
	if len(botToken) == 0 {
		log.Fatal("BOT_TOKEN is missing")
	}
	port := os.Getenv("PORT")
	if len(port) == 0 {
		log.Fatal("Port is not set")
	}

	log.Println("Starting bot...")
	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = false

	log.Printf("Authorized on account %s", bot.Self.UserName)

	_, err = bot.SetWebhook(tgbotapi.NewWebhook("https://whoamirobot.herokuapp.com/" + bot.Token))
	if err != nil {
		log.Fatal(err)
	}
	info, err := bot.GetWebhookInfo()
	if err != nil {
		log.Fatal(err)
	}
	if info.LastErrorDate != 0 {
		log.Printf("Telegram callback failed: %s", info.LastErrorMessage)
	}

	updates := bot.ListenForWebhook("/" + bot.Token)
	go http.ListenAndServe(net.JoinHostPort("", port), nil)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		log.Printf("[%s] %s", strconv.Itoa(update.Message.From.ID), update.Message.Text)

		if update.Message.Text == "/me" || update.Message.Text == "/whoami" {
			text = ("chat id: " + strconv.Itoa(update.Message.From.ID) +
				"\nusername: " + update.Message.From.UserName +
				"\nname: " + update.Message.From.FirstName + " " + update.Message.From.LastName)

		} else if update.Message.Text == "/start" {
			text = "Hi! I know all about you!\nSend me /whoami command"

		} else {
			text = "send /me or /whoami for get info"

		}
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
		msg.ReplyToMessageID = update.Message.MessageID

		if _, err := bot.Send(msg); err != nil {
			log.Println(err)
		}
	}
}
