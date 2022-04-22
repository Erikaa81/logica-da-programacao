//Crie uma função que recebe uma string e um caractere, e apague todas as ocorrências desses caractere na string.

package main

import (
	"fmt"
)

func main() {

	frase := "Esta calor"


	for i := 0; i < len(frase); i++ {
		if frase[i] == 'a' {
			continue
		}
	
		fmt.Printf("%c", frase[i])
	}

}
	