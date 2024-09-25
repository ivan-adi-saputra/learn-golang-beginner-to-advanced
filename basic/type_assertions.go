package main

import "fmt"
// type assertions biasanya digunakan ketika ingin merubah type data
// biasanya kalo ketemu dengan type data interface kosong / any 

func typeAssertions() any {
	return "Ok"
}

func main() {
	types := typeAssertions()
	switch val := types.(type){
		case string:
			fmt.Println("String", val)
			break
		case int: 
			fmt.Println("int", val)
			break
		default:
			fmt.Println("Unknow")
	}
}