//Crie uma função que calcule o IMC de uma pessoa e informe se ela está acima, abaixo ou no peso ideal, conforme na tabela abaixo

package main

import (
	"fmt"
	"math"
)

func main() {

	peso := 63.0
	altura := 1.70

	imc := peso / (math.Pow(altura, 2))

	if imc < 18.5 {
		fmt.Println("Baixo peso")

	} else if imc < 25 {
		fmt.Println("Normal")

	} else if imc <= 30 {
		fmt.Println("Sobrepeso")

	} else {
		fmt.Println("Obesidade")
	}
}
