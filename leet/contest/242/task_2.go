package _42

import "math"

// todo 1
func minSpeedOnTime(dist []int, hour float64) int {
	n := len(dist)
	if hour <= float64(n) {
		return -1
	}

	max := 0
	for i, d := range dist {
		t := hour - float64(i+1)

		speed := int(math.Ceil(float64(d) / t))

		if speed > max {
			max = speed
		}
	}

	return 0
}
