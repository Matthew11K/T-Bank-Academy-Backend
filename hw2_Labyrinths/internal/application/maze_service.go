package application

import (
	"fmt"
	"os"

	"github.com/central-university-dev/backend_academy_2024_project_2-go-Matthew11K/internal/config"
	"github.com/central-university-dev/backend_academy_2024_project_2-go-Matthew11K/internal/domain"
	"github.com/central-university-dev/backend_academy_2024_project_2-go-Matthew11K/internal/ui"
	"golang.org/x/exp/slog"
)

type MazeGenerator interface {
	Generate(height, width int, swampFreq, sandFreq, coinFreq float64) (*domain.Maze, error)
}

type MazeSolver interface {
	Solve(maze *domain.Maze, start, end domain.Coordinate) ([]domain.Coordinate, error)
}

type MazeService struct {
	generator MazeGenerator
	solver    MazeSolver
	logger    *slog.Logger
}

func NewMazeService(generator MazeGenerator, solver MazeSolver, logger *slog.Logger) *MazeService {
	return &MazeService{
		generator: generator,
		solver:    solver,
		logger:    logger,
	}
}

func (s *MazeService) SetupMaze(cfg *config.Config) (
	maze *domain.Maze,
	start domain.Coordinate,
	end domain.Coordinate,
	err error,
) {
	maze, err = s.generator.Generate(cfg.Height, cfg.Width, cfg.SwampFrequency, cfg.SandFrequency, cfg.CoinFrequency)
	if err != nil {
		err = fmt.Errorf("не удалось сгенерировать лабиринт: %w", err)
		return
	}

	start, end, err = maze.SetupStartEnd(cfg)
	if err != nil {
		return
	}

	return
}

func (s *MazeService) SolveMaze(maze *domain.Maze, start, end domain.Coordinate) ([]domain.Coordinate, error) {
	return s.solver.Solve(maze, start, end)
}

func (s *MazeService) Start(cfg *config.Config) {
	maze, start, end, err := s.SetupMaze(cfg)
	if err != nil {
		s.logger.Error("Ошибка настройки лабиринта", slog.Any("error", err))
		os.Exit(1)
	}

	path, err := s.SolveMaze(maze, start, end)
	if err != nil {
		if _, ok := err.(*domain.NoPathFoundError); ok {
			ui.DisplayNoPath()
		} else {
			s.logger.Error("Не удалось решить лабиринт", slog.Any("error", err))
			os.Exit(1)
		}
	} else {
		ui.DisplayMazeWithPath(maze, path)
	}
}
