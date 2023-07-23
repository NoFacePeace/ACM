package array

func minNumberOfHours(initialEnergy int, initialExperience int, energy []int, experience []int) int {
	hours := 0
	for _, v := range energy {
		hours += v
	}
	if hours < initialEnergy {
		hours = 0
	} else {
		hours = hours - initialEnergy + 1
	}
	for _, v := range experience {
		if initialExperience > v {
			initialExperience += v
			continue
		}
		hours = hours + v - initialExperience + 1
		initialExperience = initialExperience + v - initialExperience + 1 + v
	}
	if hours < 0 {
		return 0
	}
	return hours
}
