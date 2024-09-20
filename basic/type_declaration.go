package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	// type biasanya digunakan untuk alias 
	type KTP string

	var MyKtp KTP = "1231321"
	fmt.Println(MyKtp)

	var NoKtp = "123123123"
	conversionNo := KTP(NoKtp)
	fmt.Println(conversionNo)
}