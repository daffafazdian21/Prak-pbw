package main

import (
	"fmt"
	"os"
)

func main() {
	var err error

	// Mengubah izin direktori
	err = os.Chmod("PEBEWE", 2001) // Change permissions to 0009
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Izin 'PEBEWE' telah diubah menjadi 2003")
}
