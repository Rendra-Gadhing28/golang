package main

import (
	"fmt"
)

func main(){
	mapEx()
	Nilai()
}

func mapEx(){
	m := make(map[string]string)
	m["nama"] = "Rendra Gadhing Pamungkas"
	m["kelas"] = "11 PPLG 3"

	fmt.Println("map : ", m)
	fmt.Println("nama :", m["nama"])
	fmt.Println("kelas :", m["kelas"])

	m["hobi"] = "fotografi dan berusaha lebih baik lagi"
	fmt.Println("hobi : ", m["hobi"])
}

func Nilai(){
	siswa := make(map[string]int)
	siswa["rendra"] = 90
	siswa["arka"] = 95
	siswa["ungaran"] = 100

	fmt.Println(siswa)
}