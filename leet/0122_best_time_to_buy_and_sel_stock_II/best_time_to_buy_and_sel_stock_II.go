package _122_best_time_to_buy_and_sel_stock_II

func maxProfit(prices []int) int {
	profit := 0
	lastBuy := -1

	for i := 0; i < len(prices); i++ {
		peak := (i == len(prices)-1 || prices[i+1] < prices[i]) && (i == 0 || prices[i-1] < prices[i])
		if peak {
			if lastBuy == -1 {
				continue
			}
			profit += prices[i]
			lastBuy = -1
			continue
		}

		pit := (i == len(prices)-1 || prices[i+1] > prices[i]) && (i == 0 || prices[i-1] > prices[i])
		if pit {
			profit -= prices[i]
			lastBuy = prices[i]
		}
	}

	if profit < 0 {
		return 0
	}

	if lastBuy > 0 {
		profit -= lastBuy
	}

	return profit
}
