package domain_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/central-university-dev/backend_academy_2024_project_3-go-Matthew11K/internal/domain"
)

func TestLogRecord_MatchesFilter(t *testing.T) {
	record := &domain.LogRecord{
		Request:       "GET /index.html HTTP/1.1",
		HTTPUserAgent: "Mozilla/5.0",
	}

	assert.True(t, record.MatchesFilter(domain.FieldMethod, "GET"))
	assert.False(t, record.MatchesFilter(domain.FieldMethod, "POST"))

	assert.True(t, record.MatchesFilter(domain.FieldAgent, "Mozilla.*"))
	assert.False(t, record.MatchesFilter(domain.FieldAgent, "Chrome.*"))
}

func TestLogRecord_GetMethod(t *testing.T) {
	testCases := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Valid GET request",
			request:  "GET /index.html HTTP/1.1",
			expected: "GET",
		},
		{
			name:     "Valid POST request",
			request:  "POST /api/data HTTP/1.1",
			expected: "POST",
		},
		{
			name:     "Empty request",
			request:  "",
			expected: "Неизвестный метод",
		},
		{
			name:     "Request with extra spaces",
			request:  "  PUT   /update HTTP/1.1",
			expected: "PUT",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			record := &domain.LogRecord{
				Request: tc.request,
			}
			method := record.GetMethod()
			assert.Equal(t, tc.expected, method)
		})
	}
}
