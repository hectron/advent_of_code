package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type GammaAndEpsilonResult struct {
	GammaRate, EpsilonRate string
}

func (r *GammaAndEpsilonResult) PowerConsumption() int64 {
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

func FindMostCommonBinary(report []string, position int) byte {
	oneCount := 0
	zeroCount := 0

	for _, binaryString := range report {
		if binaryString == "" {
			continue
		}

		if binaryString[position] == '0' {
			zeroCount += 1
		} else {
			oneCount += 1
		}
	}

	if oneCount >= zeroCount {
		return '1'
	} else {
		return '0'
	}
}

func FindLeastCommonBinary(report []string, position int) byte {
	oneCount := 0
	zeroCount := 0

	for _, binaryString := range report {
		if binaryString == "" {
			continue
		}

		if binaryString[position] == '0' {
			zeroCount += 1
		} else {
			oneCount += 1
		}
	}

	if oneCount < zeroCount {
		return '1'
	}

	return '0'
}

func FindGammaAndEpsilonRates(report []string) GammaAndEpsilonResult {
	lengthOfBinaryString := len(report[0]) // assume that every binary input is of equal length/base

	gammaRate := make([]byte, lengthOfBinaryString)

	for i := 0; i < lengthOfBinaryString; i++ {
		gammaRate[i] = FindMostCommonBinary(report, i)
	}

	gammaRateString := string(gammaRate)
	epsilonRateString := FlipBinaryString(gammaRateString)

	return GammaAndEpsilonResult{gammaRateString, epsilonRateString}
}

type OxygenAndCO2Result struct {
	OxygenBinaryString, CO2BinaryString string
}

func (o *OxygenAndCO2Result) LifeSupportRating() int64 {
	oxygen, err := strconv.ParseInt(o.OxygenBinaryString, 2, 64)

	if err != nil {
		return 0
	}

	co2, err := strconv.ParseInt(o.CO2BinaryString, 2, 64)

	if err != nil {
		return 0
	}

	return oxygen * co2
}
func FindOxygenAndCO2(report []string) OxygenAndCO2Result {
	oxygen := FindOxygen(report)
	co2 := FindCO2(report)

	return OxygenAndCO2Result{oxygen, co2}
}

func FindOxygen(report []string) string {
	lengthOfBinaryString := len(report[0])

	subset := report

	for i := 0; i < lengthOfBinaryString; i++ {
		fmt.Println("----", len(subset))
		if len(subset) == 1 {
			break
		}

		commonBinary := FindMostCommonBinary(subset, i)
		fmt.Printf("Most common binary at position %d: %s", i, string(commonBinary))

		newSubset := []string{}

		for _, sub := range subset {
			if sub == "" {
				continue
			}

			if sub[i] == commonBinary {
				newSubset = append(newSubset, sub)
			}
		}

		subset = newSubset
	}

	return subset[0]
}

func FindCO2(report []string) string {
	lengthOfBinaryString := len(report[0])

	subset := report

	for i := 0; i < lengthOfBinaryString; i++ {
		if len(subset) == 1 {
			break
		}

		commonBinary := FindLeastCommonBinary(subset, i)

		newSubset := []string{}

		for _, sub := range subset {
			if sub == "" {
				continue
			}
			if sub[i] == commonBinary {
				newSubset = append(newSubset, sub)
			}
		}

		subset = newSubset
	}

	return subset[0]
}

func main() {
	report, err := loadInput()

	if err != nil {
		panic(err)
	}

	rates := FindGammaAndEpsilonRates(report)

	fmt.Printf("Got a gamma rate of %s, epsilon rate of %s, resulting in a power consumption of %d", rates.GammaRate, rates.EpsilonRate, rates.PowerConsumption())

	oxygenAndCo2Result := FindOxygenAndCO2(report)

	rating := oxygenAndCo2Result.LifeSupportRating()

	fmt.Printf("Got an oxygen level of %s, co2 level of %s, resulting in a life support rating of %d", oxygenAndCo2Result.OxygenBinaryString, oxygenAndCo2Result.CO2BinaryString, rating)
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
