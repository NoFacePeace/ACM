package hash

import "sort"

func closeStrings(word1 string, word2 string) bool {
	n1 := len(word1)
	n2 := len(word2)
	if n1 != n2 {
		return false
	}
	m1 := map[byte]int{}
	for _, c := range word1 {
		m1[byte(c)]++
	}
	m2 := map[byte]int{}
	for _, c := range word2 {
		m2[byte(c)]++
	}
	if len(m1) != len(m2) {
		return false
	}
	arr1 := []int{}
	arr2 := []int{}
	for k, v := range m1 {
		if _, ok := m2[k]; !ok {
			return false
		}
		arr1 = append(arr1, v)
		arr2 = append(arr2, m2[k])
	}
	sort.Ints(arr1)
	sort.Ints(arr2)
	for i := 0; i < len(arr1); i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}
