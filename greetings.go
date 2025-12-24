package greetings

import "fmt"

func Hello(name string) string{
	pesan := fmt.Sprintf("Hello %v. Welcome", name)
	return pesan
}