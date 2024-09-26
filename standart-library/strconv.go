package main

import (
	"fmt"
	"strconv"
)

// Sebelumnya kita sudah belajar cara konversi tipe data, misal dari int32 ke int34
// Bagaimana jika kita butuh melakukan konversi yang tipe datanya berbeda? Misal dari int ke string, atau sebaliknya
// Hal tersebut bisa kita lakukan dengan bantuan package strconv (string conversion)

// beberaoa function dalam strconv
// strconv.parseBool(string) -> Mengubah string ke bool
// strconv.parseFloat(string) -> Mengubah string ke float
// strconv.parseInt(string) -> Mengubah string ke int64
// strconv.FormatBool(bool) -> Mengubah bool ke string
// strconv.FormatFloat(float, … ) -> Mengubah float64 ke string
// strconv.FormatInt(int, … ) -> Mengubah int64 ke string


func main() {
	// data yang dikembalikan 2 
	// result, err

	// string ke int
    strInt := "123"
    numInt, _ := strconv.Atoi(strInt)
    fmt.Printf("String to int: %d\n", numInt)

    // int ke string
    numStr := 456
    strStr := strconv.Itoa(numStr)
    fmt.Printf("Int to string: %s\n", strStr) 

    // float ke string
    numFloat := 78.9
    strFloat := strconv.FormatFloat(numFloat, 'f', 2, 64)
    fmt.Printf("Float to string: %s\n", strFloat)

    // string ke float
    strFloat2 := "78.9"
    numFloat2, _ := strconv.ParseFloat(strFloat2, 64)
    fmt.Printf("String to float: %f", numFloat2)

	// string ke bool
	strBool := "true"
    boolVal, _ := strconv.ParseBool(strBool)
    fmt.Printf("String to bool: %t\n", boolVal)

    // bool ke string
    boolVal2 := true
    strBool2 := strconv.FormatBool(boolVal2)
    fmt.Printf("Bool to string: %s\n", strBool2)

}