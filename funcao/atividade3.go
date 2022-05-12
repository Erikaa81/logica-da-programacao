//Crie uma função que calcule o IMC de uma pessoa e informe se ela está acima, abaixo ou no peso ideal, conforme na tabela abaixo

package main

import (
	"fmt"
	"math"
)

const (
	PESO_BAIXO = "Peso Baixo"
	NORMAL     = "Normal"
	SOBREPESO  = "Sobrepeso"
	OBESIDADE  = "Obesidade"
)

func main() {

	peso := 63.0
	altura := 1.70

	fmt.Println(imc(peso, altura))

}

func imc(peso float64, altura float64) (float64, string) {

	imc := peso / (math.Pow(altura, 2))
	res_1 := math.Round(imc*100) / 100

	if imc < 18.5 {
		return res_1, PESO_BAIXO

	} else if imc < 25 {
		return res_1, NORMAL

	} else if imc <= 30 {
		return res_1, SOBREPESO

	} else {
		return res_1, OBESIDADE
	}
}
