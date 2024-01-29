package main

import "fmt"

// any
func printList(list ...interface{}) {
	for _, v := range list {
		fmt.Println(v)
	}
}

func printListAnly(list ...any) {
	for _, v := range list {
		fmt.Println(v)
	}
}

func main() {

	printList("Nicolas", 232, true, 2.4444)
	printListAnly("sads", 233, 2333)
}
