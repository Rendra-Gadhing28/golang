package crud

import (
	"endpoint/db"
	"fmt"
	"log"
)

type Bahasa struct{
	ID int `json:"id"`
	USERNAME string `json:"username"`
	DESKRIPSI string `json:"deskripsi"`
	KATEGORI string `json:"kategori"` 
}

func FetchData() ([]Bahasa, error){
	db := db.Database()
	q := "SELECT id,username, deskripsi, kategori FROM bahasa"
	data,err := db.Query(q)
	if err != nil {
		fmt.Println("Data tidak berhasil diambil")
	}
	defer data.Close()

	var users []Bahasa
	for data.Next(){
		var user Bahasa
		data.Scan(&user.ID, &user.USERNAME, &user.DESKRIPSI, &user.KATEGORI)
		if err != nil{
			log.Fatal(err)
		}
		users = append(users, user)
	}
	return users, nil
}