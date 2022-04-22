//Criar um algoritmo que, dado 10 valores informados pelo usuário, calcule e escreva a média aritmética desses valores lidos.

package main

import "fmt"


func main() {
	var nota int
	somatorioNotas := 0 +nota
	count := 0
	for count < 10 {

		fmt.Print("Informe a nota: ")
		fmt.Scan(&nota)

		somatorioNotas = somatorioNotas + nota
		count = count + 1
	}
	media:= somatorioNotas/10
		fmt.Println("A média das notas é:", media)
	}

