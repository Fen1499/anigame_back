package pkg

import (
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func OpenConn() *sql.DB {
	println("TENTANDO CONECTAR")
	db, err := sql.Open("pgx", "postgres://postgres:postgres@localhost:5432/anigame")
	if err != nil {
		fmt.Println("ERRO: ", err)
	} else {
		fmt.Println("DE BOAS")
	}

	return db
}

func GetDrivers() []string {
	return sql.Drivers()
}
