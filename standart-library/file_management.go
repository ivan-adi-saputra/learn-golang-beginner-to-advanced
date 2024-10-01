package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// File Management
// Di package os, terdapat File Management, namun sengaja ditunda pembahasannya, karena kita harus tahu dulu tentang IO
// Saat kita membuat atau membaca file menggunakan Package os, struct File merupakan implementasi dari io.Reader dan io.Writer
// Oleh karena itu, kita bisa melakukan baca dan tulis terhadap File tersebut menggunakan Package io / bufio

// Open File
// Untuk membuat / membaca File, kita bisa menggunakan os.OpenFile(name, flag, permission)
// name berisikan nama file, bisa absolute atau relative / local
// flag merupakan penanda file, apakah untuk membaca, menulis, dan lain-lain
// permission, merupakan permission yang diperlukan ketika membuat file, bisa kita simulasikan disini : https://chmod-calculator.com/

// createNewFile membuat file baru atau membuka file jika sudah ada, kemudian menulis pesan ke dalamnya.
// Parameter:
// - name: Nama file yang akan dibuat atau ditulis.
// - message: Pesan yang akan ditulis ke file.
// Return:
// - error: Mengembalikan error jika terjadi kesalahan dalam pembuatan file atau penulisan.
func createNewFile(name string, message string) error {
	// Membuka atau membuat file dengan nama yang diberikan. Mode O_CREATE membuat file jika belum ada,
    // O_WRONLY membuka file dalam mode tulis, dan O_TRUNC mengosongkan file jika sudah ada.
	file, err := os.OpenFile(name, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
    
    // Jika terjadi kesalahan dalam membuka atau membuat file, kembalikan error.
    if err != nil {
        return err
    }
    
    // Menutup file secara otomatis setelah fungsi selesai dieksekusi.
    defer file.Close()

    // Menulis string message ke dalam file.
    _, err = file.WriteString(message)
    
    // Jika terjadi kesalahan dalam penulisan file, kembalikan error.
    if err != nil {
        return err
    }

    // Mengembalikan nil jika berhasil membuat file dan menulis pesan tanpa kesalahan.
    return nil
}

// readFile adalah fungsi yang menerima nama file (name string) dan mengembalikan
// isi file sebagai string serta error jika terjadi kesalahan
func readFile(name string) (string, error) {

    // Membuka file dengan nama yang diberikan dalam mode baca saja (os.O_RDONLY).
    // Jika file tidak bisa dibuka, fungsi akan mengembalikan error
    file, err := os.OpenFile(name, os.O_RDONLY, 0666)
	if err != nil {
        return "", err // Mengembalikan error jika file gagal dibuka
    }
	
	// Menunda penutupan file hingga fungsi selesai dieksekusi untuk
	// memastikan file ditutup meskipun ada error dalam proses selanjutnya
	defer file.Close()

	// Membuat reader baru dari file menggunakan bufio.NewReader,
	// yang memungkinkan kita membaca file baris per baris
	reader := bufio.NewReader(file)

	// Variabel untuk menampung isi file yang akan dibaca
	var message string
	
	// Melakukan perulangan untuk membaca file baris demi baris
	for {
		// Membaca satu baris dari file. Fungsi ReadLine mengembalikan 3 nilai:
		// line (baris yang dibaca), isPrefix (apakah baris terlalu panjang), err (error jika ada)
		line, _, err := reader.ReadLine()

		// Jika error adalah io.EOF (End of File), hentikan perulangan karena
		// seluruh file sudah terbaca
		if err == io.EOF {
            break
        }

		// Menambahkan baris yang dibaca ke variabel message
		message += string(line)
	}
	
	// Mengembalikan isi file yang sudah dibaca sebagai string
	return message, nil
}

func addToFile(name string, message string) error {
	// Membaca file dan menulis pesan baru ke file
    file, err := os.OpenFile(name, os.O_APPEND|os.O_WRONLY, 0644)
    if err!= nil {
        return err
    }
    defer file.Close()

    _, err = file.WriteString(message + "\n")
    if err!= nil {
        return err
    }

    return nil
}

func main() {
	// membuar file baru
	err := createNewFile("example.txt", "Hello, this is an example text.")
    // jika terjadi error, tampilkan pesan error
    if err!= nil {
        fmt.Println("Error creating file:", err)
    } else {
        fmt.Println("File created successfully!")
    }

	// membaca file
	message, err := readFile("example.txt")
    // jika terjadi error, tampilkan pesan error
    if err!= nil {
        fmt.Println("Error reading file:", err)
    } else {
        fmt.Println("File content:", message)
    }

	// menambahkan pesan baru ke file
	addToFile("example.txt", "This is an additional line.")
}