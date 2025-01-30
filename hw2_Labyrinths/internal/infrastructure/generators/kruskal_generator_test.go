package generators_test

import (
	"testing"

	"github.com/central-university-dev/backend_academy_2024_project_2-go-Matthew11K/internal/domain"
	"github.com/central-university-dev/backend_academy_2024_project_2-go-Matthew11K/internal/infrastructure/generators"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestKruskalGenerator_Generate(t *testing.T) {
	generator := generators.NewKruskalGenerator()
	maze, err := generator.Generate(10, 10, 0.1, 0.1, 0.1)
	require.NoError(t, err)
	require.NotNil(t, maze)
	assert.Equal(t, 21, maze.Height)
	assert.Equal(t, 21, maze.Width)

	hasPassage := false

	for i := 0; i < maze.Height; i++ {
		for j := 0; j < maze.Width; j++ {
			if maze.Grid[i][j].Type != domain.Wall {
				hasPassage = true
				break
			}
		}
	}

	assert.True(t, hasPassage, "В лабиринте должны быть проходы")
}
