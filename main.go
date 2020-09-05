package main

import (
	"context"
	"fmt"
	"os"

	"github.com/hanami125/todo-app/src/lib/mysql"
)

func main() {
	ctx := context.Background()

	dbConn, err := mysql.NewSQLClient(ctx, &mysql.DBConfig{
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Database: os.Getenv("DB_DATABASE"),
	})
	if err != nil {
		panic("fails db connection")
	}

	fmt.Printf("success db connection %v", dbConn)
}
