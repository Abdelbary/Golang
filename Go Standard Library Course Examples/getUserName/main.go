package main

import "fmt"

func main() {
	var firstName string
	var lastName string
	fmt.Println("Please Entre Your Name")
	fmt.Scanf("%s %s", &firstName, &lastName)
	fmt.Printf("Hello %s %s", firstName, lastName)
}
