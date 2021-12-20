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
