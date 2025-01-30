package infrastructure

import (
	"regexp"
	"strconv"

	"github.com/central-university-dev/backend_academy_2024_project_3-go-Matthew11K/internal/application"
	"github.com/central-university-dev/backend_academy_2024_project_3-go-Matthew11K/internal/domain"
)

type LogParser struct {
	regex *regexp.Regexp
}

func NewLogParser() application.LogParser {
	pattern := `^(?P<remote_addr>[^ ]+) - (?P<remote_user>[^ ]+) \[` +
		`(?P<time_local>[^\]]+)\] "(?P<request>[^"]+)" (?P<status>\d{3}) ` +
		`(?P<body_bytes_sent>\d+) "(?P<http_referer>[^"]*)" "` +
		`(?P<http_user_agent>[^"]*)"$`
	regex := regexp.MustCompile(pattern)

	return &LogParser{
		regex: regex,
	}
}

func (lp *LogParser) Parse(line string) (*domain.LogRecord, error) {
	matches := lp.regex.FindStringSubmatch(line)
	if matches == nil {
		return nil, &domain.ErrInvalidLogFormat{}
	}

	record := domain.NewLogRecord()

	for i, name := range lp.regex.SubexpNames() {
		if i == 0 || name == "" {
			continue
		}

		switch name {
		case string(domain.FieldRemoteAddr):
			record.RemoteAddr = matches[i]
		case string(domain.FieldRemoteUser):
			record.RemoteUser = matches[i]
		case string(domain.FieldTimeLocal):
			time, err := domain.ParseNginxTime(matches[i])
			if err != nil {
				return nil, err
			}

			record.Time = time
		case string(domain.FieldRequest):
			record.Request = matches[i]
		case string(domain.FieldStatus):
			status, err := strconv.Atoi(matches[i])
			if err != nil {
				return nil, &domain.ErrParseInt{Field: string(domain.FieldStatus), Err: err}
			}

			record.Status = status
		case string(domain.FieldBodyBytesSent):
			bytesSent, err := strconv.Atoi(matches[i])
			if err != nil {
				return nil, &domain.ErrParseInt{Field: string(domain.FieldBodyBytesSent), Err: err}
			}

			record.BodyBytesSent = bytesSent
		case string(domain.FieldHTTPReferer):
			record.HTTPReferer = matches[i]
		case string(domain.FieldHTTPUserAgent):
			record.HTTPUserAgent = matches[i]
		}
	}

	return record, nil
}
