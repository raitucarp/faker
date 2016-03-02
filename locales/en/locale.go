package locale_en

import (
	"github.com/raitucarp/faker/locales/en/address"
	"github.com/raitucarp/faker/locales/en/commerce"
	"github.com/raitucarp/faker/locales/en/company"
	"github.com/raitucarp/faker/locales/en/date"
	"github.com/raitucarp/faker/locales/en/finance"
	"github.com/raitucarp/faker/locales/en/hacker"
	"github.com/raitucarp/faker/locales/en/internet"
	"github.com/raitucarp/faker/locales/en/lorem"
	"github.com/raitucarp/faker/locales/en/name"
	"github.com/raitucarp/faker/locales/en/phone"
)

type Definitions struct {
	Address  address.Definitions
	Commerce commerce.Definitions
	Company  company.Definitions
	Date     date.Definitions
	Finance  finance.Definitions
	Hacker   hacker.Definitions
	Name     name.Definitions
	Internet internet.Definitions
	Lorem    lorem.Definitions
	Phone    phone.Definitions
}

func Export() Definitions {
	d := Definitions{
		Name:     name.Export(),
		Address:  address.Export(),
		Commerce: commerce.Export(),
		Company:  company.Export(),
		Date:     date.Export(),
		Finance:  finance.Export(),
		Hacker:   hacker.Export(),
		Internet: internet.Export(),
		Lorem:    lorem.Export(),
		Phone:    phone.Export(),
	}

	return d
}
