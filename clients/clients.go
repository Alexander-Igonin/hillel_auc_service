package clients

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Clients struct {
	Postgres *sql.DB
}

func NewClients() *Clients {
	connStr := "user=postgres dbname=postgres password=12345 host=localhost port=5432 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	return &Clients{
		Postgres: db,
	}
}
