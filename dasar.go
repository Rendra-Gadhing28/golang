package main
import "fmt"

func main(){
	fmt.Println("Hello, world"[0])

	nama := "arka mudya"
	fmt.Println(nama)
	//tipe data int ada 8,16,32,64 
	nama = "Ezar hama"
	fmt.Println(nama)

	var(
		first = "Rendra"
		last = "Gadhing"

		a = first[0]
		b = last[0]
		ss = string(a)
	)
	fmt.Println(first,last, a, b,ss)

	const pertama = "Belajar golang"
	const kedua = "Golang oke"

	fmt.Println(pertama, kedua)
	//uint 8, 16,32,64

	var nilai32 int32 = 32768
	var nilai64 int64 = int64(nilai32)
	var nilai16 int16 = int16(nilai32)
	var nilai8 int8 = int8(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai16)
	fmt.Println(nilai8)
	//tipe data alias rune int32
	//tipe data int = minimal int32
	//uint = minimal uint32
}