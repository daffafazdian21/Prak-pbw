package main

import (
	"fmt"
	"os"
)

func main() {

	var err error
	fileinfo, err := os.Stat("PEBEWE")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	if fileinfo.IsDir() {
		fmt.Println("PEBEWE adalah sebuah direktori")
	} else {
		fmt.Println("PEBEWE adalah seubah file")
	}
}
