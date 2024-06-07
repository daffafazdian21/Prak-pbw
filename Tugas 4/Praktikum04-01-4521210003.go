package main

import "fmt"

type person struct {
	name string
	age  int // Corrected typo (space before int)
}

func main() {
	fmt.Println(person{"daffa", 21})            // Use struct literal
	fmt.Println(person{name: "ghitbah", age: 17}) // Named fields for clarity
	fmt.Println(person{"Clarissa", 0})         // Default age for missing field
	fmt.Println(person{name: "agus", age: 40})

	s := person{name: "Bobo", age: 90}
	fmt.Println(s.name) // Print name

	// Access and modify age using a pointer
	sp := &s
	sp.age = 51
	fmt.Println("sp.age:", sp.age)
	fmt.Println("s.age: ", s.age) // Modified age is reflected here
}