package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	envErr := godotenv.Load()
	if envErr != nil {
		panic("Erro ao carregar variáveis de ambiente")
	}
	env := os.Getenv("MSSQL_DB_SERVER")
	fmt.Println(env)
}
