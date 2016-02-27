package locale_en

import (
	"github.com/raitucarp/faker/locales/en/address"
	"github.com/raitucarp/faker/locales/en/name"
)

type definitions struct {
	Name    interface{}
	Address interface{}
}

func Export() interface{} {
	d := definitions{
		Name:    name.Export(),
		Address: address.Export(),
	}

	return d
}
