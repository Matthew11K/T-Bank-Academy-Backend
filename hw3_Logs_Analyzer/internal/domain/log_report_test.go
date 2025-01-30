package domain_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/central-university-dev/backend_academy_2024_project_3-go-Matthew11K/internal/domain"
)

func TestLogReport_Format(t *testing.T) {
	testCases := []struct {
		name          string
		records       []*domain.LogRecord
		format        string
		expectedParts []string
		hasError      bool
	}{
		{
			name: "Markdown Format with multiple records",
			records: []*domain.LogRecord{
				{
					Request:       "GET /index.html HTTP/1.1",
					Status:        200,
					BodyBytesSent: 1000,
					RemoteAddr:    "192.168.1.1",
					HTTPUserAgent: "Mozilla/5.0",
				},
				{
					Request:       "POST /api/data HTTP/1.1",
					Status:        201,
					BodyBytesSent: 1500,
					RemoteAddr:    "192.168.1.2",
					HTTPUserAgent: "curl/7.64.1",
				},
			},
			format: "markdown",
			expectedParts: []string{
				"#### Общая информация",
				"|        Метрика        |     Значение |",
				"|:---------------------:|-------------:|",
				"|  Количество запросов  |       2 |",
				"| Средний размер ответа |         1250.00b |",
				"|   95p размера ответа  |         1500b |",
				"#### Распределение методов запросов",
				"|   Метод   | Количество |",
				"|:---------:|-----------:|",
				"| GET | 1 |",
				"| POST | 1 |",
				"#### Топ User-Agent'ов",
				"|  User-Agent  | Количество |",
				"|:------------:|-----------:|",
				"| Mozilla/5.0 | 1 |",
				"| curl/7.64.1 | 1 |",
				"#### Запрашиваемые ресурсы",
				"|     Ресурс      | Количество |",
				"|:---------------:|-----------:|",
				"| GET /index.html HTTP/1.1 | 1 |",
				"| POST /api/data HTTP/1.1 | 1 |",
				"#### Коды ответа",
				"| Код |          Имя          | Количество |",
				"|:---:|:---------------------:|-----------:|",
				"| 200 |                    OK | 1 |",
				"| 201 |               Created | 1 |",
				"#### Активные клиенты",
				"|    IP-адрес    | Количество |",
				"|:--------------:|-----------:|",
				"| 192.168.1.1 | 1 |",
				"| 192.168.1.2 | 1 |",
			},
			hasError: false,
		},
		{
			name: "AsciiDoc Format with single record",
			records: []*domain.LogRecord{
				{
					Request:       "GET /home HTTP/1.1",
					Status:        200,
					BodyBytesSent: 500,
					RemoteAddr:    "192.168.1.3",
					HTTPUserAgent: "Safari/537.36",
				},
			},
			format: "adoc",
			expectedParts: []string{
				"=== Общая информация",
				"|Метрика|Значение|",
				"|---|---|",
				"|Количество запросов|1|",
				"|Средний размер ответа|500.00b|",
				"|95p размера ответа|500b|",
				"=== Распределение методов запросов",
				"|Метод|Количество|",
				"|---|---|",
				"|GET|1|",
				"=== Топ User-Agent'ов",
				"|User-Agent|Количество|",
				"|---|---|",
				"|Safari/537.36|1|",
				"=== Запрашиваемые ресурсы",
				"|Ресурс|Количество|",
				"|---|---|",
				"|GET /home HTTP/1.1|1|",
				"=== Коды ответа",
				"|Код|Имя|Количество|",
				"|---|---|---|",
				"|200|OK|1|",
				"=== Активные клиенты",
				"|IP-адрес|Количество|",
				"|---|---|",
				"|192.168.1.3|1|",
			},
			hasError: false,
		},
		{
			name:          "Invalid Format",
			records:       []*domain.LogRecord{},
			format:        "invalid",
			expectedParts: []string{},
			hasError:      true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			report := domain.NewLogReport()
			for _, record := range tc.records {
				report.AddRecord(record)
			}

			output, err := report.Format(domain.ReportFormatType(tc.format))

			if tc.hasError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)

				for _, part := range tc.expectedParts {
					assert.Contains(t, output, part)
				}
			}
		})
	}
}
