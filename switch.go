package main
import ( 
	"fmt"

)
func main(){
	Switch()
	Hari()
}

func Switch(){
	angka := 10
	switch angka {
	case 1 :
		fmt.Println("tampilkan angka 1")
	case 2 :
		fmt.Println("tampilkan angka 2")
	case 3 :
		fmt.Println("tampilkan angka 3")

	default:
		fmt.Println("angka anomali")
	}
}

func Hari(){
	hari := "Senin"

	switch hari {
	case "Sabtu", "Minggu":
		fmt.Println("Weekend")
	case "Senin", "Selasa", "Rabu", "Kamis", "Jumat":
		fmt.Println("Weekday")
	default:
		fmt.Println("Hari tidak valid")
	}
	
}