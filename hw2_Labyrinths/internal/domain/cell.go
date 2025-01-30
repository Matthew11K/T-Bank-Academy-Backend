package domain

import (
	"fmt"
	"math"
)

type CellType int

const (
	Wall CellType = iota
	Passage
	Start
	End
	Path
	Swamp
	Sand
	Coin
)

type Cell struct {
	Coordinate
	Type CellType
}

func (c *Cell) Symbol() (string, error) {
	switch c.Type {
	case Wall:
		return "#", nil
	case Passage:
		return " ", nil
	case Swamp:
		return "~", nil
	case Sand:
		return ":", nil
	case Coin:
		return "$", nil
	case Path:
		return ".", nil
	case Start:
		return "S", nil
	case End:
		return "E", nil
	default:
		return "", fmt.Errorf("неизвестный тип клеток в (%d, %d)", c.Row, c.Col)
	}
}

func (c Cell) GetMovementCostAStar() float64 {
	switch c.Type {
	case Wall:
		return math.Inf(1)
	case Swamp:
		return 2.0
	case Sand:
		return 1.5
	case Coin:
		return 0.5
	case Passage, Start, End, Path:
		return 1.0
	default:
		return 1.0
	}
}

func (c Cell) GetMovementCostDijkstra() float64 {
	return c.GetMovementCostAStar()
}
