package main

import (
	"encoding/base64"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	// base64
	value := "Ivan Adi Saputra"

	encoded := base64.StdEncoding.EncodeToString([]byte(value))
	fmt.Println(encoded)

	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		// karena decoded byte maka konversi string
		fmt.Println(string(decoded))
	}

	// csv reader
	csvString := "ivan,adi,saputra\n" + 
		"list,tiana\n" + 
		"fil,iv"
	
	// Ini membuat pembaca (reader) CSV dengan menggunakan fungsi csv.NewReader.
	// strings.NewReader(csvString) digunakan untuk mengubah string csvString menjadi io.Reader, yang kemudian bisa dibaca oleh csv.NewReader.
	// reader akan digunakan untuk membaca baris-baris CSV dari csvString.
	reader := csv.NewReader(strings.NewReader(csvString))

	for {
		// Setiap kali loop berjalan, baris ini mencoba membaca satu baris dari input CSV (reader.Read()).
		// Fungsi Read() mengembalikan dua nilai:
		// record: berisi array dari string yang merupakan elemen-elemen dalam baris CSV tersebut.
		// err: berisi error jika ada kesalahan saat membaca (misalnya, mencapai akhir file).
		record, err := reader.Read()
		// Mengecek apakah error yang dihasilkan adalah io.EOF, yang berarti telah mencapai akhir input (EOF = End Of File).
		if err == io.EOF {
			break
		}

		fmt.Println(record)
	}

	// menulis csv
	writer := csv.NewWriter(os.Stdout)

	// Menulis baris 
	_ = writer.Write([]string{"Ivan", "Adi", "Saputra"})
	_ = writer.Write([]string{"List", "Tiana"})

	// Memastikan data ditulis ke output
	writer.Flush()
}