package _1774_closest_dessert_cost

import "math"

func closestCost3(baseCosts []int, toppingCosts []int, target int) int {
	toppingCosts = append(toppingCosts, toppingCosts...)

	diff, price := math.MinInt64, 0
	for _, base :

}

func abs(num int) int {
	if num < 0 {
		return -num
	}

	return num
}
