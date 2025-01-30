package utils

import (
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/central-university-dev/backend_academy_2024_project_3-go-Matthew11K/internal/domain"
)

func ResolveFiles(pattern string) ([]string, error) {
	matches, err := filepath.Glob(pattern)
	if err != nil {
		return nil, &domain.ErrFilePattern{Err: err}
	}

	var files []string

	for _, match := range matches {
		fi, err := os.Stat(match)
		if err != nil {
			return nil, &domain.ErrFileStat{Path: match, Err: err}
		}

		if fi.Mode().IsRegular() {
			files = append(files, match)
		}
	}

	return files, nil
}

func IsURL(str string) bool {
	if strings.HasPrefix(str, "http://") || strings.HasPrefix(str, "https://") {
		_, err := url.ParseRequestURI(str)
		return err == nil
	}

	return false
}
