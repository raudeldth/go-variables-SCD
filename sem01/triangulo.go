package main

import "fmt"

func main() {
	var base float32
	var altura float32

	fmt.Print("Ingrese la base del triangulo: ")
	fmt.Scan(&base)

	fmt.Print("Ingrese la altura del triangulo: ")
	fmt.Scan(&altura)

	area := base * altura / 2

	fmt.Println("El area del triangulo es: ", area)
}
