package kata

func add(a, b int) int      { return a + b }
func subtract(a, b int) int { return a - b }
func multiply(a, b int) int { return a * b }
func divide(a, b int) int   { return a / b }

func Arithmetic(a int, b int, operator string) int {
	resultsMap := map[string]func(a, b int) int{
		"add":      add,
		"subtract": subtract,
		"multiply": multiply,
		"divide":   divide,
	}
	return resultsMap[operator](a, b)
}
