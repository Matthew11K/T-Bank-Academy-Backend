package infrastructure

import (
	"fmt"
	"io"

	"golang.org/x/exp/slog"
)

type IOAdapter struct {
	writer io.Writer
	logger *slog.Logger
}

func NewIOAdapter(writer io.Writer, logger *slog.Logger) *IOAdapter {
	return &IOAdapter{
		writer: writer,
		logger: logger,
	}
}

func (a *IOAdapter) Output(content string) error {
	_, err := a.writer.Write([]byte(content))
	if err != nil {
		a.logger.Error("Ошибка при выводе результата", slog.Any("error", err))
		return fmt.Errorf("ошибка вывода: %w", err)
	}

	return nil
}
