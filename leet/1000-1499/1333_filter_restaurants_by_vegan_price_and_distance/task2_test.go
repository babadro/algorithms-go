package _1333_filter_restaurants_by_vegan_price_and_distance

import "testing"

func TestFilter(t *testing.T) {
	input := [][]int{
		{1, 4, 1, 40, 10},
		{2, 8, 0, 50, 5},
		{3, 8, 1, 30, 4},
		{4, 10, 0, 10, 3},
		{5, 1, 1, 15, 1},
	}
	t.Log(filterRestaurants(input, 1, 50, 10))

	input2 := [][]int{
		{1, 4, 1, 40, 10},
		{2, 8, 0, 50, 5},
		{3, 8, 1, 30, 4},
		{4, 10, 0, 10, 3},
		{5, 1, 1, 15, 1},
	}

	t.Log(filterRestaurants(input2, 0, 50, 10))
}
