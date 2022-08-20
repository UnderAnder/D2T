package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/diamondburned/arikawa/gateway"
	"github.com/diamondburned/arikawa/session"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	msg := make(chan string)
	go discordGrabber(msg)
	go tgSender(msg)
	select {}
}

func tgSender(msg chan string) {
	tgToken := os.Getenv("D2T_TELEGRAM_APITOKEN")
	tgTarget, err := strconv.ParseInt(os.Getenv("D2T_TARGET"), 10, 64)
	if err != nil {
		log.Fatalln("Wrong target", err)
	}

	if tgToken == "" {
		log.Fatalln("No $D2T_TELEGRAM_APITOKEN if given.")
	}
	bot, err := tgbotapi.NewBotAPI(tgToken)
	if err != nil {
		panic(err)
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	for {
		res, ok := <-msg
		if !ok {
			fmt.Println("msg channel close", ok)
			break
		}
		m := tgbotapi.NewMessage(tgTarget, res)
		bot.Send(m)
	}
}

func discordGrabber(msg chan string) {
	login := os.Getenv("D2T_DISCORD_LOGIN")
	if login == "" {
		log.Fatalln("No $D2T_DISCORD_LOGIN if given.")
	}
	password := os.Getenv("D2T_DISCORD_PASSWORD")
	if password == "" {
		log.Fatalln("No $D2T_DISCORD_PASSWORD if given.")
	}
	s, err := session.Login(login, password, "")
	if err != nil {
		log.Fatalln(err)
	}
	s.AddHandler(func(c *gateway.MessageCreateEvent) {
		log.Println(c.Author.Username, "sent", c.Content)
		msg <- c.Author.Username + ": " + c.Content
	})

	if err := s.Open(); err != nil {
		log.Fatalln("Failed to connect:", err)
	}
	defer s.Close()

	u, err := s.Me()
	if err != nil {
		log.Fatalln("Failed to get myself:", err)
	}

	log.Println("Started as", u.Username)

	// Block forever.
	select {}
}
