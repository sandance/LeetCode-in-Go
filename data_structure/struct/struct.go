package main

import (
	"fmt"
	"strings"
)

type Person struct {
	firstName string
	lastName  string
}

func upPerson(p *Person) {
	p.firstName = strings.ToUpper(p.firstName)
	p.lastName = strings.ToUpper(p.lastName)
}

func main() {
	var pers1 Person
	pers1.firstName = "Chris"
	pers1.lastName = "Kennedy"
	upPerson(&pers1)

	pers2 := new(Person)
	pers2.firstName = "Chris"
	pers2.lastName = "WoodWard"
	upPerson(pers2)

	fmt.Println(pers2)
}
