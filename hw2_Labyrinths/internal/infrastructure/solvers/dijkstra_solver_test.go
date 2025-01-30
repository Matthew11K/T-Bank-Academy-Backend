package solvers_test

import (
	"testing"

	"github.com/central-university-dev/backend_academy_2024_project_2-go-Matthew11K/internal/domain"
	domain_test "github.com/central-university-dev/backend_academy_2024_project_2-go-Matthew11K/internal/domain/test"
	"github.com/central-university-dev/backend_academy_2024_project_2-go-Matthew11K/internal/infrastructure/solvers"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDijkstraSolver_Solve(t *testing.T) {
	solver := solvers.NewDijkstraSolver()
	maze := domain_test.GetTestMaze()

	start := domain.Coordinate{Row: 0, Col: 0}
	end := domain.Coordinate{Row: 2, Col: 2}

	path, err := solver.Solve(maze, start, end)
	require.NoError(t, err)
	require.NotNil(t, path)

	assert.Equal(t, start, path[0], "Путь должен начинаться с начальной координаты")
	assert.Equal(t, end, path[len(path)-1], "Путь должен заканчиваться в конечной координате")
}
