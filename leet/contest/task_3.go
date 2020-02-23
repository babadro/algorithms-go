package contest

import "math"

// TODO
func closestDivisors(num int) []int {
	sqrt := math.Sqrt(float64(num))
	ceiled := math.Ceil(sqrt)
	num1, num2 := int(ceiled), int(ceiled)
	if check(num, num1, num2) {
		return []int{num1, num2}
	}
	lastTimeWasNum1 := false
	for num1 > 0 && num2 > 0 && num1 < num && num2 < num {
		if lastTimeWasNum1 {
			lastTimeWasNum1 = false
			num2++
		} else {
			lastTimeWasNum1 = true
			num1--
		}
		if check(num, num1, num2) {
			return []int{num1, num2}
		}
	}
	return []int{0, 0}
}

func check(target, divisor1, divisor2 int) bool {
	res := divisor1 * divisor2
	return res == target+1 || res == target+2
}
