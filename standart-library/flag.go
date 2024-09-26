package main

import (
	"flag"
	"fmt"
)

func main() {
	// Package flag berisikan fungsionalitas untuk memparsing command line argument

	// (name flag, default value, deskripsi)
	var username *string = flag.String("username", "root", "database untuk username")
	var password *string = flag.String("password", "root", "database untuk password")
	var host *string = flag.String("host", "localhost", "database untuk host")
	var port *int = flag.Int("port", 3306, "database untuk port")

	// Parsing argument yang diinputkan user
	flag.Parse()

    // Menampilkan hasil parsing
    fmt.Printf("Username: %s\nPassword: %s\nHost: %s\nPort: %d\n", *username, *password, *host, *port)

	// menjalankan go run flag.go -username=ivansaputra -password="admin 123" -host=123.08.123 -port=3000
	// output:
	// Username: ivansaputra
	// Password: admin 123  
	// Host: 123.08.123     
	// Port: 3000
}