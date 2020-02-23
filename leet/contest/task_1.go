package contest

import (
	"log"
	"time"
)

func daysBetweenDates(date1 string, date2 string) int {
	d1, err := time.Parse("2006-01-02", date1)
	if err != nil {
		log.Fatal(err)
	}
	d2, err := time.Parse("2006-01-02", date2)
	if err != nil {
		log.Fatal(err)
	}
	res := int(d1.Sub(d2).Hours() / 24)
	if res < 0 {
		res = -res
	}
	return res
}
