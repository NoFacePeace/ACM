package math

func numOfBurgers(tomatoSlices int, cheeseSlices int) []int {

	if (tomatoSlices-2*cheeseSlices)%2 != 0 {
		return []int{}
	}
	x := (tomatoSlices - 2*cheeseSlices) / 2
	y := cheeseSlices - x
	if x < 0 {
		return []int{}
	}
	if y < 0 {
		return []int{}
	}
	return []int{x, y}
}
