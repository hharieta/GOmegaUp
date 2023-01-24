package main

import (
	"fmt"
)

func main() {
	var array [2]string
	array[0] = "Manzana"
	array[1] = "Pera"
	fmt.Println(array)

	frutas := [3]string{"Manzana", "Pera", "Banana"}
	fmt.Println(frutas)

	var frutasArray [2][2]string
	frutasArray[0][0] = "Manzana"
	frutasArray[0][1] = "Golden"
	frutasArray[1][0] = "Uva"
	frutasArray[1][1] = "Moscatel"

	fmt.Println(frutasArray)

	// SLICE
	var verduras = []string{"calabaza", "cebolla", "papa"}
	verduras = append(verduras, "chicharos")

	fmt.Println(verduras[0:3])
	fmt.Println("Verduras items: ", len(verduras))

}
