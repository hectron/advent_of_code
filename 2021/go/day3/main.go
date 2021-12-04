package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Result struct {
	GammaRate, EpsilonRate string
}

func (r *Result) PowerConsumption() int64 {
	gamma, err := strconv.ParseInt(r.GammaRate, 2, 64)

	if err != nil {
		return 0
	}

	epsilon, err := strconv.ParseInt(r.EpsilonRate, 2, 64)

	if err != nil {
		return 0
	}

	return gamma * epsilon
}

func FlipBinaryString(binaryString string) string {
	returnVal := make([]byte, len(binaryString))

	for idx := range binaryString {
		switch binaryString[idx] {
		case '0':
			returnVal[idx] = '1'
		case '1':
			returnVal[idx] = '0'
		}
	}

	return string(returnVal)
}

func FindRates(report []string) Result {
	lengthOfBinaryString := len(report[0]) // assume that every binary input is of equal length/base

	gammaRate := make([]byte, lengthOfBinaryString)

	for i := 0; i < lengthOfBinaryString; i++ {
		zeroCount := 0
		oneCount := 0

		for _, binaryString := range report {
			if binaryString == "" {
				continue
			}
			if binaryString[i] == '0' {
				zeroCount += 1
			} else {
				oneCount += 1
			}
		}

		if oneCount > zeroCount {
			gammaRate[i] = '1'
		} else {
			gammaRate[i] = '0'
		}
	}

	gammaRateString := string(gammaRate)
	epsilonRateString := FlipBinaryString(gammaRateString)

	return Result{gammaRateString, epsilonRateString}
}

func main() {
	report, err := loadInput()

	if err != nil {
		panic(err)
	}

	rates := FindRates(report)

	fmt.Printf("Got a gamma rate of %d, epsilon rate of %d, resulting in a power consumption of %d", rates.GammaRate, rates.EpsilonRate, rates.PowerConsumption())
}

func loadInput() ([]string, error) {
	var (
		report []string
		err    error
	)

	data, err := os.ReadFile("input.txt")

	if err != nil {
		fmt.Println("Unable to read input.txt")
		return report, err
	}

	dataAsString := string(data)
	report = strings.Split(dataAsString, "\n")

	return report, err
}
