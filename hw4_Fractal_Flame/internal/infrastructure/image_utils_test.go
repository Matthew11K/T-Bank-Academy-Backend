package infrastructure_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/central-university-dev/backend_academy_2024_project_4-go-Matthew11K/internal/domain"
	"github.com/central-university-dev/backend_academy_2024_project_4-go-Matthew11K/internal/infrastructure"
)

func TestImageSaverImpl_Save_Success(t *testing.T) {
	fi := domain.NewFractalImage(10, 10)
	fi.UpdatePixel(5, 5, [3]float64{1.0, 0.0, 0.0})

	saver := &infrastructure.ImageSaverImpl{}

	tempDir := t.TempDir()
	filename := filepath.Join(tempDir, "test_output.png")

	defer os.Remove(filename)

	err := saver.Save(fi, filename, 2.2)

	require.NoError(t, err)

	info, err := os.Stat(filename)
	require.NoError(t, err)
	require.True(t, info.Size() > 0, "Файл должен быть не пустым")
}
