package locales

import (
	"github.com/raitucarp/faker/locales/en"
)

const (
	EN = "EN"
)

type Definitions struct {
	EN locale_en.Definitions
}

func Export() Definitions {
	def := Definitions{
		EN: locale_en.Export(),
	}
	return def
}
