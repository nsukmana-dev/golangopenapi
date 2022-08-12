package app

import (
	"database/sql"
	"golangopenapi/helper"
	"time"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/openapigolang")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(15)
	db.SetMaxOpenConns(60)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
