package io_test

import (
	"testing"

	"github.com/es-debug/backend-academy-2024-go-template/internal/domain"

	ioMocks "github.com/es-debug/backend-academy-2024-go-template/internal/infrastructure/io/mocks"
)

func TestOutputWriterMock(t *testing.T) {
	mockOutput := new(ioMocks.OutputWriter)

	mockOutput.On("WriteOutput", "тестовое сообщение").Return().Once()

	mockOutput.WriteOutput("тестовое сообщение")

	mockOutput.AssertExpectations(t)
}

func TestDisplayGameStateMock(t *testing.T) {
	mockOutput := new(ioMocks.OutputWriter)
	game := &domain.Game{
		Word:           &domain.Word{Value: "тест"},
		GuessedLetters: []string{"т", "е"},
		AttemptsLeft:   5,
		MaxAttempts:    7,
	}

	mockOutput.On("DisplayGameState", game).Return().Once()

	mockOutput.DisplayGameState(game)

	mockOutput.AssertExpectations(t)
}
