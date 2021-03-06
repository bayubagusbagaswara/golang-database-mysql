# Golang Database MySQL

## Agenda

- Pengenalan Golang Database
- Package Database
- Membuat Koneksi Database
- Eksekusi Perintah SQL
- SQL Injection
- Prepare Statement
- Database Transaction

## Pengenalan Database

- Bahasa pemrograman Go-Lang secara default memiliki sebuah package bernama `database`
- Package database adalah package yang berisikan kumpulan standard `interface` yang menjadi standard untuk berkomunikasi ke database
- Hal ini menjadikan kode program yang kita buat untuk mengakses jenis database apapun bisa menggunakan kode yang sama
- Yang `berbeda hanya kode SQL` yang perlu kita gunakan sesuai dengan database yang kita gunakan

## Cara Kerja Package Database

- Diagram alir

  ![Diagram_Alir_Package_Database](img/cara-kerja-package-database.jpg)

- Database Interface (package database) hanya berisi kontrak, tetap membutuhkan Database Driver
- Database driver nya perlu kita install terlebih dahulu, hanya install library nya saja

## MySQL

- Kita akan menggunakan MySQL sebagai Database Management System

## Menambah Database Driver

- Terlebih dahulu kita wajib menambahkan driver databasenya
- Tanpa driver database, maka package database di Go-Lang tidak mengerti apapun, karena hanya berisi kontrak interface saja
- Menambahkan module database mysql dengan perintah `go get -u github.com/go-sql-driver/mysql`
- harus include dulu driver dari mysql nya di import file go nya, seperti ini `"github.com/go-sql-driver/mysql"`

## Membuat Koneksi ke Database

- Hal pertama yang akan kita lakukan ketika aplikasi yang menggunakan database adalah melakukan koneksi ke databasenya
- Untuk melakukan koneksi ke database di Golang, kita bisa membuat object `sql.DB` menggunakan function `sql.Open(driver, dataSourceName)`
- Untuk menggunakan database MySQL, kita bisa menggunakan driver `"mysql"`
- Sedangkan untuk dataSourceName, tiap database biasanya punya cara penulisan masing-masing. Misalnya di MySQL, kita bisa menggunakan dataSourceName seperti dibawah ini:
  - `username:password@tcp(host:port)/database_name`
- Jika `object sql.DB` sudah tidak digunakan lagi, disarankan untuk menutupnya menggunakan function `Close()`

## Database Pooling

- Object `sql.DB` di Golang sebenarnya bukanlah sebuah koneksi ke database
- Melainkan `sebuah pool` ke database, atau dikenal dengan konsep `Database Pooling`
- Di dalam sql.DB, Golang melakukan management koneksi ke database secara otomatis. Hal ini menjadikan kita tidak perlu melakukan management koneksi database secara manual
- Jadi database pooling ini adalah `kumpulan dari object koneksi`
- Dengan kemampuan database pooling ini, kita bisa menentukan `jumlah minimal dan maksimal koneksi` yang dibuat oleh Golang, sehingga tidak membanjiri koneksi ke database, karena biasanya ada batas maksimal koneksi yang bisa ditangani oleh database yang kita gunakan

## Pengaturan Database Pooling

- Pengaturan database pooling

  ![Pengaturan_Database_Pooling](img/pengaturan-database-pooling.jpg)

## Eksekusi Perintah SQL

- Saat membuat aplikasi menggunakan database, sudah pastik kita ingin berkomunikasi dengan database menggunakan perintah SQL
- Di Golang juga menyediakan function yang bisa kita gunakan untuk mengirim perintah SQL ke database menggunakan function `(DB) ExecContext(context, sql, params)`
- Ketika mengirim perintah SQL, kita butuh mengirimkan context, dan seperti yang sudah pernah kita pelajari di `Golang Context`. Dengan menggunakan context, kita bisa mengirimkan sinyal cancel, jika kita ingin membatalkan pengiriman perintah SQL nya

## Query SQL

- Untuk operasi SQL yang tidak membutuhkan hasil, kita bisa menggnakan perintah Exec. Namun jika kita membutuhkan `result`, seperti `SELECT SQL`, kita bisa menggunakan function yang berbeda
- Function untuk melakukan query ke database, bisa menggunakan function `(DB) QueryContext(context, sql, params)`

