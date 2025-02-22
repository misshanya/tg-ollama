package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/go-telegram/bot"
	"github.com/misshanya/tg-ollama/internal/config"
	"github.com/misshanya/tg-ollama/internal/handlers"
	"github.com/misshanya/tg-ollama/internal/services"
	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	cfg := config.NewConfig()

	ollamaClient := openai.NewClient(option.WithBaseURL(cfg.OllamaURL))
	ollamaService := services.NewOllamaService(ollamaClient, cfg.OllamaModel)
	ollamaHandler := handlers.NewOllamaHandler(ollamaService)

	opts := []bot.Option{
		bot.WithDefaultHandler(ollamaHandler.SendMessage),
	}

	b, err := bot.New(cfg.BotToken, opts...)
	if err != nil {
		panic(err)
	}

	log.Println("Starting bot")
	b.Start(ctx)
}
