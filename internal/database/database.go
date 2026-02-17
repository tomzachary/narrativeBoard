package database

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

func Run(query func(conn *pgx.Conn) (any, error)) (any, error) {
	conn, err := pgx.Connect(context.Background(), os.Getenv("POSTGRES_URL"))
	if err != nil {
		fmt.Printf("Error while connection to the database: %v\n", err)
		return nil, err
	}
	defer conn.Close(context.Background())

	fmt.Println("Connection Success!")
	res, err := query(conn)
	return res, err

}
