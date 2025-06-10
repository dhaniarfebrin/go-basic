package main

import "fmt"

func main() {
	type NoKTP string

	var noKtp NoKTP = "1234567890"
	fmt.Println("Nomor KTP:", noKtp)
	fmt.Println(NoKTP("9876543210")) // Convert string to NoKTP type
}
