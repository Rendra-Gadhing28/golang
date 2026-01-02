package main
import "fmt"

func main(){
	names := [...]string{
		"rendra","gadhing","pamungkas","arka","mudya","ezar",}
	slice:= names[1:4]//diambil dari index ke 3 dan diakhiri sebelum nilai high

	fmt.Println(len(slice))

	fmt.Println(cap(slice))

	tambah := append([]int{
		1,2},3,4,)
	fmt.Println(tambah)

	s:= make([]int, 5,10,)
	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Println(cap(s))


	s = append(s,10,20,30,40,50,60,70,)
	fmt.Println("s setelah append",s)

	slice2 := make([]int,3)
	slice3 := copy(slice2,s,)
	fmt.Println("copy :", slice2)
	fmt.Println("jumlah data yang dicopy :", slice3)

	angka := []int{
		1,2,3,4,5,
	}

	slice4 := angka[1:4]
	fmt.Println("slice 4 :", slice4)
}