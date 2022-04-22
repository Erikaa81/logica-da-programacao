//Desenvolva um algoritmo que solicite o preço de três produtos e informe qual produto deve 
//ser comprado, sabendo que a decisão é sempre pelo mais barato.

package main

import "fmt"

func main() {

	produto1 := 10.0
	produto2 := 15.0
	produto3 := 7.0

	if produto1 < produto2 && produto1 < produto3 {
		fmt.Printf("Compra Produto 1")

	} else if produto2 < produto1 && produto2 < produto3 {
		fmt.Printf("Compra Produto 2")

	} else if produto3 < produto1 && produto3 < produto2 {
		fmt.Printf("Compra Produto 3")
	}
}