package config

type Config struct {
	MaxAttemptsEasy   int
	MaxAttemptsMedium int
	MaxAttemptsHard   int
}

func LoadConfig() *Config {
	return &Config{
		MaxAttemptsEasy:   10,
		MaxAttemptsMedium: 7,
		MaxAttemptsHard:   5,
	}
}
