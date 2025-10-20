package main

import (
	"database/sql"
	"log"

	"github.com/Xayz-X/ecom/cmd/api"
	"github.com/Xayz-X/ecom/config"
	"github.com/Xayz-X/ecom/db"
	"github.com/go-sql-driver/mysql"
)

func main() {
	// entry point of our app

	// create a new storage
	db, err := db.NewMySQLStorage(mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Addr:                 config.Envs.DBAddress,
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})
	if err != nil {
		log.Fatal("Database connection failed", err)
	}

	// ping to the database
	initStorage(db)

	server := api.NewAPIServer(":8080", db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal("Database ping failed", err)
	}

	log.Println("DB sucessfully connected")
}
