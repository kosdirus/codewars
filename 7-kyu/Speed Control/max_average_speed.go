package kata

func Gps(s int, x []float64) int {
	var maxSpeed float64
	if len(x) <= 1 {
		return 0
	}
	for i := 0; i < len(x)-1; i++ {
		speed := (3600 * (x[i+1] - x[i])) / float64(s)
		if speed > maxSpeed {
			maxSpeed = speed
		}
	}
	return int(maxSpeed)
}
