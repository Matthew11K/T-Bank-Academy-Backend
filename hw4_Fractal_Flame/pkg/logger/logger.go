package logger

import (
	"io"

	"golang.org/x/exp/slog"
)

func NewLogger(w io.Writer) *slog.Logger {
	handler := slog.NewJSONHandler(w, nil)
	return slog.New(handler)
}
