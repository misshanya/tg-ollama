package handlers

import (
	"context"
	"fmt"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/misshanya/tg-ollama/internal/services"
)

type OllamaHandler struct {
	service *services.OllamaService
}

func NewOllamaHandler(service *services.OllamaService) *OllamaHandler {
	return &OllamaHandler{
		service: service,
	}
}

func (h *OllamaHandler) SendMessage(ctx context.Context, b *bot.Bot, update *models.Update) {
	response, err := h.service.SendMessage(ctx, update.Message.Text)
	if err != nil {
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   "Произошла ошибка при генерации :(",
		})
		return
	}
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   fmt.Sprintf("Ответ нейросети:\n%s", response),
	})
}
