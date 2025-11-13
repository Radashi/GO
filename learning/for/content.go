package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func TrueFalse() {
	fmt.Println("You find yourself in a dimly lit cavern.")
	var command = "Walk outside"
	var exit = strings.Contains(command, "outside")
	fmt.Println("You leave the cave:", exit)
}

func Comparisions() {
	fmt.Println("There is a big sign near the entrance that reads 'No minors'")
	var age = 41
	var minor = age < 18
	fmt.Printf("At age %v, am i a minor? %v\n", age, minor)
}

func BranchingIf() {
	var command = "go inside"

	switch command {
	case "go east":
		fmt.Println("You head further up the mountain")
	case "go inside":
		fmt.Println("You enter the cave where you live out...")
	default:
		fmt.Println("Didn't quite get that")
	}
}

func GuessNumber(numero int) {
	var resultado string
	for {
		randomGuess := rand.Intn(100)
		if randomGuess == numero {
			resultado = fmt.Sprintf("Se adivino su numero, es %v", numero)
			fmt.Println(resultado)
			break
		} else {
			resultado = fmt.Sprintf("%v no era el numero", randomGuess)
			fmt.Println(resultado)
		}
	}
}

func Count() {
	var count int = 10

	for count > 0 {
		fmt.Println(count)
		time.Sleep(time.Second)
		if rand.Intn(100) == 0 {
			break
		}
		count--
	}
	if count == 0 {
		fmt.Println("liftoff!")
	} else {
		fmt.Println("Launch failed")
	}
}
