package crud
import (
	
	"fmt"
	"log"
	"DBCONNCECT/db"
	
)
func UpdateData(){
	db:= db.ConnDB()
	q := "UPDATE bahasa SET deskripsi = ? WHERE id = ?"
	result, err := db.Exec(q, "Bahasa pemrograman yang gercep cepat, efisien dan mudah untuk pemula, jossoke", 1000)
	if err != nil {
		log.Fatal(err)
	}
	defer result.RowsAffected()
	fmt.Println("baris yang diubah", result)
	fmt.Println("data berhasil di update")
} 
