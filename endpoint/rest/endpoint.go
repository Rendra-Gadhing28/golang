package rest

import(
	"net/http"
	"endpoint/handler"
)

func SetupRoute(){
	http.HandleFunc("/bahasa", handler.GetUser)
	// http.HandleFunc("/bahasa/add")
	http.ListenAndServe(":8000", nil)
}