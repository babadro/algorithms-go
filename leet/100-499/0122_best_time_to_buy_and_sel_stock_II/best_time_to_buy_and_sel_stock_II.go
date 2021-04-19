package _122_best_time_to_buy_and_sel_stock_II

import "fmt"

// doesn't work
func maxProfit2(prices []int) int {
	if len(prices) == 1 {
		return 0
	}

	profit := 0
	lastPrice := 0
	for i := 0; i < len(prices); i++ {
		for i < len(prices)-1 && prices[i] == prices[i+1] {
			i++
			continue
		}
		timeToBuy := (i == 0 && prices[1] > prices[0]) || (i < len(prices)-1 && lastPrice > prices[i] && prices[i+1] > prices[i])
		if timeToBuy {
			profit -= prices[i]
			fmt.Println("timeToBuy, i=", i)
		} else if timeToSell := (i == len(prices)-1 && lastPrice < prices[i]) || (i > 0 && i < len(prices)-1 && lastPrice < prices[i] && prices[i+1] < prices[i]); timeToSell {
			profit += prices[i]
			fmt.Println("timeToSell, i=", i)
		}

		lastPrice = prices[i]
	}

	return profit
}
