package array

func numTilings(n int) int {
	first := 0
	second := 0
	third := 0
	fourth := 1
	mod := int(1e9 + 7)
	for i := 1; i <= n; i++ {
		f := fourth % mod
		s := (first + third) % mod
		t := (first + second) % mod
		fo := (first + second + third + fourth) % mod
		first, second, third, fourth = f, s, t, fo
	}
	return fourth
}
