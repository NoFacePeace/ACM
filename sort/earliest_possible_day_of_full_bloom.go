package sort

import "sort"

func earliestFullBloom(plantTime []int, growTime []int) int {
	n := len(plantTime)
	if n == 0 {
		return 0
	}
	plants := []plant{}
	for i := 0; i < n; i++ {
		p := plant{
			plantTime: plantTime[i],
			growTime:  growTime[i],
		}
		plants = append(plants, p)
	}
	sort.Slice(plants, func(a, b int) bool {
		return plants[a].growTime > plants[b].growTime
	})
	ans := plants[0].plantTime + plants[0].growTime
	time := plants[0].plantTime
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	for i := 1; i < n; i++ {
		time += plants[i].plantTime
		ans = max(ans, time+plants[i].growTime)
	}
	return ans
}

type plant struct {
	plantTime int
	growTime  int
}
