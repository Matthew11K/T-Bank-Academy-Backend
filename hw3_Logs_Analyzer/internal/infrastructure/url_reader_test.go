package infrastructure_test

import (
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/central-university-dev/backend_academy_2024_project_3-go-Matthew11K/internal/domain"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/central-university-dev/backend_academy_2024_project_3-go-Matthew11K/internal/domain/mocks"
	"github.com/central-university-dev/backend_academy_2024_project_3-go-Matthew11K/internal/infrastructure"
)

func TestURLReader_Scanner(t *testing.T) {
	mockClient := new(mocks.HTTPClient)

	response := &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(strings.NewReader("test log line\n")),
	}
	mockClient.On("Get", "http://example.com/logs").Return(response, nil)

	factory := infrastructure.NewURLReaderFactory(mockClient)
	reader, err := factory.NewReader("http://example.com/logs")
	require.NoError(t, err)

	scanner, err := reader.Scanner()
	require.NoError(t, err)
	defer reader.Close()

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	assert.NoError(t, scanner.Err())
	assert.Equal(t, []string{"test log line"}, lines)
}

func TestURLReader_Scanner_HTTPError(t *testing.T) {
	mockClient := new(mocks.HTTPClient)

	mockClient.On("Get", "http://example.com/logs").Return(nil, errors.New("http error"))

	factory := infrastructure.NewURLReaderFactory(mockClient)
	reader, err := factory.NewReader("http://example.com/logs")
	require.NoError(t, err)

	_, err = reader.Scanner()
	assert.Error(t, err)
	assert.IsType(t, &domain.ErrURLFetch{}, err)
}
