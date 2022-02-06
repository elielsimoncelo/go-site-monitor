package main

import (
	"fmt"
	"reflect"
)

func main() {
	var nome string = "Jose"
	var idade = 18 // inferencia de tipo
	//var versao float32 = 1.1
	//var versao = 1.1 // inferencia de tipo / variavel
	versao := 1.1 // auto inferencia de tipo e declaracao de variavel

	fmt.Println("Ola", nome, "sua idade e", idade)
	fmt.Println("Este programa esta na versao: ", versao)
	fmt.Println("O tipo da variavel versao e", reflect.TypeOf(versao))

	fmt.Println("[1] Iniciar o monitoramento")
	fmt.Println("[2] Exibir logs")
	fmt.Println("[0] sair do programa")

	var comando int

	//fmt.Scanf("%d", &comando) // definindo com o modificador %d para definir o tipo de entrada
	fmt.Scan(&comando) // sem a necessidade do modificador auto inferencia do tipo

	fmt.Println("O endereco da variavel comando e", &comando)
	fmt.Println("O comando escolhido foi", comando)
}
