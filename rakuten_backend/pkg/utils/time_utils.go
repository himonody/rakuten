package utils

import (
	"time"
)

func EndOfDay() string {
	now := time.Now()
	return time.Date(
		now.Year(), now.Month(), now.Day(),
		23, 59, 59, 0, now.Location(),
	).Format("2006-01-02 15:04:05")
}

func StartOfDay() string {
	now := time.Now()
	return time.Date(
		now.Year(), now.Month(), now.Day(),
		0, 0, 0, 0, now.Location(),
	).Format("2006-01-02 15:04:05")
}
