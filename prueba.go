package main

import (
	"fmt"
)

func main() {
	var myCoffe = Coffe{"espresso", 5.22, true, 1}

	var name string = "Gatovsky"
	var number int
	number = 5 + 6
	var goVersion float64
	goVersion = 1.8
	const PI float32 = 3.141516
	fmt.Println("Hola mundo, soy ", name, "Saludos Golang", goVersion, number, PI)

	fmt.Println("My Coffe: ", myCoffe)
	fmt.Println("Price: ", myCoffe.price)
}

type Coffe struct {
	name   string
	price  float64
	suggar bool
	milk   int
}
