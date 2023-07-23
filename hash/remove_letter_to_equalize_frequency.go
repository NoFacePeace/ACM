package hash

func equalFrequency(word string) bool {
	m := map[byte]int{}
	for i := 0; i < len(word); i++ {
		m[word[i]]++
	}
	for k, v := range m {
		cnt := v - 1
		status := true
		for k1, v1 := range m {
			if k1 == k {
				continue
			}
			if cnt == 0 {
				cnt = v1
			}
			if v1 != cnt {
				status = false
				break
			}
		}
		if status {
			return true
		}
	}
	return false
}
