package main

import "fmt"

func main() {
	names := []string{"ivan", "adi", "saputra"}
	for index, name := range names {
		fmt.Printf("Index: %d, Name: %s\n", index, name)
	}

	// jika ingin hanya print nama-nama saja tanpa index
	for _, name := range names {
		fmt.Printf(" Name: %s\n", name)
	}
}
// for range -> digunakan loopin map, slice, atau array yang lebih mudah