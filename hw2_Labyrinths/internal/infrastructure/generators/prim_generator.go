package generators

import (
	"math/rand"
	"time"

	"github.com/central-university-dev/backend_academy_2024_project_2-go-Matthew11K/internal/domain"
)

type PrimGenerator struct {
	rnd *rand.Rand
}

func NewPrimGenerator() *PrimGenerator {
	return &PrimGenerator{
		rnd: rand.New(rand.NewSource(time.Now().UnixNano())), // #nosec G404
	}
}

type Wall struct {
	Between *domain.Cell
	Next    *domain.Cell
}

func (g *PrimGenerator) Generate(height, width int, swampFreq, sandFreq, coinFreq float64) (*domain.Maze, error) {
	maze := g.initializeMaze(height, width)

	startRow, startCol := g.rnd.Intn(height), g.rnd.Intn(width)
	startCell := maze.Grid[startRow][startCol]
	startCell.Type = domain.Passage

	walls := g.getAdjacentWalls(maze, startCell)

	for len(walls) > 0 {
		idx := g.rnd.Intn(len(walls))
		wall := walls[idx]

		if wall.Next.Type == domain.Wall {
			wall.Between.Type = domain.Passage
			wall.Next.Type = domain.Passage

			walls = append(walls, g.getAdjacentWalls(maze, wall.Next)...)
		}

		walls = append(walls[:idx], walls[idx+1:]...)
	}

	g.assignCellTypes(maze, swampFreq, sandFreq, coinFreq)

	return maze, nil
}

func (g *PrimGenerator) initializeMaze(height, width int) *domain.Maze {
	maze := &domain.Maze{
		Height: height,
		Width:  width,
		Grid:   make([][]*domain.Cell, height),
	}
	for i := 0; i < height; i++ {
		maze.Grid[i] = make([]*domain.Cell, width)

		for j := 0; j < width; j++ {
			cellType := domain.Wall
			maze.Grid[i][j] = &domain.Cell{
				Coordinate: domain.Coordinate{Row: i, Col: j},
				Type:       cellType,
			}
		}
	}

	return maze
}

func (g *PrimGenerator) getAdjacentWalls(maze *domain.Maze, cell *domain.Cell) []Wall {
	walls := []Wall{}
	directions := []struct {
		dRow int
		dCol int
	}{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	}

	for _, dir := range directions {
		newRow := cell.Row + dir.dRow*2
		newCol := cell.Col + dir.dCol*2

		if newRow >= 0 && newRow < maze.Height && newCol >= 0 && newCol < maze.Width {
			betweenRow := cell.Row + dir.dRow
			betweenCol := cell.Col + dir.dCol

			betweenCell := maze.Grid[betweenRow][betweenCol]
			nextCell := maze.Grid[newRow][newCol]

			if nextCell.Type == domain.Wall {
				walls = append(walls, Wall{
					Between: betweenCell,
					Next:    nextCell,
				})
			}
		}
	}

	return walls
}

func (g *PrimGenerator) assignCellTypes(maze *domain.Maze, swampFreq, sandFreq, coinFreq float64) {
	for i := 0; i < maze.Height; i++ {
		for j := 0; j < maze.Width; j++ {
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
				default:
					cell.Type = domain.Passage
				}
			}
		}
	}
}
