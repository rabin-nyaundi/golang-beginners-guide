package main

import "fmt"

type ContactInfo struct {
	zipCode int
	email   string
}

type Person struct {
	firstName string
	lastName  string
	conatct   ContactInfo
}

func main() {

	var alex Person
	name := "rabin"
	fmt.Println(&name)

	alex.firstName = "Alex"
	alex.lastName = "Alia"
	alex.conatct.zipCode = 40200
	alex.conatct.email = "alex@alex.com"

	alex.updatePerson("ppppppppppppppppppppppp")
	alex.print()
}

func (p *Person) updatePerson(newFirstName string) {
	p.firstName = newFirstName
}

func (p Person) print() {
	fmt.Printf("%+v", p)
}
