package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Movement struct {
	Direction string
	Unit      int
}

func (m *Movement) SignedUnit() int {
	if m.Direction == "up" {
		return m.Unit * -1
	}

	return m.Unit
}

type Position struct {
	Horizontal, Depth int
}

func GetFinalPosition(movements []Movement) Position {
	finalPosition := Position{0, 0}

	for _, movement := range movements {
		if movement.Direction == "forward" {
			finalPosition.Horizontal += movement.Unit
		} else {
			finalPosition.Depth += movement.SignedUnit()
		}
	}

	return finalPosition
}

func loadInput() ([]Movement, error) {
	var (
		movements []Movement
		err       error
	)

	data, err := os.ReadFile("input.txt")

	if err != nil {
		fmt.Println("Unable to read input.txt")
		return movements, err
	}

	dataAsString := string(data)
	inputAsStrings := strings.Split(dataAsString, "\n")
	movements = make([]Movement, len(inputAsStrings)-1)

	for idx, m := range inputAsStrings {
		cleanString := strings.TrimSpace(m)

		if cleanString != "" {
			lineInParts := strings.Split(cleanString, " ")

			if len(lineInParts) != 2 {
				msg := fmt.Sprintf("Expected to get a line with a direction and number, but got %s", cleanString)
				err = errors.New(msg)

				return movements, err
			}

			unit, err := strconv.Atoi(lineInParts[1])

			if err != nil {
				return movements, err
			}

			movements[idx] = Movement{lineInParts[0], unit}
		}
	}

	return movements, err
}

func main() {
	movements, err := loadInput()

	if err != nil {
		panic(err)
	}

	position := GetFinalPosition(movements)

	msg := fmt.Sprintf("Final position is a depth of %d and a horizonal position of %d", position.Depth, position.Horizontal)
	fmt.Println(msg)

	fmt.Printf("Final depth * position = %d\n", (position.Depth * position.Horizontal))
}
