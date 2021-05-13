package _1854_maximum_population_year

// tptl. passed. best solution. easy (medium for me). inspired by https://leetcode.com/problems/maximum-population-year/discuss/1200399/Go-golang-solution
func maximumPopulation2(logs [][]int) int {
	years := make([]int, 101)
	for _, log := range logs {
		idxBirth, idxDeath := log[0]-1950, log[1]-1950
		years[idxBirth]++
		years[idxDeath]--
	}

	maxPopulation, currPopulation, idx := 0, 0, 0
	for i, delta := range years {
		currPopulation += delta
		if currPopulation > maxPopulation {
			idx, maxPopulation = i, currPopulation
		}
	}

	return idx + 1950
}

// passed but it's O(n^2).
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
