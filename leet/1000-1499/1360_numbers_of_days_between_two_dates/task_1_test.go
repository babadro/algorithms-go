package _1360_numbers_of_days_between_two_dates

import "testing"

func TestDaysBetweenDays(t *testing.T) {
	t.Log(daysBetweenDates("2019-06-29", "2019-06-30"))
	t.Log(daysBetweenDates("2020-01-15", "2019-12-31"))
}
