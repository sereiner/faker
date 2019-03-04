package address

import (
	"reflect"
	"regexp"

	"github.com/sereiner/faker/base"
)

type IAddress interface {
	CityPrefix() string
	State() string
	StateAbbr() string
	CitySuffix() string
	StreetSuffix() string
	BuildingNumber() string
	City() string
	StreetName() string
	StreetAddress() string
	Postcode() string
	Address() string
	Country() string
	Latitude() float32
	Longitude() float32
}

type Address struct {
}

//CityPrefix 城市前缀
func (a *Address) CityPrefix() string {
	return base.RandomElement(cityPrefix)
}

func (a *Address) State() string {
	return base.RandomElement(state)
}

func (a *Address) StateAbbr() string {
	return base.RandomElement(stateAbbr)
}

func (a *Address) CitySuffix() string {
	return base.RandomElement(citySuffix)
}

func (a *Address) StreetSuffix() string {
	return base.RandomElement(streetSuffix)
}

func (a *Address) BuildingNumber() string {
	format := base.RandomElement(buildingNumber)
	return base.Numerify(format)
}

func (a *Address) City() string {
	format := base.RandomElement(cityFormats)
	return a.parse(format)
}

func (a *Address) parse(format string) (str string) {
	re, _ := regexp.Compile(base.Rex)
	all := re.FindAll([]byte(format), -1)
	for _, v := range all {
		mtV := reflect.ValueOf(&a).Elem()
		s := mtV.MethodByName(string(v)).Call(nil)[0]
		str += s.Interface().(string) + " "
	}
	return str
}

func (a *Address) LastName() string {
	return base.RandomElement(lastName)
}

func (a *Address) StreetName() string {
	format := base.RandomElement(streetNameFormats)
	return a.parse(format)
}

func (a *Address) StreetAddress() string {
	format := base.RandomElement(streetAddressFormats)
	return a.parse(format)
}

func (a *Address) Postcode() string {
	format := base.RandomElement(postcode)
	return base.Numerify(format)
}

func (a *Address) Address() string {
	format := base.RandomElement(addressFormats)
	return a.parse(format)
}

func (a *Address) Country() string {
	return base.RandomElement(country)

}

func (a *Address) Latitude() float32 {
	return float32(base.RandInt(-90, 90)) + base.RandomFloat32()
}

func (a *Address) Longitude() float32 {
	return float32(base.RandInt(-180, 180)) + base.RandomFloat32()
}
