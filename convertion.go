package main

import (
	"fmt"
	"strconv"
)

func main(){
	var text string = "123"
	number,err := strconv.Atoi(text)
	if err != nil {
		fmt.Println("pesan error :", err.Error())
	} else {
		fmt.Println("Angka : ", number)
	}

	booleanToString()
}

func booleanToString() {
	truth := true 
	str := strconv.FormatBool(truth)
	fmt.Println("Boolean ke string :", str)

	val, _ := strconv.ParseBool("true")
	fmt.Println("String ke boolean : ", val)
}