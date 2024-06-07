package main

import "fmt"

func main() {
	fmt.Println("Bubble Sort")
	var arrayNumber = [20]int{9, 5, 3, 7, 2, 1, 8, 4, 6} // Misalnya, array yang akan diurutkan
	fmt.Println("Sebelum dilakukan pengurutan:")
	fmt.Println(arrayNumber)

	// Bubble Sort
	for i := 0; i < len(arrayNumber); i++ {
		for j := 0; j < len(arrayNumber)-1; j++ {
			if arrayNumber[j] > arrayNumber[j+1] {
				// Tukar nilai
				temp := arrayNumber[j]
				arrayNumber[j] = arrayNumber[j+1]
				arrayNumber[j+1] = temp
			}
		}
	}

	fmt.Println("Setelah dilakukan pengurutan:")
	fmt.Println(arrayNumber)
}
