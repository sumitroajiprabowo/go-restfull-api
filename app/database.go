package app

import (
	"database/sql"
	"time"

	"github.com/sumitroajiprabowo/go-restfull-api/helper"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "root:kmzway87aa@tcp(localhost:3306)/golang_restfull_api")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}