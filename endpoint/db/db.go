package db

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

func Database() *sql.DB{
	dsn := "root:@tcp(127.0.0.1:3306)/bahasa"
	db, err := sql.Open("mysql",dsn )
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Server mysql tidak aktif, coba lagi")
	}

	fmt.Println("databases berhasil connect")
	return db
}