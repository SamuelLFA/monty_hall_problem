package main

import "testing"

func TestChooseWrongDoor(t *testing.T) {
	tt := []struct {
		choosedDoor int
		rightDoor   int
		wrongDoor   int
	}{
		{0, 0, 1},
		{1, 1, 2},
		{2, 2, 0},

		{0, 1, 2},
		{1, 0, 2},
		{2, 0, 1},
		{0, 2, 1},
		{1, 2, 0},
		{2, 1, 0},
	}

	for _, tc := range tt {
		wrongDoor := chooseWrongDoor(tc.choosedDoor, tc.rightDoor)
		if wrongDoor != tc.wrongDoor {
			t.Errorf("Wrong door should be %d, got %d", tc.wrongDoor, wrongDoor)
		}
	}
}

func TestChooseAnotherDoor(t *testing.T) {
	tt := []struct {
		choosedDoor int
		wrongDoor   int
		anotherDoor int
	}{
		{0, 1, 2},
		{1, 0, 2},
		{2, 0, 1},
		{0, 2, 1},
		{1, 2, 0},
		{2, 1, 0},
	}

	for _, tc := range tt {
		anotherDoor := chooseAnotherDoor(tc.choosedDoor, tc.wrongDoor)
		if anotherDoor != tc.anotherDoor {
			t.Errorf("Another door should be %d, got %d", tc.anotherDoor, anotherDoor)
		}
	}
}
