package golangdatabasemysql

import (
	"context"
	"fmt"
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
		var id, name, email string
		var balance int32
		var rating float64
		var birthDate, createdAt time.Time
		var married bool
		// urutan recordnya harus sesuai urutan parameter (kolom) di script Query nya
		err = rows.Scan(&id, &name, &email, &balance, &rating, &birthDate, &married, &createdAt)
		if err != nil {
			panic(err)
		}
		fmt.Println("=====================")
		fmt.Println("Id:", id)
		fmt.Println("Name:", name)
		fmt.Println("Email:", email)
		fmt.Println("Balance:", balance)
		fmt.Println("Rating:", rating)
		fmt.Println("Birth Date:", birthDate)
		fmt.Println("Married:", married)
		fmt.Println("Created At:", createdAt)
	}

}
