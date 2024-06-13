package main

import "fmt"

const ebulicaoK = 373.15

func main(){

	var tempK = ebulicaoK
	var tempC = tempK - 273.15
  
	
	fmt.Println("A temperatura de ebulição da água em °C =", tempC)
	fmt.Println("A temperatura de ebulição da água em °K =", tempK)
}