package config

type Config struct {
	Width          int
	Height         int
	GeneratorType  string
	SolverType     string
	SwampFrequency float64
	SandFrequency  float64
	CoinFrequency  float64
	StartRow       int
	StartCol       int
	EndRow         int
	EndCol         int
}
