package dp

func sumSubarrayMins(arr []int) (ans int) {
	const mod int = 1e9 + 7
	n := len(arr)
	monoStack := []int{}
	dp := make([]int, n)
	for i, x := range arr {
		for len(monoStack) > 0 && arr[monoStack[len(monoStack)-1]] > x {
			monoStack = monoStack[:len(monoStack)-1]
		}
		k := i + 1
		if len(monoStack) > 0 {
			k = i - monoStack[len(monoStack)-1]
		}
		dp[i] = k * x
		if len(monoStack) > 0 {
			dp[i] += dp[i-k]
		}
		ans = (ans + dp[i]) % mod
		monoStack = append(monoStack, i)
	}
	return
}
