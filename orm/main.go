package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/lunny/log"
	"github.com/sereiner/faker/base"
	"github.com/sereiner/faker/company"
	"github.com/sereiner/faker/internet"
	"github.com/sereiner/faker/person"
	"github.com/sereiner/faker/phone"
	"github.com/sereiner/faker/text"
)

type SingleTable struct {
	ID          int
	Key1        string
	Key2        int
	Key3        string
	KeyPart1    string
	KeyPart2    string
	KeyPart3    string
	CommonField string
}

var seqs = func() func() int {
	index := base.RandomDigitNotNull()
	return func() int {
		index = index + 1
		return index
	}
}

func main() {
	db, err := gorm.Open("mysql", "root:root@/test?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()
	if err != nil {
		log.Error(err)
	}
	seq := seqs()

	for i := 0; i < 100000; i++ {
		single := SingleTable{}
		single.Key1 = person.LastName()
		single.Key2 = seq()
		single.Key3 = company.Company()
		single.KeyPart1 = internet.Email()
		single.KeyPart2 = company.JobTitle()
		single.KeyPart3 = phone.PhoneNumber()
		single.CommonField = text.RealText(5)

		if err = db.Create(&single).Error; err != nil {
			log.Error(err)
		}
	}

}
