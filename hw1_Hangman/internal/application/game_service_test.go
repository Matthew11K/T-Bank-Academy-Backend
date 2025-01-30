package application_test

import (
	"errors"
	"testing"

	"log/slog"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/es-debug/backend-academy-2024-go-template/internal/application"
	"github.com/es-debug/backend-academy-2024-go-template/internal/application/mocks"
	"github.com/es-debug/backend-academy-2024-go-template/internal/config"
	"github.com/es-debug/backend-academy-2024-go-template/internal/domain"
	ioMocks "github.com/es-debug/backend-academy-2024-go-template/internal/infrastructure/io/mocks"
	visualizationMocks "github.com/es-debug/backend-academy-2024-go-template/internal/infrastructure/visualization/mocks"
)

func TestGameService_StartGame_Win(t *testing.T) {
	wordServiceMock := new(mocks.WordServiceInterface)
	inputReaderMock := new(ioMocks.InputReader)
	outputWriterMock := new(ioMocks.OutputWriter)
	visualizerMock := new(visualizationMocks.HangmanVisualizer)
	logger := slog.New(slog.NewTextHandler(nil, nil))

	cfg := &config.Config{
		MaxAttemptsEasy:   10,
		MaxAttemptsMedium: 7,
		MaxAttemptsHard:   5,
	}

	gameService := application.NewGameService(
		wordServiceMock,
		inputReaderMock,
		outputWriterMock,
		visualizerMock,
		cfg,
		logger,
	)

	word := &domain.Word{
		Value:      "тест",
		Category:   "категория",
		Difficulty: "easy",
		Hint:       "подсказка",
	}

	wordServiceMock.On("GetRandomWord", "", "").Return(word, nil)

	inputReaderMock.On("ReadInput").Return("т", nil).Once()
	inputReaderMock.On("ReadInput").Return("е", nil).Once()
	inputReaderMock.On("ReadInput").Return("с", nil).Once()

	outputWriterMock.On("WriteOutput", mock.Anything).Return()
	outputWriterMock.On("DisplayGameState", mock.Anything).Return()
	visualizerMock.On("DisplayHangman", mock.Anything, mock.Anything).Return()

	err := gameService.StartGame("", "")
	assert.NoError(t, err)

	wordServiceMock.AssertExpectations(t)
	inputReaderMock.AssertExpectations(t)
	outputWriterMock.AssertExpectations(t)
	visualizerMock.AssertExpectations(t)
}

func TestGameService_StartGame_ErrorWordService(t *testing.T) {
	wordServiceMock := new(mocks.WordServiceInterface)
	inputReaderMock := new(ioMocks.InputReader)
	outputWriterMock := new(ioMocks.OutputWriter)
	visualizerMock := new(visualizationMocks.HangmanVisualizer)
	logger := slog.New(slog.NewTextHandler(nil, nil))

	cfg := &config.Config{
		MaxAttemptsEasy:   10,
		MaxAttemptsMedium: 7,
		MaxAttemptsHard:   5,
	}

	gameService := application.NewGameService(
		wordServiceMock,
		inputReaderMock,
		outputWriterMock,
		visualizerMock,
		cfg,
		logger,
	)

	wordServiceMock.On("GetRandomWord", "", "").Return(nil, errors.New("ошибка получения слова"))

	err := gameService.StartGame("", "")
	assert.Error(t, err)
	assert.Equal(t, "ошибка при получении слова: ошибка получения слова", err.Error())

	wordServiceMock.AssertExpectations(t)
}

func TestGameService_StartGame_InvalidParams(t *testing.T) {
	wordServiceMock := new(mocks.WordServiceInterface)
	inputReaderMock := new(ioMocks.InputReader)
	outputWriterMock := new(ioMocks.OutputWriter)
	visualizerMock := new(visualizationMocks.HangmanVisualizer)
	logger := slog.New(slog.NewTextHandler(nil, nil))

	cfg := &config.Config{
		MaxAttemptsEasy:   10,
		MaxAttemptsMedium: 7,
		MaxAttemptsHard:   5,
	}

	gameService := application.NewGameService(
		wordServiceMock,
		inputReaderMock,
		outputWriterMock,
		visualizerMock,
		cfg,
		logger,
	)

	wordServiceMock.On("GetRandomWord", "unknown", "easy").Return(nil, errors.New("слово не найдено"))

	err := gameService.StartGame("unknown", "easy")
	assert.Error(t, err)
	wordServiceMock.AssertExpectations(t)
}
