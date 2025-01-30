package solvers

import (
	"fmt"

	"github.com/central-university-dev/backend_academy_2024_project_2-go-Matthew11K/internal/domain"
)

type BFSSolver struct{}

func NewBFSSolver() *BFSSolver {
	return &BFSSolver{}
}

func (s *BFSSolver) Solve(maze *domain.Maze, start, end domain.Coordinate) ([]domain.Coordinate, error) {
	if !start.IsValid(maze.Height, maze.Width) || !end.IsValid(maze.Height, maze.Width) {
		return nil, fmt.Errorf("недопустимые координаты: %w", &domain.InvalidCoordinateError{})
	}

	visited := make([][]bool, maze.Height)
	for i := range visited {
		visited[i] = make([]bool, maze.Width)
	}

	queue := []domain.Coordinate{start}
	cameFrom := make(map[domain.Coordinate]*domain.Coordinate)
	visited[start.Row][start.Col] = true

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current == end {
			return maze.ReconstructPath(cameFrom, current), nil
		}

		for _, offset := range []domain.Coordinate{
			{Row: -1, Col: 0},
			{Row: 1, Col: 0},
			{Row: 0, Col: -1},
			{Row: 0, Col: 1},
		} {
			neighbor := domain.Coordinate{Row: current.Row + offset.Row, Col: current.Col + offset.Col}

			if neighbor.IsValid(maze.Height, maze.Width) {
				cellType := maze.Grid[neighbor.Row][neighbor.Col].Type
				if cellType == domain.Wall {
					continue
				}

				if !visited[neighbor.Row][neighbor.Col] {
					queue = append(queue, neighbor)
					visited[neighbor.Row][neighbor.Col] = true
					cameFrom[neighbor] = &current
				}
			}
		}
	}

	return nil, fmt.Errorf("путь не найден: %w", &domain.NoPathFoundError{})
}
