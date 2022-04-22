//Escrever um algoritmo para salvar informações pertinentes a empresas. Elas devem ser cadastradas com cnpj, nome, data de inauguração, valor de faturamento mensal e endereço. 
//O endereço deve ser composto por: rua, bairro, cep, número, estado e país. Algumas validações devem ser feitas, como:
//CNPJ tem que ser válido
//data de inauguração não deve ser superior a data atual


package estruturas

import (
	"fmt"
	"time"
)

type Empresa struct {
	cnpj                   string
	nome                   string
	dataDeInauguracao      string
	valorFaturamentoMensal float64
	endereco               Endereco
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
	var empresa Empresa
	var endereco Endereco

	empresa.cnpj = "33333333333333"
	empresa.nome = "Padaria"
	empresa.dataDeInauguracao = "2022-03-12"
	empresa.valorFaturamentoMensal = 78899

	if len(empresa.cnpj) != 14 {
		fmt.Printf("CNPJ invalido\n")
		return
	}
	if empresa.dataDeInauguracao >= time.Now().String() {

		fmt.Printf("Data invalida\n")
		return
	}

	endereco.rua = " Rua da Esperanca"
	endereco.bairro = "Centro"
	endereco.cep = 1877373
	endereco.numero = 231
	endereco.estado = "sp"
	endereco.pais = "Brasil"
	empresa.endereco = endereco

	fmt.Printf("CNPJ %s, Nome %s, data %v, valor %v \nEndereço: %s, %s, Cep %v, numero %v,  %s, %s \n", empresa.cnpj, empresa.nome, empresa.dataDeInauguracao, empresa.valorFaturamentoMensal,
		endereco.rua, endereco.bairro, endereco.cep, endereco.numero, endereco.estado, endereco.pais)
	fmt.Print(empresa.endereco.cep)
}
