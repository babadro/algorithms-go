package _1744_can_you_eat_your_favorite_candy_on_your_favorite_day

// tptl. passed.
func canEat(candiesCount []int, queries [][]int) []bool {
	sum := make([]int, len(candiesCount))
	sum[0] = candiesCount[0]
	for i := 1; i < len(sum); i++ {
		sum[i] = candiesCount[i] + sum[i-1]
	}

	res := make([]bool, len(queries))
	for i := 0; i < len(res); i++ {
		favType, favDay, dailyCap := queries[i][0], queries[i][1], queries[i][2]
		if favDay >= sum[favType] {
			res[i] = false
			continue
		}

		if favType > 0 && (favDay+1)*dailyCap < sum[favType-1]+1 {
			res[i] = false
			continue
		}

		res[i] = true
	}

	return res
}
