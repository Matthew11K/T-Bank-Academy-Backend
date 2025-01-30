package domain_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/es-debug/backend-academy-2024-go-template/internal/domain"
)

func TestNewGame(t *testing.T) {
	word := &domain.Word{Value: "тест"}
	game := domain.NewGame(word, 5)

	assert.Equal(t, word, game.Word)
	assert.Equal(t, 5, game.MaxAttempts)
	assert.Equal(t, 5, game.AttemptsLeft)
	assert.Empty(t, game.GuessedLetters)
}

func TestGame_IsWin(t *testing.T) {
	word := &domain.Word{Value: "тест"}
	game := domain.NewGame(word, 5)

	game.GuessedLetters = []string{"т", "е", "с"}
	assert.True(t, game.IsWin())
}

func TestGame_IsGameOver(t *testing.T) {
	word := &domain.Word{Value: "тест"}
	game := domain.NewGame(word, 1)

	game.AttemptsLeft = 0
	assert.True(t, game.IsGameOver())

	game.AttemptsLeft = 1
	game.GuessedLetters = []string{"т", "е", "с", "т"}
	assert.True(t, game.IsGameOver())
}

func TestGame_RevealLetter(t *testing.T) {
	word := &domain.Word{Value: "тест"}
	game := domain.NewGame(word, 5)

	correct, err := game.RevealLetter("т")
	assert.NoError(t, err)
	assert.True(t, correct)
	assert.Contains(t, game.GuessedLetters, "т")

	correct, err = game.RevealLetter("т")
	assert.Error(t, err)
	assert.False(t, correct)

	_, ok := err.(*domain.ErrLetterAlreadyGuessed)

	assert.True(t, ok)

	correct, err = game.RevealLetter("а")
	assert.NoError(t, err)
	assert.False(t, correct)
	assert.NotContains(t, game.GuessedLetters, "а")
}

func TestGame_GetDisplayWord(t *testing.T) {
	word := &domain.Word{Value: "тест"}
	game := domain.NewGame(word, 5)

	game.GuessedLetters = []string{"т"}
	displayWord := game.GetDisplayWord()
	assert.Equal(t, "т _ _ т", displayWord)
}
