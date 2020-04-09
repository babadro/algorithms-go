package _121_best_time_to_buy_and_sell_stock

import (
	"math"
)

func maxProfit(prices []int) int {
	max, min, maxProfit := 0, math.MaxInt32, 0
	for _, price := range prices {
		if price < min {
			min = price
			max = 0
			continue
		}
		if price > max {
			max = price
			if profit := max - min; profit > maxProfit {
				maxProfit = profit
			}
		}
	}
	return maxProfit
}
