package main //main utama
import (
	"fmt"
	"log"
	GO2 "example.com/greetings"
)

func main(){
	//
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{
		"Cikal", "Prameswari", "Putri",
	}

	nama, err := GO2.Ho(names)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(nama)
}