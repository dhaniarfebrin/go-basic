package main

import "fmt"

func main() {
	// Declare a string variable named 'name' with zero value (empty string)
	var name string

	// Assign initial value to the name variable
	name = "Dhaniar Febrin Wahyudi"
	fmt.Println(name)

	// Reassign a new value to the existing name variable
	name = "Dhaniar Febrin"
	fmt.Println(name)

	// Reassign another new value to the name variable
	name = "Dhani"
	fmt.Println(name)

	// Declare and initialize an integer variable using short declaration syntax
	age := 20
	fmt.Println(age)

	// Reassign a new value to the age variable
	age = 21
	fmt.Println(age)

	// Declare multiple variables in a single var block with initialization
	var (
		firstName = "Dhaniar"        // String variable for first name
		lastName  = "Febrin Wahyudi" // String variable for last name
	)

	// Print the firstName variable
	fmt.Println(firstName)
	// Print the lastName variable
	fmt.Println(lastName)
}
