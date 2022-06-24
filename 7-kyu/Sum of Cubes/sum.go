package kata

func SumCubes(n int) int {
	var result int
	for i := 1; i <= n; i++ {
		result += i * i * i
	}
	return result
}
