package dp

func smallestSufficientTeam(req_skills []string, people [][]string) []int {
	n, m := len(req_skills), len(people)
	mp := map[string]int{}
	for i := 0; i < len(req_skills); i++ {
		mp[req_skills[i]] = i
	}
	dp := make([]int, 1<<n)
	for i := range dp {
		dp[i] = m
	}
	dp[0] = 0
	preSkill := make([]int, 1<<n)
	prePeople := make([]int, 1<<n)
	for i := 0; i < m; i++ {
		cur := 0
		for _, v := range people[i] {
			cur |= 1 << mp[v]
		}
		for prev := 0; prev < (1 << n); prev++ {
			comb := prev | cur
			if dp[comb] > dp[prev]+1 {
				dp[comb] = dp[prev] + 1
				preSkill[comb] = prev
				prePeople[comb] = i
			}
		}
	}
	res := []int{}
	i := (1 << n) - 1
	for i > 0 {
		res = append(res, prePeople[i])
		i = preSkill[i]
	}
	return res
}
