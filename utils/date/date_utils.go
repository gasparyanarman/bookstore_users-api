package date

import (
	"time"
)

const (
	layout = "2006-01-02T15:04:05Z"
)

func GetNowString() string {
	now := time.Now()
	return now.Format(layout)
}
