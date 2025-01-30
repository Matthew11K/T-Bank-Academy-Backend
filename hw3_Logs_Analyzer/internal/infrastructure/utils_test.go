package infrastructure_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/central-university-dev/backend_academy_2024_project_3-go-Matthew11K/internal/infrastructure"
)

func TestFileResolver_ResolveFiles(t *testing.T) {
	t.Run("Существующий_файл", func(t *testing.T) {
		tempDir, err := os.MkdirTemp("", "test_logs")
		require.NoError(t, err)
		defer os.RemoveAll(tempDir)

		tempFile, err := os.CreateTemp(tempDir, "test_log_*.log")
		require.NoError(t, err)

		tempFileName := tempFile.Name()
		tempFile.Close()

		resolver := infrastructure.NewFileResolver()

		files, err := resolver.ResolveFiles(tempFileName)
		require.NoError(t, err)
		assert.Len(t, files, 1)
		assert.Equal(t, tempFileName, files[0])
	})

	t.Run("Несуществующий_файл", func(t *testing.T) {
		resolver := infrastructure.NewFileResolver()

		files, err := resolver.ResolveFiles("non_existing_file.log")
		require.NoError(t, err)
		assert.Len(t, files, 0)
	})

	t.Run("Некорректный_шаблон", func(t *testing.T) {
		resolver := infrastructure.NewFileResolver()

		_, err := resolver.ResolveFiles("[invalid_pattern")
		assert.Error(t, err)
	})
}
