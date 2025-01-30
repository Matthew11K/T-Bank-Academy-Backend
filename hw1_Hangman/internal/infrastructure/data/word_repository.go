package data

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"os"

	"github.com/es-debug/backend-academy-2024-go-template/internal/domain"
)

type WordRepository interface {
	GetRandomWord(category, difficulty string) (*domain.Word, error)
}

type InMemoryWordRepository struct {
	words []domain.Word
}

func NewInMemoryWordRepository(filePath string) (*InMemoryWordRepository, error) {
	repo := &InMemoryWordRepository{}
	err := repo.loadWordsFromFile(filePath)

	if err != nil {
		return nil, err
	}

	return repo, nil
}

func (repo *InMemoryWordRepository) loadWordsFromFile(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("ошибка при открытии файла: %w", err)
	}
	defer file.Close()

	byteValue, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("ошибка при чтении файла: %w", err)
	}

	err = json.Unmarshal(byteValue, &repo.words)
	if err != nil {
		return fmt.Errorf("ошибка при парсинге JSON: %w", err)
	}

	return nil
}

func (repo *InMemoryWordRepository) GetRandomWord(category, difficulty string) (*domain.Word, error) {
	var filteredWords []domain.Word

	for _, word := range repo.words {
		if (category == "" || word.Category == category) && (difficulty == "" || word.Difficulty == difficulty) {
			filteredWords = append(filteredWords, word)
		}
	}

	if len(filteredWords) == 0 {
		return nil, fmt.Errorf("слово не найдено: %w", &domain.ErrWordNotFound{
			Category:   category,
			Difficulty: difficulty,
		})
	}

	randomIndexBig, err := rand.Int(rand.Reader, big.NewInt(int64(len(filteredWords))))
	if err != nil {
		return nil, fmt.Errorf("ошибка при генерации случайного числа: %w", err)
	}

	randomIndex := int(randomIndexBig.Int64())

	return &filteredWords[randomIndex], nil
}
