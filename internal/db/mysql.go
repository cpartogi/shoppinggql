package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" // mysql driver
	"github.com/joho/godotenv"
)

// mysql variable
var Db *sql.DB

//init db session
func InitDB() {
	var envs map[string]string
	envs, err := godotenv.Read(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbhost := envs["DBHOST"]
	dbport := envs["DBPORT"]
	dbname := envs["DBNAME"]
	dbuser := envs["DBUSER"]
	dbpassword := envs["DBPASSWORD"]

	mysqlInfo := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbuser, dbpassword, dbhost, dbport, dbname)

	db, err := sql.Open("mysql", mysqlInfo)
	if err != nil {
		log.Panic(err)
	}

	if err = db.Ping(); err != nil {
		log.Panic(err)
	}
	Db = db

	log.Println("Connection to mysql was sucessfull!!")
}
