package kata

import (
	"strings"
)

func ToCamelCase(s string) string {
	var res string
	strSlice := strings.Split(s, "-")
	if len(strSlice) <= 1 {
		strSlice = strings.Split(s, "_")
	}

	for i := range strSlice {
		if i == 0 {
			res += strSlice[i]
			continue
		}
		strSliceRunes := []rune(strSlice[i])
		if len(strSliceRunes) != 0 {
			strSliceRunes[0] = []rune(strings.ToUpper(string(strSliceRunes)))[0]
		}

		res += string(strSliceRunes)
	}

	return res
}

// ToCamelCase1 Another way to solve this kata
func ToCamelCase1(s string) string {
	var res string
	strSlice := strings.Split(s, "-")
	if len(strSlice) <= 1 {
		strSlice = strings.Split(s, "_")
	}

	for i := range strSlice {
		if i == 0 {
			res += strSlice[i]
			continue
		}

		if firstSymbol := strSlice[i][0]; firstSymbol >= 65 && firstSymbol <= 90 {
			res += strSlice[i]
		} else if firstSymbol >= 97 && firstSymbol <= 122 {
			res += string(firstSymbol-32) + strSlice[i][1:]
		}
	}
	return res
}
