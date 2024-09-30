package main

import (
	"fmt"
	"regexp"
)

// regexp digunakan untuk melakukan pencarian regular exspresion
// function package regexp
// regexp.MustCompile(string) -> membuat regexp
// regexp.MatchString(string, string) -> mengembalikan true jika string mengandung regexp
// regexp.FindString(string, string) -> mengembalikan string yang mengandung regexp
// regexp.FindStringIndex(string, string) -> mengembalikan index awal dan akhir string yang mengandung regexp
// regexp.FindAllString(string, int, string) -> mengembalikan slice string yang mengandung regexp

func main() {
	// Buat regex yang mengandung huruf 'i' diikuti oleh huruf vokal 'a'
	// huruf kecil atau besar sangat berpengaruh
	regex := regexp.MustCompile(`i([a-z])an`)

	fmt.Println(regex.MatchString("ivan"))
	fmt.Println(regex.FindAllString("ivan ipan saputra", 10))
}