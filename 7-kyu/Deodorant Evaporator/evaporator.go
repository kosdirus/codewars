package kata

func Evaporator(content float64, evapPerDay int, threshold int) int {
	contentThreshold := content * float64(threshold) / 100
	var i int
	for ; content >= contentThreshold; i++ {
		content -= content * float64(evapPerDay) / 100
	}
	return i
}
