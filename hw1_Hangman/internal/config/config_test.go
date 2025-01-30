package config_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/es-debug/backend-academy-2024-go-template/internal/config"
)

func TestLoadConfig(t *testing.T) {
	cfg := config.LoadConfig()
	assert.NotNil(t, cfg)
	assert.Equal(t, 10, cfg.MaxAttemptsEasy)
	assert.Equal(t, 7, cfg.MaxAttemptsMedium)
	assert.Equal(t, 5, cfg.MaxAttemptsHard)
}
