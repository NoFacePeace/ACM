package hash

func maxRepOpt1(text string) int {
	m := map[byte]int{}
	s := []byte{}
	arr := []int{}
	cnt := 0
	var c byte
	for i := 0; i < len(text); i++ {
		m[text[i]]++
		if cnt == 0 {
			c = text[i]
			cnt++
			s = append(s, c)
			continue
		}
		if c == text[i] {
			cnt++
			continue
		}
		arr = append(arr, cnt)
		c = text[i]
		s = append(s, c)
		cnt = 1
	}
	arr = append(arr, cnt)
	max := 0
	for i := 0; i < len(s); i++ {
		if arr[i] > max {
			max = arr[i]
		}
		if i-1 >= 0 && m[s[i]] > arr[i] {
			if arr[i]+1 > max {
				max = arr[i] + 1
			}

		}
		if i+1 >= len(s) {
			continue
		}
		if arr[i+1] > 1 || i+2 >= len(s) || s[i] != s[i+2] {
			if m[s[i]] == arr[i] {
				continue
			}
			cnt := arr[i] + 1
			if cnt > max {
				max = cnt
			}
			continue
		}
		cnt := arr[i] + arr[i+2]
		if m[s[i]] > cnt {
			cnt += 1
		}
		if cnt > max {
			max = cnt
		}
	}
	return max
}
