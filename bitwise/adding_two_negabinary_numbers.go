package bitwise

func addNegabinary(arr1 []int, arr2 []int) []int {
	arr := []int{}
	i := len(arr1) - 1
	j := len(arr2) - 1
	for i >= 0 && j >= 0 {
		arr = append(arr, arr1[i]+arr2[j])
		i--
		j--
	}
	for i >= 0 {
		arr = append(arr, arr1[i])
		i--
	}
	for j >= 0 {
		arr = append(arr, arr2[j])
		j--
	}
	arr = append(arr, 0)
	arr = append(arr, 0)
	for i := 0; i < len(arr)-2; i++ {
		for arr[i] >= 2 {
			if arr[i+1] > 0 {
				arr[i] -= 2
				arr[i+1] -= 1
			} else {
				arr[i] -= 2
				arr[i+1] += 1
				arr[i+2] += 1
			}
		}
	}
	ans := []int{}
	for i := len(arr) - 1; i >= 0; i-- {
		ans = append(ans, arr[i])
	}
	for len(ans) > 0 {
		if ans[0] != 0 {
			break
		}
		ans = ans[1:]
	}
	if len(ans) == 0 {
		ans = append(ans, 0)
	}
	return ans
}
