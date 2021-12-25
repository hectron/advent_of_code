package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Coordinate struct {
	X, Y int
}

type Basin struct {
	Coordinates []Coordinate
}

func (b *Basin) Size() int {
	return len(b.Coordinates)
}

func FindLowPointCoordinates(heatmap [][]int) []Coordinate {
	coordinates := []Coordinate{}

	maxY := len(heatmap)
	maxX := len(heatmap[0])

	for y := 0; y < maxY; y++ {
		for x := 0; x < maxX; x++ {
			currentVal := heatmap[y][x]

			neighborVals := []int{}

			if y == 0 {
				if x == 0 {
					neighborVals = []int{
						heatmap[y][x+1],
						heatmap[y+1][x],
					}
				} else if x < maxX-1 {
					neighborVals = []int{
						heatmap[y][x-1],
						heatmap[y+1][x],
						heatmap[y][x+1],
					}
				} else {
					neighborVals = []int{
						heatmap[y][x-1],
						heatmap[y+1][x],
					}
				}
			} else if y < maxY-1 {
				if x == 0 {
					neighborVals = []int{
						heatmap[y-1][x],
						heatmap[y][x+1],
						heatmap[y+1][x],
					}
				} else if x < maxX-1 {
					neighborVals = []int{
						heatmap[y][x-1],
						heatmap[y-1][x],
						heatmap[y+1][x],
						heatmap[y][x+1],
					}
				} else {
					neighborVals = []int{
						heatmap[y][x-1],
						heatmap[y-1][x],
						heatmap[y+1][x],
					}
				}
			} else {
				if x == 0 {
					neighborVals = []int{
						heatmap[y-1][x],
						heatmap[y][x+1],
					}
				} else if x < maxX-1 {
					neighborVals = []int{
						heatmap[y][x-1],
						heatmap[y-1][x],
						heatmap[y][x+1],
					}
				} else {
					neighborVals = []int{
						heatmap[y][x-1],
						heatmap[y-1][x],
					}
				}
			}

			lowPoint := true

			for _, n := range neighborVals {
				if n <= currentVal {
					lowPoint = false
					break
				}
			}

			if lowPoint {
				coordinates = append(coordinates, Coordinate{X: x, Y: y})
			}
		}
	}

	return coordinates
}

func CalculateRiskLevel(heatmap [][]int, coordinates []Coordinate) int {
	sum := 0

	for _, coordinate := range coordinates {
		point := heatmap[coordinate.Y][coordinate.X]
		sum += (point + 1)
	}

	return sum
}

func loadInput() ([][]int, error) {
	var (
		result [][]int
		err    error
	)

	filename := "input.txt"
	data, err := os.ReadFile(filename)

	if err != nil {
		return result, err
	}

	dataInParts := strings.Split(string(data), "\n")

	result = make([][]int, len(dataInParts)-1)

	for idx, line := range dataInParts {
		lineOfNumbers := []int{}

		for _, c := range line {
			num, err := strconv.Atoi(string(c))

			if err != nil {
				return result, err
			}

			lineOfNumbers = append(lineOfNumbers, num)
		}

		if len(lineOfNumbers) > 0 {
			result[idx] = lineOfNumbers
		}
	}

	return result, err
}

func FindBasins(heatmap [][]int, lowPoints []Coordinate) []Basin {
	basins := []Basin{}

	for _, lowPoint := range lowPoints {
		neighbors := FindBiggerNeighbors(heatmap, lowPoint)

		uniqueCoords := []Coordinate{}
		keys := map[Coordinate]bool{}

		for _, coord := range neighbors {
			if _, exists := keys[coord]; !exists {
				keys[coord] = true
				uniqueCoords = append(uniqueCoords, coord)
			}
		}

		SortCoordinates(uniqueCoords)

		basins = append(basins, Basin{Coordinates: uniqueCoords})
	}

	return basins
}

func SortCoordinates(coords []Coordinate) {
	sort.SliceStable(coords, func(i, j int) bool {
		return coords[i].X <= coords[j].X && coords[i].Y <= coords[j].Y
	})
}

func FindBiggerNeighbors(heatmap [][]int, coord Coordinate) []Coordinate {
	coordinates := []Coordinate{}

	currentValue := heatmap[coord.Y][coord.X]
	traverseUp := coord.Y > 0
	traverseDown := coord.Y < len(heatmap)-1
	traverseLeft := coord.X > 0
	traverseRight := coord.X < len(heatmap[0])-1

	if traverseUp {
		neighbor := heatmap[coord.Y-1][coord.X]

		if neighbor > currentValue && neighbor != 9 {
			coordinates = append(coordinates, FindBiggerNeighbors(heatmap, Coordinate{X: coord.X, Y: coord.Y - 1})...)
		}
	}

	if traverseDown {
		neighbor := heatmap[coord.Y+1][coord.X]

		if neighbor > currentValue && neighbor != 9 {
			coordinates = append(coordinates, FindBiggerNeighbors(heatmap, Coordinate{X: coord.X, Y: coord.Y + 1})...)
		}
	}

	if traverseLeft {
		neighbor := heatmap[coord.Y][coord.X-1]

		if neighbor > currentValue && neighbor != 9 {
			coordinates = append(coordinates, FindBiggerNeighbors(heatmap, Coordinate{X: coord.X - 1, Y: coord.Y})...)
		}
	}

	if traverseRight {
		neighbor := heatmap[coord.Y][coord.X+1]

		if neighbor > currentValue && neighbor != 9 {
			coordinates = append(coordinates, FindBiggerNeighbors(heatmap, Coordinate{X: coord.X + 1, Y: coord.Y})...)
		}
	}

	coordinates = append(coordinates, coord)

	return coordinates
}

func main() {
	heatmap, err := loadInput()

	if err != nil {
		panic(err)
	}

	lowPoints := FindLowPointCoordinates(heatmap)
	riskLevel := CalculateRiskLevel(heatmap, lowPoints)
	fmt.Printf("The risk level is %d.\n", riskLevel)

	basins := FindBasins(heatmap, lowPoints)

	fmt.Printf("Found %d basins. Reducing to the largest 3\n", len(basins))

	basinSizes := make([]int, len(basins))

	for idx, basin := range basins {
		basinSizes[idx] = basin.Size()
	}

	sort.Ints(basinSizes)

	biggestBasinSizes := basinSizes[len(basinSizes)-3:]

	answer := 1
	for _, basinSize := range biggestBasinSizes {
		answer *= basinSize
	}

	fmt.Printf("Multiplication of the top three basins %d * %d * %d = %d\n", biggestBasinSizes[0], biggestBasinSizes[1], biggestBasinSizes[2], answer)
}
