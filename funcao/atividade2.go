//Crie uma função que recebe uma string e um caractere, e apague todas as ocorrências desses caractere na string.

package main

import "fmt"

func main() {
	frase := "Esta calor"

	fmt.Printf("%s", removerCaracter(frase))

}

func removerCaracter(frase string) string {

	novaFrase := ""
	for i := 0; i < len(frase); i++ {
		if frase[i] != 'a' {
			novaFrase = novaFrase + string(frase[i])
		}
	}

	return string(novaFrase)

}
