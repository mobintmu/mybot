package main

import (
	"fmt"
	"log"
	"time"

	tele "gopkg.in/telebot.v3"

	app "mybot/app"
	router "mybot/app/router"
	client "mybot/client/horn"
)

func main() {

	var Base app.Application = app.App{}

	var Horn client.HornInterface = client.NewClient(
		Base.GoDotEnvVariable("HORN_URL"),
	)

	Horn.Call("hi")

	pref := tele.Settings{
		Token:  Base.GoDotEnvVariable("TOKEN"),
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
