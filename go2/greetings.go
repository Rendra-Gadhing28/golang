package GO2
import (
	"fmt"
	"errors"
	"math/rand"
)

func Hello(name string) (string,error) {

	if name == "" {
		return "", errors.New("Nama masih kosong njr") //handle error 
	}
	form := Random() // sebuah fungsi random
	pesan :=fmt.Sprintf(form, name)
	return pesan, nil	
}

func Ho(name []string)(map[string] string, error){
	mesages := make(map[string]string) // make an array object
	for _, name := range name {
		mesage, err := Hello(name)
		if err != nil {
			return nil, err
		}
		mesages[name] = mesage
	}
	return mesages, nil
}

func Random()string{
	format:= []string{
		"Hai, %v. Tadaima",
		"Semoga harimu baik baik saja, %v",
		"Semangatt, %v",
	}

	return format[rand.Intn(len(format))]
}