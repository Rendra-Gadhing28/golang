package crud
import (
	"fmt"
	"log"
	"DBCONNCECT/db"
)

func DeleteData(){
	db:= db.ConnDB()
	q := "Delete FROM bahasa WHERE id = ?"
	result, err := db.Exec(q, 1002)
	defer result.RowsAffected()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("data berhasil di hapus")
}
