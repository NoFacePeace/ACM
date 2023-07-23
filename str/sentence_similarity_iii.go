package str

import "strings"

func areSentencesSimilar(sentence1 string, sentence2 string) bool {
	if len(sentence1) < len(sentence2) {
		sentence1, sentence2 = sentence2, sentence1
	}
	arr1 := strings.Split(sentence1, " ")
	arr2 := strings.Split(sentence2, " ")
	left := 0
	for ; left < len(arr2); left++ {
		if arr1[left] != arr2[left] {
			break
		}
	}
	if left == len(arr2) {
		return true
	}
	right := 1
	for ; len(arr2)-right >= left; right++ {
		if arr1[len(arr1)-right] != arr2[len(arr2)-right] {
			break
		}
	}
	if left+right-1 == len(arr2) {
		return true
	}
	return false
}
