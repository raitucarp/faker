package faker

import (
	"encoding/json"
	"encoding/xml"
	"math/rand"
	"strconv"
)

type City struct {
	Prefix string
	Suffix string
}

type Street struct {
	Name    string
	Address string
	Prefix  string
	Suffix  string
}

type Address struct {
	ZipCode     string
	City        City
	Street      Street
	Secondary   string
	County      string
	Country     string
	CountryCode string
	State       string
	StateAbbr   string
	Latitude    string
	Longitude   string
}

func (addr *Address) Fake() {
	addr.CitySuffix_()
	addr.CityPrefix_()
	addr.ZipCode_()
}

func (addr *Address) ZipCode_() string {
	// get last name list
	list := getData("Address", "PostCode")

	var format string
	if len(list) > 1 {
		format = sample(list)
	} else {
		format = list[0]
	}

	addr.ZipCode = replaceSymbols(format)
	return addr.ZipCode
}

func (addr *Address) CityPrefix_() string {
	// get last name list
	list := getData("Address", "CityPrefix")
	addr.City.Prefix = sample(list)
	return addr.City.Prefix
}

func (addr *Address) City_() string {
	// get last name list
	list := getData("Address", "City")
	city := sample(list)
	result, err := Fake(city)
	if err != nil {
		panic(err)
	}
	return result
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
	return replaceSymbolsWithNumber(sample(format), '#')
}

func (addr *Address) StreetAddress_() string {
	addr.Street.Address = addr.BuildingNumber_() + " " + addr.StreetName_()
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
	addr.Secondary = replaceSymbolsWithNumber(secondaryAddress, '#')
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
	addr.Latitude = strconv.FormatFloat(latitude, 'f', 4, 64)
	return addr.Latitude

}

func (addr *Address) Longitude_() string {
	// (faker.random.number(360 * 10000) / 10000.0 - 180.0).toFixed(4);
	rnd := rand.Intn(360 * 10000)
	longitude := (float64(rnd) / 10000.0) - 180.0
	addr.Longitude = strconv.FormatFloat(longitude, 'f', 4, 64)
	return addr.Longitude
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
