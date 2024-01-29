package main

import "fmt"

func sum[T int | float64](nums ...T) T {
	var total T
	for _, v := range nums {
		total += v
	}
	return total
}

func main() {
	fmt.Println(sum(12, 12, 12, 12, 12))

	fmt.Println(sum(12, 12, 12.23, 12.23, 12))
}
