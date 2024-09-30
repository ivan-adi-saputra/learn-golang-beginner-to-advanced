package main

import (
	"fmt"
	"time"
)

// time

func main() {
	now := time.Now()
	// Local -> waktu berdasarkan computer ini
	fmt.Println(now.Local())

	utc := time.Date(2000, time.August, 19, 17, 45, 45, 0, time.UTC)
	// UTC -> waktu yang diawali pada 00:00:00 UTC pada tanggal 1 Agustus 2000
	fmt.Println(utc.UTC())

    // Timezone
    timezone, _ := time.LoadLocation("Asia/Jakarta")
    fmt.Println(now.In(timezone))

    // Duration
    duration := time.Duration(1) * time.Minute
    fmt.Println(duration)

    // Sleep
    time.Sleep(2 * time.Second)
    fmt.Println("Hello, World!")
    fmt.Println(time.Now())
    // output: Hello, World! 2022-01-01 19:02:01.000986 +07:00
    // output: 2022-01-01 19:02:03.000986 +07:00
    // output: 2022-01-0

	// duration
	duration1 := time.Second * 100 // 100 detik
	duration2 := time.Millisecond * 10
	duration3 := duration1 - duration2
	fmt.Println(duration3) // 990000000 ns = 99 seconds
}