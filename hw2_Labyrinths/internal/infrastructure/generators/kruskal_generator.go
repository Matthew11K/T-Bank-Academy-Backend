package generators

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"math"

	"github.com/central-university-dev/backend_academy_2024_project_2-go-Matthew11K/internal/domain"
)

type KruskalGenerator struct{}

func NewKruskalGenerator() *KruskalGenerator {
	return &KruskalGenerator{}
}

type Edge struct {
	Row int
	Col int
}

func (g *KruskalGenerator) Generate(height, width int, swampFreq, sandFreq, coinFreq float64) (*domain.Maze, error) {
	mazeHeight, mazeWidth := height*2+1, width*2+1
	maze := g.initializeMaze(mazeHeight, mazeWidth)

	sets := g.initializeSets(height, width)
	edges := g.generateEdges(height, width)

	if err := g.shuffleEdges(edges); err != nil {
		return nil, fmt.Errorf("Generate: %w", err)
	}

	maze = g.carvePassages(maze, edges, sets, height, width)

	if err := g.assignSpecialCells(maze, swampFreq, sandFreq, coinFreq); err != nil {
		return nil, fmt.Errorf("Generate: %w", err)
	}

	return maze, nil
}

func (g *KruskalGenerator) initializeMaze(height, width int) *domain.Maze {
	maze := &domain.Maze{
		Height: height,
		Width:  width,
		Grid:   make([][]*domain.Cell, height),
	}
	for i := 0; i < height; i++ {
		maze.Grid[i] = make([]*domain.Cell, width)

		for j := 0; j < width; j++ {
			cellType := domain.Wall
			if i%2 == 1 && j%2 == 1 {
				cellType = domain.Passage
			}

			maze.Grid[i][j] = &domain.Cell{
				Coordinate: domain.Coordinate{Row: i, Col: j},
				Type:       cellType,
			}
		}
	}

	return maze
}

func (g *KruskalGenerator) initializeSets(height, width int) [][]int {
	sets := make([][]int, height)
	for i := range sets {
		sets[i] = make([]int, width)
		for j := range sets[i] {
			sets[i][j] = i*width + j
		}
	}

	return sets
}

func (g *KruskalGenerator) generateEdges(height, width int) []Edge {
	edges := []Edge{}

	for i := 0; i < height; i++ {
		for j := 0; j < width-1; j++ {
			edges = append(edges, Edge{Row: i*2 + 1, Col: j*2 + 2})
		}
	}

	for i := 0; i < height-1; i++ {
		for j := 0; j < width; j++ {
			edges = append(edges, Edge{Row: i*2 + 2, Col: j*2 + 1})
		}
	}

	return edges
}

func (g *KruskalGenerator) shuffleEdges(edges []Edge) error {
	n := len(edges)

	for i := n - 1; i > 0; i-- {
		j, err := g.randInt(i + 1)
		if err != nil {
			return fmt.Errorf("shuffleEdges: %w", err)
		}

		edges[i], edges[j] = edges[j], edges[i]
	}

	return nil
}

func (g *KruskalGenerator) carvePassages(maze *domain.Maze, edges []Edge, sets [][]int, height, width int) *domain.Maze {
	for _, edge := range edges {
		row, col := edge.Row, edge.Col

		var cell1, cell2 domain.Coordinate
		if row%2 == 1 {
			cell1 = domain.Coordinate{Row: row, Col: col - 1}
			cell2 = domain.Coordinate{Row: row, Col: col + 1}
		} else {
			cell1 = domain.Coordinate{Row: row - 1, Col: col}
			cell2 = domain.Coordinate{Row: row + 1, Col: col}
		}

		set1 := sets[cell1.Row/2][cell1.Col/2]
		set2 := sets[cell2.Row/2][cell2.Col/2]

		if set1 != set2 {
			maze.Grid[row][col].Type = domain.Passage

			g.mergeSets(sets, set1, set2, height, width)
		}
	}

	return maze
}

func (g *KruskalGenerator) mergeSets(sets [][]int, oldSet, newSet, height, width int) {
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if sets[i][j] == oldSet {
				sets[i][j] = newSet
			}
		}
	}
}

func (g *KruskalGenerator) assignSpecialCells(maze *domain.Maze, swampFreq, sandFreq, coinFreq float64) error {
	for i := 1; i < maze.Height; i += 2 {
		for j := 1; j < maze.Width; j += 2 {
			cell := maze.Grid[i][j]
			randValue, err := g.randFloat()

			if err != nil {
				return err
			}

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

	return nil
}

func (g *KruskalGenerator) randInt(maxVal int) (int, error) {
	if maxVal <= 0 {
		return 0, &domain.ErrInvalidMaxValue{}
	}

	var b [8]byte
	_, err := rand.Read(b[:])

	if err != nil {
		return 0, fmt.Errorf("randInt: %w", err)
	}

	r := binary.BigEndian.Uint64(b[:])
	randVal := r % uint64(maxVal)

	result, err := g.safeUint64ToInt(randVal)
	if err != nil {
		return 0, fmt.Errorf("randInt: %w", err)
	}

	return result, nil
}

func (g *KruskalGenerator) randFloat() (float64, error) {
	var b [8]byte
	_, err := rand.Read(b[:])

	if err != nil {
		return 0, fmt.Errorf("randFloat: %w", err)
	}

	u := binary.BigEndian.Uint64(b[:])
	f := math.Float64frombits(u)

	if f < 0.0 || f >= 1.0 {
		f = math.Mod(f, 1.0)
	}

	return f, nil
}

func (g *KruskalGenerator) safeUint64ToInt(u uint64) (int, error) {
	if u > uint64(math.MaxInt) {
		return 0, fmt.Errorf("значение %d превышает максимально допустимое для int", u)
	}

	return int(u), nil
}
