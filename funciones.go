package main

import "fmt"

func main() {
	fmt.Println(saludo())
	fmt.Println(numeros())
	fmt.Println(ganado(2))
	alumnos("Panchito", "Fulanito", "Menganito")
}

func saludo() (string, string) {
	var a string = "Hola"
	var b string = " mundo"

	return a, b
}

func numeros() (a int, b float32) {
	a = 3
	b = 0.141516
	return
}

// Closure
func ganado(numero int) (int, string) {
	vacas := func() int {
		return numero * 3
	}

	return vacas() + 1, " vacas"
}

func alumnos(alumno ...string) {
	for _, valor := range alumno {
		fmt.Println(valor)
	}
}
