package phone

import "github.com/sereiner/faker/base"

var operators = []string{
	"134", "135", "136", "137", "138", "139", "147", "150", "151", "152", "157", "158", "159", "175", "178", "182", "183", "184", "187", "188", // China Mobile
	"130", "131", "132", "145", "155", "156", "177", "178", "179", "176", "185", "186", // China Unicom
	"133", "153", "170", "171", "177", "180", "181", "189", // China Telecom
	"170", "171", // virtual operators
}

var formats = []string{
	"########",
}

func PhoneNumber() string {
	op := base.RandomElement(operators)
	form := base.RandomElement(formats)
	return op + base.Numerify(form)
}
