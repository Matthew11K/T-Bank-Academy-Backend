package application_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/es-debug/backend-academy-2024-go-template/internal/application"
	"github.com/es-debug/backend-academy-2024-go-template/internal/application/mocks"
	"github.com/es-debug/backend-academy-2024-go-template/internal/domain"
)

func TestWordService_GetRandomWord(t *testing.T) {
	mockRepo := new(mocks.WordRepository)
	wordService := application.NewWordService(mockRepo)

	expectedWord := &domain.Word{Value: "тест"}

	mockRepo.On("GetRandomWord", "", "").Return(expectedWord, nil)

	word, err := wordService.GetRandomWord("", "")
	assert.NoError(t, err)
	assert.Equal(t, expectedWord, word)

	mockRepo.AssertExpectations(t)
}

func TestWordService_GetRandomWord_Error(t *testing.T) {
	mockRepo := new(mocks.WordRepository)
	wordService := application.NewWordService(mockRepo)

	mockRepo.On("GetRandomWord", "unknown", "hard").Return(nil, errors.New("слово не найдено"))

	word, err := wordService.GetRandomWord("unknown", "hard")
	assert.Error(t, err)
	assert.Nil(t, word)

	mockRepo.AssertExpectations(t)
}

func TestWordService_GetHint(t *testing.T) {
	wordService := application.NewWordService(nil)

	word := &domain.Word{Hint: "подсказка"}

	hint := wordService.GetHint(word)
	assert.Equal(t, "подсказка", hint)
}
