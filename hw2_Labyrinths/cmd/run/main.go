package main

import (
	"flag"
	"os"

	"github.com/central-university-dev/backend_academy_2024_project_2-go-Matthew11K/internal/application"
	"github.com/central-university-dev/backend_academy_2024_project_2-go-Matthew11K/internal/config"
	"github.com/central-university-dev/backend_academy_2024_project_2-go-Matthew11K/internal/infrastructure/generators"
	"github.com/central-university-dev/backend_academy_2024_project_2-go-Matthew11K/internal/infrastructure/solvers"

	"golang.org/x/exp/slog"
)

func parseFlags() config.Config {
	width := flag.Int("width", 20, "Ширина лабиринта")
	height := flag.Int("height", 20, "Высота лабиринта")
	generatorType := flag.String("generator", "prim", "Алгоритм генерации лабиринта (prim|kruskal|wilson)")
	solverType := flag.String("solver", "bfs", "Алгоритм решения лабиринта (bfs|astar|dijkstra)")
	swampFrequency := flag.Float64("swampFreq", 0.05, "Частота болот (0.0 - 1.0)")
	sandFrequency := flag.Float64("sandFreq", 0.05, "Частота песка (0.0 - 1.0)")
	coinFrequency := flag.Float64("coinFreq", 0.05, "Частота монет (0.0 - 1.0)")
	startRow := flag.Int("start-row", -1, "Начальная строка лабиринта (по умолчанию зависит от генератора)")
	startCol := flag.Int("start-col", -1, "Начальный столбец лабиринта (по умолчанию зависит от генератора)")
	endRow := flag.Int("end-row", -1, "Конечная строка лабиринта (по умолчанию зависит от генератора)")
	endCol := flag.Int("end-col", -1, "Конечный столбец лабиринта (по умолчанию зависит от генератора)")
	flag.Parse()

	return config.Config{
		Width:          *width,
		Height:         *height,
		GeneratorType:  *generatorType,
		SolverType:     *solverType,
		SwampFrequency: *swampFrequency,
		SandFrequency:  *sandFrequency,
		CoinFrequency:  *coinFrequency,
		StartRow:       *startRow,
		StartCol:       *startCol,
		EndRow:         *endRow,
		EndCol:         *endCol,
	}
}

func main() {
	cfg := parseFlags()
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{}))

	generator, err := generators.InitGenerator(cfg.GeneratorType)
	if err != nil {
		logger.Error("Ошибка инициализации генератора", slog.String("error", err.Error()))
		os.Exit(1)
	}

	solver, err := solvers.InitSolver(cfg.SolverType)
	if err != nil {
		logger.Error("Ошибка инициализации решателя", slog.String("error", err.Error()))
		os.Exit(1)
	}

	mazeService := application.NewMazeService(generator, solver, logger)

	mazeService.Start(&cfg)
}
