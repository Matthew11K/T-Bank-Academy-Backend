package infrastructure

import (
	"bufio"
	"os"

	"github.com/central-university-dev/backend_academy_2024_project_3-go-Matthew11K/internal/domain"
)

type FileOpener interface {
	Open(name string) (*os.File, error)
}

type DefaultFileOpener struct{}

func (d DefaultFileOpener) Open(name string) (*os.File, error) {
	return os.Open(name)
}

type FileReader struct {
	fileOpener FileOpener
	file       *os.File
	source     string
}

type FileReaderFactory struct {
	fileOpener FileOpener
}

func NewFileReaderFactory() domain.ReaderFactory {
	return &FileReaderFactory{
		fileOpener: DefaultFileOpener{},
	}
}

func (f *FileReaderFactory) NewReader(path string) (domain.LogReader, error) {
	return &FileReader{
		source:     path,
		fileOpener: f.fileOpener,
	}, nil
}

func (fr *FileReader) Scanner() (domain.Scanner, error) {
	file, err := fr.fileOpener.Open(fr.source)
	if err != nil {
		return nil, &domain.ErrFileOpen{Err: err}
	}

	fr.file = file

	return &BufioScanner{Scanner: bufio.NewScanner(file)}, nil
}

func (fr *FileReader) Close() error {
	if fr.file != nil {
		return fr.file.Close()
	}

	return nil
}

func (fr *FileReader) Source() string {
	return fr.source
}

type BufioScanner struct {
	*bufio.Scanner
}

func (bs *BufioScanner) Scan() bool {
	return bs.Scanner.Scan()
}

func (bs *BufioScanner) Text() string {
	return bs.Scanner.Text()
}

func (bs *BufioScanner) Err() error {
	return bs.Scanner.Err()
}
