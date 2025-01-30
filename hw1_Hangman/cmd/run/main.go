package main

import (
	"flag"
	"log/slog"
	"os"

	"github.com/es-debug/backend-academy-2024-go-template/internal/application"
	"github.com/es-debug/backend-academy-2024-go-template/internal/config"
	"github.com/es-debug/backend-academy-2024-go-template/internal/infrastructure/data"
	"github.com/es-debug/backend-academy-2024-go-template/internal/infrastructure/io"
	"github.com/es-debug/backend-academy-2024-go-template/internal/infrastructure/visualization"
	"github.com/es-debug/backend-academy-2024-go-template/internal/logs"
)

func main() {
	logger := logs.NewLogger()

	cfg := config.LoadConfig()

	category := flag.String("category", "", "Категория слов")
	difficulty := flag.String("difficulty", "", "Уровень сложности (easy, medium, hard)")
	flag.Parse()

	filePath := "internal/infrastructure/data/words.json"
	wordRepo, err := data.NewInMemoryWordRepository(filePath)

	if err != nil {
		logger.Error("Ошибка при создании репозитория слов", slog.Any("error", err))
		os.Exit(1)
	}

	wordService := application.NewWordService(wordRepo)
	inputReader := io.NewConsoleInputReader()
	outputWriter := io.NewConsoleOutputWriter()
	visualizer := visualization.NewConsoleHangmanVisualizer()
	gameService := application.NewGameService(wordService, inputReader, outputWriter, visualizer, cfg, logger)

	err = gameService.StartGame(*category, *difficulty)
	if err != nil {
		logger.Error("Ошибка при запуске игры", slog.Any("error", err))
		os.Exit(1)
	}
}
