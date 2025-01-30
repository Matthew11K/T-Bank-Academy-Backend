package data_test

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/es-debug/backend-academy-2024-go-template/internal/domain"
	"github.com/es-debug/backend-academy-2024-go-template/internal/infrastructure/data"
)

func TestNewInMemoryWordRepository_Success(t *testing.T) {
	file, err := os.CreateTemp("", "words*.json")
	assert.NoError(t, err)
	defer os.Remove(file.Name())

	words := []domain.Word{
		{Value: "test", Category: "category1", Difficulty: "easy", Hint: "hint1"},
		{Value: "example", Category: "category2", Difficulty: "medium", Hint: "hint2"},
	}
	err = json.NewEncoder(file).Encode(words)
	assert.NoError(t, err)
	file.Close()

	repo, err := data.NewInMemoryWordRepository(file.Name())
	assert.NoError(t, err)
	assert.NotNil(t, repo)

	word, err := repo.GetRandomWord("category1", "easy")
	assert.NoError(t, err)
	assert.NotNil(t, word)
	assert.Equal(t, "test", word.Value)
}
