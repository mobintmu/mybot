package router

import (
	tele "gopkg.in/telebot.v3"
)

func NewRouter(bot *tele.Bot) *RouterStruct {
	return &RouterStruct{
		BotInstance: bot,
	}
}

func (r RouterStruct) Start() {

	r.BotInstance.Handle("/hello", func(c tele.Context) error {
		return c.Send("Hello!")
	})

}
