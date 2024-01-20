package greedy

func distMoney(money int, children int) int {
	if money < children {
		return -1
	}
	if money == children {
		return 0
	}
	money -= children
	cnt := 0
	for money >= 7 && cnt < children {
		money -= 7
		cnt++
	}
	if money == 3 && children-cnt == 1 {
		return cnt - 1
	}
	if money > 0 && cnt == children {
		return cnt - 1
	}
	return cnt
}
