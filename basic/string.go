package main

import "fmt"

func main() {
    fmt.Println("This is a Go program.")
    fmt.Println("This is my first Go program.")

	fmt.Println(len("This is a Go program.")) // output: 32 
    fmt.Println("This is my first Go program."[0]) // output: 84 -> dalam bentuk byte
    fmt.Println("This is my first Go program."[1]) // output: 104 -> dalam bentuk byte
}

// function untuk string

// len("string") -> Menghitung jumlah karakter di String
// “string”[number] -> Mengambil karakter pada posisi yang ditentukan