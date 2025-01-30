package application_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/central-university-dev/backend_academy_2024_project_2-go-Matthew11K/internal/application"
	"github.com/central-university-dev/backend_academy_2024_project_2-go-Matthew11K/internal/config"
	"github.com/central-university-dev/backend_academy_2024_project_2-go-Matthew11K/internal/domain"
	"github.com/central-university-dev/backend_academy_2024_project_2-go-Matthew11K/internal/domain/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/exp/slog"
)

func TestMazeService_SetupMaze_Success(t *testing.T) {
	mockGenerator := new(mocks.MazeGenerator)
	mockSolver := new(mocks.MazeSolver)
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{}))

	service := application.NewMazeService(mockGenerator, mockSolver, logger)

	cfg := config.Config{
		Height:         10,
		Width:          10,
		GeneratorType:  "prim",
		SwampFrequency: 0.1,
		SandFrequency:  0.1,
		CoinFrequency:  0.1,
		StartRow:       -1,
		StartCol:       -1,
		EndRow:         -1,
		EndCol:         -1,
	}

	expectedMaze := &domain.Maze{
		Height: 21,
		Width:  21,
		Grid:   make([][]*domain.Cell, 21),
	}
	for i := 0; i < 21; i++ {
		expectedMaze.Grid[i] = make([]*domain.Cell, 21)
		for j := 0; j < 21; j++ {
			expectedMaze.Grid[i][j] = &domain.Cell{
				Coordinate: domain.Coordinate{Row: i, Col: j},
				Type:       domain.Wall,
			}
		}
	}

	mockGenerator.On("Generate", 10, 10, 0.1, 0.1, 0.1).Return(expectedMaze, nil)

	expectedStart := domain.Coordinate{Row: 0, Col: 0}
	expectedEnd := domain.Coordinate{Row: 20, Col: 20}

	expectedMaze.SetCellType(expectedStart, domain.Start)
	expectedMaze.SetCellType(expectedEnd, domain.End)

	maze, start, end, err := service.SetupMaze(&cfg)
	require.NoError(t, err)
	require.NotNil(t, maze)
	assert.Equal(t, expectedMaze, maze)
	assert.Equal(t, expectedStart, start)
	assert.Equal(t, expectedEnd, end)

	mockGenerator.AssertExpectations(t)
}

func TestMazeService_SetupMaze_GeneratorError(t *testing.T) {
	mockGenerator := new(mocks.MazeGenerator)
	mockSolver := new(mocks.MazeSolver)
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{}))

	service := application.NewMazeService(mockGenerator, mockSolver, logger)

	cfg := config.Config{
		Height:         10,
		Width:          10,
		GeneratorType:  "prim",
		SwampFrequency: 0.1,
		SandFrequency:  0.1,
		CoinFrequency:  0.1,
		StartRow:       -1,
		StartCol:       -1,
		EndRow:         -1,
		EndCol:         -1,
	}

	mockGenerator.On("Generate", 10, 10, 0.1, 0.1, 0.1).Return((*domain.Maze)(nil), fmt.Errorf("ошибка генерации"))

	maze, start, end, err := service.SetupMaze(&cfg)
	require.Error(t, err)
	assert.Nil(t, maze)
	assert.Equal(t, domain.Coordinate{}, start)
	assert.Equal(t, domain.Coordinate{}, end)

	mockGenerator.AssertExpectations(t)
}

func TestMazeService_SolveMaze_Success(t *testing.T) {
	mockGenerator := new(mocks.MazeGenerator)
	mockSolver := new(mocks.MazeSolver)
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{}))

	service := application.NewMazeService(mockGenerator, mockSolver, logger)

	maze := &domain.Maze{
		Height: 3,
		Width:  3,
		Grid: [][]*domain.Cell{
			{
				{Coordinate: domain.Coordinate{Row: 0, Col: 0}, Type: domain.Start},
				{Coordinate: domain.Coordinate{Row: 0, Col: 1}, Type: domain.Wall},
				{Coordinate: domain.Coordinate{Row: 0, Col: 2}, Type: domain.End},
			},
			{
				{Coordinate: domain.Coordinate{Row: 1, Col: 0}, Type: domain.Passage},
				{Coordinate: domain.Coordinate{Row: 1, Col: 1}, Type: domain.Wall},
				{Coordinate: domain.Coordinate{Row: 1, Col: 2}, Type: domain.Passage},
			},
			{
				{Coordinate: domain.Coordinate{Row: 2, Col: 0}, Type: domain.Passage},
				{Coordinate: domain.Coordinate{Row: 2, Col: 1}, Type: domain.Passage},
				{Coordinate: domain.Coordinate{Row: 2, Col: 2}, Type: domain.Passage},
			},
		},
	}

	start := domain.Coordinate{Row: 0, Col: 0}
	end := domain.Coordinate{Row: 2, Col: 2}
	expectedPath := []domain.Coordinate{
		{Row: 0, Col: 0},
		{Row: 1, Col: 0},
		{Row: 2, Col: 0},
		{Row: 2, Col: 1},
		{Row: 2, Col: 2},
	}

	mockSolver.On("Solve", maze, start, end).Return(expectedPath, nil)

	path, err := service.SolveMaze(maze, start, end)
	require.NoError(t, err)
	require.NotNil(t, path)
	assert.Equal(t, expectedPath, path)

	mockSolver.AssertExpectations(t)
}
