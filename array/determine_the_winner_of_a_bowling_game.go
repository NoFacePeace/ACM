package array

func isWinner(player1 []int, player2 []int) int {
	get := func(arr []int) int {
		score := 0
		cnt := 2
		for _, v := range arr {
			if cnt < 2 {
				score += 2 * v
			} else {
				score += v
			}
			if v == 10 {
				cnt = 0
			} else {
				cnt++
			}
		}
		return score
	}
	score1 := get(player1)
	score2 := get(player2)
	if score1 == score2 {
		return 0
	}
	if score1 > score2 {
		return 1
	}
	return 2
}
