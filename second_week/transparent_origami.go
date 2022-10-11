package algorithmsunimi

import "fmt"

type position struct {
	x, y int
}

type instruction struct {
	coordinates []position
	foldAlong   position
}

func transparentOrigami() {
	instruct := instruction{
		coordinates: []position{
			{6, 10}, {0, 14}, {9, 10}, {0, 3}, {10, 4}, {4, 11}, {6, 0}, {6, 12}, {4, 1},
			{0, 13}, {10, 12}, {3, 4}, {3, 0}, {8, 4}, {1, 10}, {2, 14}, {8, 10}, {9, 0},
		},
		foldAlong: position{5, 7},
	}
	maxY := 14
	maxX := 10
	distinctCoordinates := map[position]bool{}
	for _, coord := range instruct.coordinates {
		if coord.y > instruct.foldAlong.y {
			coord.y = maxY - coord.y //"folding" along Y (rotating along the Y axis)
		}
		if coord.x > instruct.foldAlong.x {
			coord.x = maxX - coord.x //"folding" along X (rotating along the X axis)
		}
		distinctCoordinates[coord] = true //after coordinates re-assignment we have to remove unecessary copies
	}
	for i := 0; i < instruct.foldAlong.y; i++ {
		for j := 0; j < instruct.foldAlong.x; j++ {
			if distinctCoordinates[position{x: i, y: j}] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}
