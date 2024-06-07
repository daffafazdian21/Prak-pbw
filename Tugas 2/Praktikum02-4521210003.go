package main

import (
	"fmt"
)

// Struct untuk merepresentasikan data mahasiswa
type Mahasiswa struct {
	Nama    string
	NPM     string
	Jurusan string
}

func () {
	// Membuat map untuk menyimpan data mahasiswa dengan NPM sebagai kunci
	dataMahasiswa := make(map[string]Mahasiswa)

	// Menambahkan data mahasiswa ke dalam map
	dataMahasiswa["4521210003"] = Mahasiswa{"daffa", "4521210003", "Teknik Informatika"}
	dataMahasiswa["4521210001"] = Mahasiswa{"alam", "4521210001", "Sistem Informasi"}

	// Menampilkan daftar nama mahasiswa dengan perulangan
	fmt.Println("Daftar Nama Mahasiswa:")
	for _, mahasiswa := range dataMahasiswa {
		fmt.Println(mahasiswa.Nama)
	}

	// Mencari data mahasiswa berdasarkan NPM
	cariNPM := "4521210003"
	if mahasiswa, found := dataMahasiswa[cariNPM]; found {
		fmt.Printf("\nData Mahasiswa dengan NPM %s ditemukan:\n", cariNPM)
		fmt.Printf("Nama: %s\nNPM: %s\nJurusan: %s\n", mahasiswa.Nama, mahasiswa.NPM, mahasiswa.Jurusan)
	} else {
		fmt.Printf("\nData Mahasiswa dengan NPM %s tidak ditemukan.\n", cariNPM)
	}
}
