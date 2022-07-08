package _1774_closest_dessert_cost

import "math"

// tptl. passed
func closestCost3(baseCosts []int, toppingCosts []int, target int) int {
	best := math.MaxInt64
	for _, base := range baseCosts {
		rec(toppingCosts, 0, base, target, &best)
	}

	return best
}

func rec(toppingCosts []int, i, cost, target int, best *int) {
	if cost >= target || i == len(toppingCosts) {
		currDiff, bestDiff := abs(cost-target), abs(*best-target)
		if currDiff < bestDiff {
			*best = cost
		} else if currDiff == bestDiff {
			*best = min(*best, cost)
		}
	} else {
		rec(toppingCosts, i+1, cost, target, best)                   // skip curr topping
		rec(toppingCosts, i+1, cost+toppingCosts[i], target, best)   // take one curr topping
		rec(toppingCosts, i+1, cost+2*toppingCosts[i], target, best) // take two curr toppings
	}

}

func abs(num int) int {
	if num < 0 {
		return -num
	}

	return num
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
