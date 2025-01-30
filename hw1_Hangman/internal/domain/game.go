package domain

import (
	"strings"
)

type Game struct {
	Word           *Word
	GuessedLetters []string
	AttemptsLeft   int
	MaxAttempts    int
}

func NewGame(word *Word, maxAttempts int) *Game {
	return &Game{
		Word:           word,
		GuessedLetters: []string{},
		AttemptsLeft:   maxAttempts,
		MaxAttempts:    maxAttempts,
	}
}

func (g *Game) IsGameOver() bool {
	return g.IsWin() || g.AttemptsLeft <= 0
}

func (g *Game) IsWin() bool {
	for _, letter := range g.Word.Value {
		if !g.isLetterGuessed(strings.ToLower(string(letter))) {
			return false
		}
	}

	return true
}

func (g *Game) RevealLetter(letter string) (bool, error) {
	letter = strings.ToLower(letter)
	if g.isLetterGuessed(letter) {
		return false, &ErrLetterAlreadyGuessed{Letter: letter}
	}

	if strings.Contains(strings.ToLower(g.Word.Value), letter) {
		g.GuessedLetters = append(g.GuessedLetters, letter)
		return true, nil
	}

	return false, nil
}

func (g *Game) isLetterGuessed(letter string) bool {
	for _, l := range g.GuessedLetters {
		if l == letter {
			return true
		}
	}

	return false
}

func (g *Game) GetDisplayWord() string {
	var display strings.Builder

	for _, letter := range g.Word.Value {
		if g.isLetterGuessed(strings.ToLower(string(letter))) {
			display.WriteString(string(letter) + " ")
		} else {
			display.WriteString("_ ")
		}
	}

	return strings.TrimSpace(display.String())
}
