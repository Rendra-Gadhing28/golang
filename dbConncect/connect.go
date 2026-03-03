package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func connDB() *sql.DB {
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
func main() {

	db := connDB()
	defer db.Close()
	// addData(db)
	showData(db)
	// updateData(db)
	deleteData(db)
}

func addData(db *sql.DB) {
	q := "INSERT INTO bahasa (username, deskripsi, kategori) VALUES(?,?,?)"
	result, err := db.Exec(q, "PHP", "Bahasa pemrograman yang mudah dipelajari, frameworknya developer friendly karena dokumentasi lengkap dan banyak tools", "mid")
	if err != nil {
		log.Fatalln("tipe data tidak sama woi, gagal insert", err)
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Data berhasil ditambahkan dengan Id ke -", lastID)
}

func showData(db *sql.DB) {
	var query string = "SELECT id, username, deskripsi, kategori FROM bahasa ORDER BY id DESC"
	result, err := db.Query(query)
	if err != nil {
		fmt.Println("Tidak ada data")
		log.Fatal(err)
	}
	defer result.Close()

	for result.Next() {
		var id int
		var username, deskripsi, kategori string

		err := result.Scan(&id, &username, &deskripsi, &kategori)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("ID", id)
		fmt.Println("Username:", username)
		fmt.Println("Deskripsi:", deskripsi)
		fmt.Println("Kategori:", kategori)
	}
}

func updateData(db *sql.DB){
	q := "UPDATE bahasa SET deskripsi = ? WHERE id = ?"
	result, err := db.Exec(q, "Bahasa pemrograman yang gercep cepat, efisien dan mudah untuk pemula", 1000)
	if err != nil {
		log.Fatal(err)
	}
	defer result.RowsAffected()
	fmt.Println("baris yang diubah", result)
	fmt.Println("data berhasil di update")
} 

func deleteData(db *sql.DB){
	q := "Delete FROM bahasa WHERE id = ?"
	result, err := db.Exec(q, 1002)
	defer result.RowsAffected()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("data berhasil di hapus")
}
