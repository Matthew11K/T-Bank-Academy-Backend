package domain_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/central-university-dev/backend_academy_2024_project_3-go-Matthew11K/internal/domain"
)

func TestParseTime_ISO8601(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected time.Time
		hasError bool
	}{
		{
			name:     "RFC3339",
			input:    "2024-11-07T15:04:05Z",
			expected: time.Date(2024, 11, 7, 15, 4, 5, 0, time.UTC),
			hasError: false,
		},
		{
			name:     "RFC3339 with nanoseconds",
			input:    "2024-11-07T15:04:05.123456789Z",
			expected: time.Date(2024, 11, 7, 15, 4, 5, 123456789, time.UTC),
			hasError: false,
		},
		{
			name:     "Date only",
			input:    "2024-11-07",
			expected: time.Date(2024, 11, 7, 0, 0, 0, 0, time.UTC),
			hasError: false,
		},
		{
			name:     "Invalid format",
			input:    "invalid-time-format",
			expected: time.Time{},
			hasError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tw, err := domain.ParseTime(tc.input)
			if tc.hasError {
				assert.Error(t, err, "Input: %s", tc.input)
			} else {
				assert.NoError(t, err, "Input: %s", tc.input)
				assert.True(t, tw.Time.Equal(tc.expected), "Input: %s", tc.input)
			}
		})
	}
}

func TestParseNginxTime(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected time.Time
		hasError bool
	}{
		{
			name:     "Valid NGINX time UTC",
			input:    "31/Aug/2024:12:00:00 +0000",
			expected: time.Date(2024, 8, 31, 12, 0, 0, 0, time.FixedZone("", 0)),
			hasError: false,
		},
		{
			name:     "Valid NGINX time with negative timezone",
			input:    "01/Jan/2023:23:59:59 -0500",
			expected: time.Date(2023, 1, 1, 23, 59, 59, 0, time.FixedZone("", -5*3600)),
			hasError: false,
		},
		{
			name:     "Invalid NGINX time",
			input:    "Invalid/Nginx/Time/Format",
			expected: time.Time{},
			hasError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tw, err := domain.ParseNginxTime(tc.input)
			if tc.hasError {
				assert.Error(t, err, "Input: %s", tc.input)
			} else {
				assert.NoError(t, err, "Input: %s", tc.input)
				assert.True(t, tw.Time.Equal(tc.expected), "Input: %s", tc.input)
			}
		})
	}
}
