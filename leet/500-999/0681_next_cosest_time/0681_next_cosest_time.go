package _681_next_cosest_time

import "strconv"

// #bnsrg medium passed
// todo 2 check shorter solutions or not bruteforce solutions
func nextClosestTime(time string) string {
	var allowed [10]bool

	allowed[time[0]-'0'] = true
	allowed[time[1]-'0'] = true
	allowed[time[3]-'0'] = true
	allowed[time[4]-'0'] = true

	hours, _ := strconv.Atoi(time[0:2])
	minutes, _ := strconv.Atoi(time[3:])

	for {
		if minutes < 59 {
			minutes++
		} else {
			minutes = 0

			if hours < 23 {
				hours++
			} else {
				hours = 0
			}
		}

		var digit1, digit2, digit3, digit4 int

		if digit1 = hours / 10; !allowed[digit1] {
			continue
		}

		if digit2 = hours % 10; !allowed[digit2] {
			continue
		}

		if digit3 = minutes / 10; !allowed[digit3] {
			continue
		}

		if digit4 = minutes % 10; !allowed[digit4] {
			continue
		}

		b := []byte("00:00")

		b[0] = '0' + byte(digit1)
		b[1] = '0' + byte(digit2)
		b[3] = '0' + byte(digit3)
		b[4] = '0' + byte(digit4)

		return string(b)
	}
}
