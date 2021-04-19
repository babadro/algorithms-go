package _121_best_time_to_buy_and_sell_stock

import (
	"math"
)

func maxProfit(prices []int) int {
	min, maxProfit := math.MaxInt32, 0
	for _, price := range prices {
		if price < min {
			min = price
		} else if profit := price - min; profit > maxProfit {
			maxProfit = profit
		}
	}
	return maxProfit
}
