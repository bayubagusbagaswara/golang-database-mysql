package golangdatabasemysql

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

// fungsi dari github.com/go-sql-driver/mysql agar nanti bisa dipanggil method init nya

func TestEmpty(t *testing.T) {

}

// bikin test open connection database
func TestOpenConnection(t *testing.T) {
	// username = root
	// password = root
	// open koneksi db, masukkan parameter driver db, username, password, host, port nama databasenya
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/golang_database")
	if err != nil {
		panic((err))
	}
	// close koneksi db
	defer db.Close()

}
