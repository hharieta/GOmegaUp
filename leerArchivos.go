package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

//struct País
type Pais struct {
	NOMBRE     string
	POBLACION  int
	SUPERFICIE int
}

// struct Paises
type Paises struct {
	Paises []Pais `json:"paises"`
}

func main() {
	escribirArchivo("file.txt")
	LeerArchivo("file.txt")
	leerJson("paises.json")
}

func mostrarError(e error) {
	if e != nil {
		panic(e)
	}
}

func LeerArchivo(file string) {
	datos, errLectura := ioutil.ReadFile(file)
	mostrarError(errLectura)
	fmt.Println(string(datos))
}

func escribirArchivo(file string) {
	contenido := []byte("esto es una cadena de texto")
	datos := ioutil.WriteFile(file, contenido, 0755)
	mostrarError(datos)
}

func leerJson(json_f string) {
	datos, errLectura := os.Open(json_f)
	mostrarError(errLectura)

	byteVal, _ := ioutil.ReadAll(datos)
	var paises Paises

	errLectura = json.Unmarshal(byteVal, &paises)
	mostrarError(errLectura)

	for i := 0; i < len(paises.Paises); i++ {
		fmt.Println("*", paises.Paises[i].NOMBRE, paises.Paises[i].POBLACION, paises.Paises[i].SUPERFICIE)
	}

}
