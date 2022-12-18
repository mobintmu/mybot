package main

import (
	"fmt"
	"log"
	"time"

	tele "gopkg.in/telebot.v3"

	app "mybot/app"
	router "mybot/app/router"
)

func main() {

	var Base app.Application = app.App{}

	token := Base.GoDotEnvVariable("TOKEN")

	pref := tele.Settings{
		Token:  token,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	bot, err := tele.NewBot(pref)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
		return
	}

	var routerHandler router.RouterInterface = router.NewRouter(bot)

	routerHandler.Start()

	bot.Start()
}
