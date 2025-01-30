package infrastructure

import (
	"bufio"
	"fmt"
	"net/http"

	"github.com/central-university-dev/backend_academy_2024_project_3-go-Matthew11K/internal/domain"
)

type HTTPClient interface {
	Get(url string) (*http.Response, error)
}

type URLReader struct {
	httpClient HTTPClient
	resp       *http.Response
	source     string
}

type URLReaderFactory struct {
	httpClient HTTPClient
}

func NewURLReaderFactory(client HTTPClient) domain.ReaderFactory {
	if client == nil {
		client = &http.Client{}
	}

	return &URLReaderFactory{
		httpClient: client,
	}
}

func (u *URLReaderFactory) NewReader(url string) (domain.LogReader, error) {
	return &URLReader{
		source:     url,
		httpClient: u.httpClient,
	}, nil
}

func (ur *URLReader) Scanner() (domain.Scanner, error) {
	resp, err := ur.httpClient.Get(ur.source)
	if err != nil {
		return nil, &domain.ErrURLFetch{Err: err}
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("не удалось получить данные: статус %d", resp.StatusCode)
	}

	ur.resp = resp

	return &BufioScanner{Scanner: bufio.NewScanner(resp.Body)}, nil
}

func (ur *URLReader) Close() error {
	if ur.resp != nil {
		return ur.resp.Body.Close()
	}

	return nil
}

func (ur *URLReader) Source() string {
	return ur.source
}
