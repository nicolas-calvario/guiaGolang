package animal

import "fmt"

type IAnimal interface {
	Sonido()
}

type Perro struct {
	Nombre string
}

func (p *Perro) Sonido() {
	fmt.Println(p.Nombre + " Hace guau guau")
}

// ---------------------------
type Gato struct {
	Nombre string
}

func (g *Gato) Sonido() {
	fmt.Println(g.Nombre + " Hace Miauuu")
}

func HaceSonido(animal IAnimal) {
	animal.Sonido()
}
