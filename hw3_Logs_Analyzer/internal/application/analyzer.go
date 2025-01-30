package application

import (
	"fmt"

	"golang.org/x/exp/slog"

	"github.com/central-university-dev/backend_academy_2024_project_3-go-Matthew11K/internal/domain"
	"github.com/central-university-dev/backend_academy_2024_project_3-go-Matthew11K/internal/utils"
)

type IOAdapter interface {
	Output(content string) error
}

type LogParser interface {
	Parse(line string) (*domain.LogRecord, error)
}

type FileResolver interface {
	ResolveFiles(pattern string) ([]string, error)
}

type ReaderFactoryProvider interface {
	GetFactory(readerType domain.ReaderType) domain.ReaderFactory
}

type Config struct {
	StartTime   *domain.TimeWrapper
	EndTime     *domain.TimeWrapper
	FilterField domain.FilterField
	FilterValue string
}

type Analyzer struct {
	logger        *slog.Logger
	parser        LogParser
	fileResolver  FileResolver
	readerFactory ReaderFactoryProvider
	report        *domain.LogReport
	config        Config
	io            IOAdapter
}

func NewAnalyzer(
	logger *slog.Logger,
	parser LogParser,
	fileResolver FileResolver,
	readerFactory ReaderFactoryProvider,
	config Config,
	io IOAdapter,
) *Analyzer {
	return &Analyzer{
		logger:        logger,
		parser:        parser,
		fileResolver:  fileResolver,
		readerFactory: readerFactory,
		report:        domain.NewLogReport(),
		config:        config,
		io:            io,
	}
}

func (a *Analyzer) Analyze(paths []string, format string) error {
	if len(paths) == 0 {
		return &domain.ErrNoPathsProvided{}
	}

	readers, err := a.getReaders(paths)
	if err != nil {
		return err
	}

	for _, reader := range readers {
		a.logger.Info("Чтение логов", slog.String("source", reader.Source()))

		err := a.processLogs(reader)
		if err != nil {
			return fmt.Errorf("ошибка при обработке логов: %w", err)
		}
	}

	output, err := a.report.Format(domain.ReportFormatType(format))
	if err != nil {
		return fmt.Errorf("ошибка при форматировании отчета: %w", err)
	}

	if err := a.io.Output(output); err != nil {
		return fmt.Errorf("ошибка при выводе результата: %w", err)
	}

	return nil
}

func (a *Analyzer) processLogs(reader domain.LogReader) error {
	defer func() {
		if err := reader.Close(); err != nil {
			a.logger.Error("Ошибка при закрытии reader", slog.Any("error", err))
		}
	}()

	scanner, err := reader.Scanner()
	if err != nil {
		return fmt.Errorf("ошибка при создании сканера: %w", err)
	}

	for scanner.Scan() {
		line := scanner.Text()

		record, err := a.parser.Parse(line)
		if err != nil {
			a.logger.Warn("Ошибка парсинга строки", slog.String("line", line), slog.Any("error", err))
			continue
		}

		if !record.Validate(a.config.StartTime, a.config.EndTime, a.config.FilterField, a.config.FilterValue) {
			continue
		}

		a.report.AddRecord(record)
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("ошибка чтения: %w", err)
	}

	return nil
}

func (a *Analyzer) getReaders(paths []string) ([]domain.LogReader, error) {
	var readers []domain.LogReader

	for _, path := range paths {
		var factory domain.ReaderFactory
		if utils.IsURL(path) {
			factory = a.readerFactory.GetFactory(domain.URLReaderType)

			reader, err := factory.NewReader(path)
			if err != nil {
				return nil, fmt.Errorf("ошибка при создании URLReader: %w", err)
			}

			readers = append(readers, reader)
		} else {
			files, err := a.fileResolver.ResolveFiles(path)
			if err != nil {
				return nil, fmt.Errorf("ошибка при обработке пути %s: %w", path, err)
			}

			for _, file := range files {
				factory := a.readerFactory.GetFactory(domain.FileReaderType)

				reader, err := factory.NewReader(file)
				if err != nil {
					return nil, fmt.Errorf("ошибка при создании FileReader: %w", err)
				}

				readers = append(readers, reader)
			}
		}
	}

	return readers, nil
}
