package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func CountRollingIncrements(numberSequence []int) int {
	count := 0

	for i := 1; i < len(numberSequence); i++ {
		if numberSequence[i] > numberSequence[i-1] {
			count += 1
		}
	}

	return count
}

func CountRollingIncrementsOfThree(numberSequence []int) int {
	count := 0
	lastSum := 0

	for i := 2; i < len(numberSequence); i++ {
		sum := numberSequence[i-2] + numberSequence[i-1] + numberSequence[i]

		if i > 2 && sum > lastSum {
			count += 1
		}

		lastSum = sum
	}

	return count
}

func loadInput() ([]int, error) {
	var (
		ints []int
		err  error
	)
	data, err := os.ReadFile("input.txt")

	if err != nil {
		fmt.Println("Unable to read input.txt")
		return ints, err
	}

	dataAsString := string(data)
	inputAsStrings := strings.Split(dataAsString, "\n")
	ints = make([]int, len(inputAsStrings)-1)

	for idx, d := range inputAsStrings {
		cleanString := strings.TrimSpace(d)

		if cleanString != "" {
			i, err := strconv.Atoi(strings.TrimSpace(d))

			if err != nil {
				fmt.Printf("something went wrong reading input at index %d\n", idx)
				return ints, err
			}

			ints[idx] = i
		}
	}

	return ints, err
}

func main() {
	ints, err := loadInput()

	if err != nil {
		panic(err)
	}

	counts := CountRollingIncrements(ints)
	fmt.Printf("Found a total of %d instances where a depth measurement increases\n", counts)

	counts = CountRollingIncrementsOfThree(ints)
	fmt.Printf("Found a total of %d sequences of three numbers where a depth measurement increases\n", counts)
}
