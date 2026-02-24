package main

import (
	"fmt"
)

type base struct{
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

type container struct{
	base
	str string
}

func main(){
	co := container{
		base: base{
			num: 5,
		},
		str: "Cikal",
	}

	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)
	fmt.Println(co.base.num)
	fmt.Println(co.describe())
	

	type describes interface{
		describe() string
	}

	var s describes = co
	fmt.Println(s.describe())
}