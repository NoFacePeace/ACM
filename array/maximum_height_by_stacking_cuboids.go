package array

import "sort"

func maxHeight(cuboids [][]int) int {
	l := len(cuboids)
	for i := 0; i < l; i++ {
		sort.Slice(cuboids[i], func(a, b int) bool {
			return cuboids[i][a] < cuboids[i][b]
		})
	}
	sort.Slice(cuboids, func(i, j int) bool {
		if cuboids[i][2] < cuboids[j][2] {
			return true
		}
		if cuboids[i][2] > cuboids[j][2] {
			return false
		}
		if cuboids[i][1] < cuboids[j][1] {
			return true
		}
		if cuboids[i][1] > cuboids[j][1] {
			return false
		}
		if cuboids[i][0] <= cuboids[j][0] {
			return true
		}
		return false
	})
	ans := 0
	stack := [][]int{}
	var dfs func(i, height int)
	dfs = func(i, height int) {
		if i < 0 {
			if height > ans {
				ans = height
			}
			return
		}
		if len(stack) == 0 {
			stack = append(stack, cuboids[i])
			dfs(i-1, height+cuboids[i][2])
			stack = stack[:len(stack)-1]
			dfs(i-1, height)
			return
		}
		if cuboids[i][2]*(i+1)+height < ans {
			return
		}
		s := stack[len(stack)-1]
		if cuboids[i][0] <= s[0] && cuboids[i][1] <= s[1] && cuboids[i][2] <= s[2] {
			stack = append(stack, cuboids[i])
			dfs(i-1, height+cuboids[i][2])
			stack = stack[:len(stack)-1]
			dfs(i-1, height)
			return
		}
		dfs(i-1, height)
	}
	dfs(l-1, 0)
	return ans
}
