package main

import (
	"reflect"
	"testing"
)

func TestFlipBinaryString(t *testing.T) {
	t.Run("test flipping binary", func(t *testing.T) {
		testCases := []struct {
			Got, Want string
		}{
			{"01001", "10110"},
			{"00100", "11011"},
			{"11110", "00001"},
			{"10110", "01001"},
			{"10111", "01000"},
			{"10101", "01010"},
			{"01111", "10000"},
			{"00111", "11000"},
			{"11100", "00011"},
			{"10000", "01111"},
			{"11001", "00110"},
			{"00010", "11101"},
			{"01010", "10101"},
		}

		for _, tc := range testCases {
			want := tc.Want

			got := FlipBinaryString(tc.Got)

			if want != got {
				t.Errorf("got %s, want %s", got, want)
			}
		}
	})
}

func TestFindRates(t *testing.T) {
	t.Run("test finding the gamma and epsilon rate", func(t *testing.T) {
		report := []string{
			"01001",
			"00100",
			"11110",
			"10110",
			"10111",
			"10101",
			"01111",
			"00111",
			"11100",
			"10000",
			"11001",
			"00010",
			"01010",
		}

		got := FindRates(report)
		want := Result{"10110", "01001"}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func TestResultPowerConsumption(t *testing.T) {
	t.Run("It calculates the correct power consumption rate", func(t *testing.T) {
		report := Result{"10110", "01001"}
		want := int64(198)
		got := report.PowerConsumption()

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}
