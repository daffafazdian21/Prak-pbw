package main

import (
	"fmt"
)

func () {
	// Meminta input nama dari pengguna
	fmt.Print("Masukkan nama Anda: ")
	var nama string
	fmt.Scanln(&nama)

	// Meminta input usia dari pengguna
	fmt.Print("Masukkan usia Anda: ")
	var usia int
	fmt.Scanln(&usia)

	// Menentukan kategori usia
	var kategoriUsia string
	if usia < 18 {
		kategoriUsia = "Anak-anak"
	} else {
		kategoriUsia = "Dewasa"
	}

	// Menampilkan pesan selamat datang dan kategori usia
	fmt.Printf("Selamat datang, %s! Anda termasuk dalam kategori usia %s.\n", nama, kategoriUsia)
}