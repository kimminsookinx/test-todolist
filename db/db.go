/*
	does not require persistence, so we implement it to be simple
	https://github.com/Massad/gin-boilerplate/blob/master/db/db.go
*/

package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-gorp/gorp"
	_ "github.com/go-sql-driver/mysql"
)

type DB struct {
	*sql.DB
}

var db *gorp.DbMap

func Init() {
	dbinfo := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_ADDRESS"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
	fmt.Print(dbinfo + "\n")
	//dbinfo := "user:user@tcp(127.0.0.1:3006)/"
	var err error
	db, err = Connect(dbinfo)
	if err != nil {
		log.Fatal(err)
	}

}

func Connect(datasSourceName string) (*gorp.DbMap, error) {
	db, err := sql.Open("mysql", datasSourceName)
	if err != nil {
		fmt.Printf("open error")
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		fmt.Printf("ping error")
		return nil, err
	}
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{}}
	return dbmap, nil
}

func GetDB() *gorp.DbMap {
	return db
}
