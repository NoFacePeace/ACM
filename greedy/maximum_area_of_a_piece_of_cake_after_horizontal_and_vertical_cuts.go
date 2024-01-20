package greedy

import "sort"

func maxArea(h int, w int, horizontalCuts []int, verticalCuts []int) int {
	horizontalCuts = append(horizontalCuts, 0, h)
	verticalCuts = append(verticalCuts, 0, w)
	sort.Ints(horizontalCuts)
	sort.Ints(verticalCuts)
	mh, mw := 0, 0
	i := 0
	for _, v := range horizontalCuts {
		if v-i > mh {
			mh = v - i
		}
		i = v
	}
	i = 0
	for _, v := range verticalCuts {
		if v-i > mw {
			mw = v - i
		}
		i = v
	}
	mod := int(1e9) + 7
	return mh * mw % mod
}
