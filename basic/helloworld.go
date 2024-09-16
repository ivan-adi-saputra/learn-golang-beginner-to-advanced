package main

import "fmt"

func main() {
	fmt.Println("Hello World")
}

// note: dalam 1 package tidak boleh menggunakan nama func yang sama. 
// ex: 
// file helloworld.go memiliki func dengan nama main
// file sample.go memiliki func dengan nama main
// maka nanti akan menyebabkan error