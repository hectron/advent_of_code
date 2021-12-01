package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func CountRollingIncrements(numberSequence []int) int {
	count := 0
	lastNumber := numberSequence[0]

	for _, number := range numberSequence {
		if number > lastNumber {
			count += 1
		}

		lastNumber = number
	}

	return count
}

func main() {
	data, err := os.ReadFile("input.txt")

	if err != nil {
		fmt.Println("Unable to read input.txt")
		panic(err)
	}

	dataAsString := string(data)
	inputAsStrings := strings.Split(dataAsString, "\n")
	ints := make([]int, len(inputAsStrings)-1)

	for idx, d := range inputAsStrings {
		cleanString := strings.TrimSpace(d)

		if cleanString != "" {
			i, err := strconv.Atoi(strings.TrimSpace(d))

			if err != nil {
				fmt.Printf("something went wrong reading input at index %d\n", idx)
				panic(err)
			}

			ints[idx] = i
		}
	}

	counts := CountRollingIncrements(ints)

	fmt.Printf("Found a total of %d instances where a depth measurement increases", counts)
}
