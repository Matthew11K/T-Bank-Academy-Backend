package generators

import (
	"math/rand"
	"time"

	"github.com/central-university-dev/backend_academy_2024_project_2-go-Matthew11K/internal/domain"
)

type WilsonGenerator struct {
	rnd *rand.Rand
}

func NewWilsonGenerator() *WilsonGenerator {
	return &WilsonGenerator{
		rnd: rand.New(rand.NewSource(time.Now().UnixNano())), // #nosec G404
	}
}

var directions = []domain.Coordinate{
	{Row: -2, Col: 0},
	{Row: 2, Col: 0},
	{Row: 0, Col: -2},
	{Row: 0, Col: 2},
}

func (g *WilsonGenerator) Generate(height, width int, swampFreq, sandFreq, coinFreq float64) (*domain.Maze, error) {
	mazeHeight := height*2 + 1
	mazeWidth := width*2 + 1

	maze := g.initializeMaze(mazeHeight, mazeWidth)

	cells := []domain.Coordinate{}

	for i := 1; i < mazeHeight; i += 2 {
		for j := 1; j < mazeWidth; j += 2 {
			cells = append(cells, domain.Coordinate{Row: i, Col: j})
		}
	}

	visited := make(map[domain.Coordinate]bool)
	initialIdx := g.rnd.Intn(len(cells))
	initialCell := cells[initialIdx]
	visited[initialCell] = true
	maze.Grid[initialCell.Row][initialCell.Col].Type = domain.Passage

	for len(visited) < len(cells) {
		unvisitedCells := g.getUnvisitedCells(cells, visited)
		if len(unvisitedCells) == 0 {
			break
		}

		currentIdx := g.rnd.Intn(len(unvisitedCells))
		current := unvisitedCells[currentIdx]

		path := []domain.Coordinate{current}

		for !visited[current] {
			dir := directions[g.rnd.Intn(len(directions))]
			next := domain.Coordinate{Row: current.Row + dir.Row, Col: current.Col + dir.Col}

			if next.Row < 1 || next.Row >= mazeHeight-1 || next.Col < 1 || next.Col >= mazeWidth-1 {
				continue
			}

			if idx := g.indexOf(path, next); idx != -1 {
				path = path[:idx+1]
			} else {
				path = append(path, next)
			}

			current = next
		}

		g.carvePath(maze, path, visited)
	}

	g.assignCellTypes(maze, swampFreq, sandFreq, coinFreq)

	return maze, nil
}

func (g *WilsonGenerator) initializeMaze(height, width int) *domain.Maze {
	maze := &domain.Maze{
		Height: height,
		Width:  width,
		Grid:   make([][]*domain.Cell, height),
	}
	for i := 0; i < height; i++ {
		maze.Grid[i] = make([]*domain.Cell, width)
		for j := 0; j < width; j++ {
			maze.Grid[i][j] = &domain.Cell{
				Coordinate: domain.Coordinate{Row: i, Col: j},
				Type:       domain.Wall,
			}
		}
	}

	return maze
}

func (g *WilsonGenerator) getUnvisitedCells(cells []domain.Coordinate, visited map[domain.Coordinate]bool) []domain.Coordinate {
	var unvisited []domain.Coordinate

	for _, cell := range cells {
		if !visited[cell] {
			unvisited = append(unvisited, cell)
		}
	}

	return unvisited
}

func (g *WilsonGenerator) carvePath(maze *domain.Maze, path []domain.Coordinate, visited map[domain.Coordinate]bool) {
	for i := 0; i < len(path)-1; i++ {
		c1 := path[i]
		c2 := path[i+1]

		maze.Grid[c1.Row][c1.Col].Type = domain.Passage
		maze.Grid[c2.Row][c2.Col].Type = domain.Passage

		wallRow := (c1.Row + c2.Row) / 2
		wallCol := (c1.Col + c2.Col) / 2
		maze.Grid[wallRow][wallCol].Type = domain.Passage

		visited[c1] = true
		visited[c2] = true
	}
}

func (g *WilsonGenerator) assignCellTypes(maze *domain.Maze, swampFreq, sandFreq, coinFreq float64) {
	for i := 1; i < maze.Height; i += 2 {
		for j := 1; j < maze.Width; j += 2 {
			cell := maze.Grid[i][j]
			if cell.Type == domain.Passage {
				randValue := g.rnd.Float64()

				switch {
				case randValue < swampFreq:
					cell.Type = domain.Swamp
				case randValue < swampFreq+sandFreq:
					cell.Type = domain.Sand
				case randValue < swampFreq+sandFreq+coinFreq:
					cell.Type = domain.Coin
				}
			}
		}
	}
}

func (g *WilsonGenerator) indexOf(path []domain.Coordinate, coord domain.Coordinate) int {
	for i, c := range path {
		if c == coord {
			return i
		}
	}

	return -1
}
