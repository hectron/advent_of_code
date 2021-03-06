package main

import (
	"testing"
)

func TestCountRollingIncrements(t *testing.T) {
	t.Run("It returns the number of rolling increments", func(t *testing.T) {
		input := []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}
		want := 7

		got := CountRollingIncrements(input)

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}

func TestCountRollingIncrementsOfThree(t *testing.T) {
	t.Run("It returns the number of rolling increments", func(t *testing.T) {
		input := []int{607, 618, 618, 617, 647, 716, 769, 792}
		want := 5

		got := CountRollingIncrementsOfThree(input)

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}
