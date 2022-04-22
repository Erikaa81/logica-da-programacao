//Criar um fluxograma e algoritmo que dado dois números informados pelo usuário (um menor e um maior, respectivamente - garanta isso!), 
//seja verificado a quantidade de valores divisíveis por 5 que há no intervalo.

package main

import "fmt"

func main() {
	var numero1 int
	var numero2 int

	fmt.Println("Informe um numero:")
	fmt.Scanln(&numero1)
	fmt.Print("Informe outro numero:")
	fmt.Scan(&numero2)

	for numero1 > numero2 {
		fmt.Println("infomre outro valor para o numero1:")
		fmt.Scanln(&numero1)

	}
	divisiveis := 0

	for n := numero1; n <= numero2; n++ {
		if n%5 == 0 {

			divisiveis += 1

		}
	}

	fmt.Printf("No intervalo existem %v numeros divisiveis por 5.\n", divisiveis)
}
