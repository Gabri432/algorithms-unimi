package algorithmsunimi

func lanthernFishes(days int) []int {
	fishAges := []int{3, 4, 3, 1, 2}
	for d := 0; d < days; d++ {
		for i := 0; i < len(fishAges); i++ {
			if fishAges[i] > 0 {
				fishAges[i]--
			} else if fishAges[i] == 0 {
				fishAges[i] = 6
				fishAges = append(fishAges, 9)
			}
		}
	}
	return fishAges
}
