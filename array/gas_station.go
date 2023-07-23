package array

func canCompleteCircuit(gas []int, cost []int) int {
	for i := 0; i < len(gas); {
		balance := 0
		cnt := 0
		for {
			mod := (i + cnt) % len(gas)
			balance += gas[mod]
			if balance < cost[mod] {
				break
			}
			balance -= cost[mod]
			cnt++
			if cnt == len(gas) {
				break
			}
		}
		if cnt == len(gas) {
			return i
		}
		i = i + cnt + 1
	}
	return -1
}
