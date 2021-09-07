package main

import (
	"fmt"

	"datatypes/organization"
)

func main() {
	p := organization.NewPerson("Mahmoud", "Abdelbary")
	err := p.SetTwitterHandler("@mahmoud.abdelbary")

	if err != nil {
		fmt.Printf("error accqured in setTwitterHandler func : %s\n", err.Error())
	}

	fmt.Println(p.GetTwitterHandler())
	fmt.Println(p.FullName())
	fmt.Println(p.ID())
}
