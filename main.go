package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/diamondburned/arikawa/gateway"
	"github.com/diamondburned/arikawa/session"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/ilyakaznacheev/cleanenv"
)

func main() {
	var cfg Config
	err := cleanenv.ReadConfig("config.yml", &cfg)
	if err != nil {
		log.Println("Config err:", err)
	}

	msg := make(chan string)
	go discordGrabber(msg, cfg)
	go tgSender(msg, cfg)
	select {}
}

func tgSender(msg chan string, cfg Config) {
	bot, err := tgbotapi.NewBotAPI(cfg.TelegramApiToken)
	if err != nil {
		panic(err)
	}

	bot.Debug = true
	log.Printf("Telegram: Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	for {
		res, ok := <-msg
		if !ok {
			fmt.Println("msg channel close", ok)
			break
		}
		m := tgbotapi.NewMessage(cfg.TelegramChannel, res)
		bot.Send(m)
	}
}

func discordGrabber(msg chan string, cfg Config) {
	s, err := session.Login(cfg.DiscordLogin, cfg.DiscordPassword, "")
	if err != nil {
		log.Fatalln(err)
	}
	s.AddHandler(func(c *gateway.MessageCreateEvent) {
		channel, err := s.Channel(c.ChannelID)
		if err != nil {
			log.Println("Can't get Discord channel:", err)
		}
		log.Println(channel.Name, c.ChannelID.String())
		log.Println(c.Author.Username, "sent", c.Content)
		for _, ch := range cfg.DiscordWhiteList {
			chID, err := strconv.ParseInt(c.ChannelID.String(), 10, 64)
			if err != nil {
				log.Println("can't parse ChannelID", err)
			}
			if chID == ch {
				log.Println("WHITE LIST")
				msg <- c.Author.Username + ": " + c.Content
			}
		}
	})

	if err := s.Open(); err != nil {
		log.Fatalln("Failed to connect:", err)
	}
	defer s.Close()

	u, err := s.Me()
	if err != nil {
		log.Fatalln("Failed to get myself:", err)
	}

	log.Println("Discord: Started as", u.Username)

	// Block forever.
	select {}
}
