package main

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/central-university-dev/backend_academy_2024_project_4-go-Matthew11K/internal/application"
	"github.com/central-university-dev/backend_academy_2024_project_4-go-Matthew11K/internal/infrastructure"
	"github.com/central-university-dev/backend_academy_2024_project_4-go-Matthew11K/pkg/logger"
)

var (
	width      int
	height     int
	iterations int
	threads    int
	output     string
	gamma      float64
	symmetry   int
	configFile string
	points     int
	rootCmd    = &cobra.Command{
		Use:   "run",
		Short: "Запускает генерацию изображения",
		Run:   run,
	}
)

func initFlags() {
	rootCmd.Flags().IntVarP(&width, "width", "w", 800, "Ширина изображения")
	rootCmd.Flags().IntVarP(&height, "height", "H", 600, "Высота изображения")
	rootCmd.Flags().IntVarP(&iterations, "iterations", "i", 1000000, "Количество итераций")
	rootCmd.Flags().IntVarP(&threads, "threads", "t", 1, "Количество потоков")
	rootCmd.Flags().StringVarP(&output, "output", "o", "output.png", "Имя выходного файла")
	rootCmd.Flags().Float64VarP(&gamma, "gamma", "g", 2.2, "Гамма-коррекция")
	rootCmd.Flags().IntVarP(&symmetry, "symmetry", "s", 1, "Параметр симметрии")
	rootCmd.Flags().StringVarP(&configFile, "config", "c", "config.json", "Файл конфигурации трансформаций")
	rootCmd.Flags().IntVarP(&points, "points", "p", 1000, "Количество точек для генерации фрактала")
}

func main() {
	initFlags()

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func run(_ *cobra.Command, _ []string) {
	loggerInstance := logger.NewLogger(os.Stdout)

	transformationConfigs, err := application.LoadTransformationConfigs(configFile)
	if err != nil {
		loggerInstance.Error("Ошибка при загрузке конфигурации", "error", err)
		os.Exit(1)
	}

	config := application.Config{
		Width:           width,
		Height:          height,
		Iterations:      iterations,
		Points:          points,
		Threads:         threads,
		Output:          output,
		Gamma:           gamma,
		Symmetry:        symmetry,
		Transformations: transformationConfigs,
	}

	imageSaver := &infrastructure.ImageSaverImpl{}

	app := application.NewApplication(&config, imageSaver, loggerInstance)
	if err := app.Run(); err != nil {
		loggerInstance.Error("Ошибка при выполнении приложения", "error", err)
		os.Exit(1)
	}

	loggerInstance.Info("Изображение успешно сгенерировано", "output", output)
}
