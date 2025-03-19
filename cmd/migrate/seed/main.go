package main

import (
	"log"

	"github.com/karthikbhandary2/Social/internal/db"
	"github.com/karthikbhandary2/Social/internal/env"
	"github.com/karthikbhandary2/Social/internal/store"
)

func main() {

	addr := env.GetString("DB_ADDR", "postgres://Karthik:1234@localhost/social?sslmode=disable")
	conn, err := db.New(addr, 3, 3, "15m")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	store := store.NewStorage(conn)
	db.Seed(store)
}