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

func TestLanternFishes(t *testing.T) {
	fishes := lanthernFishes(18)
	if len(fishes) != 26 {
		t.Fatalf("Expected 26 fishes after 18 days, got %d.", len(fishes))
	}
	expectedFishAges := []int{6, 0, 6, 4, 5, 6, 0, 1, 1, 2, 6, 0, 1, 1, 1, 2, 2, 3, 3, 4, 6, 7, 8, 8, 8, 8}
	for i := 0; i < len(expectedFishAges); i++ {
		if fishes[i] != expectedFishAges[i] {
			t.Fatalf("Expected fish numer %d to be %d days old, got %d days of age.", i, expectedFishAges[i], fishes[i])
		}
	}
	fishes = lanthernFishes(80)
	if len(fishes) != 5934 {
		t.Fatalf("Expecte 5934 fishes, got %d.", len(fishes))
	}

}

func BenchmarkLanternFishes18(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lanthernFishes(18)
	}
}

func BenchmarkLanternFishes80(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lanthernFishes(80)
	}
}
