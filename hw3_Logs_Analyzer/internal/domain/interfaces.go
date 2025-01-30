package domain

type LogReader interface {
	Scanner() (Scanner, error)
	Close() error
	Source() string
}

type ReaderFactory interface {
	NewReader(source string) (LogReader, error)
}

type Scanner interface {
	Scan() bool
	Text() string
	Err() error
}
