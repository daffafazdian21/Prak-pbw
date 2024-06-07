package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func cekKategoriUsia(usia int) string {
	if usia < 18 {
		return "anak-anak"
	} else if usia < 60 {
		return "dewasa"
	} else {
		return "lansia"
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Masukkan nama Anda: ")
	nama, _ := reader.ReadString('\n')
	nama = nama[:len(nama)-1] // Hapus newline character

	fmt.Print("Masukkan usia Anda: ")
	usiaStr, _ := reader.ReadString('\n')
	usiaStr = usiaStr[:len(usiaStr)-1] // Hapus newline character
	usia, _ := strconv.Atoi(usiaStr)

	kategoriUsia := cekKategoriUsia(usia)
	fmt.Printf("Selamat datang, %s! Anda termasuk kategori %s.\n", nama, kategoriUsia)
}