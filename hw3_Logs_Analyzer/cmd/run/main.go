package main

import (
	"os"

	"github.com/spf13/cobra"
	"golang.org/x/exp/slog"

	"github.com/central-university-dev/backend_academy_2024_project_3-go-Matthew11K/internal/application"
	"github.com/central-university-dev/backend_academy_2024_project_3-go-Matthew11K/internal/domain"
	"github.com/central-university-dev/backend_academy_2024_project_3-go-Matthew11K/internal/infrastructure"
	"github.com/central-university-dev/backend_academy_2024_project_3-go-Matthew11K/pkg/logger"
)

func main() {
	var paths []string

	var from, to, format string

	var filterField, filterValue string

	var rootCmd = &cobra.Command{
		Use:   "analyzer",
		Short: "NGINX лог анализатор",
		Run: func(_ *cobra.Command, _ []string) {
			logger := logger.NewLogger()
			parser := infrastructure.NewLogParser()
			fileResolver := infrastructure.NewFileResolver()
			unifiedFactory := infrastructure.NewUnifiedReaderFactory(
				infrastructure.NewFileReaderFactory(),
				infrastructure.NewURLReaderFactory(nil),
			)
			io := infrastructure.NewIOAdapter(os.Stdout, logger)

			var config application.Config
			var err error

			if from != "" {
				config.StartTime, err = domain.ParseTime(from)
				if err != nil {
					logger.Error("Неверный формат времени 'from'", slog.Any("error", err))
					os.Exit(1)
				}
			}

			if to != "" {
				config.EndTime, err = domain.ParseTime(to)
				if err != nil {
					logger.Error("Неверный формат времени 'to'", slog.Any("error", err))
					os.Exit(1)
				}
			}

			config.FilterField = domain.FilterField(filterField)
			config.FilterValue = filterValue

			analyzer := application.NewAnalyzer(
				logger,
				parser,
				fileResolver,
				unifiedFactory,
				config,
				io,
			)

			err = analyzer.Analyze(paths, format)
			if err != nil {
				logger.Error("Ошибка при анализе логов", slog.Any("error", err))
				os.Exit(1)
			}
		},
	}

	rootCmd.Flags().StringSliceVar(&paths, "path", nil, "Путь к лог-файлам")
	rootCmd.Flags().StringVar(&from, "from", "", "Начальная дата в формате ISO8601")
	rootCmd.Flags().StringVar(&to, "to", "", "Конечная дата в формате ISO8601")
	rootCmd.Flags().StringVar(&format, "format", "markdown", "Формат вывода (markdown или adoc)")
	rootCmd.Flags().StringVar(&filterField, "filter-field", "", "Поле для фильтрации (agent или method)")
	rootCmd.Flags().StringVar(&filterValue, "filter-value", "", "Значение для фильтрации")

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
