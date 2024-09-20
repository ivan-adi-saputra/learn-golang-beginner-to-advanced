package main

import "fmt"

func main() {
	// pertama
	var names [3]string

	names[0] = "Ivan"
	names[1] = "Adi"
	names[2] = "Saputra"

	fmt.Println(names) // output: [Ivan Adi Saputra]
	fmt.Println(names[1]) // output: Adi
	fmt.Println(names[2]) // output: Saputra

	// kedua
	// ... digunakan untuk isi array yang tidak terhingga
	// ... harus di deklarasi isi
	values := [...]int{
		10,
		20,
        30,
        30,
        30,
	}
	fmt.Println(values) // output: [10 20 30]
	fmt.Println(len(values))
}