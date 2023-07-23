package multiply

func findLongestSubarray(array []string) []string {
	l := len(array)
	m := map[int]int{}
	cnt := 0
	start, end := -1, -1
	for i := 0; i < l; i++ {
		if array[i] >= "a" && array[i] <= "z" {
			cnt += 1
		} else if array[i] >= "A" && array[i] <= "Z" {
			cnt += 1
		} else {
			cnt -= 1
		}
		if cnt == 0 {
			start = -1
			end = i + 1
			continue
		}
		if val, ok := m[cnt]; ok {
			if i+1-val > end-start {
				end = i + 1
				start = val
			}
		} else {
			m[cnt] = i
		}
	}
	if start == -1 && end == -1 {
		return []string{}
	}
	return array[start+1 : end]
}
