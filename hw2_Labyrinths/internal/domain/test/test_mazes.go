package domain_test

import (
	"github.com/central-university-dev/backend_academy_2024_project_2-go-Matthew11K/internal/domain"
)

func GetTestMaze() *domain.Maze {
	grid := [][]*domain.Cell{
		{
			{Type: domain.Passage}, {Type: domain.Wall}, {Type: domain.Passage}, {Type: domain.Wall}, {Type: domain.Passage},
		},
		{
			{Type: domain.Passage}, {Type: domain.Wall}, {Type: domain.Passage}, {Type: domain.Wall}, {Type: domain.Passage},
		},
		{
			{Type: domain.Passage}, {Type: domain.Passage}, {Type: domain.Passage}, {Type: domain.Wall}, {Type: domain.Passage},
		},
		{
			{Type: domain.Wall}, {Type: domain.Wall}, {Type: domain.Passage}, {Type: domain.Wall}, {Type: domain.Passage},
		},
		{
			{Type: domain.Passage}, {Type: domain.Passage}, {Type: domain.Passage}, {Type: domain.Passage}, {Type: domain.Passage},
		},
	}

	maze := &domain.Maze{
		Height: 5,
		Width:  5,
		Grid:   grid,
	}

	maze.Grid[0][0].Type = domain.Start
	maze.Grid[4][4].Type = domain.End

	return maze
}
