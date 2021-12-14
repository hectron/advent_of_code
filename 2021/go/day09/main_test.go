package main

import (
	"reflect"
	"testing"
)

func TestFindLowPointsInArray(t *testing.T) {
	t.Run("It finds the correct low points in a sample array", func(t *testing.T) {
		input := [][]int{
			{2, 1, 9, 9, 9, 4, 3, 2, 1, 0},
			{3, 9, 8, 7, 8, 9, 4, 9, 2, 1},
			{9, 8, 5, 6, 7, 8, 9, 8, 9, 2},
			{8, 7, 6, 7, 8, 9, 6, 7, 8, 9},
			{9, 8, 9, 9, 9, 6, 5, 6, 7, 8},
		}
		want := []int{1, 0, 5, 5}

		got := FindLowPointsInArray(input)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}
