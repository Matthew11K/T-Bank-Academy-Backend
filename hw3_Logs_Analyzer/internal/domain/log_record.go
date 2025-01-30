package domain

import (
	"regexp"
	"strings"
)

type FilterField string

const (
	FieldAgent  FilterField = "agent"
	FieldMethod FilterField = "method"
)

type LogField string

const (
	FieldRemoteAddr    LogField = "remote_addr"
	FieldRemoteUser    LogField = "remote_user"
	FieldTimeLocal     LogField = "time_local"
	FieldRequest       LogField = "request"
	FieldStatus        LogField = "status"
	FieldBodyBytesSent LogField = "body_bytes_sent"
	FieldHTTPReferer   LogField = "http_referer"
	FieldHTTPUserAgent LogField = "http_user_agent"
)

type LogRecord struct {
	RemoteAddr    string
	RemoteUser    string
	Time          *TimeWrapper
	Request       string
	Status        int
	BodyBytesSent int
	HTTPReferer   string
	HTTPUserAgent string
}

func NewLogRecord() *LogRecord {
	return &LogRecord{}
}

func (lr *LogRecord) MatchesFilter(field FilterField, value string) bool {
	if field == "" || value == "" {
		return true
	}

	switch field {
	case FieldAgent:
		matched, _ := regexp.MatchString(value, lr.HTTPUserAgent)
		return matched
	case FieldMethod:
		parts := strings.Split(lr.Request, " ")
		if len(parts) > 0 {
			return parts[0] == value
		}

		return false
	default:
		return true
	}
}

func (lr *LogRecord) Validate(startTime, endTime *TimeWrapper, filterField FilterField, filterValue string) bool {
	if startTime != nil && lr.Time.Before(startTime) {
		return false
	}

	if endTime != nil && lr.Time.After(endTime) {
		return false
	}

	if !lr.MatchesFilter(filterField, filterValue) {
		return false
	}

	return true
}
