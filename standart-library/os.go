package main

import (
	"fmt"
	"os"
)

func main() { 
	// package Args digunakan untuk mengambil argumen yang dikirimkan
	// ex: go run os.go ivan adi saputra
	// maka akan menjasi array ["ivan", "adi", "saputra"]
	args := os.Args
	for _, arg := range args {
		fmt.Printf("Arg %s\n", arg)
	}

	// hostname digunakan untuk mengambil hostname
	hostname, err := os.Hostname()
    if err!= nil {
        fmt.Println("Error getting hostname:", err)
    }
    fmt.Printf("Hostname: %s\n", hostname)
	// output: Hostname: DESKTOP-STMCJQV

    // env digunakan untuk mengambil environment variable
    // ex: go run os.go ENV USER
    envUser := os.Getenv("USER")
    fmt.Printf("User: %s\n", envUser)
    // jika environment variable USER belum di set, maka akan menampilkan ""
 }