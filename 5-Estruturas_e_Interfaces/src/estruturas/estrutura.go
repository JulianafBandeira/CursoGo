package main

import "fmt"

type pessoa struct{
	nome string
	idade int
}

func main(){

	fmt.Println(pessoa{"juliana", 23})	

}