package organization

import (
	"errors"
	"fmt"
	"strings"
)

type Identifiable interface {
	ID() string
}

type Person struct {
	firstName      string
	lastName       string
	twitterHandler string
}

func (p Person) ID() string {
	return "123"
}

func NewPerson(firstName string, lastName string) Person {
	return Person{firstName, lastName, ""}
}

func (p Person) FullName() string {
	return fmt.Sprintf("%s %s\n", p.firstName, p.lastName)
}

func (p *Person) SetTwitterHandler(twitterHandler string) error {
	if !strings.HasPrefix(twitterHandler, "@") {
		return errors.New("twitter handler must start with @ symbol")
	}

	p.twitterHandler = twitterHandler
	return nil
}

func (p Person) GetTwitterHandler() string {
	return p.twitterHandler
}
