package embed

import (
	"embed"
	_ "embed"
	"fmt"
	"io/fs"
	"io/ioutil"
	"testing"
)

// file ke string

//go:embed 1_introduction.txt
var introduction string

func TestString(t *testing.T){
	fmt.Println(introduction)
}

// Embed File ke Byte[]
// Selain ke tipe data String, embed file juga bisa dilakukan ke variable tipe data []byte
// Ini cocok sekali jika kita ingin melakukan embed file dalam bentuk binary, seperti gambar dan lain-lain

//go:embed motor.jpg
var motor []byte

func TestByte(t *testing.T){
    ioutil.WriteFile("motor_new.jpg", motor, fs.ModePerm)
}

// Embed Multiple Files
// Kadang ada kebutuhan kita ingin melakukan embed beberapa file sekaligus
// Hal ini juga bisa dilakukan menggunakan embed package
// Kita bisa menambahkan komentar //go:embed lebih dari satu baris
// Selain itu variable nya bisa kita gunakan tipe data embed.FS

//go:embed files/a.txt
//go:embed files/b.txt
var embeddedFiles embed.FS

func TestEmbeddedFiles(t *testing.T){
    dataA, _ := embeddedFiles.ReadFile("files/a.txt")
    dataB, _ := embeddedFiles.ReadFile("files/b.txt")

    fmt.Println(string(dataA))
    fmt.Println(string(dataB))
}

// Patch Matcher
// Selain manual satu per satu, kita bisa mengguakan patch matcher untuk membaca multiple file yang kita inginkan
// Ini sangat cocok ketika misal kita punya pola jenis file yang kita inginkan untuk kita baca
// Caranya, kita perlu menggunakan path matcher seperti pada package function path.Match

//go:embed files/*.txt
var path embed.FS

func TestPathMatcher(t *testing.T){
	files, _ := path.ReadDir("files")

    for _, file := range files {
        if file.IsDir() {
            continue
        }

        data, _ := path.ReadFile("files/" + file.Name())
        fmt.Println(string(data))
    }
}

// Hasil Embed di Compile
// Perlu diketahui, bahwa hasil embed yang dilakukan oleh package embed adalah permanent dan data file yang dibaca disimpan dalam binary file golang nya
// Artinya bukan dilakukan secara realtime membaca file yang ada diluar
// Hal ini menjadikan jika binary file golang sudah di compile, kita tidak butuh lagi file external nya, dan bahkan jika diubah file external nya, isi variable nya tidak akan berubah lagi

// Compile
// go build

