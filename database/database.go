package database

import (
	"database/sql"
	"time"
)

// Database Pooling
// sql.DB di Golang sebenarnya bukanlah sebuah koneksi ke database
// Melainkan sebuah pool ke database, atau dikenal dengan konsep Database Pooling
// Di dalam sql.DB, Golang melakukan management koneksi ke database secara otomatis. Hal ini menjadikan kita tidak perlu melakukan management koneksi database secara manual
// Dengan kemampuan database pooling ini, kita bisa menentukan jumlah minimal dan maksimal koneksi yang dibuat oleh Golang, sehingga tidak membanjiri koneksi ke database, karena biasanya ada batas maksimal koneksi yang bisa ditangani oleh database yang kita gunakan

// Pengaturan Database Pooling
// (DB) SetMaxIdleConns(number) = Pengaturan berapa jumlah koneksi minimal yang dibuat
// (DB) SetMaxOpenConns(number) = Pengaturan berapa jumlah koneksi maksimal yang dibuat
// (DB) SetConnMaxIdleTime(duration) = Pengaturan berapa lama koneksi yang sudah tidak digunakan akan dihapus
//  (DB) SetConnMaxLifetime(duration) = Pengaturan berapa lama koneksi boleh digunakan

func GetConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/bwastore-laravel?parseTime=true")
	if err!= nil {
        panic(err.Error())
    }

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}