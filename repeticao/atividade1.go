//Criar um algoritmo que imprime a soma de todos os números pares entre 100 e 200
package main

import (
	"fmt"
)

func main() {
	somaPares := 0
	for n := 100; n <= 200; n++ {
		if n%2 == 0 {

			somaPares += n
		}

	}
	fmt.Println(somaPares)
}
