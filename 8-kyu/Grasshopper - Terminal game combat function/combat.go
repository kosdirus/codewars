package kata

func combat(health, damage float64) (res float64) {
	if health-damage < 0 {
		return 0
	}
	return health - damage
}
