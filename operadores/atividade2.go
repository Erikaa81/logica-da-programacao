
//Elabore um algoritmo para exibir o dia, o mês e o ano, dado uma data inteira informada no formato ddmmaaaa.

package main

import "fmt"

func main() {
	data := 1032022

	dia := data / 1000000
	mes := data / 10000 % 10
	ano := data % 10000

	fmt.Printf("dia: %.2d mes: %.2d ano: %d", dia, mes, ano)
}