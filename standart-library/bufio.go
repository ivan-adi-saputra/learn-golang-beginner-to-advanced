package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// IO atau singkatan dari Input Output, merupakan fitur di Golang yang digunakan sebagai standard untuk proses Input Output
// Di Golang, semua mekanisme input output pasti mengikuti standard package io

// Reader
// Untuk membaca input, Golang menggunakan kontrak interface bernama Reader yang terdapat di package io

// Writer
// Untuk menulis ke output, Golang menggunakan kontrak interface bernama Writer yang terdapat di package io

// Implementasi IO
// Penggunaan dari IO sendiri di Golang terdapat dibanyak package, sebelumnya contohnya kita menggunakan CSV Reader dan CSV Writer
// Karena Package IO sebenarnya hanya kontrak untuk IO, untuk implementasinya kita harus lakukan sendiri
// Tapi untungnya, Golang juga menyediakan package untuk mengimplementasikan IO secara mudah, yaitu menggunakan package bufio

// bufio
// Package bufio atau singkatan dari buffered io
// Package ini digunakan untuk membuat data IO seperti Reader dan Writer

func main() {
	// Reader
	input := strings.NewReader("ivan adi \n saputra")
	reader := bufio.NewReader(input)

	for {
		line, _, err := reader.ReadLine()
        if err == io.EOF {
            break
        }
		fmt.Println(string(line))
	}
	// output
	// ivan adi
	// saputra

	// Writer
	output := bufio.NewWriter(os.Stdout)
    output.WriteString("Hello, ")
    output.WriteString("World!\n")
    output.Flush() // flush menyimpan output ke stdout
    // output: Hello, World!
    // flush() digunakan ketika kita ingin menyimpan output sebelum program selesai
}

