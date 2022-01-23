package main

import (
	"database/sql"
	"log"
	"os"

	"tmi-gin/api"
	db "tmi-gin/db/sqlc"

	_ "github.com/lib/pq"
)

func main() {

	dsn := os.Getenv("DATABASE_URL")

	if dsn == "" {
		dsn = "postgresql://postgres:12345678@localhost:5432/tmi-gin?sslmode=disable"
	}

	conn, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	defer conn.Close()

	store := db.NewStore(conn)
	server, err := api.NewServer(api.Config{}, store)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	err = server.Start(":" + port)

	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
