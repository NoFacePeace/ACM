package set

import . "github.com/NoFacePeace/ACM/packages"

func minSwapsCouples(row []int) int {
	u := NewUnionFind(len(row) / 2)
	for i := 1; i < len(row); i += 2 {
		u.Union(row[i-1]/2, row[i]/2)
	}
	return len(row)/2 - u.GetCount()
}
