//Escrever um algoritmo que vai armazenar os clientes das empresas cadastradas na atividade 1. Os clientes devem ser registrados com CPF, nome, endereço, média de consumo mensal e email. 
//O endereço deve ser composto pelos mesmos dados da questão anterior. A validação necessária é para o CPF.

package main

import (
	"fmt"
)

type Cliente struct {
	CPF                  string
	nome                 string
	mediaDeConsumoMensal float64
	endereco             Endereco
}
type Endereco struct {
	rua    string
	bairro string
	cep    int
	numero int
	estado string
	pais   string
}

func main() {
	var cliente Cliente
	var endereco Endereco

	cliente.CPF = "12333333333"
	cliente.nome = "Carlos Silva"
	cliente.mediaDeConsumoMensal = 1000.00

	if len(cliente.CPF) != 11 {
		fmt.Printf("CPF invalido")
		return
	}

	endereco.rua = " Rua da Esperanca"
	endereco.bairro = "Centro"
	endereco.cep = 1877373
	endereco.numero = 231
	endereco.estado = "SP"
	endereco.pais = "Brasil"
	cliente.endereco = endereco

	fmt.Printf("CPF %s, Nome %s, valor %v \nEndereço: %s, %s, Cep %v, numero %v, %s, %s \n", cliente.CPF, cliente.nome, cliente.mediaDeConsumoMensal,
		endereco.rua, endereco.bairro, endereco.cep, endereco.numero, endereco.estado, endereco.pais)
	fmt.Print(cliente.endereco.cep)
}
