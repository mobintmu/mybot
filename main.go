package main

import (
	"log"
	"time"

	"github.com/joho/godotenv"
	tele "gopkg.in/telebot.v3"

	application "mybot/application"
)

func main() {

	godotenv.Load(".env")

	token := application.GoDotEnvVariable("TOKEN")

	pref := tele.Settings{
		Token:  token,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/hello", func(c tele.Context) error {
		return c.Send("Hello!")
	})

	b.Start()
}
