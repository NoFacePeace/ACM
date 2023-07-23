package str

import (
	"sort"
	"strings"
)

func removeSubfolders(folder []string) []string {
	if len(folder) <= 0 {
		return folder
	}
	sort.Strings(folder)
	ans := []string{}
	ans = append(ans, folder[0])
	last := folder[0]
	match := func(s1, s2 string) bool {
		arr1 := strings.Split(s1, "/")
		arr2 := strings.Split(s2, "/")
		if len(arr1) >= len(arr2) {
			return false
		}
		for i := 0; i < len(arr1); i++ {
			if arr1[i] != arr2[i] {
				return false
			}
		}
		return true
	}
	for i := 1; i < len(folder); i++ {
		if match(last, folder[i]) {
			continue
		}
		ans = append(ans, folder[i])
		last = folder[i]
	}
	return ans
}
