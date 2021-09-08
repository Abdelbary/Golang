package main

import (
	"fmt"

	"datatypes/organization"
)

func main() {
	p := organization.NewPerson("Mahmoud", "Abdelbary", organization.NewEuropianUnionIdentifier("11-22-33", "Germany"))
	err := p.SetTwitterHandler("@mahmoud.abdelbary")

	if err != nil {
		fmt.Printf("error accqured in setTwitterHandler func : %s\n", err.Error())
	}

	fmt.Println(p.GetTwitterHandler())
	fmt.Println(p.GetTwitterHandler().GetRedirectURL())
	fmt.Println(p.FullName())
	fmt.Println(p.Country())
	fmt.Println(p.ID())
}
