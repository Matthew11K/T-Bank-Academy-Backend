package infrastructure

import (
	"github.com/central-university-dev/backend_academy_2024_project_3-go-Matthew11K/internal/application"
	"github.com/central-university-dev/backend_academy_2024_project_3-go-Matthew11K/internal/utils"
)

type FileResolver struct{}

func NewFileResolver() application.FileResolver {
	return &FileResolver{}
}

func (fr *FileResolver) ResolveFiles(pattern string) ([]string, error) {
	matches, err := utils.ResolveFiles(pattern)
	if err != nil {
		return nil, err
	}

	return matches, nil
}
