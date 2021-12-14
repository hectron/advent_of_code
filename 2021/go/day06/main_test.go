package main

import (
	"reflect"
	"testing"
)

func TestLanternSleepAndSpawn(t *testing.T) {
	type Want struct {
		Timer    int
		NewSpawn bool
	}
	testCases := []struct {
		Description string
		Timer       int
		Want        Want
	}{
		{"Sleeping when the timer is 5, sets the new timer to 4", 5, Want{4, false}},
		{"Sleeping when the timer is 1, sets the new timer 0", 1, Want{0, false}},
		{"Sleeping when the timer is 0, sets the new timer 6", 0, Want{6, true}},
	}

	for _, tc := range testCases {
		t.Run(tc.Description, func(t *testing.T) {
			fish := LanternFish{tc.Timer}

			gotBool := fish.Sleep()

			got := fish.Timer

			if got != tc.Want.Timer {
				t.Errorf("got %d, want %d", got, tc.Want.Timer)
			}

			if gotBool != tc.Want.NewSpawn {
				t.Errorf("got %v, want %v", gotBool, tc.Want.NewSpawn)
			}
		})
	}
}

func TestParseInput(t *testing.T) {
	t.Run("It returns the correct number of LanternFish with proper timers", func(t *testing.T) {
		input := "3,4,3,1,2"

		want := []LanternFish{
			{3},
			{4},
			{3},
			{1},
			{2},
		}

		got := ParseInput(input)

		if len(got) != len(want) {
			t.Errorf("got len %d, want len %d", len(got), len(want))
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func TestObserveForDays(t *testing.T) {
	testCases := []struct {
		Description string
		Days        int
		Want        []LanternFish
	}{
		{"Observing for a single day", 1, []LanternFish{{2}, {3}, {2}, {0}, {1}}},
		{"Observing for a single birth", 2, []LanternFish{{1}, {2}, {1}, {6}, {0}, {8}}},
		{"Observing for multiple births", 4, []LanternFish{{6}, {0}, {6}, {4}, {5}, {6}, {7}, {8}, {8}}},
		{"Observing for 18 days", 18, []LanternFish{{6}, {0}, {6}, {4}, {5}, {6}, {0}, {1}, {1}, {2}, {6}, {0}, {1}, {1}, {1}, {2}, {2}, {3}, {3}, {4}, {6}, {7}, {8}, {8}, {8}, {8}}},
	}

	for _, tc := range testCases {
		t.Run(tc.Description, func(t *testing.T) {
			inputTxt := "3,4,3,1,2"
			input := ParseInput(inputTxt)

			got := ObserveForDays(input, tc.Days)

			if len(got) != len(tc.Want) {
				t.Errorf("got len %d, want len %d", len(got), len(tc.Want))
			}

			if !reflect.DeepEqual(got, tc.Want) {
				t.Errorf("got %v, want %v", got, tc.Want)
			}
		})
	}
}
