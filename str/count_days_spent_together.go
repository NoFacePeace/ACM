package str

import (
	"strconv"
	"strings"
)

func countDaysTogether(arriveAlice string, leaveAlice string, arriveBob string, leaveBob string) int {
	var arrive, leave string
	if arriveAlice > arriveBob {
		arrive = arriveAlice
	} else {
		arrive = arriveBob
	}
	if leaveAlice > leaveBob {
		leave = leaveBob
	} else {
		leave = leaveAlice
	}
	if arrive > leave {
		return 0
	}
	arr := strings.Split(arrive, "-")
	aM, _ := strconv.Atoi(arr[0])
	aD, _ := strconv.Atoi(arr[1])
	arr = strings.Split(leave, "-")
	lM, _ := strconv.Atoi(arr[0])
	lD, _ := strconv.Atoi(arr[1])
	days := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	if aM == lM {
		return lD - aD + 1
	}
	d := days[aM-1] - aD + 1
	for i := aM + 1; i < lM; i++ {
		d += days[i-1]
	}
	d += lD
	return d
}
