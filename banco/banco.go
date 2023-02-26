package banco

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Conexao() (*sql.DB, error) {

	strConexao := "USUARIODOBANCO:@SENHADOBANCO@/stock_sys?charset=utf8&parseTime=true&loc=Local"

	db, erro := sql.Open("mysql", strConexao)
	if erro != nil {
		fmt.Println("Erro de Conexão com o Banco")
		log.Fatal(erro)
	}

	if erro = db.Ping(); erro != nil {
		fmt.Println("Erro! A solicitação ping não pôde encontrar o host de destino.")
		return nil, erro
	}

	return db, nil

}
