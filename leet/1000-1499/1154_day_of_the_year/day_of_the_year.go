package _1154_day_of_the_year

import "strconv"

// passed. tptl. trick - 1900 is not e leap year.
func dayOfYear(date string) int {
	year, _ := strconv.Atoi(date[:4])
	month, _ := strconv.Atoi(date[5:7])
	day, _ := strconv.Atoi(date[8:])

	count := 0
	for i := 1; i < month; i++ {
		count += daysCount(year, i)
	}

	return count + day
}

func daysCount(year, month int) int {
	if month == 2 {
		if year > 1900 && year%4 == 0 {
			return 29
		}

		return 28
	}

	if (month <= 7 && month%2 == 1) || (month > 7 && month%2 == 0) {
		return 31
	}

	return 30
}
