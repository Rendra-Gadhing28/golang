package crud

import (
	"fmt"
	"log"
	"DBCONNCECT/db"
)

func ShowData() {
	db := db.ConnDB()
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
