package kata

func SmallestIntegerFinder(numbers []int) int {
	var smallest = numbers[0]
	for _, v := range numbers {
		if smallest > v {
			smallest = v
		}
	}
	return smallest
}
