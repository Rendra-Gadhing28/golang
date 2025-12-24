package main
import "fmt"

func main(){
	names := [...]string{
		"rendra","gadhing","pamungkas","arka","mudya","ezar",}
	slice:= names[:]//diambil dari index ke 3 dan diakhiri sebelum nilai high

	fmt.Println(slice)

}