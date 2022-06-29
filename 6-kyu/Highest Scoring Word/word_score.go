package kata

import "strings"

func High(s string) string {
	var diffBetweenLetterAndScore int32 = 96
	stringSlice := strings.Split(s, " ")
	letterScores := make([]struct {
		word    string
		index   int
		wordSum int32
	}, len(stringSlice))
	var maxScore int32
	var result string

	for i := range letterScores {
		letterScores[i].index = i
		letterScores[i].word = stringSlice[i]

		for _, v := range letterScores[i].word {
			letterScores[i].wordSum += v - diffBetweenLetterAndScore
		}
		if maxScore < letterScores[i].wordSum {
			maxScore = letterScores[i].wordSum
		}
	}
	for _, v := range letterScores {
		if v.wordSum == maxScore {
			result = v.word
			break
		}
	}

	return result
}
