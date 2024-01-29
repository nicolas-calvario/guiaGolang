package main

import "fmt"

func main() {
	// fmt.Println(sum(2, 2, 2, 3232, 32, 3))
	// fmt.Println(sum(2, 2, 2))
	// fmt.Println(NumText("nicolas", 1212, 12, 1, 21, 21, 2, 12, 12))
	// imprimirDatos(12, true, 23.423, "Nicolas", 1.23)

	fmt.Println(Factorial(5))

	//fuciones anonimas
	saludo := func(s string) {
		fmt.Printf("hola: %s\n", s)
	}

	saludo("Nicolas")

	fmt.Println(aplicar(duplicar, 5))
	fmt.Println(aplicar(triplicar, 5))

	FunOrdenSuperior()
	Course()
}

// funciona variadicas
func sum(nums ...int) int {
	var total int
	for _, v := range nums {
		total += v
	}
	return total
}

func NumText(s string, num ...int) int {
	var total int
	for _, v := range num {
		total += v
	}
	fmt.Println(s)
	return total
}
func imprimirDatos(datos ...interface{}) {
	for _, v := range datos {
		fmt.Printf("%T- %v\n", v, v)
	}
}

//fuciones recursivas

func Factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * Factorial(n-1)
}

// -----------
func duplicar(n int) int {
	return n * 2
}
func triplicar(n int) int {
	return n * 3
}

func aplicar(f func(int) int, n int) int {
	return f(n)
}

//funciones de orden superior

func doble(f func(int) int, x int) int {
	return f(x * 2)
}

func AddOne(x int) int {
	return x + 1
}

func FunOrdenSuperior() {
	r := doble(AddOne, 3)

	fmt.Println("Resultado: ", r)
}

//--Course

func Course() {
	nextint := incrementar()

	fmt.Println(nextint())
	fmt.Println(nextint())

	fmt.Println(nextint())

	fmt.Println(nextint())

	fmt.Println(nextint())

}

func incrementar() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
