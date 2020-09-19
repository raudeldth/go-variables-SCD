package main

import "fmt"

func main() {
	var lado int32
	fmt.Print("Ingrese un lado del cuadrado: ")
	fmt.Scanf("%d", &lado)

	area := lado * lado

	fmt.Println("El area del cuadrado es: ", area)
}
