package packages

type UnionFind struct {
	parents []int
	sizes   []int
	count   int
}

func NewUnionFind(n int) *UnionFind {
	parents := make([]int, n)
	sizes := make([]int, n)
	for i := 0; i < n; i++ {
		parents[i] = i
		sizes[i] = 1
	}
	return &UnionFind{
		parents: parents,
		sizes:   sizes,
		count:   n,
	}
}

func (u *UnionFind) Find(x int) int {
	if u.parents[x] == x {
		return x
	}
	u.parents[x] = u.Find(u.parents[x])
	return u.parents[x]
}

func (u *UnionFind) Union(x, y int) {
	xp, yp := u.Find(x), u.Find(y)
	if xp != yp {
		if u.sizes[xp] > u.sizes[yp] {
			u.parents[yp], u.sizes[xp] = xp, u.sizes[xp]+u.sizes[yp]
		} else {
			u.parents[xp], u.sizes[yp] = yp, u.sizes[xp]+u.sizes[yp]
		}
	}
}

func (u *UnionFind) GetSize(x int) int {
	return u.sizes[x]
}

func (u *UnionFind) GetCount() int {
	return u.count
}
