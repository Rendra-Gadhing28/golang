package crud
import (
	"fmt"
	"log"
	"DBCONNCECT/db"
)

func AddData() {
	db:= db.ConnDB()
	q := "INSERT INTO bahasa (username, deskripsi, kategori) VALUES(?,?,?)"
	result, err := db.Exec(q, "Python", "Bahasa pemrograman yang mudah dipelajari, dan sangat banyak digunakan dalam teknologi AI", "high")
	if err != nil {
		log.Fatalln("tipe data tidak sama woi, gagal insert", err)
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Data berhasil ditambahkan dengan Id ke -", lastID)
}