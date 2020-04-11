package database

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func Init() *gorm.DB {
	user, userErr := os.LookupEnv("DB_USER")
	pass, passErr := os.LookupEnv("DB_PASS")
	host, hostErr := os.LookupEnv("DB_HOST")
	dbName, dbNameErr := os.LookupEnv("DB_NAME")

	if !userErr {
		log.Fatal("DB User not specified in .env")
	}

	if !passErr {
		log.Fatal("DB Pass not specified in .env")
	}

	if !hostErr {
		log.Fatal("Host not specified in .env")
	}

	if !dbNameErr {
		log.Fatal("Db Name not specified in .env")
	}

	credentials := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local", user, pass, host, dbName)

	db, err := gorm.Open("mysql", credentials)

	if err != nil {
		log.Fatal(err)
	}

	return db
}
