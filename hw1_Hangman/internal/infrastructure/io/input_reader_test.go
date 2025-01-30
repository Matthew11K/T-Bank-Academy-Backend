package io_test

import (
	"os"
	"testing"

	ioMocks "github.com/es-debug/backend-academy-2024-go-template/internal/infrastructure/io/mocks"

	"github.com/stretchr/testify/assert"

	"github.com/es-debug/backend-academy-2024-go-template/internal/infrastructure/io"
)

func TestConsoleInputReader_ReadInput_Success(t *testing.T) {
	mockInput := new(ioMocks.InputReader)

	mockInput.On("ReadInput").Return("тестовый ввод", nil).Once()

	input, err := mockInput.ReadInput()

	assert.NoError(t, err)
	assert.Equal(t, "тестовый ввод", input)
	mockInput.AssertExpectations(t)
}

func TestConsoleInputReader_ReadInput_Error(t *testing.T) {
	oldStdin := os.Stdin

	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("не удалось создать pipe: %v", err)
	}

	os.Stdin = r

	w.Close()

	defer func() {
		os.Stdin = oldStdin

		r.Close()
	}()

	inputReader := io.NewConsoleInputReader()
	_, err = inputReader.ReadInput()
	assert.Error(t, err)
}

func TestConsoleInputReader_ReadInput_EmptyInput(t *testing.T) {
	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }()

	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("не удалось создать pipe: %v", err)
	}

	os.Stdin = r

	_, err = w.WriteString("\n")
	if err != nil {
		t.Fatalf("не удалось записать в pipe: %v", err)
	}

	w.Close()

	inputReader := io.NewConsoleInputReader()
	input, err := inputReader.ReadInput()
	assert.NoError(t, err)
	assert.Empty(t, input)
}
