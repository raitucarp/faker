package locales

import (
	"github.com/raitucarp/faker/locales/en"
)

const (
	EN = "EN"
)

type definitions struct {
	EN interface{}
}

func Export() interface{} {
	def := definitions{
		EN: locale_en.Export(),
	}
	return def
}
