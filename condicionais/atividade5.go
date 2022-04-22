//Em um país imaginário denominado Lisarb, todos os habitantes ficam felizes em pagar seus impostos, pois sabem que nele não
//existem políticos corruptos e os recursos arrecadados são utilizados em benefício da população, sem qualquer desvio. A moeda deste país é o Rombus, cujo símbolo é o R$.
//Leia um valor com duas casas decimais, equivalente ao salário de uma pessoa de Lisarb. 
//Em seguida, calcule e mostre o valor que esta pessoa deve pagar de Imposto de Renda, segundo a tabela abaixo.
// R$ 0 - 2000.00 :isento   R$ 2000.01 - 3000.00: 8%  R$ 3000.01 -4500: 18% R$ > 4500 : 28%
//Lembre que, se o salário for R$ 3002.00, a taxa que incide é de 8% apenas sobre R$ 1000.00, pois a faixa de salário que fica de R$ 0.00 até R$ 2000.00 é isenta de Imposto de Renda.
// + 18% sobre R$ 2.00, o que resulta em R$ 80.36 no total. O valor deve ser impresso com duas casas decimais.

package main

import (
	"fmt"
)

func main() {
	salario := 3002.00
	valor := salario - 2000.00
	taxa8 := valor * 8 / 100
	taxa18 := ((valor-1000.00)*18/100 + 80)
	taxa28 := ((valor-2500.00)*28/100 + 80 + 270)

	if salario <= 2000.00 {
		fmt.Printf("Isento de taxa")

	} else if salario <= 3000.00 {
		fmt.Printf("taxa : R$ %.2f ", taxa8)

	} else if salario <= 4500.00 {

		fmt.Printf("taxa : R$ %.2f", taxa18)
	} else {

		fmt.Printf("taxa : R$ %.2f", taxa28)
	}

}