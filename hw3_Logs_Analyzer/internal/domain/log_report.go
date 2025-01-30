package domain

import (
	"fmt"
	"math"
	"net/http"
	"sort"
	"strings"

	"golang.org/x/exp/constraints"
)

type ReportFormat struct {
	HeaderGeneral       string
	HeaderMethods       string
	HeaderUserAgents    string
	HeaderResources     string
	HeaderStatusCodes   string
	HeaderActiveClients string

	TableGeneral       string
	TableMethods       string
	TableUserAgents    string
	TableResources     string
	TableStatusCodes   string
	TableActiveClients string

	MetricRequests        string
	MetricAverageResponse string
	MetricPercentile95    string
	MethodFormat          string
	UserAgentFormat       string
	ResourceFormat        string
	StatusCodeFormat      string
	ActiveClientFormat    string
}

type ReportFormatType string

const (
	MarkdownFormat ReportFormatType = "markdown"
	AdocFormat     ReportFormatType = "adoc"
)

var (
	markdownFormat = &ReportFormat{
		HeaderGeneral:       "#### Общая информация\n\n",
		HeaderMethods:       "\n#### Распределение методов запросов\n\n",
		HeaderUserAgents:    "\n#### Топ User-Agent'ов\n\n",
		HeaderResources:     "\n#### Запрашиваемые ресурсы\n\n",
		HeaderStatusCodes:   "\n#### Коды ответа\n\n",
		HeaderActiveClients: "\n#### Активные клиенты\n\n",

		TableGeneral:       "|        Метрика        |     Значение |\n|:---------------------:|-------------:|\n",
		TableMethods:       "|   Метод   | Количество |\n|:---------:|-----------:|\n",
		TableUserAgents:    "|  User-Agent  | Количество |\n|:------------:|-----------:|\n",
		TableResources:     "|     Ресурс      | Количество |\n|:---------------:|-----------:|\n",
		TableStatusCodes:   "| Код |          Имя          | Количество |\n|:---:|:---------------------:|-----------:|\n",
		TableActiveClients: "|    IP-адрес    | Количество |\n|:--------------:|-----------:|\n",

		MetricRequests:        "|  Количество запросов  |       %d |\n",
		MetricAverageResponse: "| Средний размер ответа |         %.2fb |\n",
		MetricPercentile95:    "|   95p размера ответа  |         %db |\n",
		MethodFormat:          "| %s | %d |\n",
		UserAgentFormat:       "| %s | %d |\n",
		ResourceFormat:        "| %s | %d |\n",
		StatusCodeFormat:      "| %d | %21s | %d |\n",
		ActiveClientFormat:    "| %s | %d |\n",
	}

	adocFormat = &ReportFormat{
		HeaderGeneral:       "=== Общая информация\n\n",
		HeaderMethods:       "\n=== Распределение методов запросов\n\n",
		HeaderUserAgents:    "\n=== Топ User-Agent'ов\n\n",
		HeaderResources:     "\n=== Запрашиваемые ресурсы\n\n",
		HeaderStatusCodes:   "\n=== Коды ответа\n\n",
		HeaderActiveClients: "\n=== Активные клиенты\n\n",

		TableGeneral:       "|Метрика|Значение|\n|---|---|\n",
		TableMethods:       "|Метод|Количество|\n|---|---|\n",
		TableUserAgents:    "|User-Agent|Количество|\n|---|---|\n",
		TableResources:     "|Ресурс|Количество|\n|---|---|\n",
		TableStatusCodes:   "|Код|Имя|Количество|\n|---|---|---|\n",
		TableActiveClients: "|IP-адрес|Количество|\n|---|---|\n",

		MetricRequests:        "|Количество запросов|%d|\n",
		MetricAverageResponse: "|Средний размер ответа|%.2fb|\n",
		MetricPercentile95:    "|95p размера ответа|%db|\n",
		MethodFormat:          "|%s|%d|\n",
		UserAgentFormat:       "|%s|%d|\n",
		ResourceFormat:        "|%s|%d|\n",
		StatusCodeFormat:      "|%d|%s|%d|\n",
		ActiveClientFormat:    "|%s|%d|\n",
	}
)

type LogReport struct {
	TotalRequests       int
	Resources           map[string]int
	StatusCodes         map[int]int
	ResponseSizes       []int
	AverageResponseSize float64
	Percentile95        int
	Clients             map[string]int
	Methods             map[string]int
	UserAgents          map[string]int
}

func NewLogReport() *LogReport {
	return &LogReport{
		Resources:   make(map[string]int),
		StatusCodes: make(map[int]int),
		Clients:     make(map[string]int),
		Methods:     make(map[string]int),
		UserAgents:  make(map[string]int),
	}
}

