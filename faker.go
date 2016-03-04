package faker

import "github.com/raitucarp/faker/locales"

type Gender int

const (
	Male Gender = iota
	Female
)

var gender Gender
var data interface{}

// Faker is complete structure for all type.
type Faker struct {
	Name     Name
	Address  Address
	Commerce Commerce
	Company  Company
	Date     Date
	Finance  Finance
	Hacker   Hacker
	Image    Image
	Internet Internet
	Lorem    Lorem
	Phone    Phone
	// temp data handling
	data interface{}
}

// Fake Generate random data, fake it all.
func (f *Faker) Fake() {
	f.Name.Fake()
	f.Address.Fake()
	f.Commerce.Fake()
	f.Company.Fake()
	f.Finance.Fake()
	f.Hacker.Fake()
	f.Image.Fake()
	f.Internet.Fake()
	f.Lorem.Fake()
	f.Phone.Fake()

}

func (f *Faker) generate() {
	err := SetLocale(locales.EN)
	if err != nil {
		panic(err)
	}
}

// New Create new Faker
func New() (f Faker) {
	f.generate()

	return
}