## Rows (Baris)

- Hasil Query function adalah sebuah data structs `sql.Rows`
- Rows digunakan untuk melakukan iterasi terhadap hasil dari query
- Kita bisa menggunakan function `(Rows) Next() (boolean)` untuk melakukan iterasi terhadap data hasil query. Jika return data false, artinya sudah tidak ada data lagi didalam result
- Untuk membaca tiap data, kita bisa menggunakan `(Rows) Scan(columns...)`
- Dan jangan lupa, setelah menggunakan Rows, jangan lupa untuk menutupnya menggunakan `(Rows) Close()`

## Tipe Data Column

- Sebelumnya kita hanya membuat table dengan tipe data di kolom nya berupa VARCHAR
- Untuk VARCHAR di database, biasanya kita gunakan String di Golang
- Bagaimana dengan tipe data yang lain?
- Ada representasinya di Golang, misal tipe data timestamp, date dan lain-lain

## Mapping Tipe Data

- Berikut mapping tipe data dari Golang ke Database

  ![Mapping_Tipe_Data](img/mapping-tipe-data.jpg)

## Error Tipe Data Date

- Biasanya golang tidak bisa menerima balikan data dari database dengan tipe data Date

  ![Error_Tipe_Data_Date](img/error-tipe-data-date.jpg)

- Secara deafult, Driver MySQL untuk Golang akan melakukan query tipe data DATE, DATETIME, TIMESTAMP menjadi `[]byte atau []uint8`. Dimana ini bisa dikonversi menjadi `String`, lalu di parsing menjadi `time.Time`
- Namun, hal ini merepotkan jika dilakukan manual, kita bisa meminta Driver MySQL untuk Golang secara otomatis melakukan parsing dengan menambahkan parameter `parseTime=true`

## Nullable Type

- Golang database tidak mengerti dengan tipe data NULL di database
- Oleh karena itu, khusus untuk kolom yang bisa NULL di database, akan jadi masalah jika kita melakukan Scan secara bulat-bulat menggunakan tipe data representasinya di Golang

## Error Data Null

- Konversi secara otomatis data NULL tidak didukung oleh Driver MySQL Golang
- Oleh karena itu, khusus tipe kolom yang bisa NULL, kita perlu menggunakan tipe data yang ada dalam `package.sql`

  ![Tipe_Data_Nullable](img/tipe-data-nullable.jpg)

## SQL Dengan Parameter

- Saat membuat aplikasi, kita tidak mungkin akan melakukan hardcode perintah SQL di kode Golang kita
- Biasanya kita akan menerima input data dari user, lalu membuat perintah SQL dari input user, dan mengirimnya menggunakan perintah SQL

## SQL Injection

- SQL Injection adalah sebuah teknik yang menyalahgunakan sebuah celah keamanan yang terjadi dalam lapisan basis data sebuah aplikasi
- Biasanya SQL Injection dilakukan dengan mengirim input dari user dengan perintah yang salah, sehingga menyebabkan hasil SQL yang kita buat menjadi tidak valid
- SQL Injection sangat berbahaya, jika sampai kita salah membuat SQL, bisa jadi data kita tidak aman. Misalnya bisa login tanpa username dan password nya

## Solusinya?

- Jangan membuat query SQL secara manual dengan menggabungkan String secara bulat-bulat
- Jika kita membutuhkan parameter ketika membuat SQL, kita bisa menggunakan `function Execute atau Query dengan parameter`

## SQL dengan Parameter

- Sekarang kita sudah tahu bahayanya SQL Injection, jika menggabungkan string ketika membuat query
- Jika ada kebutuhan seperti itu, sebenarnya function Exec dan Query memiliki parameter tambahan yang bisa kita gunakan untuk mensubtitusi parameter dari function tersebut ke SQL query yang kita buat
- Untuk menandai sebuah SQL membutuhkan parameter, kita bisa gunakan karakter `? (tanda tanya)`
- Contoh SQL:
  - SELECT username FROM user WHERE username = ? AND password = ? LIMIT 1
  - INSERT INTO user(username, password) VALUES (?,?)
  - Dan lain-lain

