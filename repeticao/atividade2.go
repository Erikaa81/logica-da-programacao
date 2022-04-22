//Criar um algoritimo que calcula um fatorial de um número inteiro positivo.
package main

import (
	"fmt"
)

func main() {
	numero := 4
	fatorial := 1
	for numero > 0 {
		fatorial = fatorial * numero
		numero = numero - 1

	}
	fmt.Println(fatorial)
}
