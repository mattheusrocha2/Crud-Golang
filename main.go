package main

import (
	"fmt"
	"stock/serve"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("Sistema de Login")

	fmt.Println("Selecione uma opção: ")
	fmt.Println("1º Cadastrar Usuário. ")
	fmt.Println("2º Listar todos Usuários.")
	fmt.Println("3º Listar um Usuário.")
	fmt.Println("4º Atualizar Cadastro de Usuário.")
	fmt.Println("5º Deletar Cadastro de Usuário.")
	var resp int
	fmt.Scan(&resp)

	if resp == 1 {
		serve.CadUsuario()
	} else if resp == 2 {
		serve.ListarUsuarios()
	} else if resp == 3 {
		serve.ListarCadastro()
	} else if resp == 4 {
		serve.AtualizarCadastro()
	} else if resp == 5 {
		serve.DeletarCadastro()
	} else {
		fmt.Println("Digite umas das opções acima!")
	}
}
