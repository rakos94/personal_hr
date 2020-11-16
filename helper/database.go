package helper

import (
	"errors"
	"time"
)

// Timestamp ...
type Timestamp time.Time

// UnmarshalJSON ...
func (t *Timestamp) UnmarshalJSON(data []byte) error {
	layout := "2006-01-02 15:04:05"

	ts, err := validateDataTime(layout, data)
	*t = Timestamp(ts)

	return err
}

// Date ...
type Date time.Time

// UnmarshalJSON ...
func (t *Date) UnmarshalJSON(data []byte) error {
	layout := "2006-01-02"

	ts, err := validateDataTime(layout, data)
	*t = Date(ts)

	return err
}

func validateDataTime(layout string, data []byte) (time.Time, error) {
	var timeStr = string(data)
	if timeStr == "null" || timeStr == `""` {
		return time.Time{}, errors.New("Data is Null")
	}
	if len(timeStr) > 0 && timeStr[0] == '"' {
		timeStr = timeStr[1:]
	}
	if len(timeStr) > 0 && timeStr[len(timeStr)-1] == '"' {
		timeStr = timeStr[:len(timeStr)-1]
	}

	ts, err := time.Parse(layout, timeStr)
	if err != nil {
		return time.Time{}, err
	}

	return ts, nil
}
