package _30

import (
	"math"
	"sort"
)

// todo 1 doesn't work
func closestCost(baseCosts []int, toppingCosts []int, target int) int {
	allToppings := append(toppingCosts, toppingCosts...)
	sort.Ints(allToppings)
	closestPrice := 0
	smallestDiff := math.MaxInt64

	for _, price := range baseCosts {
		for tail, head := -1, -1; tail <= head && head < len(allToppings); {
			diff := price - target
			if price < target {
				diff = target - price
			}

			if diff == 0 {
				return target
			} else if diff < smallestDiff {
				closestPrice, smallestDiff = price, diff
			} else if diff == smallestDiff && price < closestPrice {
				closestPrice = price
			}

			if price > target {
				tail++
				price -= allToppings[tail]
			} else {
				head++
				if head < len(allToppings) {
					price += allToppings[head]
				}
			}
		}
	}

	return closestPrice
}
