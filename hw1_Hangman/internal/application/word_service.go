package application

import (
	"fmt"

	"github.com/es-debug/backend-academy-2024-go-template/internal/domain"
)

type WordService struct {
	wordRepo WordRepository
}

type WordServiceInterface interface {
	GetRandomWord(category, difficulty string) (*domain.Word, error)
	GetHint(word *domain.Word) string
}

var _ WordServiceInterface = (*WordService)(nil)

type WordRepository interface {
	GetRandomWord(category, difficulty string) (*domain.Word, error)
}

func NewWordService(repo WordRepository) *WordService {
	return &WordService{
		wordRepo: repo,
	}
}

func (s *WordService) GetRandomWord(category, difficulty string) (*domain.Word, error) {
	word, err := s.wordRepo.GetRandomWord(category, difficulty)
	if err != nil {
		return nil, fmt.Errorf("ошибка в WordService при получении случайного слова: %w", err)
	}

	return word, nil
}

func (s *WordService) GetHint(word *domain.Word) string {
	return word.Hint
}
