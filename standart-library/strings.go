package main

import (
	"fmt"
	"strings"
)

// Package strings adalah package yang berisikan function-function untuk memanipulasi tipe data String

// beberapa function string
//  strings.Trim(string, cutset) ->  Memotong cutset di awal dan akhir string
// strings.ToLower(string) -> Membuat semua karakter string menjadi lower case
// strings.ToUpper(string) -> Membuat semua karakter string menjadi upper case
// strings.Split(string, separator) -> Memotong string berdasarkan separator
// strings.Contains(string, search) -> Mengecek apakah string mengandung string lain
// strings.ReplaceAll(string, from, to) -> Mengubah semua string dari from ke to


func main() {
	fmt.Println(strings.Trim("Ivan Adi", "Iv")) // an Adi
	fmt.Println(strings.ToLower("Hello, World!")) // hello, world!
	fmt.Println(strings.ToUpper("hello, world!")) // HELLO, WORLD!
	fmt.Println(strings.Split("Ivan, Adi, Saputra", ", ")) // [Ivan Adi Saputra]
	fmt.Println(strings.Contains("Ivan Adi Saputra", "Adi")) // true
	fmt.Println(strings.ReplaceAll("Ivan Adi Saputra", "Adi", "Andi")) // Ivan Andi Saputra
}