package main

import "fmt"

func main() {
	fmt.Println("Satu = ", 1)
	fmt.Println("Dua = ", 2)
	fmt.Println("Tiga koma empat = ", 3.4)
}

// terdapat 2 jenis type data number yaitu integer & floating point

// 1. type data integer
// type data | nilai minimum        | nilai maksimum
// -------------------------------------------------
// int8      | -128                 | 128
// int16     | -32768               | 32768
// int32     | -2147483648          | 2147483648
// int64     | -9223372036854775808 | 2147483648

// 2. type data integer
// type data  | nilai minimum   | nilai maksimum
// -------------------------------------------------
// uint8      | 0               | 255
// uint16     | 0               | 65535
// uint32     | 0          		| 4294967295
// uint64     | 0 				| 18446744073709551615

// type data floating point
// type data  | nilai minimum   | nilai maksimum
// -------------------------------------------------
// float32     | 1.18×10−38     | 3.4×1038
// float64     | 2.23×10−308    | 1.80×1030
// complex64   | complex numbers with float32 real and imaginary parts.
// complex128  | complex numbers with float64 real and imaginary parts.

// Alias
// type data  | alias untuk
// ------------------------
// byte     | uint8
// rune     | int32
// int      | Minimal int32
// uint     | Minimal uint32