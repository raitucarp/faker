package faker

import (
	"encoding/json"
	"encoding/xml"
	"math/rand"
	"reflect"
	"strconv"
	"strings"
)

type City struct {
	Prefix string `json:"prefix"`
	Suffix string `json:"suffix"`
}

type Street struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Prefix  string `json:"prefix"`
	Suffix  string `json:"suffix"`
}

type Geo struct {
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

type Address struct {
	ZipCode     string `json:"zip_code"`
	City        City   `json:"city"`
	Street      Street `json:"street"`
	Secondary   string `json:"secondary_address"`
	County      string `json:"county"`
	Country     string `json:"country"`
	CountryCode string `json:"country_code"`
	State       string `json:"state"`
	StateAbbr   string `json:"state_abbr"`
	Geo         Geo    `json:"geo"`
}

func (addr *Address) Fake() {
	addr.CitySuffix_()
	addr.CityPrefix_()
	addr.ZipCode_()
}

func (addr *Address) ZipCode_(params ...interface{}) string {
	var format string

	kinds := kindOf(params...)
	if len(kinds) >= 1 && kinds[0] == reflect.String {
		format = params[0].(string)
	}

	// get last name list
	list := getData("Address", "PostCode")

	if len(list) > 1 {
		format = sample(list)
	} else {
		format = list[0]
	}

	addr.ZipCode = replaceSymbols(format)
	return addr.ZipCode
}

func (addr *Address) City_(params ...interface{}) string {
	var format string
	// get last name list
	list := getData("Address", "City")

	kinds := kindOf(params...)
	if len(kinds) >= 1 && kinds[0] == reflect.String {
		format = params[0].(string)
	} else {
		format = sample(list)
	}

	result, err := Parse(format)
	if err != nil {
		panic(err)
	}
	return result
}

func (addr *Address) CityPrefix_() string {
	// get last name list
	list := getData("Address", "CityPrefix")
	addr.City.Prefix = sample(list)
	return addr.City.Prefix
}

func (addr *Address) CitySuffix_() string {
	// get last name list
	list := getData("Address", "CitySuffix")
	addr.City.Suffix = sample(list)
	return addr.City.Suffix
}

func (addr *Address) StreetName_() string {
	suffix := addr.StreetSuffix_()

	if suffix != "" {
		suffix = " " + suffix
	}

	name := Name{}
	randCandidate := []string{
		name.LastName_() + suffix,
		name.FirstName_() + suffix,
	}

	addr.Street.Name = sample(randCandidate)

	return addr.Street.Name
}

func (addr *Address) BuildingNumber_() string {
	format := getData("Address", "BuildingNumber")
	return replaceSymbolWithNumber(sample(format), '#')
}

func (addr *Address) StreetAddress_(params ...interface{}) string {
	var address string
	useFullAddress := false

	kinds := kindOf(params...)
	if len(kinds) >= 1 && kinds[0] == reflect.Bool {
		useFullAddress = params[0].(bool)
	}

	r := rand.Intn(3)
	symbol := strings.Repeat("#", 5-r)
	address = symbol + " " + addr.StreetName_()

	if useFullAddress {
		address = address + " " + addr.SecondaryAddress_()
	}

	addr.Street.Address = address
	return addr.Street.Address
}

func (addr *Address) StreetPrefix_() string {
	prefix := getData("Address", "StreetPrefix")
	addr.Street.Prefix = sample(prefix)
	return addr.Street.Prefix
}

func (addr *Address) StreetSuffix_() string {
	suffix := getData("Address", "StreetSuffix")
	addr.Street.Suffix = sample(suffix)
	return addr.Street.Suffix
}

func (addr *Address) SecondaryAddress_() string {
	formats := getData("Address", "SecondaryAddress")
	secondaryAddress := sample(formats)
	addr.Secondary = replaceSymbolWithNumber(secondaryAddress, '#')
	return addr.Secondary
}

func (addr *Address) County_() string {
	county := getData("Address", "County")
	addr.County = sample(county)
	return addr.County
}

func (addr *Address) Country_() string {
	country := getData("Address", "Country")
	addr.Country = sample(country)
	return addr.Country
}

func (addr *Address) CountryCode_() string {
	countryCode := getData("Address", "CountryCode")
	addr.CountryCode = sample(countryCode)
	return addr.CountryCode
}

func (addr *Address) State_() string {
	state := getData("Address", "State")
	addr.State = sample(state)
	return addr.State
}

func (addr *Address) StateAbbr_() string {
	stateAbbr := getData("Address", "StateAbbr")
	addr.StateAbbr = sample(stateAbbr)
	return addr.StateAbbr
}

func (addr *Address) Latitude_() string {
	rnd := rand.Intn(180 * 10000)
	latitude := (float64(rnd) / 10000.0) - 90.0
	addr.Geo.Latitude = strconv.FormatFloat(latitude, 'f', 4, 64)
	return addr.Geo.Latitude

}

func (addr *Address) Longitude_() string {
	// (faker.random.number(360 * 10000) / 10000.0 - 180.0).toFixed(4);
	rnd := rand.Intn(360 * 10000)
	longitude := (float64(rnd) / 10000.0) - 180.0
	addr.Geo.Longitude = strconv.FormatFloat(longitude, 'f', 4, 64)
	return addr.Geo.Longitude
}

func (addr *Address) ToJSON() (s string, err error) {
	result, err := json.Marshal(addr)
	if err != nil {
		return
	}
	return string(result), err
}

func (addr *Address) ToXML() (s string, err error) {
	result, err := xml.Marshal(addr)

	if err != nil {
		return
	}
	return string(result), err
}
