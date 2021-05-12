package _1854_maximum_population_year

// passed but it's O(n^2). todo 1 https://leetcode.com/problems/maximum-population-year/discuss/1200399/Go-golang-solution
func maximumPopulation(logs [][]int) int {
	max, year := 0, 0
	for i := 1950; i < 2050; i++ {
		counter := 0
		for _, log := range logs {
			if i >= log[0] && i < log[1] {
				counter++
			}
		}

		if counter > max {
			max, year = counter, i
		}
	}

	return year
}
