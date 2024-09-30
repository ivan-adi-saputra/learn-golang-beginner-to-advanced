package main

import (
	"fmt"
	"reflect"
)

// reflect digunakan untuk melihat struktur code pada saat aplikasi sedang berjalan
// reflection ketika kita ingin membuat library yang general sehingga mudah digunakan

type Sample struct {
	Name string `required:"true" max:"10"`
}

type Person struct {
	Name string
}

func readFile(value any) {
	valueType := reflect.TypeOf(value)
	
	fmt.Println("Type name ", valueType.Name())

	// NumField digunakan untuk menghitung panjang
	for i := 0; i < valueType.NumField(); i++ {
		field := valueType.Field(i)
        fmt.Println(field.Name, " with type ", field.Type)
		fmt.Println(field.Tag.Get("max"))
		fmt.Println(field.Tag.Get("required"))
	}
}

// membuar validation
func isValid(value any) (result bool) {
	valueType := reflect.TypeOf(value)
	result = true

	for i := 0; i < valueType.NumField(); i++ {
		field := valueType.Field(i)
		if field.Tag.Get("required") == "true" {
		// 	reflect.ValueOf(value): Mengambil representasi nilai dari value dalam bentuk tipe reflect.Value. Ini memungkinkan kita untuk bekerja dengan nilai secara dinamis (runtime) menggunakan refleksi.
		// .Field(i): Mengakses field ke-i dari struct yang direpresentasikan oleh value. Ini mengacu pada field di posisi ke-i dalam struct.
		// .Interface(): Mengambil nilai dari field yang diakses sebelumnya dan mengubahnya kembali menjadi interface{}. Pada dasarnya, ini mengembalikan nilai asli dari field tersebut sebagai tipe interface{}.
			data := reflect.ValueOf(value).Field(i).Interface() 
			result = data != ""
			if result == false {
				return result
			}
		}
	}
	return result
}

func main() {
	sample := Sample{Name: "Ivan Adi"}
    readFile(sample)

    person := Person{Name: "Saputra"}
    readFile(person)

	personIsValid := Sample{""}
	fmt.Println(isValid(personIsValid))
}