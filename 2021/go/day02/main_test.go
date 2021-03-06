package main

import (
	"reflect"
	"testing"
)

func TestGetFinalPosition(t *testing.T) {
	t.Run("It returns a valid final position", func(t *testing.T) {
		input := []Movement{
			{"forward", 5},
			{"down", 5},
			{"forward", 8},
			{"up", 3},
			{"down", 8},
			{"forward", 2},
		}

		want := Position{15, 10, 0}

		got := GetFinalPosition(input)

		if !reflect.DeepEqual(want, got) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func TestGetFinalPositionWithAim(t *testing.T) {
	t.Run("it returns a valid final position", func(t *testing.T) {
		input := []Movement{
			{"forward", 5},
			{"down", 5},
			{"forward", 8},
			{"up", 3},
			{"down", 8},
			{"forward", 2},
		}

		want := Position{15, 60, 10}

		got := GetFinalPositionWithAim(input)

		if !reflect.DeepEqual(want, got) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}
