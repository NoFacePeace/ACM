package hash

import (
	"sort"
	"strings"
)

func topStudents(positive_feedback []string, negative_feedback []string, report []string, student_id []int, k int) []int {
	p := map[string]bool{}
	n := map[string]bool{}
	for _, v := range positive_feedback {
		p[v] = true
	}
	for _, v := range negative_feedback {
		n[v] = true
	}
	m := map[int]int{}
	for i := 0; i < len(report); i++ {
		arr := strings.Split(report[i], " ")
		m[student_id[i]] = 0
		for _, v := range arr {
			if p[v] {
				m[student_id[i]] += 3
			}
			if n[v] {
				m[student_id[i]] -= 1
			}
		}
	}
	arr := [][2]int{}
	for k, v := range m {
		arr = append(arr, [2]int{k, v})
	}
	sort.Slice(arr, func(a, b int) bool {
		if arr[a][1] == arr[b][1] {
			return arr[a][0] < arr[b][0]
		}
		return arr[a][1] > arr[b][1]
	})
	ans := []int{}
	for i := 0; i < k; i++ {
		if i >= len(arr) {
			break
		}
		ans = append(ans, arr[i][0])
	}
	return ans
}
