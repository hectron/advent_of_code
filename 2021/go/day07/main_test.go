package main

import (
	"testing"
)

func TestCalculateCheapestRoute(t *testing.T) {
	t.Run("It returns the most fuel efficient route to align numbers", func(t *testing.T) {
		input := []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}

		wantCost := 37
		wantPosition := 2

		gotCost, gotPosition := CalculateCheapestCostAndRoute(input)

		if gotPosition != wantPosition {
			t.Errorf("got %d, want %d", gotPosition, wantPosition)
		}

		if gotCost != wantCost {
			t.Errorf("got %d, want %d", gotCost, wantCost)
		}
	})
}

func TestCalculateCostOfRoute(t *testing.T) {
	testCases := []struct {
		Description    string
		Position, Want int
	}{
		{"Moving to index 2 results in a cost of 37", 2, 37},
		{"Moving to index 1 results in a cost of 41", 1, 41},
		{"Moving to index 3 results in a cost of 39", 3, 39},
		{"Moving to index 10 results in a cost of 71", 10, 71},
	}

	for _, tc := range testCases {
		t.Run(tc.Description, func(t *testing.T) {
			input := []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}

			got := CostOfMovingTo(input, tc.Index)

			if got != tc.Want {
				t.Errorf("got %d, want %d", got, tc.Want)
			}
		})
	}
}
