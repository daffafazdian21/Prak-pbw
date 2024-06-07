package main

import "fmt"

type Customer struct {
	Name    string
	Address string
	Age     int
}

// Improved sayHello method
func (c Customer) sayHello(name string) {
	fmt.Printf("Hello %s, my name is %s (%d years old)\n", name, c.Name, c.Age)
}

func main() {
	// Create an empty Customer struct
	var eko Customer

	fmt.Println("Empty eko:", eko) // Output: Empty eko: {  }

	// Assign values to eko
	eko.Name = "Eko Kurniawan"
	eko.Address = "Indonesia"
	eko.Age = 30

	fmt.Println("eko after assigning values:", eko) // Output: eko after assigning values: { Eko Kurniawan Indonesia 30 }
	fmt.Println("eko.Name:", eko.Name)              // Output: eko.Name: Eko Kurniawan
	fmt.Println("eko.Address:", eko.Address)        // Output: eko.Address: Indonesia
	fmt.Println("eko.Age:", eko.Age)                // Output: eko.Age: 30

	// Create joko with struct literal (preferred)
	joko := Customer{
		Name:    "Joko",
		Address: "Indonesia",
		Age:     30,
	}
	fmt.Println("joko:", joko) // Output: joko: { Joko Indonesia 30 }

	// Create budi using an incorrect constructor approach (not idiomatic Go)
	// Use struct literal instead:
	// budi := Customer{"Budi", "Indonesia", 30}
	var budi Customer // Declare an empty Customer struct
	budi.Name = "Budi"
	budi.Address = "Indonesia"
	budi.Age = 30
	fmt.Println("budi:", budi) // Output: budi: { Budi Indonesia 30 }

	// Call sayHello methods
	budi.sayHello("Agus") // Output: Hello Agus, my name is Budi (30 years old)
	eko.sayHello("Agus")  // Output: Hello Agus, my name is Eko Kurniawan (30 years old)
	joko.sayHello("Agus") // Output: Hello Agus, my name is Joko (30 years old)
}