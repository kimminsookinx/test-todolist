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
	dbinfo := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		os.Getenv("TODO_DB_USER"), os.Getenv("TODO_DB_PASS"),
		os.Getenv("TODO_DB_ADDRESS"), os.Getenv("TODO_DB_PORT"), os.Getenv("TODO_DB_NAME"))
	var err error
	db, err = connect(dbinfo)
	if err != nil {
		log.Fatal(err)
	}

}

func connect(datasSourceName string) (*gorp.DbMap, error) {
	db, err := sql.Open("mysql", datasSourceName)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{}}
	return dbmap, nil
}

func GetDB() *gorp.DbMap {
	return db
}
