package str

import (
	"sort"
	"strconv"
	"strings"
)

func alertNames(keyName []string, keyTime []string) []string {
	m := map[string][]string{}
	for i := 0; i < len(keyName); i++ {
		m[keyName[i]] = append(m[keyName[i]], keyTime[i])
	}
	ans := []string{}
	for _, v := range m {
		sort.Slice(v, func(a, b int) bool {
			arr1 := strings.Split(v[a], ":")
			arr2 := strings.Split(v[b], ":")
			h1, _ := strconv.Atoi(arr1[0])
			m1, _ := strconv.Atoi(arr1[1])
			h2, _ := strconv.Atoi(arr2[0])
			m2, _ := strconv.Atoi(arr2[1])
			if h1 < h2 {
				return true
			}
			if h1 == h2 {
				return m1 < m2
			}
			return false
		})
	}
	inHour := func(s1, s2 string) bool {
		arr1 := strings.Split(s1, ":")
		arr2 := strings.Split(s2, ":")
		h1, _ := strconv.Atoi(arr1[0])
		m1, _ := strconv.Atoi(arr1[1])
		h2, _ := strconv.Atoi(arr2[0])
		m2, _ := strconv.Atoi(arr2[1])
		if h1 > h2 {
			return false
		}
		if h1+2 <= h2 {
			return false
		}
		if h1 == h2 {
			return true
		}
		if 60-m1+m2 <= 60 {
			return true
		}
		return false
	}
	for k, v := range m {
		for i := 0; i < len(v)-2; i++ {
			if inHour(v[i], v[i+2]) {
				ans = append(ans, k)
				break
			}
		}
	}
	sort.Slice(ans, func(a, b int) bool {
		return ans[a] < ans[b]
	})
	return ans
}
