package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5"
	"os"
)

const (
	host     = "localhost"
	port     = "8080"
	user     = "admin"
	password = "password"
	dbname   = "json"
)

func main() {
	json_data := "{\"name\":\"Jopa\"}"
	urlExample := "postgres://" + user + ":" + password + "@" + host + ":" + port + "/" + dbname
	conn, err := pgx.Connect(context.Background(), urlExample)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	/*	var name string
		var weight int64*/
	/*err = conn.QueryRow(context.Background(), "select name, weight from widgets where id=$1", 42).Scan(&name, &weight)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}*/
	//conn.
	conn.Exec(context.Background(), "INSERT INTO cache (user_data) values ($1)", json_data)
	/*fmt.Println(name, weight)*/
}
