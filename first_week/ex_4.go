package algorithmsunimi

func ex4() int {
	myList := []int{4, 5, 1, 3, 6, 2}
	countToLeft := 0
	for i := 1; i < len(myList); i++ {
		for j := 0; j < len(myList); j++ {
			if myList[j] == i+1 && i < j {
				countToLeft++
			}
		}
	}
	return countToLeft
}
