package _122_best_time_to_buy_and_sel_stock_II

// best solution.
// tptl
// todo 1 read solutions explanation
func maxProfit(prices []int) int {
	maxProfit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			maxProfit += prices[i] - prices[i-1]
		}
	}
	return maxProfit
}
