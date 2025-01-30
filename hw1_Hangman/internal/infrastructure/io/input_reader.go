package io

import (
	"bufio"
	"os"
	"strings"
)

type InputReader interface {
	ReadInput() (string, error)
}

type ConsoleInputReader struct {
}

func NewConsoleInputReader() *ConsoleInputReader {
	return &ConsoleInputReader{}
}

func (r *ConsoleInputReader) ReadInput() (string, error) {
	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')

	if err != nil {
		return "", err
	}

	input = strings.TrimSpace(input)

	return input, nil
}
