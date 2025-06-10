package main

import "fmt"

func main() {
	// Declare a constant named 'pi' with a value of 3.14
	const pi = 3.14

	// Print the value of the constant pi
	fmt.Println("Value of pi:", pi)

	// Attempting to change the value of a constant will result in an error
	// Uncommenting the next line will cause a compilation error
	// pi = 3.14159 // This line will cause an error

	// Declare multiple constants in a single const block
	const (
		gravity = 9.81  // Acceleration due to gravity in m/s^2
		e       = 2.718 // Euler's number
	)

	// Print the values of the constants
	fmt.Println("Gravity:", gravity)
	fmt.Println("Euler's number:", e)
}
