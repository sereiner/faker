package main

import (
	"fmt"

	"github.com/sereiner/faker/address"
)

func main() {
	a := new(address.Address)

	fmt.Println(a.CityPrefix())
	a.Hello()
	a.P.Hello()
}
