package main

import (
	"reflect"
	"testing"
)

func TestMarkNumberOnBoard(t *testing.T) {
	t.Run("It marks the number if it is on the board", func(t *testing.T) {
		board := Board{
			[][]Entry{
				{
					{22, false}, {13, false}, {17, false}, {11, false}, {0, false},
				},
				{
					{8, false}, {2, false}, {23, false}, {4, false}, {24, false},
				},
				{
					{21, false}, {9, false}, {14, false}, {16, false}, {7, false},
				},
				{
					{6, false}, {10, false}, {3, false}, {18, false}, {5, false},
				},
				{
					{1, false}, {12, false}, {20, false}, {15, false}, {19, false},
				},
			},
		}

		marked := board.MarkNumber(7)

		if !marked {
			t.Errorf("No number was marked on the board")
		}

		got := board.Entries[2][4].Called
		want := true

		if got != want {
			t.Errorf("%d was not marked correctly - %v", 7, board.Entries[2][4])
		}
	})

	t.Run("It does not mark any numbers that do not match", func(t *testing.T) {
		board := Board{
			[][]Entry{
				{
					{22, false}, {13, false}, {17, false}, {11, false}, {0, false},
				},
				{
					{8, false}, {2, false}, {23, false}, {4, false}, {24, false},
				},
				{
					{21, false}, {9, false}, {14, false}, {16, false}, {7, false},
				},
				{
					{6, false}, {10, false}, {3, false}, {18, false}, {5, false},
				},
				{
					{1, false}, {12, false}, {20, false}, {15, false}, {19, false},
				},
			},
		}

		number := 99
		marked := board.MarkNumber(number)

		if marked {
			t.Errorf("Marked %d, but did not expect to", number)
		}

		for rowIndex, row := range board.Entries {
			for entryIdx, entry := range row {
				if entry.Called {
					t.Errorf("%d was marked, but was not expecting to. Row: %d, index: %d", entry.Value, rowIndex, entryIdx)
				}
			}
		}
	})
}

func TestBoardIsWinner(t *testing.T) {
	testCases := []struct {
		msg   string
		input []int
		want  bool
	}{
		{"It returns true if an entire row was called", []int{6, 10, 3, 18, 5}, true},
		{"It returns true if an entire column was called", []int{11, 4, 16, 18, 15}, true},
		{"It returns false if less than five numbers are called", []int{11, 4, 3, 18}, false},
		{"It returns false on a diagonal match", []int{22, 2, 14, 18, 19}, false},
	}

	for _, tc := range testCases {
		t.Run(tc.msg, func(t *testing.T) {
			input := []string{
				"22 13 17 11  0",
				"8  2 23  4 24",
				"21  9 14 16  7",
				"6 10  3 18  5",
				"1 12 20 15 19",
			}

			board, err := BuildBoard(input)

			if err != nil {
				t.Errorf("Got error building board: %s", err)
			}

			got := board.IsWinner()
			want := false

			if got != want {
				t.Errorf("Expected the board not to be a winner, but it was")
			}

			for _, num := range tc.input {
				board.MarkNumber(num)
			}

			got = board.IsWinner()

			if got != tc.want {
				t.Errorf("got %t, want %t", got, tc.want)
			}
		})
	}
}

func TestSumOfNumbersNotCalled(t *testing.T) {
	t.Run("It returns the sum of every number not called", func(t *testing.T) {
		input := []string{
			"14 21 17 24  4",
			"10 16 15  9 19",
			"18  8 23 26 20",
			"22 11 13  6  5",
			"2  0 12  3  7",
		}

		board, err := BuildBoard(input)

		if err != nil {
			t.Errorf("Got error building board: %s", err)
		}

		for _, num := range []int{7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24} {
			board.MarkNumber(num)
		}

		got := board.NonMatchSum()
		want := 188

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}

func TestBuildBoard(t *testing.T) {
	t.Run("It builds a board", func(t *testing.T) {
		input := []string{
			"22 13 17 11  0",
			"8  2 23  4 24",
			"21  9 14 16  7",
			"6 10  3 18  5",
			"1 12 20 15 19",
		}
		want := Board{
			[][]Entry{
				{
					{22, false}, {13, false}, {17, false}, {11, false}, {0, false},
				},
				{
					{8, false}, {2, false}, {23, false}, {4, false}, {24, false},
				},
				{
					{21, false}, {9, false}, {14, false}, {16, false}, {7, false},
				},
				{
					{6, false}, {10, false}, {3, false}, {18, false}, {5, false},
				},
				{
					{1, false}, {12, false}, {20, false}, {15, false}, {19, false},
				},
			},
		}

		got, err := BuildBoard(input)

		if err != nil {
			t.Errorf("Got error building board: %s", err)
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func TestEvaluatingMultipleBoards(t *testing.T) {
	t.Run("It finds the first matching board, and the last number called", func(t *testing.T) {
		bingoSequence := []int{7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24, 10, 16, 13, 6, 15, 25, 12, 22, 18, 20, 8, 19, 3, 26, 1}

		boards := []Board{}

		input := [][]string{
			{
				"22 13 17 11  0",
				" 8  2 23  4 24",
				"21  9 14 16  7",
				" 6 10  3 18  5",
				" 1 12 20 15 19",
			},
			{
				" 3 15  0  2 22",
				" 9 18 13 17  5",
				"19  8  7 25 23",
				"20 11 10 24  4",
				"14 21 16 12  6",
			},
			{
				"14 21 17 24  4",
				"10 16 15  9 19",
				"18  8 23 26 20",
				"22 11 13  6  5",
				" 2  0 12  3  7",
			},
		}

		for _, i := range input {
			board, err := BuildBoard(i)

			if err != nil {
				t.Errorf("Got an error building a board: %v", err)
			}

			boards = append(boards, board)
		}

		board, winningNumber := EvaluateBoards(boards, bingoSequence)

		if !board.IsWinner() {
			t.Errorf("Got a board, but is not a winner")
		}

		if winningNumber != 24 {
			t.Errorf("Expected winning number of %d, got %d", 24, winningNumber)
		}

		got := board.NonMatchSum() * winningNumber
		want := 4512

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}

func TestFindingLastBoardToWin(t *testing.T) {
	t.Run("It returns the last board to win and the number last called", func(t *testing.T) {
		bingoSequence := []int{7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24, 10, 16, 13, 6, 15, 25, 12, 22, 18, 20, 8, 19, 3, 26, 1}

		boards := []Board{}

		input := [][]string{
			{
				"22 13 17 11  0",
				" 8  2 23  4 24",
				"21  9 14 16  7",
				" 6 10  3 18  5",
				" 1 12 20 15 19",
			},
			{
				" 3 15  0  2 22",
				" 9 18 13 17  5",
				"19  8  7 25 23",
				"20 11 10 24  4",
				"14 21 16 12  6",
			},
			{
				"14 21 17 24  4",
				"10 16 15  9 19",
				"18  8 23 26 20",
				"22 11 13  6  5",
				" 2  0 12  3  7",
			},
		}

		for _, i := range input {
			board, err := BuildBoard(i)

			if err != nil {
				t.Errorf("Got an error building a board: %v", err)
			}

			boards = append(boards, board)
		}

		_, winningNumber := FindLastBoardToWin(boards, bingoSequence)

		if winningNumber != 13 {
			t.Errorf("Expected 13 to be the winning number, got %d", winningNumber)
		}
	})
}
