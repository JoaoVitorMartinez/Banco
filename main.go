package main

import (
	"fmt"

	".Banco/cliente"
	".Banco/conta"
)

func main() {
	fmt.Println("Starting...")

	Teste := cliente.Cliente{Nome: "Joao", Cpf: "15615615", Telefone: "98888-8888"}
	Teste2 := cliente.Cliente{Nome: "Joao", Cpf: "15615615", Telefone: "98888-8888"}

	fmt.Println(conta.Conta{NumeroConta: 0001, Agencia: 001, Titular: Teste})

	zconta := conta.Conta{NumeroConta: 0001, Agencia: 001, Titular: Teste}
	outraConta := conta.Conta{NumeroConta: 0001, Agencia: 001, Titular: Teste2}

	fmt.Println(zconta.Transferir(199.00, &outraConta))
	fmt.Println(outraConta)
	fmt.Println(zconta)

	teste3 := new(cliente.Cliente)
	fmt.Println(teste3)
}
