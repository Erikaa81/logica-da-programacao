//Crie um programa que leia 3 valores inteiros e verifique se eles podem formar  um triângulo.
//Se for possível formar um triângulo, escreva uma mensagem informando se é um triângulo equilátero, isósceles ou escaleno.

package main

import "fmt"

func main() {

	a := 12
	b := 15
	c := 15

	if a > b+c || b > c+a || c > b+a {
		fmt.Println("Não forma triangulo")

	} else if a == b && a == c {
		fmt.Println("Triangulo Equilátero")

	} else if a != b && a != c && b != c {
		fmt.Println("Triangulo Escaleno")

	} else {
		fmt.Println("Triangulo Isósceles")
	}
}
