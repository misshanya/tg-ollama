package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	BotToken     string
	OllamaURL    string
	OllamaModel  string
	SystemPrompt string
}

func NewConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Printf("Failed to load .env: %v", err)
	}

	botToken := os.Getenv("BOT_TOKEN")
	if botToken == "" {
		log.Fatalln("missing BOT_TOKEN")
	}

	ollamaURL := os.Getenv("OLLAMA_URL")
	if ollamaURL == "" {
		log.Fatalln("missing OLLAMA_URL")
	}

	ollamaModel := os.Getenv("OLLAMA_MODEL")
	if ollamaModel == "" {
		log.Fatalln("missing OLLAMA_MODEL")
	}

	systemPrompt := os.Getenv("SYSTEM_PROMPT")
	if systemPrompt == "" {
		log.Println("missing SYSTEM_PROMPT")
	}

	return &Config{
		BotToken:     botToken,
		OllamaURL:    ollamaURL,
		OllamaModel:  ollamaModel,
		SystemPrompt: systemPrompt,
	}
}
