package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	fmt.Println("Concurrencia")

	//canales
	// canal := make(chan int)
	// canal <- 15      //Como se envian los datos por el canar
	// valor := <-canal // como se asignan valores de un canal
	// fmt.Println(valor, " :valor del canar")
	star := time.Now()
	apis := []string{
		"https://www.facebook.com/",
		"https://www.openstreetmap.org/#map=12/18.7688/-98.8842",
		"https://www.yeeoutube.com/",
		"https://www.google.com/maps",
		"https://talento.evaluatest.com/",
	}

	ch := make(chan string) // se crea el canar con el tipo de dato
	for _, api := range apis {
		go CheckApi(api, ch) // se ejecuta la concurrencia con, y se envia el canal para que haya sincronia
	}

	for i := 0; i < len(apis); i++ {
		fmt.Println(<-ch) // se imprime el calor del canar
	}

	// para ver que si funciona
	//time.Sleep((1 * time.Second))

	delaysed := time.Since(star)
	fmt.Printf("Se ha termiando la ejecucion en %v segudos!\n", delaysed.Seconds())

}

func CheckApi(api string, ch chan string) {
	if _, err := http.Get(api); err != nil {
		ch <- fmt.Sprintf("Errror: !%s esta cahido\n", api) // se asigana el valo de canal
		return
	}
	ch <- fmt.Sprintf("Success: !%s esta fincionando Ok\n", api) // se asigana el valo de canal
}
