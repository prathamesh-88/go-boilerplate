package database

import (
	"database/sql"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

var Db *sql.DB

func Connect() {
	var err error

	connStr := "postgres://postgres:tamanekar@localhost:5432/go_auth?sslmode=disable"
	Db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error while connecting to database", err.Error())
	}

	log.Println("Connected to database!")
	driver, err := postgres.WithInstance(Db, &postgres.Config{})
	if err != nil {
		log.Fatal("Error while initializing migrations", err)
	}
	m, err := migrate.NewWithDatabaseInstance("file://migrations", "go_auth", driver)
	if err != nil {
		log.Fatal("Error while running migrations: ", err)
	}

	m.Up()
	log.Println("Migrations were sucessful!")

}

func RunQuery(query string, args ...any) *sql.Rows {
	// Add any validation if required
	rows, err := Db.Query(query, args...)

	if err != nil {
		log.Fatal(err.Error())
	}

	return rows
}

func CloseConnection() {
	err := Db.Close()
	if err != nil {
		log.Println("Error while closing DB connection: ", err.Error())
	}
}
