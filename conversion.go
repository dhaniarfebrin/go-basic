package main

import "fmt"

func main() {
	// about conversion
	var i int = 10
	var f float64 = float64(i) // Convert int to float64
	var u uint = uint(i)       // Convert int to uint
	fmt.Println("Integer:", i)
	fmt.Println("Float64:", f)
	fmt.Println("Unsigned Integer:", u)

	// about type conversion
	var str string = "123"
	var num int
	fmt.Sscanf(str, "%d", &num) // Convert string to int
	fmt.Println("Converted String to Int:", num)

	var name = "Dhaniar Febrin Wahyudi"
	var d = name[0]
	var dString = string(d) // Convert byte to string
	fmt.Println("First character as string:", dString)
}
