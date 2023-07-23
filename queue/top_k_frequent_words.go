package queue

func topKFrequent(words []string, k int) []string {
	m := map[string]int{}
	for i := 0; i < len(words); i++ {
		m[words[i]]++
	}
	type pair struct {
		s string
		c int
	}
	arr := []pair{}
	for k, v := range m {
		arr = append(arr, pair{k, v})
	}
	tmp := []pair{}
	for _, v := range arr {
		if len(tmp) < k {
			tmp = append(tmp, v)
			for i := len(tmp) - 1; i > 0; i-- {
				if tmp[i].c < tmp[i-1].c {
					break
				}
				if tmp[i].c == tmp[i-1].c && tmp[i].s > tmp[i-1].s {
					break
				}
				tmp[i], tmp[i-1] = tmp[i-1], tmp[i]
			}
			continue
		}
		if tmp[len(tmp)-1].c > v.c {
			continue
		}
		if tmp[len(tmp)-1].c == v.c && tmp[len(tmp)-1].s < v.s {
			continue
		}
		tmp[len(tmp)-1] = v
		for i := len(tmp) - 1; i > 0; i-- {
			if tmp[i].c < tmp[i-1].c {
				break
			}
			if tmp[i].c == tmp[i-1].c && tmp[i].s > tmp[i-1].s {
				break
			}
			tmp[i], tmp[i-1] = tmp[i-1], tmp[i]
		}
	}
	ans := []string{}
	for _, v := range tmp {
		ans = append(ans, v.s)
	}
	return ans
}
