package handlers

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type GeneralHandler struct {
}

func NewGeneralHandler() *GeneralHandler {
	return &GeneralHandler{}
}

func (h *GeneralHandler) Start(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Привет!",
	})
}
