package main

import "fmt"

func main() {
	var fah float64

	fmt.Print("Ingrese la temperatura en Fahrenheit: ")
	fmt.Scan(&fah)

	celsius := (fah - 32) * 5 / 9

	fmt.Println("La temperatura convertida a Celsius es: ", celsius)

}
