package golangdatabasemysql

import (
	"database/sql"
	"time"
)

// balikan dari function adalah pointer *sql.DB
func GetConnection() *sql.DB {

	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/golang_database?parseTime=true")
	if err != nil {
		panic((err))
	}

	// ini tidak di close disini

	// buat Database Pooling
	db.SetMaxIdleConns(10)                 // minimum jumlah koneksi yang dibuat
	db.SetMaxOpenConns(10)                 // maksimal koneksi
	db.SetConnMaxIdleTime(5 * time.Minute) // waktu tunggu koneksi akan dihapus (jika tidak digunakan)
	// kalo udah 60 menit, kita buat koneksi baru
	db.SetConnMaxLifetime(60 * time.Minute)

	// balikan object db
	return db
}
