package faker

import "github.com/raitucarp/faker/locales"

type Gender int

const (
	Male Gender = iota
	Female
	RandomGender
)

var gender Gender

type Faker struct {
	Name    Name
	Address Address

	// temp data handling
	data interface{}
}

type Definitions map[string]interface{}

var data interface{}

func (f *Faker) generate() {
	err := SetLocale(locales.EN)
	if err != nil {
		panic(err)
	}
}

func New() (f Faker) {
	f.generate()

	return
}

/*
func SetGender(g Gender) {
	gender = g
}*/
