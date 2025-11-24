package main

import (
	"fmt"
	"github.com/pterm/pterm"
	"os"
	"os/exec"
	"time"
)

func clearTerminal() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func Maps(user string, password int) bool {
	// Limpieza de la pantalla
	clearTerminal()

	// Definicion del map y los datos
	users := make(map[string]map[string]int)
	users["minerva"] = make(map[string]int)
	users["alexander"] = make(map[string]int)
	users["minerva"]["password"] = 301123
	users["alexander"]["password"] = 171204

	// checker
	var check bool

	// condicional
	if users[user]["password"] == password {
		check = true
	} else {
		check = false
	}
	return check
}

func Carta() {
	clearTerminal()

	// Colores de la terminal
	primary := pterm.NewStyle(pterm.FgLightBlue, pterm.Bold)

	// Carta
	cargando := []string{
		"=============================================]\n",
	}
	fmt.Print("[Cargando: ")
	for _, linea := range cargando {
		for _, letra := range linea {
			time.Sleep(200 * time.Millisecond)
			fmt.Print(string(letra))
		}
	}
	clearTerminal()
	carta := []string{
		"Carta dedicada para mi flaquita bonita.\n\n",
		"En esta carta te escribo lo mucho que te quiero\n",
		"y lo feliz que me hace tenerte, eres la razon por\n",
		"la que le hecho muchas ganas a la escuela todo el dia\n",
	}

	for _, linea := range carta {
		for _, letra := range linea {
			time.Sleep(200 * time.Millisecond)
			primary.Print(string(letra))
		}
	}
}

func main() {
	// Capturado de los ususarios y contrasenas
	clearTerminal()
	var user string
	var password int
	fmt.Print("User: ")
	fmt.Scan(&user)
	fmt.Print("Password: ")
	fmt.Scan(&password)

	if Maps(user, password) == true {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println("Contrasena correcta")

	} else {
		time.Sleep(100 * time.Millisecond)
		fmt.Println("Contrasena incorrecta")
	}

	Carta()
}
