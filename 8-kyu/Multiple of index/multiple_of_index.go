package kata

func multipleOfIndex(ints []int) []int {
	var result []int
	for i, v := range ints {
		if i != 0 {
			if v%i == 0 {
				result = append(result, v)
			}
		}
	}
	return result
}
