package main

import "fmt"

// Go não possui while
func main() {

	i := 0

	for i < 20 {
		fmt.Println("i é menor que 20")
		i++
	}
}
