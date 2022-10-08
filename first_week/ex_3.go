package algorithmsunimi

func ex3() []int {
	myList := []int{6, 1, 5, 2, 3, 4, 7, 9, 8}
	for i := 0; i < len(myList); i++ {
		for j := 0; j < len(myList); j++ {
			if myList[i] > myList[j] {
				myList[i], myList[j] = myList[j], myList[i]
			}
		}
	}
	return myList
}
