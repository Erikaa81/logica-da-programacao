//Elabore um fluxograma e o algoritmo para calcular a área de um círculo, considerando a seguinte fórmula:
//A = π . r²

package main

import (
	"fmt"
	"math"
)

func main() {
raio:= 3.0
area := math.Pi * math.Pow(raio, 2)
fmt.Println(area)


}