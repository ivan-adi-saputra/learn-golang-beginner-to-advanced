package main

import "fmt"

func main() {
	// variabel
	// pertama
	var name string
	name = "Ivan"
	fmt.Println("Hello, my name is", name)
	name = "Saputra"
	fmt.Println("Hello, my name is", name)

	// kedua
	name1 := "Ivan"
	fmt.Println("Hello, my name is", name1)
	name1 = "Saputra" // ketika name sudah di deklarasi tidak boleh menggunakan := karena di anggap membuat variabel baru
	fmt.Println("Hello, my name is", name1)

	// ketiga
	var (
		firstName = "Ivan"
        lastName = "Saputra"
	)
	fmt.Println("Hello, my name is", firstName, lastName)
	// notes: ketika membuat variabel itu wajib digunakan karena ketika tidak digunakan maka akan terjadi error

	// constant 
	// penggunaan sama seperti variabel tetapi dengan const
	const name2 = "Ivan Saputra"
	fmt.Println("Hello, my name is", name2)
}