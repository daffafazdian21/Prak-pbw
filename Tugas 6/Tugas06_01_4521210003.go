package main

import (
	"errors" // Import the standard errors package
	"fmt"
)

// Define custom error types for better error handling
var (
	ValidationError = errors.New("validation error")
	NotFoundError   = errors.New("not found error")
)

// GetBuilding function takes an ID string and returns a building or an error
func GetBuilding(id string) (*Building, error) { // Return pointer to Building struct (optional)
	if id == "" {
		return nil, ValidationError // Return nil and ValidationError
	}

	if id != "daffa" {
		return nil, NotFoundError // Return nil and NotFoundError
	}

	// Simulate successful building retrieval (replace with actual logic)
	building := &Building{ID: id, Name: "affa"} // Create a Building instance
	return building, nil                        // Return the building and nil error
}

// Building struct (optional)
type Building struct {
	ID   string
	Name string
}

func main() {
	building, err := GetBuilding("daffa") // Call GetBuilding and handle potential errors

	if err != nil {
		switch err { // Use switch statement for more efficient error handling
		case ValidationError:
			fmt.Println("validation error")
		case NotFoundError:
			fmt.Println("not found error")
		default:
			fmt.Println("unknown error:", err) // Include the underlying error for debugging
		}
	} else {
		fmt.Printf("Building retrieved: %+v\n", building) // Print building details (optional)
	}
}
