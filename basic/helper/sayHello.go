package helper // biasanya nama sesuai dengan folder

// Access Modifier
// Di bahasa pemrograman lain, biasanya ada kata kunci yang bisa digunakan untuk menentukan access modifier terhadap suatu function atau variable
// Di Go-Lang, untuk menentukan access modifier, cukup dengan nama function atau variable
// Jika nama nya diawali dengan hurup besar, maka artinya bisa diakses dari package lain, jika dimulai dengan hurup kecil, artinya tidak bisa diakses dari package lain
// ex: 
var version string = "1.0.0" // tidak bisa di akses di package lain karena awal memakai huruf kecil tetapi masih bisa di akses di package yang sama
var Application string = "Golang menyenangkan bukan!" // bisa di akses di package mana aja 

// untuk function huruf pertama harus Kapital
func SayHello(name string) string {
	return "Hi, " + name
}