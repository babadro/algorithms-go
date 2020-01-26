package _1333_filter_restaurants_by_vegan_price_and_distance

import "sort"

/*
type Restaurant struct {
	id int
	rating int
	veanFriendly int
	price int
	distance int
}
*/

type Restaurant []int

func filterRestaurants(restaurants [][]int, veganFriendly int, maxPrice int, maxDistance int) []int {
	res := make([]Restaurant, 0, len(restaurants)/2)
	for _, r := range restaurants {
		if filter(r, veganFriendly, maxPrice, maxDistance) {
			res = append(res, r)
		}
	}
	sort.Slice(res, func(i, j int) bool {
		if res[i][1] > res[j][1] {
			return true
		}
		if res[i][1] == res[j][1] {
			return res[i][0] > res[j][0]
		}
		return false
	})
	ids := make([]int, len(res))
	for i := range res {
		ids[i] = res[i][0]
	}
	return ids
}

func filter(rest Restaurant, veganFriendly int, maxPrice int, maxDistance int) bool {
	if veganFriendly == 1 && rest[2] == 0 {
		return false
	}
	if rest[3] > maxPrice {
		return false
	}
	if rest[4] > maxDistance {
		return false
	}
	return true
}
