package faker

import "reflect"

type Phone struct {
	Number string
	Format string
}

func (phone *Phone) Numberʹ(params ...interface{}) string {
	format := phone.Formatʹ()

	types := kindOf(params...)
	if len(types) >= 1 && types[0] == reflect.String {
		format = params[0].(string)
	}

	phone.Number = replaceSymbolWithNumber(format, '#')
	phone.Format = format
	return phone.Number
}

func (phone *Phone) Formatʹ() string {
	format := getData("Phone", "PhoneFormat")
	phone.Format = sample(format)
	return phone.Format
}

// Fake Generate random data for all field
func (p *Phone) Fake() {
	p.Formatʹ()
	p.Numberʹ()
}

// ToJSON Encode name and its fields to JSON.
func (p *Phone) ToJSON() (string, error) {
	return ToJSON(p)
}

// ToXML Encode name and its fields to XML.
func (p *Phone) ToXML() (string, error) {
	return ToJSON(p)
}
