package main

import (
	"fmt"
	"slices"
)

// package slices digunakan untuk memanipulasi data di slice


func main() {
	names := []string{"Ivan", "Adi", "Saputra"}
	number := []int{5, 3, 8, 10}

	// digunakan untuk mencari terkecil
	fmt.Println(slices.Min(names)) // output: Adi
	fmt.Println(slices.Min(number)) // output: 3
	// digunakan untuk mencari terbesar
	fmt.Println(slices.Max(names)) // output: Saputra
    fmt.Println(slices.Max(number)) // output: 10
	// digunakan untuk mencari dalam slice
	fmt.Println(slices.Contains(names, "Ivan")) // output: true
	fmt.Println(slices.Contains(number, 5)) // output: true
	// digunakan untuk mencari index
	fmt.Println(slices.Index(names, "Adi")) // output: 1
    fmt.Println(slices.Index(number, 5)) // output: 0
}