package chapter10

import (
	"fmt"
	"strings"
)

type Person struct {
	firstName string
	lastName  string
}

//noinspection GoUnusedFunction
func upPerson(p Person) {
	p.firstName = strings.ToUpper(p.firstName)
	p.lastName = strings.ToUpper(p.lastName)
}

//noinspection GoUnusedFunction
func lowerPerson(p *Person) {
	p.firstName = strings.ToLower(p.firstName)
	p.lastName = strings.ToLower(p.lastName)
}

//noinspection GoUnusedFunction
func PersonTest() {
	// 1- struct as a value type:
	var pers1 Person
	pers1.firstName = "Chris"
	pers1.lastName = "Woodward"
	upPerson(pers1)
	fmt.Printf("The name of the person is %s %s\n", pers1.firstName, pers1.lastName)
	lowerPerson(&pers1)
	fmt.Printf("The name of the person is %s %s\n", pers1.firstName, pers1.lastName)

	// 2 - struct as a pointer:
	pers2 := new(Person)
	pers2.firstName = "Chris"
	pers2.lastName = "Woodward"
	upPerson(*pers2)
	fmt.Printf("The name of the person is %s %s\n", pers2.firstName, pers2.lastName)
	// 3 - struct as a literal:
	pers3 := &Person{"Chris", "Woodward"}
	upPerson(*pers3)
	fmt.Printf("The name of the person is %s %s\n", pers3.firstName, pers3.lastName)
}
