package golangdatabasemysql

import (
	"context"
	"fmt"
	"testing"
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
