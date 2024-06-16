package main

import "fmt"

type retangulo struct{
	cumprimento, altura int
}

func (r * retangulo) area() int{
	return r.cumprimento * r.altura
}

func(r retangulo) perimetro() int{
	return 2*r.cumprimento + 2*r.altura
}

func main() {
	r := retangulo{cumprimento: 10, altura: 5}
	
	fmt.Println("Área: ", r.area())
	fmt.Println("Perímetro: ", r.perimetro())
}