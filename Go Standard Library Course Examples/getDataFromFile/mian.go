package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("data.csv")

	if err != nil {
		log.Fatal("error openning file data.csv")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		var age int
		var name string
		n, err := fmt.Sscanf(scanner.Text(), "%s is %d years old", &name, &age)

		if err != nil {
			panic(err)
		}
		if n == 2 {
			fmt.Printf("name %s is %d years old\n", name, age)
		}
	}
}
