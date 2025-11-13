package util

import "time"

func ParseDate(dateLayout, date string) (time.Time, error) {
	return time.Parse(dateLayout, date)
}