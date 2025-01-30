package application

import (
	"encoding/json"
	"os"

	"github.com/central-university-dev/backend_academy_2024_project_4-go-Matthew11K/internal/domain"
)

func LoadTransformationConfigs(filename string) (domain.Transformations, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, WrapError(err, "ошибка открытия файла конфигурации")
	}
	defer file.Close()

	var configs []TransformationConfig

	decoder := json.NewDecoder(file)

	if err := decoder.Decode(&configs); err != nil {
		return nil, WrapError(err, "ошибка декодирования файла конфигурации")
	}

	transformations := make([]domain.Transformation, 0, len(configs))

	for _, cfg := range configs {
		variation, err := domain.NewVariation(cfg.Type)
		if err != nil {
			return nil, WrapError(err, "ошибка создания вариации")
		}

		transformation := domain.Transformation{
			Affine:    cfg.Affine,
			Variation: variation,
			Color:     cfg.Color,
			Weight:    cfg.Weight,
		}
		transformations = append(transformations, transformation)
	}

	return transformations, nil
}
