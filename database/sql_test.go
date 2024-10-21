package database

import (
	"context"
	"fmt"
	"testing"
)

// Eksekusi Perintah SQL
// Saat membuat aplikasi menggunakan database, sudah pasti kita ingin berkomunikasi dengan database menggunakan perintah SQL
// Di Golang juga menyediakan function yang bisa kita gunakan untuk mengirim perintah SQL ke database menggunakan function (DB) ExecContext(context, sql, params)
// Ketika mengirim perintah SQL, kita butuh mengirimkan context, dan seperti yang sudah pernah kita pelajari di course Golang Context, dengan context, kita bisa mengirim sinyal cancel jika kita ingin membatalkan pengiriman perintah SQL nya

func TestExecSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	query := "INSERT INTO carts (products_id, users_id) VALUES (12, 1)"
	_, err := db.ExecContext(ctx, query)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success for insert data in database")
}

// Query SQL
// Untuk operasi SQL yang tidak membutuhkan hasil, kita bisa menggunakan perintah Exec, namun jika kita membutuhkan result, seperti SELECT SQL, kita bisa menggunakan function yang berbeda
// Function untuk melakukan query ke database, bisa menggunakan function (DB) QueryContext(context, sql, params)

// Rows
// Hasil Query function adalah sebuah data structs sql.Rows
// Rows digunakan untuk melakukan iterasi terhadap hasil dari query
// Kita bisa menggunakan function (Rows) Next() (boolean) untuk melakukan iterasi terhadap data hasil query, jika return data false, artinya sudah tidak ada data lagi didalam result
// Untuk membaca tiap data, kita bisa menggunakan (Rows) Scan(columns...)
// Dan jangan lupa, setelah menggunakan Rows, jangan lupa untuk menutupnya menggunakan (Rows) Close()

func TestQuery(t *testing.T)  {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	query := "SELECT products_id, users_id FROM carts"
	rows, err := db.QueryContext(ctx, query)
	if err!= nil {
        panic(err)
    }
	defer rows.Close()
	
    for rows.Next() {
		var productId, userId int
		err = rows.Scan(&productId, &userId)
		if err!= nil {
            panic(err)
        }
		fmt.Printf("Product ID: %d, User ID: %d\n", productId, userId)
	}
}

// Tipe Data Column
// Sebelumnya kita hanya membuat table dengan tipe data di kolom nya berupa VARCHAR
// Untuk VARCHAR di database, biasanya kita gunakan String di Golang
// Bagaimana dengan tipe data yang lain?
// Apa representasinya di Golang, misal tipe data timestamp, date dan lain-lain

// Mapping Tipe Data
//  VARCHAR, CHAR = string
// INT, BIGINT = int32, int64
// FLOAT, DOUBLE = float32, float64
// BOOLEAN = bool
// DATE, DATETIME, TIME, TIMESTAMP = time.Time

// Error Tipe Data Date
// Secara default, Driver MySQL untuk Golang akan melakukan query tipe data DATE, DATETIME, TIMESTAMP menjadi []byte / []uint8. Dimana ini bisa dikonversi menjadi String, lalu di parsing menjadi time.Time
// Namun hal ini merepotkan jika dilakukan manual, kita bisa meminta Driver MySQL untuk Golang secara otomatis melakukan parsing dengan menambahkan parameter parseDate=true

// Nullable Type
// Golang database tidak mengerti dengan tipe data NULL di database
// Oleh karena itu, khusus untuk kolom yang bisa NULL di database, akan jadi masalah jika kita melakukan Scan secara bulat-bulat menggunakan tipe data representasinya di Golang

// Error Data Null
// Konversi secara otomatis NULL tidak didukung oleh Driver MySQL Golang
// Oleh karena itu, khusus tipe kolom yang bisa NULL, kita perlu menggunakan tipe data yang ada dalam package sql

// Tipe Data Nullable
//  string = database/sql.NullString
// bool = database/sql.NullBool
// float64 = database/sql.NullFloat64
// int32 = database/sql.NullInt32
// int64 = database/sql.NullInt64
// time.Time = database/sql.NullTime

