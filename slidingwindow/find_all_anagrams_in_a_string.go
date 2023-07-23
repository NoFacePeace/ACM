package slidingwindow

func findAnagrams(s string, p string) []int {
	sl := len(s)
	pl := len(p)
	ans := []int{}
	if sl < pl {
		return ans
	}
	arr := make([]int, 26)
	for _, v := range p {
		arr[v-'a']++
	}
	arr1 := make([]int, 26)
	for i := 0; i < pl; i++ {
		arr1[s[i]-'a']++
	}
	equal := func(a, b []int) bool {
		for i := 0; i < 26; i++ {
			if a[i] != b[i] {
				return false
			}
		}
		return true
	}
	i := 0
	for i <= sl-pl {
		if equal(arr, arr1) {
			ans = append(ans, i)
		}
		arr1[s[i]-'a']--
		arr1[s[i+pl]-'a']++
		i++
	}
	return ans
}
