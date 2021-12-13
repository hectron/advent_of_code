package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestParseInput(t *testing.T) {
	t.Run("It returns a slice of tuples", func(t *testing.T) {
		input := `0,9 -> 5,9
		8,0 -> 0,8
		9,4 -> 3,4
		2,2 -> 2,1
		7,0 -> 7,4
		6,4 -> 2,0
		0,9 -> 2,9
		3,4 -> 1,4
		0,0 -> 8,8
		5,5 -> 8,2
		`

		want := []Movement{
			{From: Point{0, 9}, To: Point{5, 9}},
			{From: Point{8, 0}, To: Point{0, 8}},
			{From: Point{9, 4}, To: Point{3, 4}},
			{From: Point{2, 2}, To: Point{2, 1}},
			{From: Point{7, 0}, To: Point{7, 4}},
			{From: Point{6, 4}, To: Point{2, 0}},
			{From: Point{0, 9}, To: Point{2, 9}},
			{From: Point{3, 4}, To: Point{1, 4}},
			{From: Point{0, 0}, To: Point{8, 8}},
			{From: Point{5, 5}, To: Point{8, 2}},
		}

		got := ParseInput(strings.Split(input, "\n"))

		if !reflect.DeepEqual(want, got) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func TestGenerateDiagram(t *testing.T) {
	t.Run("It generates a map with counts", func(t *testing.T) {
		input := []Movement{
			{From: Point{0, 9}, To: Point{5, 9}},
			{From: Point{8, 0}, To: Point{0, 8}},
			{From: Point{9, 4}, To: Point{3, 4}},
			{From: Point{2, 2}, To: Point{2, 1}},
			{From: Point{7, 0}, To: Point{7, 4}},
			{From: Point{6, 4}, To: Point{2, 0}},
			{From: Point{0, 9}, To: Point{2, 9}},
			{From: Point{3, 4}, To: Point{1, 4}},
			{From: Point{0, 0}, To: Point{8, 8}},
			{From: Point{5, 5}, To: Point{8, 2}},
		}

		diagram := GenerateDiagram(input)

		got := diagram[Point{0, 9}]
		want := 2

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}

func TestFindMostVisitedPoints(t *testing.T) {
	t.Run("It returns all the points that have been visited twice", func(t *testing.T) {
		input := []Movement{
			{From: Point{0, 9}, To: Point{5, 9}},
			{From: Point{8, 0}, To: Point{0, 8}},
			{From: Point{9, 4}, To: Point{3, 4}},
			{From: Point{2, 2}, To: Point{2, 1}},
			{From: Point{7, 0}, To: Point{7, 4}},
			{From: Point{6, 4}, To: Point{2, 0}},
			{From: Point{0, 9}, To: Point{2, 9}},
			{From: Point{3, 4}, To: Point{1, 4}},
			{From: Point{0, 0}, To: Point{8, 8}},
			{From: Point{5, 5}, To: Point{8, 2}},
		}

		diagram := GenerateDiagram(input)

		mostVisitedPoints := FindMostVisitedPoints(diagram)

		want := 5
		got := len(mostVisitedPoints)

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}

func TestFindLineSegment(t *testing.T) {
	testCases := []struct {
		Message  string
		Movement Movement
		Want     []Point
	}{
		{
			Message:  "It returns the line between two points if x is different and negative",
			Movement: Movement{From: Point{9, 7}, To: Point{7, 7}},
			Want:     []Point{{9, 7}, {8, 7}, {7, 7}},
		},
		{
			Message:  "It returns the line between two points if x is different and positive",
			Movement: Movement{From: Point{7, 7}, To: Point{9, 7}},
			Want:     []Point{{7, 7}, {8, 7}, {9, 7}},
		},
		{
			Message:  "It returns the line between two points if y is different and negative",
			Movement: Movement{From: Point{1, 3}, To: Point{1, 1}},
			Want:     []Point{{1, 3}, {1, 2}, {1, 1}},
		},
		{
			Message:  "It returns the line between two points if y is different and positive",
			Movement: Movement{From: Point{1, 1}, To: Point{1, 3}},
			Want:     []Point{{1, 1}, {1, 2}, {1, 3}},
		},
		{
			Message:  "It returns the same point if From and To are the same",
			Movement: Movement{From: Point{1, 1}, To: Point{1, 1}},
			Want:     []Point{{1, 1}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Message, func(t *testing.T) {
			got := tc.Movement.FindLineSegments()

			if !reflect.DeepEqual(got, tc.Want) {
				t.Errorf("got %v, want %v", got, tc.Want)
			}
		})
	}
}
