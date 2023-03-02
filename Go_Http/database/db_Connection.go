package database

import (
	"fmt"
	"log"
	"os"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"github.com/kwamekyeimonies/Go-Http/models"
)

func Database_Connection() *pg.DB {
	opts := &pg.Options{
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Addr:     os.Getenv("DB_ADDRESS"),
		Database: os.Getenv("DB_DATABASE"),
	}



	var db *pg.DB = pg.Connect(opts)

	if db == nil {
		fmt.Println("Database Connection Failed.....")
		os.Exit(100)
	}

	log.Println("Database Connected Successfully.....")

	if err := createSchema(db); err != nil {
		log.Fatal(err)
	}

	return db
}

func createSchema(db *pg.DB) error {
	models := []interface{}{
		(*models.User)(nil),
		(*models.Product)(nil),
	}

	for _, model := range models {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
			IfNotExists: true,
		})
		if err != nil {
			return err
		}
	}

	return nil
}
