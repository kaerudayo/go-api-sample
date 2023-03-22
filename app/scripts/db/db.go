package main

import (
	"fmt"
	"log"
	"os"

	"github.com/api-sample/app/pkg/db"
	"github.com/api-sample/app/scripts/db/seed"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var databaseName = os.Getenv("DB_DATABASE")

func main() {
	args := os.Args
	for _, command := range args[1:] {
		switch command {
		case "create":
			createDatabase()
		case "migrate":
			migrateDatabase()
		case "drop":
			dropDatabase()
		case "seed":
			seedDatabase()
		}
	}
}

func createDatabase() {
	db.Init(false)
	err := db.DB.Exec("CREATE SCHEMA IF NOT EXISTS " + databaseName).Error
	if err != nil {
		fmt.Printf("[createDatabase]... %s", err)
	}
}

func migrateDatabase() {
	mySQLDB := db.Init(true)

	driver, err := mysql.WithInstance(mySQLDB, &mysql.Config{})
	if err != nil {
		log.Fatal("mysql.WithInstance:", err)
		return
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://scripts/db/migrations",
		"mysql",
		driver,
	)

	if err != nil {
		log.Fatal("migrate.NewWithDatabaseInstance:", err)
	}

	if err = m.Up(); err != nil && err.Error() != "no change" {
		log.Fatal("m.Up:", err)
	}
}

func dropDatabase() {
	db.Init(false)
	err := db.DB.Exec("DROP SCHEMA IF EXISTS " + databaseName).Error
	if err != nil {
		fmt.Printf("[dropDatabase]... %s", err)
	}
}

func seedDatabase() {
	seedDB := db.Init(true)
	seed.DefaultSeed(seedDB)
}
