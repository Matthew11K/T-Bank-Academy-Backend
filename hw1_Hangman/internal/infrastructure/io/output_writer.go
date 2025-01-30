package io

import (
	"fmt"

	"github.com/es-debug/backend-academy-2024-go-template/internal/domain"
)

type OutputWriter interface {
	WriteOutput(message string)
	DisplayGameState(game *domain.Game)
}

type ConsoleOutputWriter struct {
}

func NewConsoleOutputWriter() *ConsoleOutputWriter {
	return &ConsoleOutputWriter{}
}

func (w *ConsoleOutputWriter) WriteOutput(message string) {
	fmt.Println(message)
}

func (w *ConsoleOutputWriter) DisplayGameState(game *domain.Game) {
	fmt.Println("Слово:", game.GetDisplayWord())
	fmt.Println("Оставшиеся попытки:", game.AttemptsLeft)
}
