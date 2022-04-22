//Leia um valor inteiro N. Este valor será a quantidade de valores que serão lidos em seguida. 
//Para cada valor lido, mostre uma mensagem em inglês dizendo se este valor lido é par (EVEN), ímpar (ODD), positivo (POSITIVE) ou negativo (NEGATIVE). 
//No caso do valor ser igual a zero (0), embora a descrição correta seja (EVEN NULL), pois por definição zero é par, seu programa deverá imprimir apenas NULL.

package main

import (
	"fmt"
)

func main() {
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
		if i == 0 {
			fmt.Println("NULL")
		}
		if i > 0 {
			fmt.Println("POSITIVE")
		}

		if i%2 == 0 && i > 0 {

			fmt.Println("EVEN")
		}
		if i%2 != 0 {

			fmt.Println("ODD")
		}

	}
}