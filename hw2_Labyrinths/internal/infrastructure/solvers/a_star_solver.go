package solvers

import (
	"fmt"

	"github.com/central-university-dev/backend_academy_2024_project_2-go-Matthew11K/internal/domain"
	"github.com/central-university-dev/backend_academy_2024_project_2-go-Matthew11K/pkg/priorityqueue"
)

type AStarSolver struct{}

func NewAStarSolver() *AStarSolver {
	return &AStarSolver{}
}

func (s *AStarSolver) Solve(maze *domain.Maze, start, end domain.Coordinate) ([]domain.Coordinate, error) {
	if !start.IsValid(maze.Height, maze.Width) || !end.IsValid(maze.Height, maze.Width) {
		return nil, fmt.Errorf("недопустимые координаты: %w", &domain.InvalidCoordinateError{})
	}

	cameFrom := make(map[domain.Coordinate]*domain.Coordinate)
	costSoFar := make(map[domain.Coordinate]float64)
	costSoFar[start] = 0

	pq := priorityqueue.NewPriorityQueue()
	pq.Push(&priorityqueue.Item{
		Coordinate: start,
		Priority:   0,
	})

	for pq.Len() > 0 {
		currentItem := pq.Pop().(*priorityqueue.Item)
		current := currentItem.Coordinate

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

			if !neighbor.IsValid(maze.Height, maze.Width) {
				continue
			}

			cell := maze.Grid[neighbor.Row][neighbor.Col]
			if cell.Type == domain.Wall {
				continue
			}

			movementCost := cell.GetMovementCostAStar()

			newCost := costSoFar[current] + movementCost
			if prevCost, ok := costSoFar[neighbor]; !ok || newCost < prevCost {
				costSoFar[neighbor] = newCost
				priority := newCost + neighbor.Heuristic(end)
				pq.Push(&priorityqueue.Item{
					Coordinate: neighbor,
					Priority:   priority,
				})

				cameFrom[neighbor] = &current
			}
		}
	}

	return nil, fmt.Errorf("путь не найден: %w", &domain.NoPathFoundError{})
}
