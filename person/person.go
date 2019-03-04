package person

import (
	"reflect"
	"regexp"

	"github.com/sereiner/faker/base"
)

const GENDER_MALE = "male"
const GENDER_FEMALE = "female"

type IPerson interface {
	Name(gender string) string
	FirstName(gender string) string
	FirstNameFemale() string
	FirstNameMale() string
	LastName() string
	Title(gender string) (title string)
	TitleMale() string
	TitleFemale() string
	Suffix() string
}
type Persons struct {
}

func (p *Persons) Name(gender string) string {
	var format string
	switch gender {
	case GENDER_MALE:
		format = base.RandomElement(maleNameFormats)
	case GENDER_FEMALE:
		format = base.RandomElement(femaleNameFormats)
	default:
		format = base.RandomElement(maleNameFormats)
	}
	return p.parse(format)
}

func (p *Persons) parse(format string) (str string) {
	re, _ := regexp.Compile(base.Rex)
	all := re.FindAll([]byte(format), -1)
	for _, v := range all {
		mtV := reflect.ValueOf(&p).Elem()
		s := mtV.MethodByName(string(v)).Call(nil)[0]
		str += s.Interface().(string) + " "
	}
	return str
}

func (p *Persons) FirstName(gender string) string {
	var name string
	switch gender {
	case GENDER_MALE:
		name = p.FirstNameMale()
	case GENDER_FEMALE:
		name = p.FirstNameFemale()
	default:
		name = p.FirstNameMale()
	}
	return name
}

func (p *Persons) FirstNameMale() string {
	return base.RandomElement(firstNameMale)
}

func (p *Persons) FirstNameFemale() string {
	return base.RandomElement(firstNameFemale)
}

func (p *Persons) LastName() string {
	return base.RandomElement(lastName)
}

func (p *Persons) Title(gender string) (title string) {

	switch gender {
	case GENDER_MALE:
		title = base.RandomElement(titleMale)
	case GENDER_FEMALE:
		title = base.RandomElement(titleFemale)
	default:
		title = base.RandomElement(titleMale)
	}
	return title
}

func (p *Persons) TitleMale() string {
	return base.RandomElement(titleMale)
}

func (p *Persons) TitleFemale() string {
	return base.RandomElement(titleFemale)
}

func (p *Persons) Suffix() string {
	return base.RandomElement(suffix)
}
