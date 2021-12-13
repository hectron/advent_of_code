package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type Point struct {
	X, Y int
}

type Movement struct {
	From, To Point
}

func (m *Movement) IsStraightLine() bool {
	if m.From.X == m.To.X && m.From.Y != m.To.Y {
		return true
	}

	if m.From.Y == m.To.Y && m.From.X != m.To.X {
		return true
	}

	return false
}

func (m *Movement) FindLineSegments() []Point {
	if m.From.X != m.To.X {
		distance := m.From.X - m.To.X

		difference := absoluteInt(distance) + 1
		points := make([]Point, difference)

		for i := 0; i < difference; i++ {
			if distance < 0 {
				points[i] = Point{m.From.X + i, m.From.Y}
			} else {
				points[i] = Point{m.From.X - i, m.From.Y}
			}
		}

		return points
	} else if m.From.Y != m.To.Y {
		distance := m.From.Y - m.To.Y

		difference := absoluteInt(distance) + 1
		points := make([]Point, difference)

		for i := 0; i < difference; i++ {
			if distance < 0 {
				points[i] = Point{m.From.X, m.From.Y + i}
			} else {
				points[i] = Point{m.From.X, m.From.Y - i}
			}
		}

		return points
	}

	return []Point{
		{m.From.X, m.From.Y},
	}
}

func GenerateDiagram(movements []Movement) map[Point]int {
	diagram := make(map[Point]int)

	for _, movement := range movements {
		if !movement.IsStraightLine() {
			continue
		}

		points := movement.FindLineSegments()

		for _, point := range points {
			diagram[point] += 1
		}
	}

	return diagram
}

func PrintDiagram(diagram map[Point]int) {
	maxX := 0
	maxY := 0

	for point := range diagram {
		if point.X > maxX {
			maxX = point.X
		}
		if point.Y > maxY {
			maxY = point.Y
		}
	}

	renderableDiagram := make([][]string, maxY+1)

	fmt.Printf("%d x %d diagram\n", maxX, maxY)

	for y := 0; y <= maxY; y++ {
		row := make([]string, maxX+1)

		for x := 0; x <= maxX; x++ {
			row[x] = "."
		}

		renderableDiagram[y] = row
	}

	for point, count := range diagram {
		renderableDiagram[point.Y][point.X] = fmt.Sprintf("%d", count)
	}

	for _, row := range renderableDiagram {
		fmt.Println(row)
	}
}

func FindMostVisitedPoints(diagram map[Point]int) ([]Point, int) {
	max := 0
	occurences := []Point{}

	for point, count := range diagram {
		if count > max {
			max = count
			occurences = []Point{point}
		} else if count == max {
			occurences = append(occurences, point)
		}
	}

	return occurences, max
}

func ParseInput(input []string) []Movement {
	movements := []Movement{}

	for _, line := range input {
		portions := strings.FieldsFunc(line, IsNumber)

		if len(portions) == 4 {
			fromX, err := strconv.Atoi(portions[0])

			if err != nil {
				fmt.Errorf("Unable to parse line input %s - error: %w", line, err)
				continue
			}

			fromY, err := strconv.Atoi(portions[1])

			if err != nil {
				fmt.Errorf("Unable to parse line input %s - error: %w", line, err)
				continue
			}

			toX, err := strconv.Atoi(portions[2])

			if err != nil {
				fmt.Errorf("Unable to parse line input %s - error: %w", line, err)
				continue
			}

			toY, err := strconv.Atoi(portions[3])

			if err != nil {
				fmt.Errorf("Unable to parse line input %s - error: %w", line, err)
				continue
			}

			movements = append(movements, Movement{
				From: Point{fromX, fromY},
				To:   Point{toX, toY},
			})
		}
	}

	return movements
}

func IsNumber(r rune) bool {
	return !unicode.IsNumber(r)
}

func absoluteInt(i int) int {
	if i < 0 {
		return -1 * i
	}
	return i
}

func loadInput() ([]string, error) {
	var (
		input []string
		err   error
	)
	filename := "input.txt"

	data, err := os.ReadFile(filename)

	if err != nil {
		fmt.Println("Unable to read file", filename)
		return input, err
	}

	dataAsString := string(data)
	input = strings.Split(dataAsString, "\n")

	lines := input[:len(input)-1]

	return lines, err
}

func main() {
	input, err := loadInput()

	if err != nil {
		panic(err)
	}

	movements := ParseInput(input)
	diagram := GenerateDiagram(movements)

	mostVisitedPoints, numberOfTimes := FindMostVisitedPoints(diagram)

	fmt.Printf("Found points which were the most visited (a total of %d times): %v\n", numberOfTimes, mostVisitedPoints)
}
