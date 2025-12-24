package main
import (
    "fmt"

    "sinau_go/greetings"
)

func main(name string) string{
	pesan := greetings.Hello("Rendra gadhing pamungkas")
	fmt.Println(pesan)
}	


