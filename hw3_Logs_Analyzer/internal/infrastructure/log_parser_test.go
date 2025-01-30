package infrastructure_test

import (
	"testing"

	"github.com/central-university-dev/backend_academy_2024_project_3-go-Matthew11K/internal/domain"

	"github.com/stretchr/testify/assert"

	"github.com/central-university-dev/backend_academy_2024_project_3-go-Matthew11K/internal/infrastructure"
)

func TestLogParser_Parse(t *testing.T) {
	parser := infrastructure.NewLogParser()

	line := `127.0.0.1 - - [31/Aug/2024:12:00:00 +0000] "GET /index.html HTTP/1.1" 200 1024 "-" "Mozilla/5.0"`
	record, err := parser.Parse(line)

	assert.NoError(t, err)
	assert.Equal(t, "127.0.0.1", record.RemoteAddr)
	assert.Equal(t, "-", record.RemoteUser)
	assert.Equal(t, "GET /index.html HTTP/1.1", record.Request)
	assert.Equal(t, 200, record.Status)
	assert.Equal(t, 1024, record.BodyBytesSent)
	assert.Equal(t, "-", record.HTTPReferer)
	assert.Equal(t, "Mozilla/5.0", record.HTTPUserAgent)
}

func TestLogParser_Parse_InvalidLine(t *testing.T) {
	parser := infrastructure.NewLogParser()
	line := `Invalid log line`
	_, err := parser.Parse(line)
	assert.Error(t, err)
	assert.IsType(t, &domain.ErrInvalidLogFormat{}, err)
}

func TestLogParser_Parse_MissingFields(t *testing.T) {
	parser := infrastructure.NewLogParser()
	line := `127.0.0.1 - - [31/Aug/2024:12:00:00 +0000] "GET /index.html HTTP/1.1"`
	_, err := parser.Parse(line)
	assert.Error(t, err)
	assert.IsType(t, &domain.ErrInvalidLogFormat{}, err)
}

func TestLogParser_Parse_InvalidTimeFormat(t *testing.T) {
	parser := infrastructure.NewLogParser()
	line := `127.0.0.1 - - [invalid-time] "GET /index.html HTTP/1.1" 200 1024 "-" "Mozilla/5.0"`
	_, err := parser.Parse(line)
	assert.Error(t, err)
	assert.IsType(t, &domain.ErrInvalidTimeFormat{}, err)

	expectedErrMsg := "неверный формат времени"
	assert.Contains(t, err.Error(), expectedErrMsg, "Ошибка должна содержать 'неверный формат времени'")
}
