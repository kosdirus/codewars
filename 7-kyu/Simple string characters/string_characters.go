package kata

import "unicode"

func Solve(s string) []int {
	result := []int{0, 0, 0, 0}
	for _, v := range s {
		switch {
		case unicode.IsUpper(v):
			result[0] += 1
		case unicode.IsLower(v):
			result[1] += 1
		case unicode.IsNumber(v):
			result[2] += 1
		default:
			result[3] += 1
		}
	}
	return result
}
