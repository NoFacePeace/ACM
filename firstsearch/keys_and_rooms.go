package firstsearch

func canVisitAllRooms(rooms [][]int) bool {
	visit := map[int]bool{}
	var f func(num int)
	f = func(num int) {
		if _, ok := visit[num]; ok {
			return
		}
		visit[num] = true
		for _, v := range rooms[num] {
			f(v)
		}
	}
	f(0)
	return len(visit) == len(rooms)
}
