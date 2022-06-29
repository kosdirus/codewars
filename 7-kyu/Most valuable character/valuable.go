package kata

func Solve(s string) rune {
	runeSlice := []rune(s)
	uniqueRunes := make(map[rune][]int)
	for i, v := range runeSlice {
		uniqueRunes[v] = append(uniqueRunes[v], i)
	}
	firstLastIndexDiff := make(map[rune]int)
	biggestDiffSlice := make([]rune, 0)
	var biggestDiff int
	for k, v := range uniqueRunes {
		firstLastIndexDiff[k] = v[len(v)-1] - v[0]
		if v[len(v)-1]-v[0] > biggestDiff {
			biggestDiff = v[len(v)-1] - v[0]
			biggestDiffSlice = []rune{k}
		} else if v[len(v)-1]-v[0] == biggestDiff {
			biggestDiffSlice = append(biggestDiffSlice, k)
		}
	}
	res := biggestDiffSlice[0]
	for _, v := range biggestDiffSlice {
		if v < res {
			res = v
		}
	}
	return res
}
