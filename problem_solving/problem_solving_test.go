package algorithmsunimi

import "testing"

func TestEx3(t *testing.T) {
	intList := ex3()
	expectedResult := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	for i := len(expectedResult); i > 0; i-- {
		if intList[len(expectedResult)-i] != i {
			t.Fatalf("Expected %d, got %d", i, intList[i])
		}
	}
}

func TestEx4(t *testing.T) {
	timesToTheLeft := ex4()
	if timesToTheLeft != 2 {
		t.Fatalf("Expected 2, got %d", timesToTheLeft)
	}
}
