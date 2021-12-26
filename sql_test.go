package golangdatabasemysql

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestExecSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	// bikin context dulu
	ctx := context.Background()

	// bikin perintah sql nya
	query := "INSERT INTO customer(id, name) VALUES('bayu','Bayu')"
	// balikannya ada 2 data, tapi kita tangkap errornya saja, karena execContext tidak mengembalikan data hasil query
	_, err := db.ExecContext(ctx, query)

	// cek apakah ada error atau tidak
	if err != nil {
		panic(err)
	}
	// jika success
	fmt.Println("Success insert new customer")
}

// test untuk query sql yang memiliki result (mengembalikan data hasil query)

func TestQuerySql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "SELECT id, name FROM customer"

	// ada 2 balikan datanya yakni error dan result hasil querynya
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}
	// rows harus di close
	defer rows.Close()
}

// setelah kita mendapatkan result hasil query, nah selanjutnya adalah kita dapatkan datanya per record, dengan cara iterasi Rows nya
func TestQuerySqlRows(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "SELECT id, name FROM customer"
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	// bikin perulangan
	// selama next() nya bernilai true, kita akan lakukan iterasi tiap datanya
	// caranya dengan menggunakan Scan, dan masukkan semua parameternya (sesuai dengan colomn saat query misal id dan name)

	for rows.Next() {
		var id, name string
		// pointer agar bisa menangkap hasil datanya, balikannya adalah error
		err = rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}
		// cetak id dan namenya
		fmt.Println("Id:", id)
		fmt.Println("Name:", name)
	}
}

func TestQuerySqlComplex(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "SELECT id, name, email, balance, rating, birth_date, married, created_at FROM customer"
	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	// iterasi data result
	for rows.Next() {
		var id, name string
		// karena email bisa null, pake NullString
		var email sql.NullString
		var balance int32
		var rating float64
		var birthDate sql.NullTime
		var createdAt time.Time
		var married bool
		// urutan recordnya harus sesuai urutan parameter (kolom) di script Query nya
		err = rows.Scan(&id, &name, &email, &balance, &rating, &birthDate, &married, &createdAt)
		if err != nil {
			panic(err)
		}
		fmt.Println("=====================")
		fmt.Println("Id:", id)
		fmt.Println("Name:", name)
		if email.Valid {
			fmt.Println("Email:", email.String)
		}
		fmt.Println("Balance:", balance)
		fmt.Println("Rating:", rating)
		if birthDate.Valid {
			fmt.Println("Birth Date:", birthDate.Time)
		}
		fmt.Println("Married:", married)
		fmt.Println("Created At:", createdAt)
	}
}

// test sql injection
func TestSqlInjection(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	// kita anggap username dan password adalah input dari user
	// inputan tersebut kita masukkan kedalam script query nya
	username := "admin"
	password := "admin"

	script := "SELECT username FROM user WHERE username= '" + username + "' AND password= '" + password + "' LIMIT 1"

	rows, err := db.QueryContext(ctx, script)

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	// kalau misalnya Next() berhasil, maka ada datanya di rows
	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("Sukses login", username)
	} else {
		// jika gagal
		fmt.Println("Gagal login")
	}
}

// problem SQL Injection
func TestProblemSqlInjection(t *testing.T) {
	db := GetConnection()
	defer db.Close()
	ctx := context.Background()

	// misal kita tambahkan karakter bebas di parameter untuk username
	// dan juga kita masukkan password yang salah (tidak sesuai)
	// karena SQL Injection itu akan mengubah struktur querynya
	username := "admin';#"
	password := "salah"

	script := "SELECT username FROM user WHERE username= '" + username + "' AND password= '" + password + "' LIMIT 1"

	rows, err := db.QueryContext(ctx, script)

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("Sukses login", username)
	} else {
		fmt.Println("Gagal login")
	}
}

func TestProblemSqlInjectionSafe(t *testing.T) {
	db := GetConnection()
	defer db.Close()
	ctx := context.Background()

	username := "admin"
	password := "admin"

	script := "SELECT username FROM user WHERE username= ? AND password= ? LIMIT 1"

	// parameter ketiga, masukkan parameter untuk ? di script query, ingat harus urut letak parameternya
	rows, err := db.QueryContext(ctx, script, username, password)

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("Sukses login", username)
	} else {
		fmt.Println("Gagal login")
	}
}

// test ExecConstext dengan SQL Injection safe
func TestExecSqlParameter(t *testing.T) {
	db := GetConnection()
	defer db.Close()
	ctx := context.Background()

	username := "aan"
	password := "Aan"
	// parameternya kita set dengan tanda tanya
	query := "INSERT INTO user(username, password) VALUES(?, ?)"

	// di execContext kita masukkan parameter username dan password nya, agar dimasukkan ke dalam script querynya
	_, err := db.ExecContext(ctx, query, username, password)

	if err != nil {
		panic(err)
	}
	fmt.Println("Success insert new user")
}

// test LastInsertId untuk mendapatkan ID terakhir dari auto_increment
func TestAutoIncrement(t *testing.T) {
	db := GetConnection()
	defer db.Close()
	ctx := context.Background()

	email := "bayu@mail.com"
	comment := "Test comment"
	query := "INSERT INTO comments(email, comment) VALUES(?, ?)"

	// balikan data dari ExecContext adalah result
	// result sendiri memiliki function LastInsertId
	result, err := db.ExecContext(ctx, query, email, comment)

	if err != nil {
		panic(err)
	}
	// cek id terakhir yang di auto_increment
	insertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}
	fmt.Println("Success insert new comment with id", insertId)
}

// test Prepare Statement
func TestPrepareStatement(t *testing.T) {
	db := GetConnection()
	defer db.Close()
	ctx := context.Background()

	script := "INSERT INTO comments (email, comment) VALUES(?, ?)"
	// buat prepare statement menggunakan PrepareContext,
	// balikannya adalah statementnya dan error
	statement, err := db.PrepareContext(ctx, script)

	if err != nil {
		panic(err)
	}
	// setelah selesai menggunakan statementnya, maka harus di close
	defer statement.Close()

	// perulangan untuk mengeksekusi prepare statementnya
	for i := 0; i < 10; i++ {
		email := "bayu" + strconv.Itoa(i) + "@mail.com"
		comment := "Komentar ke " + strconv.Itoa(i)

		// eksekusi menggunakan ExecContext, pakenya statement bukan db lagi
		// di ExecContext kita tidak perlu lagi membuat query, karena sudah dibuat diawal saat prepare statement
		result, err := statement.ExecContext(ctx, email, comment)
		if err != nil {
			panic(err)
		}
		// dapatkan id terakhir juga
		id, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}
		fmt.Println("Comment id", id)
	}
}
