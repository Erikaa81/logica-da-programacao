//Leia 3 valores reais (A, B e C) e verifique se eles formam ou não um triângulo. 
//Em caso positivo, calcule o perímetro do triângulo e apresente a mensagem: Perimetro = XX.X    
//Em caso negativo, calcule a área do trapézio que tem A e B como base e C como altura, mostrando a mensagem :Area = XX.X

package main

import (
	"fmt"
)

func main() {

	a := 25.20
	b := 11.50
	c := 13.30

	if a < b+c && b < c+a && c < b+a {
		perimetroTriangulo := a + b + c
		fmt.Printf("Perimetro do triangulo = %.1f", perimetroTriangulo)

	} else {
		areaTrapezio := ((a + b) * c / 2)
		fmt.Printf("Area do trapezio = %.1f", areaTrapezio)

	}

}