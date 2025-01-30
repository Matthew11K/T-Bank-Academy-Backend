package application

import (
	"errors"
	"fmt"
	"strings"

	"log/slog"

	"github.com/es-debug/backend-academy-2024-go-template/internal/config"
	"github.com/es-debug/backend-academy-2024-go-template/internal/domain"
	"github.com/es-debug/backend-academy-2024-go-template/internal/infrastructure/io"
	"github.com/es-debug/backend-academy-2024-go-template/internal/infrastructure/visualization"
)

type GameService struct {
	wordService  WordServiceInterface
	inputReader  io.InputReader
	outputWriter io.OutputWriter
	visualizer   visualization.HangmanVisualizer
	config       *config.Config
	logger       *slog.Logger
}

func NewGameService(
	wordService WordServiceInterface,
	inputReader io.InputReader,
	outputWriter io.OutputWriter,
	visualizer visualization.HangmanVisualizer,
	configuration *config.Config,
	logger *slog.Logger,
) *GameService {
	return &GameService{
		wordService:  wordService,
		inputReader:  inputReader,
		outputWriter: outputWriter,
		visualizer:   visualizer,
		config:       configuration,
		logger:       logger,
	}
}

func (gs *GameService) StartGame(category, difficulty string) error {
	word, err := gs.wordService.GetRandomWord(category, difficulty)
	if err != nil {
		return fmt.Errorf("ошибка при получении слова: %w", err)
	}

	maxAttempts := gs.getMaxAttempts(word.Difficulty)
	game := domain.NewGame(word, maxAttempts)

	gs.outputWriter.WriteOutput("Добро пожаловать в игру 'Виселица'!")
	gs.outputWriter.WriteOutput("Категория: " + word.Category)
	gs.outputWriter.WriteOutput("Подсказка: " + word.Hint)

	for !game.IsGameOver() {
		gs.visualizer.DisplayHangman(game.AttemptsLeft, game.MaxAttempts)
		gs.outputWriter.DisplayGameState(game)
		gs.outputWriter.WriteOutput("Введите букву:")
		input, err := gs.inputReader.ReadInput()

		if err != nil {
			gs.logger.Error("Ошибка ввода", slog.Any("error", err))
			continue
		}

		input = strings.TrimSpace(input)

		if len([]rune(input)) != 1 {
			gs.outputWriter.WriteOutput("Пожалуйста, введите одну букву.")
			continue
		}

		letter := strings.ToLower(input)
		correct, err := game.RevealLetter(letter)

		if err != nil {
			var alreadyGuessedErr *domain.ErrLetterAlreadyGuessed
			if errors.As(err, &alreadyGuessedErr) {
				gs.outputWriter.WriteOutput("Вы уже вводили эту букву: " + alreadyGuessedErr.Letter)
				continue
			}

			return fmt.Errorf("ошибка при попытке открыть букву: %w", err)
		}

		if correct {
			gs.outputWriter.WriteOutput("Правильно!")
		} else {
			game.AttemptsLeft--

			gs.outputWriter.WriteOutput("Неправильно!")
		}
	}

	gs.visualizer.DisplayHangman(game.AttemptsLeft, game.MaxAttempts)
	gs.outputWriter.DisplayGameState(game)

	if game.IsWin() {
		gs.outputWriter.WriteOutput("Поздравляем! Вы выиграли!")
	} else {
		gs.outputWriter.WriteOutput("Вы проиграли. Загаданное слово было: " + game.Word.Value)
	}

	return nil
}

func (gs *GameService) getMaxAttempts(difficulty string) int {
	switch difficulty {
	case "easy":
		return gs.config.MaxAttemptsEasy
	case "medium":
		return gs.config.MaxAttemptsMedium
	case "hard":
		return gs.config.MaxAttemptsHard
	default:
		return gs.config.MaxAttemptsMedium
	}
}
