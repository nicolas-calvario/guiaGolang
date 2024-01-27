package main

import (
	"fmt"
	"poo/animal"
	"poo/book"
)

func main() {
	fmt.Println("programacion orientada a objetos en go")
	//em()
	ManejoInterfaces()

}

// estructuras y metodos
func em() {
	mybook := book.NewBook("Primer libro de texto", "Nicolas Calvario", 55) //llamando al contructor

	mybook.PrintInfo() // llamando a una funcion de la estructura

	mybook.SetTitle("Primer libro de texto ED")

	//mybook.PrintInfo()
	mytexbook := book.NewTextBook("Primer libro de texto nuevo", "Nicolas Calvario", 55, "Trillas", "Principiates")
	// mytexbook.PrintInfo()

	// mytexbook.PrintInfo()

	//---------
	book.Print(mybook)
	book.Print(mytexbook)

}


func ManejoInterfaces()  {
	miperro := animal.Perro{Nombre: "Kusco"}
	miGato := animal.Gato{Nombre: "Pepinillo"}
	animal.HaceSonido(&miGato)
	animal.HaceSonido(&miperro)

	animales := []animal.IAnimal{
		&animal.Gato{Nombre: "mac"},
		&animal.Gato{Nombre: "tom"},
		&animal.Perro{Nombre: "rex"},
		&animal.Gato{Nombre: "mihci"},
	}

	for _, v := range animales {
		v.Sonido()
	}


}