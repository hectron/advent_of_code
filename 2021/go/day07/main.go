package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func CalculateCheapestCostAndRoute(positions []int) (int, int) {
	bestPosition := 0
	bestCost := 0

	for i := 0; i < len(positions); i++ {
		cost := CostOfMovingTo(positions, i)

		if bestCost == 0 || cost < bestCost {
			bestPosition = i
			bestCost = cost
		}
	}

	return bestCost, bestPosition
}

func CostOfMovingTo(positions []int, position int) int {
	cost := 0

	for j := 1; j <= len(positions); j++ {
		difference := positions[j-1] - position

		if difference < 0 {
			difference *= -1
		}

		//cost += difference
		for i := 1; i <= difference; i++ {
			cost += i
		}
	}

	return cost
}

func loadInput() ([]int, error) {
	var (
		result []int
		err    error
	)

	filename := "input.txt"
	data, err := os.ReadFile(filename)

	if err != nil {
		return result, err
	}

	dataInParts := strings.Split(string(data), ",")

	for _, strNum := range dataInParts {
		cleanStrNum := strings.TrimSpace(strNum)

		if cleanStrNum == "" {
			continue
		}

		num, err := strconv.Atoi(cleanStrNum)

		if err != nil {
			return result, err
		}

		result = append(result, num)
	}

	return result, err
}

func main() {
	input, err := loadInput()

	if err != nil {
		panic(err)
	}

	cost, position := CalculateCheapestCostAndRoute(input)

	fmt.Printf("Moving all the input to position %d is the cheapest route -- cost: %d\n", position, cost)
}