## Auto Increment

- Kadang kita membuat sebuah table dengan id auto increment
- Dan kadang pula, kita ingin mengambil data id yang sudah kita insert ke dalam MySQL
- Sebenarnya kita bisa melakukan query ulang ke database menggunakan `SELECT LAST_INSERT_ID()`
- Tapi untungnya di Golang ada cara yang lebih mudah
- Kita bisa menggunakan function `(Result) LastInsertId()` untuk mendapatkan Id terakhir yang dibuat secara auto increment
- Result adalah object yang dikembalikan ketika kita menggunakan function Exec

## Query atau Exec dengan Parameter

- Saat kita menggunakan Fuction Query atau Exec yang menggunakan parameter, sebenarnya implementasi dibawahnya ada fitur `Prepare Statement`
- Jadi tahapan pertama adalah statement nya disiapkan terlebih dahulu, kemudian
  baru diisi (statementnya) dengan parameternya
- Kadang ada kasus kita ingin melakukan beberapa hal yang sama sekaligus, tetapi hanya berbeda di parameternya. Misal insert data yang banyak secara langsung
- Pembuatan Prepare Statement bisa dilakukan dengan manual, tanpa harus menggunakan Query atau Exec dengan parameter

## Prepare Statement

- Saat kita membuat Prepare Statement, secara otomatis akan mengenali koneksi database yang digunakan
- Sehingga ketika kita mengeksekusi Prepare Statement berkali-kali, maka akan menggunakan koneksi yang sama. Hal tersebut bisa lebih efisien, karena pembuatan prepare statement nya hanya sekali diawal saja
- Jika menggunakan Query dan Exec dengan parameter, maka kita tidak bisa menjamin bahwa koneksi yang digunakan akan sama. Oleh karena itu, bisa jadi prepare statement akan selalu dibuat berkali-kali walaupun kita menggunakan SQL yang sama
- Untuk membuat Prepare Statement, kita bisa menggunakan function `(DB) PrepareContext(context, sql)`
- Prepare Statement direpresentasikan dalam struct `database/sql.Stmt`
- Sama seperti resource sql lainnya, Stmt harus di Close() jika sudah tidak digunakan lagi
- Prepare Statement bisa digunakan untuk eksekusi ataupun query

## Database Transaction

- Salah satu fitur andalan di database adalah `Transaction`

## Transaction di Golang

- Secara default, semua perintah SQL yang kita kirim menggunakan Golang akan otomatis di commit, atau istilahnya `auto commit`
- Namun, kita bisa menggunakan fitur transaksi, sehingga SQL yang kita kirim tidak secara otomatis di commit ke database
- Untuk memulai transaksi, kita bisa menggunakan function `(DB) Begin()`, dimana akan menghasilkan struct Tx yang merupakan representasi Transaction
- `Struct Tx` ini yang kita gunakan sebagai pengganti DB untuk melakukan transaksi, dimana hampir semua function di DB ada di Tx, seperti `Exec, Query, atau Prepare`
- Setelah selesai melakukan proses transaksi, kita bisa gunakan function `(Tx) Commit()` untuk melakukan commit atau `Rollback()`

## Repository Pattern

- Dalam buku Domain-Driven Design, Eric Evans menjelaskan bahwa "repository is a mechanism for encapsulating storage, retrieval, and search behavior, which emulates a collection of objects"
- Pattern Repository ini biasanya digunakan sebagai jembatan antar business logic aplikasi kita dengan semua perintah SQL ke database
- Jadi semua perintah SQL akan ditulis di Repository, sedangkan business logic kode program kita hanya cukup menggunakan Repository tersebut

## Entity / Model

- Dalam pemrograman berorientasi object, biasanya sebuah table di database akan selalu dibuat representasinya sebagai class Entity atau Model. Namun di Golang, karena tidak mengenal Class, maka kita representasikan data dalam bentuk Struct
- Ini bisa mempermudah ketika membuat kode program
- Misal ketika kita query ke Repository, dibanding mengembalikan array. Alangkah baiknya Repository melakukan konversi terlebih dahulu ke struct Entity atau Model, sehingga kita tinggal menggunakan objectnya saja
