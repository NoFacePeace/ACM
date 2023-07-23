package dp

func minHeightShelves(books [][]int, shelfWidth int) int {
	l := len(books)
	dp := make([]int, l+1)
	for i := 0; i < l; i++ {
		mh, cw := 0, 0
		for j := i; j >= 0; j-- {
			cw += books[j][0]
			if cw > shelfWidth {
				break
			}
			if books[j][1] > mh {
				mh = books[j][1]
			}
			if dp[i+1] == 0 {
				dp[i+1] = dp[j] + mh
			}
			if dp[j]+mh < dp[i+1] {
				dp[i+1] = dp[j] + mh
			}
		}
	}
	return dp[l]
}
