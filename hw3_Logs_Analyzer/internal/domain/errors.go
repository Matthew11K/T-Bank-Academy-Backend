package domain

type ErrInvalidFormat struct{}

func (e *ErrInvalidFormat) Error() string {
	return "неверный формат вывода"
}

type ErrNoPathsProvided struct{}

func (e *ErrNoPathsProvided) Error() string {
	return "не указаны пути к лог-файлам"
}

type ErrInvalidLogFormat struct{}

func (e *ErrInvalidLogFormat) Error() string {
	return "строка не соответствует формату"
}

type ErrInvalidTimeFormat struct {
	Err error
}

func (e *ErrInvalidTimeFormat) Error() string {
	return "неверный формат времени: " + e.Err.Error()
}

func (e *ErrInvalidTimeFormat) Unwrap() error {
	return e.Err
}

type ErrFileOpen struct {
	Err error
}

func (e *ErrFileOpen) Error() string {
	return "ошибка при открытии файла: " + e.Err.Error()
}

func (e *ErrFileOpen) Unwrap() error {
	return e.Err
}

type ErrURLFetch struct {
	Err error
}

func (e *ErrURLFetch) Error() string {
	return "ошибка при получении URL: " + e.Err.Error()
}

func (e *ErrURLFetch) Unwrap() error {
	return e.Err
}

type ErrParseInt struct {
	Field string
	Err   error
}

func (e *ErrParseInt) Error() string {
	return "ошибка при парсинге " + e.Field + ": " + e.Err.Error()
}

func (e *ErrParseInt) Unwrap() error {
	return e.Err
}

type ErrFilePattern struct {
	Err error
}

func (e *ErrFilePattern) Error() string {
	return "ошибка при обработке паттерна файлов: " + e.Err.Error()
}

func (e *ErrFilePattern) Unwrap() error {
	return e.Err
}

type ErrFileStat struct {
	Path string
	Err  error
}

func (e *ErrFileStat) Error() string {
	return "ошибка при получении статуса файла " + e.Path + ": " + e.Err.Error()
}

func (e *ErrFileStat) Unwrap() error {
	return e.Err
}
