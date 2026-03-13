package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"endpoint/crud"
)

func Handler(){
	http.HandleFunc("/bahasa", Cek)

	fmt.Println("server jalan di http://localhost:8080")
	http.ListenAndServe(":8000", nil)
}

func Cek(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintln(w, "Hello world!")
}

func UserHandler(w http.ResponseWriter, r *http.Request){
	if r.Method == "GET" {
		fmt.Println(w, "ini method GET request")
		//manggil function crud dari get data
		crud.FetchData()
	} else if r.Method == "POST" {
		fmt.Println(w, "ini method POST")
		//panggil method addData di crud
	} else{
		http.Error(w, "method tidak diizinkan", http.StatusMethodNotAllowed)
	}
}




func GetUser(w http.ResponseWriter, r *http.Request){
	if r.Method == "GET" {
		data,err := crud.FetchData()
	if err != nil {
		http.Error(w, "gagal mengambild data", 
	http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
	}
}