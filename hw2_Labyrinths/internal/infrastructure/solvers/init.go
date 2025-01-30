package solvers

import (
	"fmt"

	"github.com/central-university-dev/backend_academy_2024_project_2-go-Matthew11K/internal/application"
)

func InitSolver(solverType string) (application.MazeSolver, error) {
	switch solverType {
	case "bfs":
		return NewBFSSolver(), nil
	case "astar":
		return NewAStarSolver(), nil
	case "dijkstra":
		return NewDijkstraSolver(), nil
	default:
		return nil, fmt.Errorf("неизвестный тип решателя: %s", solverType)
	}
}
