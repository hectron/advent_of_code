package main

import (
	"reflect"
	"testing"
)

var input [][]int = [][]int{
	{2, 1, 9, 9, 9, 4, 3, 2, 1, 0},
	{3, 9, 8, 7, 8, 9, 4, 9, 2, 1},
	{9, 8, 5, 6, 7, 8, 9, 8, 9, 2},
	{8, 7, 6, 7, 8, 9, 6, 7, 8, 9},
	{9, 8, 9, 9, 9, 6, 5, 6, 7, 8},
}

func TestFindLowPointCoordinates(t *testing.T) {
	t.Run("It finds all the coordinates of the lowest points", func(t *testing.T) {
		want := []Coordinate{
			{X: 1, Y: 0},
			{X: 9, Y: 0},
			{X: 2, Y: 2},
			{X: 6, Y: 4},
		}

		got := FindLowPointCoordinates(input)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func TestCalculateRiskLevel(t *testing.T) {
	t.Run("It returns the proper risk level", func(t *testing.T) {
		want := 15

		coordinates := FindLowPointCoordinates(input)
		got := CalculateRiskLevel(input, coordinates)

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}

func TestFindBasins(t *testing.T) {
	t.Run("It returns all the basins", func(t *testing.T) {

		want := []Basin{
			{Coordinates: []Coordinate{{X: 0, Y: 0}, {X: 0, Y: 1}, {X: 1, Y: 0}}},
			{Coordinates: []Coordinate{{X: 5, Y: 0}, {X: 6, Y: 0}, {X: 6, Y: 1}, {X: 7, Y: 0}, {X: 8, Y: 0}, {X: 8, Y: 1}, {X: 9, Y: 0}, {X: 9, Y: 1}, {X: 9, Y: 2}}},
			{Coordinates: []Coordinate{{X: 2, Y: 1}, {X: 1, Y: 2}, {X: 0, Y: 3}, {X: 1, Y: 3}, {X: 1, Y: 4}, {X: 2, Y: 3}, {X: 3, Y: 1}, {X: 3, Y: 3}, {X: 4, Y: 1}, {X: 2, Y: 2}, {X: 3, Y: 2}, {X: 4, Y: 2}, {X: 4, Y: 3}, {X: 5, Y: 2}}},
			{Coordinates: []Coordinate{{X: 7, Y: 2}, {X: 6, Y: 3}, {X: 7, Y: 3}, {X: 8, Y: 3}, {X: 5, Y: 4}, {X: 6, Y: 4}, {X: 7, Y: 4}, {X: 8, Y: 4}, {X: 9, Y: 4}}},
		}
		lowPoints := []Coordinate{
			{X: 1, Y: 0},
			{X: 9, Y: 0},
			{X: 2, Y: 2},
			{X: 6, Y: 4},
		}

		got := FindBasins(input, lowPoints)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("\ngot %v\nwant %v\n", got, want)
		}
	})
}

// func TestFindBiggerNeighbors(t *testing.T) {
// 	t.Run("it finds neighboring values", func(t *testing.T) {
// 		want := []Coordinate{
// 			{Y: 1, X: 2},
// 			{Y: 3, X: 2},
// 			{Y: 2, X: 1},
// 			{Y: 2, X: 3},
// 		}
//
// 		got := FindBiggerNeighbors(input, Coordinate{X: 2, Y: 2})
//
// 		if !reflect.DeepEqual(got, want) {
// 			t.Errorf("got %v, want %v", got, want)
// 		}
// 	})
// }