func (lr *LogReport) AddRecord(record *LogRecord) {
	lr.TotalRequests++
	lr.Resources[record.Request]++
	lr.StatusCodes[record.Status]++
	lr.ResponseSizes = append(lr.ResponseSizes, record.BodyBytesSent)
	lr.Clients[record.RemoteAddr]++

	method := record.GetMethod()
	lr.Methods[method]++
	lr.UserAgents[record.HTTPUserAgent]++
}

func (lr *LogReport) calculateStatistics() {
	totalSize := 0

	for _, size := range lr.ResponseSizes {
		totalSize += size
	}

	if lr.TotalRequests > 0 {
		lr.AverageResponseSize = float64(totalSize) / float64(lr.TotalRequests)
	}

	sort.Ints(lr.ResponseSizes)
	n := len(lr.ResponseSizes)

	if n == 0 {
		lr.Percentile95 = 0
		return
	}

	pos := math.Ceil(0.95 * float64(n))
	index := int(pos) - 1

	if index < 0 {
		index = 0
	}

	if index >= n {
		index = n - 1
	}

	lr.Percentile95 = lr.ResponseSizes[index]
}

func (lr *LogReport) Format(format ReportFormatType) (string, error) {
	lr.calculateStatistics()

	switch format {
	case MarkdownFormat:
		return lr.formatReport(markdownFormat), nil
	case AdocFormat:
		return lr.formatReport(adocFormat), nil
	default:
		return "", &ErrInvalidFormat{}
	}
}

func (lr *LogReport) formatReport(fmtConfig *ReportFormat) string {
	var sb strings.Builder

	sb.WriteString(fmtConfig.HeaderGeneral)
	sb.WriteString(fmtConfig.TableGeneral)
	sb.WriteString(fmt.Sprintf(fmtConfig.MetricRequests, lr.TotalRequests))
	sb.WriteString(fmt.Sprintf(fmtConfig.MetricAverageResponse, lr.AverageResponseSize))
	sb.WriteString(fmt.Sprintf(fmtConfig.MetricPercentile95, lr.Percentile95))

	sb.WriteString(fmtConfig.HeaderMethods)
	sb.WriteString(fmtConfig.TableMethods)

	methods := sortMapByValue(lr.Methods)
	for _, kv := range methods {
		sb.WriteString(fmt.Sprintf(fmtConfig.MethodFormat, kv.Key, kv.Value))
	}

	sb.WriteString(fmtConfig.HeaderUserAgents)
	sb.WriteString(fmtConfig.TableUserAgents)

	userAgents := sortMapByValue(lr.UserAgents)
	for i, kv := range userAgents {
		if i >= 5 {
			break
		}

		sb.WriteString(fmt.Sprintf(fmtConfig.UserAgentFormat, kv.Key, kv.Value))
	}

	sb.WriteString(fmtConfig.HeaderResources)
	sb.WriteString(fmtConfig.TableResources)

	resources := sortMapByValue(lr.Resources)
	for _, kv := range resources {
		sb.WriteString(fmt.Sprintf(fmtConfig.ResourceFormat, kv.Key, kv.Value))
	}

	sb.WriteString(fmtConfig.HeaderStatusCodes)
	sb.WriteString(fmtConfig.TableStatusCodes)

	codes := sortMapByValue(lr.StatusCodes)
	for _, kv := range codes {
		sb.WriteString(fmt.Sprintf(fmtConfig.StatusCodeFormat, kv.Key, httpStatusText(kv.Key), kv.Value))
	}

	sb.WriteString(fmtConfig.HeaderActiveClients)
	sb.WriteString(fmtConfig.TableActiveClients)

	clients := sortMapByValue(lr.Clients)
	for _, kv := range clients {
		sb.WriteString(fmt.Sprintf(fmtConfig.ActiveClientFormat, kv.Key, kv.Value))
	}

	return sb.String()
}

func httpStatusText(code int) string {
	text := http.StatusText(code)
	if text != "" {
		return text
	}

	return "Unknown Status"
}

type KV[K comparable, V any] struct {
	Key   K
	Value V
}

func sortMapByValue[K comparable, V constraints.Ordered](m map[K]V) []KV[K, V] {
	ss := make([]KV[K, V], 0, len(m))

	for k, v := range m {
		ss = append(ss, KV[K, V]{Key: k, Value: v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})

	return ss
}

func (lr *LogRecord) GetMethod() string {
	parts := strings.Fields(lr.Request)
	if len(parts) > 0 {
		return parts[0]
	}

	return "Неизвестный метод"
}
