package db

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

func ConnDB() *sql.DB {
	dsn := "root:@tcp(127.0.0.1:3306)/bahasa"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Mysql server tidak aktif. Jalankan di laragon terlebih dahulu!")
	}
	fmt.Println("berhasil terhubung ke mysql server")
	return db
}

