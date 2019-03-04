package faker

import "github.com/sereiner/faker/person"

func NewPerson() person.IPerson {
	return &person.Persons{}
}
