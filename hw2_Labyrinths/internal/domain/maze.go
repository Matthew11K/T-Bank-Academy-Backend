package domain

import (
	"strings"

	"github.com/central-university-dev/backend_academy_2024_project_2-go-Matthew11K/internal/config"
)

type Maze struct {
	Height int
	Width  int
	Grid   [][]*Cell
}

func (m *Maze) SetupStartEnd(cfg *config.Config) (start, end Coordinate, err error) {
	userProvidedStart := cfg.StartRow >= 0 && cfg.StartCol >= 0
	userProvidedEnd := cfg.EndRow >= 0 && cfg.EndCol >= 0

	if userProvidedStart != userProvidedEnd {
		return start, end, &InvalidCoordinateError{}
	}

	if userProvidedStart && userProvidedEnd {
		start = Coordinate{Row: cfg.StartRow, Col: cfg.StartCol}
		end = Coordinate{Row: cfg.EndRow, Col: cfg.EndCol}

		if !start.IsValid(m.Height, m.Width) || !end.IsValid(m.Height, m.Width) {
			err = &InvalidCoordinateError{}
			return start, end, err
		}

		m.SetCellType(start, Start)
		m.SetCellType(end, End)
	} else {
		start, end = m.GetDefaultStartEnd(cfg.GeneratorType)
		if !start.IsValid(m.Height, m.Width) || !end.IsValid(m.Height, m.Width) {
			err = &InvalidCoordinateError{}
			return start, end, err
		}

		m.SetCellType(start, Start)
		m.SetCellType(end, End)
	}

	return start, end, nil
}

func (m *Maze) GetDefaultStartEnd(generatorType string) (start, end Coordinate) {
	switch generatorType {
	case "kruskal", "wilson":
		start = Coordinate{Row: 1, Col: 1}
		end = Coordinate{Row: m.Height - 2, Col: m.Width - 2}
	default:
		start = Coordinate{Row: 0, Col: 0}
		end = Coordinate{Row: m.Height - 1, Col: m.Width - 1}
	}

	return
}

func (m *Maze) SetCellType(coord Coordinate, cellType CellType) {
	cell := m.Grid[coord.Row][coord.Col]
	if cell.Type == Wall {
		cell.Type = Passage
	}

	cell.Type = cellType
}

func (m *Maze) String() (string, error) {
	var sb strings.Builder

	var encounteredError error

	for i := 0; i < m.Height; i++ {
		for j := 0; j < m.Width; j++ {
			symbol, err := m.Grid[i][j].Symbol()

			if err != nil {
				symbol = "?"

				if encounteredError == nil {
					encounteredError = err
				}
			}

			sb.WriteString(symbol)
		}

		sb.WriteString("\n")
	}

	return sb.String(), encounteredError
}

func (m *Maze) StringWithPath(path []Coordinate) (string, error) {
	gridCopy := make([][]*Cell, m.Height)

	for i := 0; i < m.Height; i++ {
		gridCopy[i] = make([]*Cell, m.Width)

		for j := 0; j < m.Width; j++ {
			originalCell := m.Grid[i][j]
			gridCopy[i][j] = &Cell{
				Coordinate: Coordinate{Row: originalCell.Row, Col: originalCell.Col},
				Type:       originalCell.Type,
			}
		}
	}

	for _, coord := range path {
		if gridCopy[coord.Row][coord.Col].Type == Passage ||
			gridCopy[coord.Row][coord.Col].Type == Swamp ||
			gridCopy[coord.Row][coord.Col].Type == Sand ||
			gridCopy[coord.Row][coord.Col].Type == Coin {
			gridCopy[coord.Row][coord.Col].Type = Path
		}
	}

	var sb strings.Builder

	var encounteredError error

	for i := 0; i < m.Height; i++ {
		for j := 0; j < m.Width; j++ {
			symbol, err := gridCopy[i][j].Symbol()

			if err != nil {
				symbol = "?"

				if encounteredError == nil {
					encounteredError = err
				}
			}

			sb.WriteString(symbol)
		}

		sb.WriteString("\n")
	}

	return sb.String(), encounteredError
}

func (m *Maze) ReconstructPath(cameFrom map[Coordinate]*Coordinate, current Coordinate) []Coordinate {
	path := []Coordinate{}
	for at := &current; at != nil; at = cameFrom[*at] {
		path = append([]Coordinate{*at}, path...)
	}

	return path
}
