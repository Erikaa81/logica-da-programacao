//Criar uma função que dadas 3 notas, calcule a média aritmética.

package main

import (
	"fmt"
)

func main() {

	nota1 := 10.00
	nota2 := 8.50
	nota3 := 7.00
	fmt.Printf("Média: %f", calculaMedia(nota1, nota2, nota3))

}
func calculaMedia(nota1, nota2, nota3 float64) float64 {
	return (nota1 + nota2 + nota3) / 3
}
