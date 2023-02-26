package serve

import (
	"fmt"
	"log"
	"stock/banco"
)

// TABELA LOGIN DO BANCO
type Login struct {
	ID      int
	Nome    string
	Usuario string
	Senha   string
}

func CadUsuario() {

	//Abrir conexão com o banco.
	db, erro := banco.Conexao()
	if erro != nil {
		log.Fatal("Erro ao conectar com o banco de dados!")
	}
	defer db.Close()
	fmt.Println("Banco de Dados conectado!")

	//Query para cadastrar um novo usuario no banco
	var user Login
	fmt.Print("Nome: ")
	fmt.Scan(&user.Nome)
	fmt.Print("Usuario: ")
	fmt.Scan(&user.Usuario)
	fmt.Print("Senha: ")
	fmt.Scan(&user.Senha)

	prepare, erro := db.Prepare("insert into login(nome, usuario, senha) values(?,?,?)")
	if erro != nil {
		log.Fatal("Erro Prepare Statement!")
	}
	defer prepare.Close()

	insert, erro := prepare.Exec(user.Nome, user.Usuario, user.Senha)
	if erro != nil {
		log.Fatal("Erro para cadastrar um novo usuario!")
	}

	idInserido, erro := insert.LastInsertId()
	if erro != nil {
		log.Fatal("Erro ao retornar o último ID cadastrado!")
	}
	fmt.Printf("ID %d cadastrado!", idInserido)

}

func ListarUsuarios() {
	//Abrir conexao com o banco
	db, erro := banco.Conexao()
	if erro != nil {
		log.Fatal("Erro ao Conectar com o banco de dados!")
	}
	defer db.Close()

	//Query Listar Cadastros do Banco.
	linhas, erro := db.Query("select id, nome, usuario from login")
	if erro != nil {
		log.Fatal("Erro para listar os cadastros de usuários!")
	}
	defer linhas.Close()

	var users []Login

	for linhas.Next() {
		var user Login

		if erro := linhas.Scan(&user.ID, &user.Nome, &user.Usuario); erro != nil {
			log.Fatal("Erro para scanear os dados!")
			return
		}
		users = append(users, user)
	}

	fmt.Println(users)
}

func ListarCadastro() {
	db, erro := banco.Conexao()
	if erro != nil {
		log.Fatal("Erro, conexão com o banco de dados!")
	}
	defer db.Close()

	var ID int
	fmt.Print("Digite o ID do usuário: ")
	fmt.Scan(&ID)

	linhas, erro := db.Query("select id, nome, usuario from login where id = ?", ID)
	if erro != nil {
		fmt.Print("Erro para scanear os dados do ID: ", ID)
		return
	}
	//defer linhas.Close()

	var users Login

	if linhas.Next() {
		if erro := linhas.Scan(&users.ID, &users.Nome, &users.Usuario); erro != nil {
			log.Fatal("Erro no Scan!")
		}

		if ID == 0 {
			log.Fatal("ID não cadastrado!")
		}
	}

	fmt.Println("ID:", users.ID, "Nome:", users.Nome, "Usuário:", users.Usuario)

}

func AtualizarCadastro() {

	db, erro := banco.Conexao()
	if erro != nil {
		log.Fatal("Erro ao conectar com o banco - Func AtualizarCadastro")
	}
	defer db.Close()

	var ID int
	fmt.Println("Digite o ID do Usuário: ")
	fmt.Scan(&ID)

	linhas, erro := db.Query("select * from login where id = ?", ID)
	if erro != nil {
		log.Fatal("Erro no Select da função - AtualizarCadastro")
	}
	//NÃO PRECISA PASSAR UM SLICE, POIS ESTAMOS TRAZENDO APENAS 1 CADASTRO.
	var users Login

	if linhas.Next() {
		if erro := linhas.Scan(&users.ID, &users.Nome, &users.Usuario, &users.Senha); erro != nil {
			log.Fatal("Erro no Scan função - AtualizarCadastro")
		}
	}

	if ID == 0 {
		log.Fatal("ID não cadastrado!")
		return
	}

	fmt.Println("ID:", users.ID, "Nome:", users.Nome, "Usuário:", users.Usuario)

	fmt.Println("1º Atualizar Nome.")
	fmt.Println("2º Atualizar Usuário.")
	fmt.Println("3º Atualizar Senha.")
	var resp int
	fmt.Scan(&resp)

	//variável caduser usada para receber os valores passado pelo usuário do sistema
	var caduser Login
	switch resp {
	case 1:

		fmt.Print("Digite o nome: ")
		fmt.Scan(&caduser.Nome)

		statement, erro := db.Prepare("update login set nome = ? where id = ?")
		if erro != nil {
			log.Fatal("Erro ao tentar atualizar o nome no ID: ", ID)
		}
		defer statement.Close()

		if _, erro := statement.Exec(caduser.Nome, ID); erro != nil {
			log.Fatal("Erro para executar a atualização do campo nome!")
		}

		fmt.Print("Dados Atualizados!")
	case 2:

		fmt.Print("Digite o usuario: ")
		fmt.Scan(&caduser.Usuario)

		statement, erro := db.Prepare("update login set usuario = ? where id = ?")
		if erro != nil {
			log.Fatal("Erro ao tentar atualizar o usuário no ID: ", ID)
		}
		defer statement.Close()

		if _, erro := statement.Exec(caduser.Usuario, ID); erro != nil {
			log.Fatal("Erro para executar a atualização do campo usuário!")
		}

		fmt.Print("Dados Atualizados!")
	case 3:

		fmt.Print("Digite a nova senha: ")
		fmt.Scan(&caduser.Senha)

		statement, erro := db.Prepare("update login set senha = ? where id = ?")
		if erro != nil {
			log.Fatal("Erro ao tentar atualizar o usuário no ID: ", ID)
		}
		defer statement.Close()

		if _, erro := statement.Exec(caduser.Senha, ID); erro != nil {
			log.Fatal("Erro para executar a atualização do campo senha!")
		}

		fmt.Print("Dados Atualizados!")

	default:
		fmt.Println("Digite uma das opções!")
	}
}

func DeletarCadastro() {
	db, erro := banco.Conexao()
	if erro != nil {
		log.Fatal("Erro na conexão com o banco.")
	}
	defer db.Close()

	var ID int
	fmt.Println("Digite o ID do usuário: ")
	fmt.Scan(&ID)

	linhas, erro := db.Query("select id, nome, usuario from login where id = ?", ID)
	if erro != nil {
		log.Fatal("Erro para trazer os dados do ID ", ID)
	}

	var users Login

	if linhas.Next() {
		if erro := linhas.Scan(&users.ID, &users.Nome, &users.Usuario); erro != nil {
			log.Fatal("Erro nas linhas do select.")
		}
	}

	if ID == 0 {
		fmt.Println("ID não cadastrado!")
	}

	statement, erro := db.Prepare("delete from login where id = ?")
	if erro != nil {
		log.Fatal("Erro no Prepare DELETE.")
	}
	defer statement.Close()

	if _, erro := statement.Exec(ID); erro != nil {
		log.Fatal("Erro no Exec da função DELETE")
	}

	fmt.Println("Cadastro Deletado.")

}
