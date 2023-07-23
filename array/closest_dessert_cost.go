package array

func closestCost(baseCosts []int, toppingCosts []int, target int) int {
	x := baseCosts[0]
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	for _, c := range baseCosts {
		x = min(x, c)
	}
	if x >= target {
		return x
	}
	can := make([]bool, target+1)
	ans := 2*target - x
	for _, c := range baseCosts {
		if c <= target {
			can[c] = true
			continue
		}
		ans = min(c, ans)
	}
	for _, c := range toppingCosts {
		for cnt := 0; cnt < 2; cnt++ {
			for i := target; i > 0; i-- {
				if can[i] && i+c > target {
					ans = min(ans, i+c)
				}
				if i-c > 0 {
					can[i] = can[i] || can[i-c]
				}
			}
		}
	}
	for i := 0; i <= ans-target; i++ {
		if can[target-i] {
			return target - i
		}
	}
	return ans
}
