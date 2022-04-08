package db

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var connectionPool *pgxpool.Pool

func Connect() {
	pool, err := pgxpool.Connect(context.Background(), os.Getenv("BASIS_DB_URL"))
	if err != nil {
		log.Fatalln("Unable to connect to database: ", err)
	}
	connectionPool = pool
}

func Disconnect() {
	connectionPool.Close()
}
