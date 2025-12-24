package main
import "fmt"

func main() {
	var names [3]string
	names[0]= "rendra"
	names[1] = "backend"
	names[2] = "developer"
	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [...]int{
		90,80,70,80,90,90,
	}
	fmt.Println(len(values))

	
} 