package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type LanternFish struct {
	Timer int
}

func (l *LanternFish) Sleep() bool {
	if l.Timer == 0 {
		l.Timer = 6
		return true
	} else {
		l.Timer -= 1
		return false
	}
}

//func ObserveForDays(lanternFish []LanternFish, days int) []LanternFish {
//	for d := 0; d < days; d++ {
//		fmt.Printf("Day %d\n", d+1)
//		newFishCount := 0
//
//		for idx, fish := range lanternFish {
//			hadSpawn := fish.Sleep()
//
//			if hadSpawn {
//				newFishCount += 1
//			}
//
//			lanternFish[idx] = fish
//		}
//
//		fmt.Printf("There are %d new fish\n", newFishCount)
//
//		for i := 0; i < newFishCount; i++ {
//			lanternFish = append(lanternFish, LanternFish{8})
//		}
//	}
//
//	return lanternFish
//}

func ObserveForDays(fishes []LanternFish, days int) map[int]int {
	daysToFish := map[int]int{}

	for _, fish := range fishes {
		daysToFish[fish.Timer] += 1
	}

	for d := 0; d < days; d++ {
		daysToFishesLeft := map[int]int{}
		for i := 0; i <= 8; i++ {
			if fishCount := daysToFish[i]; fishCount > 0 {
				daysToFishesLeft[i-1] = fishCount
			}
		}

		if daysToFishesLeft[-1] > 0 {
			daysToFishesLeft[6] += daysToFishesLeft[-1]
			daysToFishesLeft[8] = daysToFishesLeft[-1]
			delete(daysToFishesLeft, -1)
		}

		daysToFish = daysToFishesLeft
	}

	return daysToFish
}

func ParseInput(input string) []LanternFish {
	var (
		fishes []LanternFish
	)

	timerStrings := strings.Split(input, ",")

	for _, timerString := range timerStrings {
		timer, err := strconv.Atoi(strings.TrimSpace(timerString))

		if err != nil {
			fmt.Println(err)
			return fishes
		}

		fish := LanternFish{timer}
		fishes = append(fishes, fish)
	}

	return fishes
}

func loadInput() (string, error) {
	var (
		result string
		err    error
	)

	filename := "input.txt"
	data, err := os.ReadFile(filename)

	if err != nil {
		return result, err
	}

	return strings.TrimSpace(string(data)), err
}

func main() {
	inputTxt, err := loadInput()

	if err != nil {
		panic(err)
	}

	startFish := ParseInput(inputTxt)

	days := 256
	daysToFishCount := ObserveForDays(startFish, days)

	totalFishCount := 0

	for _, count := range daysToFishCount {
		totalFishCount += count
	}

	fmt.Printf("After %d days, there are %d fish", days, totalFishCount)
}
