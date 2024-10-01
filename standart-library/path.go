package main

import (
	"fmt"
	"path"
	"path/filepath"
)

// Package path digunakan untuk memanipulasi data path seperti path di URL atau path di File System
// Secara default Package path menggunakan slash sebagai karakter path nya, oleh karena itu cocok untuk data URL
// Namun jika ingin menggunakan untuk memanipulasi path di File System, karena Windows menggunakan backslash, maka khusus untuk File System, perlu menggunakan pacakge path/filepath

func main() {
	fmt.Println(path.Dir("Hello/main.go")) // output: Hello
	fmt.Println(path.Base("Hello/main.go")) // output: main.go
	fmt.Println(path.Ext("Hello/main.go")) // output:.go
	fmt.Println(path.Join("Hello", "world", "main.go")) // output: Hello/world/main.go

	// path/filepath digunakan untuk memanipulasi path di File System
	fmt.Println(filepath.Dir("Hello/main.go")) // output: Hello
    fmt.Println(filepath.Base("Hello/main.go")) // output: main.go
    fmt.Println(filepath.Ext("Hello/main.go")) // output:.go
	fmt.Println(filepath.IsAbs("Hello/main.go")) // output: false
	fmt.Println(filepath.IsLocal("Hello/main.go")) // output: true
    fmt.Println(filepath.Join("Hello", "world", "main.go")) // output: Hello/world/main.go
    fmt.Println(filepath.ToSlash("/Hello/world/main.go")) // output: Hello/world/main.go
    fmt.Println(filepath.FromSlash("Hello/world/main.go")) // output: /Hello/world/main.go
    fmt.Println(filepath.Split("Hello/world/main.go")) // output: (Hello/world, main.go)
    fmt.Println(filepath.SplitList("Hello/world/main.go")) //
}