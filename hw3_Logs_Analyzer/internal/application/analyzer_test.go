package application_test

import (
	"errors"
	"io"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"golang.org/x/exp/slog"

	"github.com/central-university-dev/backend_academy_2024_project_3-go-Matthew11K/internal/application"
	"github.com/central-university-dev/backend_academy_2024_project_3-go-Matthew11K/internal/domain"
	"github.com/central-university-dev/backend_academy_2024_project_3-go-Matthew11K/internal/domain/mocks"
	"github.com/central-university-dev/backend_academy_2024_project_3-go-Matthew11K/internal/infrastructure"
)

func TestAnalyzer_Analyze(t *testing.T) {
	testCases := []struct {
		name        string
		paths       []string
		from        string
		to          string
		format      string
		filterField domain.FilterField
		filterValue string
		setupMocks  func(
			_ *mocks.LogParser,
			_ *mocks.FileResolver,
			_ *mocks.ReaderFactory,
			_ *mocks.ReaderFactory,
			_ *mocks.IOAdapter,
		)
		expectedError error
	}{
		{
			name:        "No Paths Provided",
			paths:       []string{},
			from:        "",
			to:          "",
			format:      "markdown",
			filterField: domain.FieldMethod,
			filterValue: "",
			setupMocks: func(
				_ *mocks.LogParser,
				_ *mocks.FileResolver,
				_ *mocks.ReaderFactory,
				_ *mocks.ReaderFactory,
				_ *mocks.IOAdapter,
			) {
			},
			expectedError: &domain.ErrNoPathsProvided{},
		},
		{
			name:        "Error Creating URLReader",
			paths:       []string{"http://example.com/logs"},
			from:        "",
			to:          "",
			format:      "markdown",
			filterField: domain.FieldMethod,
			filterValue: "",
			setupMocks: func(
				_ *mocks.LogParser,
				_ *mocks.FileResolver,
				urlReaderFactoryMock *mocks.ReaderFactory,
				_ *mocks.ReaderFactory,
				_ *mocks.IOAdapter,
			) {
				urlReaderFactoryMock.
					On("NewReader", "http://example.com/logs").
					Return(nil, &domain.ErrURLFetch{Err: errors.New("failed to fetch URL")}).
					Once()
			},
			expectedError: &domain.ErrURLFetch{},
		},
		{
			name:        "Error Creating FileReader",
			paths:       []string{"./logs/2024-08-31.log"},
			from:        "",
			to:          "",
			format:      "markdown",
			filterField: domain.FieldMethod,
			filterValue: "",
			setupMocks: func(
				_ *mocks.LogParser,
				fileResolverMock *mocks.FileResolver,
				_ *mocks.ReaderFactory,
				fileReaderFactoryMock *mocks.ReaderFactory,
				_ *mocks.IOAdapter,
			) {
				fileResolverMock.On("ResolveFiles", "./logs/2024-08-31.log").
					Return([]string{"./logs/2024-08-31.log"}, nil).
					Once()
				fileReaderFactoryMock.On("NewReader", "./logs/2024-08-31.log").
					Return(nil, &domain.ErrFileOpen{Err: errors.New("failed to open file")}).
					Once()
			},
			expectedError: &domain.ErrFileOpen{},
		},
		{
			name:        "Error Formatting Report",
			paths:       []string{"path/to/log"},
			from:        "",
			to:          "",
			format:      "invalid-format",
			filterField: domain.FieldMethod,
			filterValue: "",
			setupMocks: func(
				parserMock *mocks.LogParser,
				fileResolverMock *mocks.FileResolver,
				_ *mocks.ReaderFactory,
				fileReaderFactoryMock *mocks.ReaderFactory,
				_ *mocks.IOAdapter,
			) {
				logReaderMock := new(mocks.LogReader)
				scannerMock := new(mocks.Scanner)

				fileResolverMock.On("ResolveFiles", "path/to/log").Return([]string{"path/to/log"}, nil).Once()
				fileReaderFactoryMock.On("NewReader", "path/to/log").Return(logReaderMock, nil).Once()
				logReaderMock.On("Scanner").Return(scannerMock, nil).Once()
				logReaderMock.On("Close").Return(nil).Once()
				logReaderMock.On("Source").Return("log_source").Once()

				scannerMock.On("Scan").Return(true).Once()
				scannerMock.On("Text").Return(`127.0.0.1 - - [31/Aug/2024:12:00:00 +0000] "GET /index.html HTTP/1.1" 200 1024 "-" "Mozilla/5.0"`).Once()
				scannerMock.On("Scan").Return(false).Once()
				scannerMock.On("Err").Return(nil).Once()

				parsedTime := &domain.TimeWrapper{Time: time.Date(2024, 8, 31, 12, 0, 0, 0, time.UTC)}
				parsedRecord := &domain.LogRecord{
					RemoteAddr:    "127.0.0.1",
					RemoteUser:    "-",
					Time:          parsedTime,
					Request:       "GET /index.html HTTP/1.1",
					Status:        200,
					BodyBytesSent: 1024,
					HTTPReferer:   "-",
					HTTPUserAgent: "Mozilla/5.0",
				}
				parserMock.On(
					"Parse",
					`127.0.0.1 - - [31/Aug/2024:12:00:00 +0000] "GET /index.html HTTP/1.1" 200 1024 "-" "Mozilla/5.0"`,
				).Return(parsedRecord, nil).Once()
			},
			expectedError: &domain.ErrInvalidFormat{},
		},
		{
			name:        "Error Closing Reader",
			paths:       []string{"path/to/log"},
			from:        "",
			to:          "",
			format:      "markdown",
			filterField: domain.FieldMethod,
			filterValue: "",
			setupMocks: func(
				parserMock *mocks.LogParser,
				fileResolverMock *mocks.FileResolver,
				_ *mocks.ReaderFactory,
				fileReaderFactoryMock *mocks.ReaderFactory,
				ioAdapterMock *mocks.IOAdapter,
			) {
				logReaderMock := new(mocks.LogReader)
				scannerMock := new(mocks.Scanner)

				fileResolverMock.On("ResolveFiles", "path/to/log").Return([]string{"path/to/log"}, nil).Once()
				fileReaderFactoryMock.On("NewReader", "path/to/log").Return(logReaderMock, nil).Once()
				logReaderMock.On("Scanner").Return(scannerMock, nil).Once()
				logReaderMock.On("Close").Return(errors.New("close error")).Once()
				logReaderMock.On("Source").Return("log_source").Once()

				scannerMock.On("Scan").Return(true).Once()
				scannerMock.On("Text").Return(`127.0.0.1 - - [31/Aug/2024:12:00:00 +0000] "GET /index.html HTTP/1.1" 200 1024 "-" "Mozilla/5.0"`).Once()
				scannerMock.On("Scan").Return(false).Once()
				scannerMock.On("Err").Return(nil).Once()

				parsedTime := &domain.TimeWrapper{Time: time.Date(2024, 8, 31, 12, 0, 0, 0, time.UTC)}
				parsedRecord := &domain.LogRecord{
					RemoteAddr:    "127.0.0.1",
					RemoteUser:    "-",
					Time:          parsedTime,
					Request:       "GET /index.html HTTP/1.1",
					Status:        200,
					BodyBytesSent: 1024,
					HTTPReferer:   "-",
					HTTPUserAgent: "Mozilla/5.0",
				}
				parserMock.On(
					"Parse",
					`127.0.0.1 - - [31/Aug/2024:12:00:00 +0000] "GET /index.html HTTP/1.1" 200 1024 "-" "Mozilla/5.0"`,
				).Return(parsedRecord, nil).Once()

				ioAdapterMock.
					On("Output", mock.AnythingOfType("string")).
					Return(nil).
					Once()
			},
			expectedError: nil,
		},
		{
			name:        "Valid Paths",
			paths:       []string{"path/to/log"},
			from:        "",
			to:          "",
			format:      "markdown",
			filterField: domain.FieldMethod,
			filterValue: "",
			setupMocks: func(
				parserMock *mocks.LogParser,
				fileResolverMock *mocks.FileResolver,
				_ *mocks.ReaderFactory,
				fileReaderFactoryMock *mocks.ReaderFactory,
				ioAdapterMock *mocks.IOAdapter,
			) {
				logReaderMock := new(mocks.LogReader)
				scannerMock := new(mocks.Scanner)

				fileResolverMock.On("ResolveFiles", "path/to/log").Return([]string{"path/to/log"}, nil).Once()
				fileReaderFactoryMock.On("NewReader", "path/to/log").Return(logReaderMock, nil).Once()
				logReaderMock.On("Scanner").Return(scannerMock, nil).Once()
				logReaderMock.On("Close").Return(nil).Once()
				logReaderMock.On("Source").Return("log_source").Once()

				scannerMock.On("Scan").Return(true).Once()
				scannerMock.On("Text").Return(`127.0.0.1 - - [31/Aug/2024:12:00:00 +0000] "GET /index.html HTTP/1.1" 200 1024 "-" "Mozilla/5.0"`).Once()
				scannerMock.On("Scan").Return(false).Once()
				scannerMock.On("Err").Return(nil).Once()

				parsedTime := &domain.TimeWrapper{Time: time.Date(2024, 8, 31, 12, 0, 0, 0, time.UTC)}
				parsedRecord := &domain.LogRecord{
					RemoteAddr:    "127.0.0.1",
					RemoteUser:    "-",
					Time:          parsedTime,
					Request:       "GET /index.html HTTP/1.1",
					Status:        200,
					BodyBytesSent: 1024,
					HTTPReferer:   "-",
					HTTPUserAgent: "Mozilla/5.0",
				}
				parserMock.On(
					"Parse",
					`127.0.0.1 - - [31/Aug/2024:12:00:00 +0000] "GET /index.html HTTP/1.1" 200 1024 "-" "Mozilla/5.0"`,
				).Return(parsedRecord, nil).Once()

				ioAdapterMock.
					On("Output", mock.AnythingOfType("string")).
					Return(nil).
					Once()
			},
			expectedError: nil,
		},
		{
			name:        "Filtering by Method",
			paths:       []string{"path/to/log"},
			from:        "",
			to:          "",
			format:      "markdown",
			filterField: domain.FieldMethod,
			filterValue: "GET",
			setupMocks: func(
				parserMock *mocks.LogParser,
				fileResolverMock *mocks.FileResolver,
				_ *mocks.ReaderFactory,
				fileReaderFactoryMock *mocks.ReaderFactory,
				ioAdapterMock *mocks.IOAdapter,
			) {
				logReaderMock := new(mocks.LogReader)
				scannerMock := new(mocks.Scanner)

				fileResolverMock.On("ResolveFiles", "path/to/log").Return([]string{"path/to/log"}, nil).Once()
				fileReaderFactoryMock.On("NewReader", "path/to/log").Return(logReaderMock, nil).Once()
				logReaderMock.On("Scanner").Return(scannerMock, nil).Once()
				logReaderMock.On("Close").Return(nil).Once()
				logReaderMock.On("Source").Return("log_source").Once()

				scannerMock.On("Scan").Return(true).Once()
				scannerMock.On("Text").Return(`127.0.0.1 - - [31/Aug/2024:12:00:00 +0000] "GET /index.html HTTP/1.1" 200 1024 "-" "Mozilla/5.0"`).Once()
				scannerMock.On("Scan").Return(true).Once()
				scannerMock.On("Text").Return(`127.0.0.1 - - [31/Aug/2024:12:05:00 +0000] "POST /api/data HTTP/1.1" 201 2048 "-" "Curl/7.64.1"`).Once()
				scannerMock.On("Scan").Return(false).Once()
				scannerMock.On("Err").Return(nil).Once()

				parsedTime1 := &domain.TimeWrapper{Time: time.Date(2024, 8, 31, 12, 0, 0, 0, time.UTC)}
				parsedRecord1 := &domain.LogRecord{
					RemoteAddr:    "127.0.0.1",
					RemoteUser:    "-",
					Time:          parsedTime1,
					Request:       "GET /index.html HTTP/1.1",
					Status:        200,
					BodyBytesSent: 1024,
					HTTPReferer:   "-",
					HTTPUserAgent: "Mozilla/5.0",
				}
				parserMock.On(
					"Parse",
					`127.0.0.1 - - [31/Aug/2024:12:00:00 +0000] "GET /index.html HTTP/1.1" 200 1024 "-" "Mozilla/5.0"`,
				).Return(parsedRecord1, nil).Once()

				parsedTime2 := &domain.TimeWrapper{Time: time.Date(2024, 8, 31, 12, 5, 0, 0, time.UTC)}
				parsedRecord2 := &domain.LogRecord{
					RemoteAddr:    "127.0.0.1",
					RemoteUser:    "-",
					Time:          parsedTime2,
					Request:       "POST /api/data HTTP/1.1",
					Status:        201,
					BodyBytesSent: 2048,
					HTTPReferer:   "-",
					HTTPUserAgent: "Curl/7.64.1",
				}
				parserMock.On(
					"Parse",
					`127.0.0.1 - - [31/Aug/2024:12:05:00 +0000] "POST /api/data HTTP/1.1" 201 2048 "-" "Curl/7.64.1"`,
				).Return(parsedRecord2, nil).Once()

				ioAdapterMock.
					On("Output", mock.AnythingOfType("string")).
					Return(nil).
					Once()
			},
			expectedError: nil,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			logger := slog.New(slog.NewTextHandler(io.Discard, &slog.HandlerOptions{}))

			parserMock := new(mocks.LogParser)
			fileResolverMock := new(mocks.FileResolver)

			urlReaderFactoryMock := new(mocks.ReaderFactory)
			fileReaderFactoryMock := new(mocks.ReaderFactory)
			ioAdapterMock := new(mocks.IOAdapter)

			unifiedFactoryMock := infrastructure.NewUnifiedReaderFactory(
				fileReaderFactoryMock,
				urlReaderFactoryMock,
			)

			tc.setupMocks(
				parserMock,
				fileResolverMock,
				urlReaderFactoryMock,
				fileReaderFactoryMock,
				ioAdapterMock,
			)

			var config application.Config

			var err error

			if tc.from != "" {
				config.StartTime, err = domain.ParseTime(tc.from)
				require.NoError(t, err)
			}

			if tc.to != "" {
				config.EndTime, err = domain.ParseTime(tc.to)
				require.NoError(t, err)
			}

			config.FilterField = tc.filterField
			config.FilterValue = tc.filterValue

			analyzer := application.NewAnalyzer(
				logger,
				parserMock,
				fileResolverMock,
				unifiedFactoryMock,
				config,
				ioAdapterMock,
			)

			err = analyzer.Analyze(tc.paths, tc.format)

			if tc.expectedError != nil {
				require.Error(t, err)

				switch expected := tc.expectedError.(type) {
				case *domain.ErrNoPathsProvided:
					var actual *domain.ErrNoPathsProvided

					assert.True(t, errors.As(err, &actual), "ожидалась ErrNoPathsProvided")
				case *domain.ErrInvalidTimeFormat:
					var actual *domain.ErrInvalidTimeFormat

					assert.True(t, errors.As(err, &actual), "ожидалась ErrInvalidTimeFormat")
				case *domain.ErrURLFetch:
					var actual *domain.ErrURLFetch

					assert.True(t, errors.As(err, &actual), "ожидалась ErrURLFetch")
				case *domain.ErrFileOpen:
					var actual *domain.ErrFileOpen

					assert.True(t, errors.As(err, &actual), "ожидалась ErrFileOpen")
				case *domain.ErrInvalidFormat:
					var actual *domain.ErrInvalidFormat

					assert.True(t, errors.As(err, &actual), "ожидалась ErrInvalidFormat")
				default:
					assert.Contains(t, err.Error(), expected.Error(), "ошибка содержит ожидаемую строку")
				}
			} else {
				require.NoError(t, err)
			}

			parserMock.AssertExpectations(t)
			fileResolverMock.AssertExpectations(t)
			fileReaderFactoryMock.AssertExpectations(t)
			urlReaderFactoryMock.AssertExpectations(t)
			ioAdapterMock.AssertExpectations(t)
		})
	}
}
