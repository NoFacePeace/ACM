package sort

import "sort"

func filterRestaurants(restaurants [][]int, veganFriendly int, maxPrice int, maxDistance int) []int {
	filter := func() [][]int {
		arr := [][]int{}
		for _, r := range restaurants {
			v, p, d := r[2], r[3], r[4]
			if v == 0 && veganFriendly == 1 {
				continue
			}
			if p > maxPrice {
				continue
			}
			if d > maxDistance {
				continue
			}
			arr = append(arr, r)
		}
		return arr
	}
	restaurants = filter()
	sort.Slice(restaurants, func(a, b int) bool {
		if restaurants[a][1] == restaurants[b][1] {
			return restaurants[a][0] > restaurants[b][0]
		}
		return restaurants[a][1] > restaurants[b][1]
	})
	ans := []int{}
	for _, v := range restaurants {
		ans = append(ans, v[0])
	}
	return ans
}
