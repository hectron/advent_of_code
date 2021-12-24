package main

import (
	"testing"
)

func TestCalculateCheapestRoute(t *testing.T) {
	t.Run("It returns the most fuel efficient route to align numbers", func(t *testing.T) {
		input := []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}

		wantCost := 168
		wantPosition := 5

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
		{"Moving to index 0 results in a cost of 290", 0, 290},
		{"Moving to index 2 results in a cost of 206", 2, 206},
		{"Moving to index 10 results in a cost of 71", 10, 311},
		{"Moving to index 5 results in a cost of 168", 5, 168},
	}

	for _, tc := range testCases {
		t.Run(tc.Description, func(t *testing.T) {
			input := []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}

			got := CostOfMovingTo(input, tc.Position)

			if got != tc.Want {
				t.Errorf("got %d, want %d", got, tc.Want)
			}
		})
	}
}
