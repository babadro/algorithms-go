package _1362_closest_divisors

import "math"

// it works
func closestDivisors(num int) []int {
	sqrt := math.Sqrt(float64(num + 2))
	limit := int(math.Ceil(sqrt))
	num1, num2 := num+1, num+2
	defaultMinDistance := (num + 2) - 1
	minDistance := defaultMinDistance
	res := []int{1, num + 2}
	for i := limit; i >= 0; i-- {
		if minDistance != defaultMinDistance {
			break
		}
		res1 := num1 / i
		remainder1 := num1 % i
		res2 := num2 / i
		remainder2 := num2 % i

		if remainder1 == 0 {
			distance := res1 - i
			if distance < 0 {
				distance = -distance
			}
			if distance < minDistance {
				minDistance = distance
				res = []int{i, res1}
			}

		}

		if remainder2 == 0 {
			distance := res2 - i
			if distance < 0 {
				distance = -distance
			}
			if distance < minDistance {
				minDistance = distance
				res = []int{i, res2}
			}
		}

	}

	return res
}
