package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type databaseLogin struct {
	usrName string
	psw     string
	addr    string
	dbName  string
}

func init() {
	err := godotenv.Load() // 载入 godotenv
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbUsrName := os.Getenv("database_usrname")
	dbPsw := os.Getenv("database_usrname")
	dbAddr := os.Getenv("database_addr")
	dbName := os.Getenv("database_dbname")
	databaseConfig := databaseLogin{usrName: dbUsrName, psw: dbPsw, addr: dbAddr, dbName: dbName}
	log.Print(databaseConfig)
}
