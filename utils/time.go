package utils

import (
	"time"
)

// GenerateCalendarWeeks generates an array of calendar weeks to a given end date
func GenerateCalendarWeeks(endDate time.Time) []int {
	var calendarWeeks []int

	_, endWeek := endDate.ISOWeek()

	for i := 1; i <= endWeek; i++ {
		calendarWeeks = append(calendarWeeks, i)
	}

	return calendarWeeks
}

func GetIsoWeek(date time.Time) int {
	_, week := date.ISOWeek()

	return week
}
