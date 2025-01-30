package ui

import (
	"fmt"

	"github.com/central-university-dev/backend_academy_2024_project_2-go-Matthew11K/internal/domain"
)

func DisplayMazeWithPath(maze *domain.Maze, path []domain.Coordinate) {
	fmt.Println("Лабиринт с путем:")
	fmt.Println(maze.StringWithPath(path))
}

func DisplayNoPath() {
	fmt.Println("Путь от старта до финиша не найден.")
}
