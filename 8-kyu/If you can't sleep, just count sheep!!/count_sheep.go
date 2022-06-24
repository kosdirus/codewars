package kata

import "fmt"

func countSheep(num int) (result string) {
	for i := 1; i <= num; i++ {
		result = fmt.Sprintf("%s%d sheep...", result, i)
	}
	return result
}
