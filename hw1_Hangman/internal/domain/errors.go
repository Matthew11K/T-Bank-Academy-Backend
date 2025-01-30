package domain

type ErrInvalidInput struct {
	Input string
}

func (e *ErrInvalidInput) Error() string {
	return "некорректный ввод: " + e.Input
}

type ErrGameOver struct{}

func (e *ErrGameOver) Error() string {
	return "игра окончена"
}

type ErrWordNotFound struct {
	Category   string
	Difficulty string
}

func (e *ErrWordNotFound) Error() string {
	return "слово не найдено для категории '" + e.Category + "' и сложности '" + e.Difficulty + "'"
}

type ErrLetterAlreadyGuessed struct {
	Letter string
}

func (e *ErrLetterAlreadyGuessed) Error() string {
	return "эта буква уже была угадана: " + e.Letter
}
