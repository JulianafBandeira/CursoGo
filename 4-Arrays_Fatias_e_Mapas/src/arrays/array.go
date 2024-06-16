package main

import "fmt"

func main() {

	var x [5]float32

	x[0] = 5.3
	x[1] = 7.2
	x[2] = 6.2
	x[3] = 8.0
	x[4] = 10.0

	var total float32 = 0

	for i := 0; i < len(x); i++ {
		total += x[i]
	}
	fmt.Println(total/float32(len(x)))
}
