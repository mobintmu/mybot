package router

import (
	tele "gopkg.in/telebot.v3"
)

// RouterInterface interface for routing
type RouterInterface interface {
	Start()
}

type RouterStruct struct {
	BotInstance *tele.Bot
}
