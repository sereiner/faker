package person

import (
	"reflect"
	"regexp"

	"github.com/sereiner/faker/base"
)

// GENDER_MALE 男性
const GENDER_MALE = "male"

// GENDER_FEMALE 女性
const GENDER_FEMALE = "female"

var globle *Persons

// Persons .
type Persons struct {
}

func init() {
	globle = &Persons{}
}

//Name .
func Name(gender string) string {
	return globle.Name(gender)
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
func FirstName(gender string) string {
	return globle.FirstName(gender)
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

func FirstNameMale() string {
	return globle.FirstNameMale()
}
func (p *Persons) FirstNameMale() string {
	return base.RandomElement(firstNameMale)
}

func FirstNameFemale() string {
	return globle.FirstNameFemale()
}

func (p *Persons) FirstNameFemale() string {
	return base.RandomElement(firstNameFemale)
}

func LastName() string {
	return globle.LastName()
}

func (p *Persons) LastName() string {
	return base.RandomElement(lastName)
}

func Title(gender string) (title string) {
	return globle.Title(gender)
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

func TitleMale() string {
	return globle.TitleMale()
}

func (p *Persons) TitleMale() string {
	return base.RandomElement(titleMale)
}

func TitleFemale() string {
	return globle.TitleFemale()
}

func (p *Persons) TitleFemale() string {
	return base.RandomElement(titleFemale)
}

func Suffix() string {
	return globle.Suffix()
}

func (p *Persons) Suffix() string {
	return base.RandomElement(suffix)
}
