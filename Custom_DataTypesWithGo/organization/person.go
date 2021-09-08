package organization

import (
	"errors"
	"fmt"
	"strings"
)

type TwitterHandler string

type Identifiable interface {
	ID() string
}
type Citizen interface {
	Identifiable
	Country() string
}
type Name struct {
	firstName string
	lastName  string
}

type Employee struct {
	Name
}

type Person struct {
	Name
	twitterHandler TwitterHandler
	Citizen
}

type socialSecurityNumber string

func NewSocialSecurityNumber(value string) Citizen {
	return socialSecurityNumber(value)
}
func (ssn socialSecurityNumber) ID() string {
	return string(ssn)
}

func (ssn socialSecurityNumber) Country() string {
	return "United States"
}

type eurpianUnionIdentifier struct {
	id      string
	country string
}

func NewEuropianUnionIdentifier(id, country string) Citizen {
	return eurpianUnionIdentifier{
		id:      id,
		country: country,
	}
}
func (euID eurpianUnionIdentifier) ID() string {
	return string(euID.id)
}

func (euID eurpianUnionIdentifier) Country() string {
	return "Eurpian Union"
}

/*
func (p *Person) ID() string {
	return "123"
}
*/
func NewPerson(firstName string, lastName string, citizen Citizen) Person {
	return Person{
		Name: Name{
			firstName: firstName,
			lastName:  lastName},
		twitterHandler: "",
		Citizen:        citizen,
	}
}

func (name *Name) FullName() string {
	return fmt.Sprintf("%s %s\n", name.firstName, name.lastName)
}

func (p *Person) SetTwitterHandler(th TwitterHandler) error {
	if !strings.HasPrefix(string(th), "@") {
		return errors.New("twitter handler must start with @ symbol")
	}

	p.twitterHandler = th
	return nil
}

func (p *Person) GetTwitterHandler() TwitterHandler {
	return p.twitterHandler
}

func (th TwitterHandler) GetRedirectURL() string {
	cleanHandler := strings.TrimPrefix(string(th), "@")

	return fmt.Sprintf("https://www.twitter.com/%s", cleanHandler)
}
