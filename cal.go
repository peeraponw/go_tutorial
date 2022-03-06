package cal

func sumOfFirst(n int) int {
	var s int = 0
	for i := 0; i <= n; i++ {
		s += i
	}
	return s

}
