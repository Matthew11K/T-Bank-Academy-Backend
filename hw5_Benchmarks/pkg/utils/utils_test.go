package utils_test

import (
	"testing"

	"github.com/central-university-dev/backend_academy_2024_project_5-go-Matthew11K/pkg/utils"

	"github.com/stretchr/testify/assert"
)

func TestCountSubstringsStringsCount(t *testing.T) {
	const (
		s      = "hello world, hello universe"
		substr = "hello"
	)

	count := utils.CountSubstringsStringsCount(s, substr)

	assert.Equal(t, 2, count)
}

func TestCountSubstringsRegex(t *testing.T) {
	s := "hello world, hello universe"
	substr := "hello"

	count, err := utils.CountSubstringsRegex(s, substr)

	assert.NoError(t, err)
	assert.Equal(t, 2, count)
}

func TestCountSubstringsManual(t *testing.T) {
	s := "hello world, hello universe"
	substr := "hello"

	count := utils.CountSubstringsManual(s, substr)

	assert.Equal(t, 2, count)
}
