package main

import "fmt"

// pointer: kemampuan membuat references ke lokasi data di memory yang sama, tanpa menduplikasikan data yang sudah ada
// contoh: default golang
// address1 := "ivan", "adi", "saputra" -> reprentasikan sebagai value
// address2 := adress1 -> sebenarnya cuma copy dari address 1
// address2 = "ivan", "adi" -> ketika dirubah maka value di address 1 tidak akan ikut berubah

// untuk menggunakan pointer menggunakan &
// contoh: pointer
// address1 := "ivan", "adi", "saputra" -> reprentasikan sebagai value
// address2 := &adress1 -> merepresentasikan sebagai address1
// address2 = "ivan", "adi" -> ketika dirubah maka value di address 1 akan ikut berubah

// operator *
// Saat kita mengubah variable pointer, maka yang berubah hanya variable tersebut.
// Semua variable yang mengacu ke data yang sama tidak akan berubah
// Jika kita ingin mengubah seluruh variable yang mengacu ke data tersebut, kita bisa menggunakan operator *

type Address struct {
	City, Province string
}

// pointer function
// menambahkan * merepresentasikan sebagai pointer
func changeAddress(address *Address) {
	address.City = "Jepara"
}

// pointer method
type Man struct {
	Name string
}
// menambahkan * merepresentasikan sebagai pointer
func (man *Man) Maried() {
	man.Name = "Mr. " + man.Name
}

func main() {
	// asterisk
	address1 := Address{"Bangsri", "Jawa Tengah"}
	address2 := &address1

	*address2 = Address{"Mlonggo", "Jawa Tengah"} // semua pada mengacu ke address 2

	fmt.Println(address1) // berubah
	fmt.Println(address2)

	// new ->  digunakan untuk membuat pointer baru
	address3 := new(Address) // data otomatis menjadi kosong
	address4 := address3 // sudah otomatis mengacu pada address 3
	address4.City = "Guyangan"
	fmt.Println(address4)

	// pointer function
	address5 := Address{}
	// menambahkan & agar yang dikirim dalam bentuk pointer
	changeAddress(&address5)
	fmt.Println(address5)

	// pointer method
	ivan := Man{"Ivan"}
	ivan.Maried()
	fmt.Println(ivan)

}