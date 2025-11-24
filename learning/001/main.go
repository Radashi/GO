package main

import (
	"fmt"
	"strings"
	_ "time"
)

func Saludo(nombre string) string {

	return fmt.Sprintf("Hola %s, bienvenido", nombre)

}

func ParImpar(numero int) string {
	if numero%2 == 0 {
		return fmt.Sprintf("%d es Numero par", numero)

	} else {
		return fmt.Sprintf("%d es Numero impar", numero)
	}
}

func Tablas(numero int) string {
	var tabla strings.Builder

	for i := 0; i < 12; i++ {
		multi := numero * i
		resultado := fmt.Sprintf("%d x %d = %d\n", numero, i, multi)
		tabla.WriteString(resultado)
	}
	return tabla.String()
}

func main() {
	var numero int
	var nombre string

	fmt.Print("Cual es tu nombre? ")
	fmt.Scan(&nombre)
	fmt.Println(Saludo(nombre))

	fmt.Print("----------------------\nNumero: ")
	fmt.Scan(&numero)
	fmt.Println(ParImpar(numero))

	fmt.Print("----------------------\nNumero: ")
	fmt.Scan(&numero)
	fmt.Println(Tablas(numero))
}