// SQL Injection
// SQL Injection adalah sebuah teknik yang menyalahgunakan sebuah celah keamanan yang terjadi dalam lapisan basis data sebuah aplikasi.
// Biasa, SQL Injection dilakukan dengan mengirim input dari user dengan perintah yang salah, sehingga menyebabkan hasil SQL yang kita buat menjadi tidak valid
// SQL Injection sangat berbahaya, jika sampai kita salah membuat SQL, bisa jadi data kita tidak aman

// Solusinya?
// Jangan membuat query SQL secara manual dengan menggabungkan String secara bulat-bulat
// Jika kita membutuhkan parameter ketika membuat SQL, kita bisa menggunakan function Execute atau Query dengan parameter yang akan kita bahas di chapter selanjutnya

// SQL Dengan Parameter
// Sekarang kita sudah tahu bahaya nya SQL Injection jika menggabungkan string ketika membuat query
// Jika ada kebutuhan seperti itu, sebenarnya function Exec dan Query memiliki parameter tambahan yang bisa kita gunakan untuk mensubtitusi parameter dari function tersebut ke SQL query yang kita buat.
// Untuk menandai sebuah SQL membutuhkan parameter, kita bisa gunakan karakter ? (tanda tanya)

// Contoh SQL
// SELECT username FROM user WHERE username = ? AND password = ? LIMIT 1
// INSERT INTO user(username, password) VALUES (?, ?)

// Auto Increment
// Kadang kita membuat sebuah table dengan id auto increment
// Dan kadang pula, kita ingin mengambil data id yang sudah kita insert ke dalam MySQL
// Sebenarnya kita bisa melakukan query ulang ke database menggunakan SELECT LAST_INSERT_ID()
// Tapi untungnya di Golang ada cara yang lebih mudah
// Kita bisa menggunakan function (Result) LastInsertId() untuk mendapatkan Id terakhir yang dibuat secara auto increment 
// Result adalah object yang dikembalikan ketika kita menggunakan function Exec

// Query atau Exec dengan Parameter
// Saat kita menggunakan Function Query atau Exec yang menggunakan parameter, sebenarnya implementasi dibawah nya menggunakan Prepare Statement
// Jadi tahapan pertama statement nya disiapkan terlebih dahulu, setelah itu baru di isi dengan parameter
// Kadang ada kasus kita ingin melakukan beberapa hal yang sama sekaligus, hanya berbeda parameternya. Misal insert data langsung banyak
// Pembuatan Prepare Statement bisa dilakukan dengan manual, tanpa harus mennggunakan Query atau Exec dengan parameter

// Prepare Statement
// Saat kita membuat Prepare Statement, secara otomatis akan mengenali koneksi database yang digunakan
// Sehingga ketika kita mengeksekusi Prepare Statement berkali-kali, maka akan menggunakan koneksi yang sama dan lebih efisien karena pembuatan prepare statement nya hanya sekali diawal saja
// Jika menggunakan Query dan Exec dengan parameter, kita tidak bisa menjamin bahwa koneksi yang digunakan akan sama, oleh karena itu, bisa jadi prepare statement akan selalu dibuat berkali-kali walaupun kita menggunakan SQL yang sama
// Untuk membuat Prepare Statement, kita bisa menggunakan function (DB) Prepare(context, sql)
// Prepare Statement direpresentasikan dalam struct database/sql.Stmt
// Sama seperti resource sql lainnya, Stmt harus di Close() jika sudah tidak digunakan lagi

// Database Transaction
// Salah satu fitur andalan di database adalah transaction
// Materi database transaction sudah saya bahas dengan tuntas di materi MySQL database, jadi silahkan pelajari di course tersebut
// Di course ini kita akan fokus bagaimana menggunakan database transaction di Golang

// Transaction di Golang
// Secara default, semua perintah SQL yang kita kirim menggunakan Golang akan otomatis di commit, atau istilahnya auto commit
// Namun kita bisa menggunakan fitur transaksi sehingga SQL yang kita kirim tidak secara otomatis di commit ke database
// Untuk memulai transaksi, kita bisa menggunakan function (DB) Begin(), dimana akan menghasilkan struct Tx yang merupakan representasi Transaction
// Struct Tx ini yang kita gunakan sebagai pengganti DB untuk melakukan transaksi, dimana hampir semua function di DB ada di Tx, seperti Exec, Query atau Prepare
// Setelah selesai proses transaksi, kita bisa gunakan function (Tx) Commit() untuk melakukan commit atau Rollback()

