package main

import "fmt"

func main() {
	// conversion integer
	var nilai32 int32 = 231234  // output: 231234
	var nilai64 int64 = int64(nilai32) // output: 231234
	var nilai16 int16 = int16(nilai32) // output: -30910 -> karena melebihi kapasitas jadi selisih akan balik ke minus

	fmt.Printf("Nilai 32-bit: %d\n", nilai32)
	fmt.Printf("Nilai 64-bit: %d\n", nilai64)
	fmt.Printf("Nilai 16-bit: %d\n", nilai16)

	// conversion string
	var name = "Ivan Adi Saputra"
	var i uint8 = name[1]
	var iString = string(i)

	fmt.Println(name)  
	fmt.Println(i)  // output: 118
	fmt.Println(iString) // output: v
}