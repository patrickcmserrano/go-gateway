package fixture

import (
	"context"
	"database/sql"
	"io/fs"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"maragu.dev/migrate"
)

func Up(migrationsDir fs.FS) *sql.DB {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal(err)
	}

	if err := migrate.Up(context.Background(), db, migrationsDir); err != nil {
		return db
	}
	return db
}

func Down(db *sql.DB, migrationsDir fs.FS) {
	if err := migrate.Down(context.Background(), db, migrationsDir); err != nil {
		db.Close()
	}
}
