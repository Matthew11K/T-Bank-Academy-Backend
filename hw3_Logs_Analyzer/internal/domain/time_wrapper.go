package domain

import (
	"time"
)

type TimeWrapper struct {
	Time time.Time
}

func (tw *TimeWrapper) Before(other *TimeWrapper) bool {
	if other == nil {
		return false
	}

	return tw.Time.Before(other.Time)
}

func (tw *TimeWrapper) After(other *TimeWrapper) bool {
	if other == nil {
		return false
	}

	return tw.Time.After(other.Time)
}

func ParseTime(value string) (*TimeWrapper, error) {
	layouts := []string{
		time.RFC3339,
		"2006-01-02T15:04:05Z07:00",
		"2006-01-02T15:04:05.999999999Z07:00",
		"2006-01-02",
	}

	var t time.Time

	var err error

	for _, layout := range layouts {
		t, err = time.Parse(layout, value)
		if err == nil {
			return &TimeWrapper{Time: t}, nil
		}
	}

	return nil, &ErrInvalidTimeFormat{Err: err}
}

func ParseNginxTime(value string) (*TimeWrapper, error) {
	layout := "02/Jan/2006:15:04:05 -0700"

	t, err := time.Parse(layout, value)
	if err != nil {
		return nil, &ErrInvalidTimeFormat{Err: err}
	}

	return &TimeWrapper{Time: t}, nil
}
