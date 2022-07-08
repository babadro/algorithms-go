package _1774_closest_dessert_cost

import (
	"math"
	"sort"
)

// passed. easy to understand
func closestCost2(baseCosts []int, toppingCosts []int, target int) int {
	allToppings := append(toppingCosts, toppingCosts...)
	closestPrice := 0
	smallestDiff := math.MaxInt64

	for _, price := range baseCosts {
		helper(price, target, &smallestDiff, &closestPrice, allToppings)
	}

	return closestPrice
}

func helper(price, target int, smallestDiff, closest *int, toppings []int) {
	if *smallestDiff == 0 {
		return
	}

	diff := price - target
	if price < target {
		diff = target - price
	}

	if diff == 0 {
		*closest, *smallestDiff = price, 0
		return
	} else if diff < *smallestDiff {
		*closest, *smallestDiff = price, diff
	} else if diff == *smallestDiff && price < *closest {
		*closest = price
	}

	if price < target && len(toppings) > 0 {
		for i := 0; i < len(toppings); i++ {
			helper(price+toppings[i], target, smallestDiff, closest, toppings[i+1:])
		}
	}
}

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
