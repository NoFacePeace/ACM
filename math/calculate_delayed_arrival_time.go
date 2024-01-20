package math

func findDelayedArrivalTime(arrivalTime int, delayedTime int) int {
	t := arrivalTime + delayedTime
	return t % 24
}
