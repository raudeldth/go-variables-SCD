package main

import (
	"fmt"
	"math"
)

func main() {
	var radio float32

	fmt.Print("Ingrese el radio del circulo: ")
	fmt.Scanf("%f", &radio)

	area := math.Pi * radio * radio

	fmt.Println("El area del cuadrado es: ", area)
}
