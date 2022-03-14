package kata

func CountSheeps(numbers []bool) int {
	var n int
	for _, v := range numbers {
		if v == true {
			n++
		}
	}
	return n
}
