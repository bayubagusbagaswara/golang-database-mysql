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
	// balikannya ada 2
	_, err := db.ExecContext(ctx, query)

	// cek apakah ada error atau tidak
	if err != nil {
		panic(err)
	}
	// jika success
	fmt.Println("Success insert new customer")

}
