package main
import (
	"fmt"
)


type Siswa struct {
	nama string
	nilai []int
}

func upgradeNilai(s *Siswa) {
	for i := range s.nilai{
		if s.nilai[i] < 100 {
			s.nilai[i] += 5
			continue
		} else {
			s.nilai[i] = 100
			break
		}
	}
}

func (s *Siswa) RataRata() float32 {
	total := 0
	for i := range s.nilai{
		total += s.nilai[i]
	}
	return float32(total) / float32(len(s.nilai))
}


func main(){
	siswa1 := Siswa{
		nama: "Cikal",
		nilai: []int{90,85,90,80,85,100},
	}
	upgradeNilai(&siswa1)
	fmt.Println("Nama Siswa: ",siswa1.nama,"\n","Nilai Siswa: ",siswa1.nilai)
	siswa1.RataRata()
	fmt.Println(`Rata Rata siswa nama`,siswa1.nama, `adalah`, siswa1.RataRata())
}