package kata

func MultiplicationTable(size int) [][]int {
	table := make([][]int, size)
	for i := range table {
		table[i] = make([]int, size)

		for ii := range table[i] {
			if i == 0 {
				table[i][ii] = ii + 1
			} else if ii == 0 {
				table[i][ii] = i + 1
			}
		}
	}

	for i := range table {
		for ii := range table[i] {
			table[i][ii] = table[i][0] * table[0][ii]
		}
	}

	return table
}
