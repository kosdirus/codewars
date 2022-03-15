package kata

func QuarterOf(month int) int {
	if month%3 == 0 {
		return month / 3
	}
	return month/3 + 1
}
