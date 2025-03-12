package main

import (
	"log"

	"github.com/karthikbhandary2/social/internal/db"
	"github.com/karthikbhandary2/social/internal/env"
	"github.com/karthikbhandary2/social/internal/store"
)


const version = "0.0.1"

func main() {
	cfg := config{
		addr: env.GetString("ADDR", ":8082"),
		db: dbConfig{
			addr:         env.GetString("DB_ADDR", "postgres://Karthik:1234@localhost/social?sslmode=disable"),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
			maxIdleTime:  env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
		env: env.GetString("ENV", "development"),
	}

	db, err := db.New(cfg.db.addr, cfg.db.maxOpenConns, cfg.db.maxIdleConns, cfg.db.maxIdleTime)
	if err != nil {
		log.Panic(err)
	}

	defer db.Close()
	log.Println("Database connection pool established")

	store := store.NewStorage(db)
	app := &application{
		config: cfg,
		store:  store,
	}

	mux := app.mount()
	log.Fatal(app.run(mux))
}
