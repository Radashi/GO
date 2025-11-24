package main

import (
	"fmt"
	"math/rand"
)

func main() {
	scoping()
}

func scoping() {
	for count := 10; count > 0; count-- {
		num := rand.Intn(10) + 1
		fmt.Println(num)
	}

	if num := rand.Intn(3); num == 0 {
		fmt.Println("Space adventurers")
	} else if num == 1 {
		fmt.Println("SpaceX")
	} else {
		fmt.Println("SpaceY")
	}

	switch num := rand.Intn(10); num {
	case 0:
		fmt.Println("Space adventurers #", num)
	case 1:
		fmt.Println("SpaceX #", num)
	case 2:
		fmt.Println("SpaceY #", num)
	default:
		fmt.Println("SpaceZ #", num)
	}
}
