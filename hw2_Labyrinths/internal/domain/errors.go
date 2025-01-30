package domain

type InvalidMazeSizeError struct{}

func (e *InvalidMazeSizeError) Error() string {
	return "недопустимый размер лабиринта"
}

type InvalidCoordinateError struct{}

func (e *InvalidCoordinateError) Error() string {
	return "недопустимая координата"
}

type NoPathFoundError struct{}

func (e *NoPathFoundError) Error() string {
	return "путь не найден"
}

type ErrInvalidMaxValue struct{}

func (e *ErrInvalidMaxValue) Error() string {
	return "максимальное значение должно быть положительным"
}
