package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type databaseLogin struct {
	UsrName string
	Psw     string
	Addr    string
	DbName  string
}

func DatabaseGenerator() databaseLogin {
	err := godotenv.Load() // 载入godotenv
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	/* .env 内容：
	database_usrname=root
	database_psw=xxxxxx
	database_addr=localhost
	database_dbname=xeth
	*/
	dbUsrName := os.Getenv("database_usrname")
	dbPsw := os.Getenv("database_psw")
	dbAddr := os.Getenv("database_addr")
	dbName := os.Getenv("database_dbname")
	databaseConfig := databaseLogin{UsrName: dbUsrName, Psw: dbPsw, Addr: dbAddr, DbName: dbName}
	log.Print(databaseConfig)
	return databaseConfig
}
