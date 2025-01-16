package main

import (
	"fmt"
	"math/rand"
)

const (
	iterations = 10000
	doors      = 3
)

func main() {
	lucky := 0
	notLucky := 0
	for i := 0; i < iterations; i++ {
		var choosedDoor = rand.Intn(doors)
		var rightDoor = rand.Intn(doors)

		wrongDoor := chooseWrongDoor(choosedDoor, rightDoor)
		fmt.Printf("choosedDoor: %d, rightDoor: %d, wrongDoor: %d\n", choosedDoor, rightDoor, wrongDoor)

		anotherDoor := chooseAnotherDoor(choosedDoor, wrongDoor)

		if choosedDoor == rightDoor {
			lucky++
		}

		if anotherDoor == rightDoor {
			notLucky++
		}
	}

	fmt.Printf("Probability of winning using the first choice: %f\n", float64(lucky)/iterations)
	fmt.Printf("Probability of winning changing the door: %f", float64(notLucky)/iterations)
}

func chooseWrongDoor(choosedDoor int, rightDoor int) int {
	// If the choosed door is the right door, choose the next door
	if choosedDoor == rightDoor {
		return (choosedDoor + 1) % 3
	}

	// If the choosed door is not the right door, choose the door that is not the choosed door and not the right door
	return 3 - choosedDoor - rightDoor
}

func chooseAnotherDoor(choosedDoor int, wrongDoor int) int {
	// If the choosed door is the wrong door, choose the door that is not the choosed door and not the wrong door
	return 3 - choosedDoor - wrongDoor
}
