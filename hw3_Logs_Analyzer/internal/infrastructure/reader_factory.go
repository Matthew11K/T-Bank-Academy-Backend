package infrastructure

import (
	"github.com/central-university-dev/backend_academy_2024_project_3-go-Matthew11K/internal/domain"
)

type UnifiedReaderFactory struct {
	fileFactory domain.ReaderFactory
	urlFactory  domain.ReaderFactory
}

func NewUnifiedReaderFactory(fileFactory, urlFactory domain.ReaderFactory) *UnifiedReaderFactory {
	return &UnifiedReaderFactory{
		fileFactory: fileFactory,
		urlFactory:  urlFactory,
	}
}

func (u *UnifiedReaderFactory) GetFactory(readerType domain.ReaderType) domain.ReaderFactory {
	switch readerType {
	case domain.FileReaderType:
		return u.fileFactory
	case domain.URLReaderType:
		return u.urlFactory
	default:
		return nil
	}
}
