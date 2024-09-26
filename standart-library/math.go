package main

import (
	"fmt"
	"math"
)

// math.Round(float64) -> Membulatkan float64 keatas atau kebawah, sesuai dengan yang paling dekat
// math.Floor(float64) -> Membulatkan float64 kebawah
// math.Ceil(float64) -> Membulatkan float64 keatas
// math.Max(float64, float64) -> Mengembalikan nilai float64 paling besar
// math.Min(float64, float64) -> Mengembalikan nilai float64 paling kecil


func main() {
	fmt.Println(math.Round(4.5))      // 5
    fmt.Println(math.Round(-4.5))     // -4
    fmt.Println(math.Floor(4.5))     // 4
    fmt.Println(math.Floor(-4.5))    // -5
    fmt.Println(math.Ceil(4.5))      // 5
    fmt.Println(math.Ceil(-4.5))     // -4
    fmt.Println(math.Max(3, 5))      // 5
    fmt.Println(math.Max(-3, -5))     // -3
    fmt.Println(math.Min(3, 5))      // 3
    fmt.Println(math.Min(-3, -5))     // -5
}