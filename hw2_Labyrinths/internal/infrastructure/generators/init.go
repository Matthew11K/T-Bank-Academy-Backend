package generators

import (
	"fmt"

	"github.com/central-university-dev/backend_academy_2024_project_2-go-Matthew11K/internal/application"
)

func InitGenerator(genType string) (application.MazeGenerator, error) {
	switch genType {
	case "prim":
		return NewPrimGenerator(), nil
	case "kruskal":
		return NewKruskalGenerator(), nil
	case "wilson":
		return NewWilsonGenerator(), nil
	default:
		return nil, fmt.Errorf("неизвестный тип генератора: %s", genType)
	}
}
