package kata

func RepeatStr(repetitions int, value string) string {
	var res string
	for ; repetitions > 0; repetitions-- {
		res = res + value
	}
	return res
}
