package handlers

import (
	"context"
	"fmt"
	"log"
	"time"

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
	startTime := time.Now()
	response, err := h.service.SendMessage(ctx, update.Message.Text)
	if err != nil {
		log.Printf("Failed to generate: %v", err)
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   "Произошла ошибка при генерации :(",
		})
		return
	}
	endTime := time.Now()

	timeDiff := endTime.Unix() - startTime.Unix()
	_, err = b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   fmt.Sprintf("Ответ нейросети занял %v сек:\n%s", timeDiff, response),
	})
	if err != nil {
		log.Printf("Failed to send message: %v", err)
	}
	log.Printf("Message from %s: %s;\nAI response: %s", update.Message.Chat.FirstName, update.Message.Text, response)
}
