package _121_best_time_to_buy_and_sell_stock

import "sort"

func maxProfit(prices []int) int {
	pairs := make([]pair, len(prices))
	for i := range prices {
		pairs[i] = pair{prices[i], i}
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].price < pairs[j].price
	})

	for j := range pairs {

	}
}

type pair struct {
	price, idx int
}
