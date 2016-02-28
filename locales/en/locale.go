package locale_en

import (
	"github.com/raitucarp/faker/locales/en/address"
	"github.com/raitucarp/faker/locales/en/commerce"
	"github.com/raitucarp/faker/locales/en/company"
	"github.com/raitucarp/faker/locales/en/name"
)

type definitions struct {
	Name     interface{}
	Address  interface{}
	Commerce interface{}
	Company  interface{}
}

func Export() interface{} {
	d := definitions{
		Name:     name.Export(),
		Address:  address.Export(),
		Commerce: commerce.Export(),
		Company:  company.Export(),
	}

	return d
}
