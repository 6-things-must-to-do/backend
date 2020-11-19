package timeUtil

import "time"

func GetDate(date time.Time) string {
	return date.Format("2006-01-02")
}
