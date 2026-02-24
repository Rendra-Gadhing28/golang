package main
import (
	"fmt"
	"time"
)

func main(){

	start := time.Now()
	
	loop := 1000000

	 for i := 0; i < loop; i++ {
		fmt.Println("Hai Cikal", "urutan yang ke - ", i)
	 }
	 elapsed := time.Since(start) // waktu selesai
	 fmt.Println("Waktu eksekusi:", elapsed)
}