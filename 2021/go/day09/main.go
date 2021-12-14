package main

func FindLowPointsInArray(heatmap [][]int) []int {
	lowPoints := []int{}
	maxY := len(heatmap)
	maxX := len(heatmap[0])

	for y := 0; y < maxY; y++ {
		yIndices := []int{}

		if y == 0 {
			yIndices = []int{y, 1}
		} else if y+1 < maxY {
			yIndices = []int{y - 1, y, y + 1}
		} else {
			yIndices = []int{y - 1, y}
		}
		for x := 0; x < maxX; x++ {
			canSubtractFromX := x > 0

		}
	}
	return []int{}
}
