package faker

import (
	"encoding/json"
	"encoding/xml"
	"math/rand"
	"reflect"
	"strconv"
	"strings"
)

// City contains two field, prefix and suffix.
type City struct {
	Prefix string `json:"prefix"`
	Suffix string `json:"suffix"`
}

// Street is a street address.
type Street struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Prefix  string `json:"prefix"`
	Suffix  string `json:"suffix"`
}

// Geo is geolocation data. It contains latitude and longitude
type Geo struct {
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

// Address is complete type of address. For example user needs random Address data.
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

// Fake fill the address field with random data.
func (addr *Address) Fake() {
	addr.CitySuffixʹ()
	addr.CityPrefixʹ()
	addr.ZipCodeʹ()
}

// ZipCodeʹ generate zip code for Address.
func (addr *Address) ZipCodeʹ(params ...interface{}) string {
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

// Cityʹ generate random city from data for Address field.
func (addr *Address) Cityʹ(params ...interface{}) string {
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

// CityPrefixʹ generate prefix for City.
func (addr *Address) CityPrefixʹ() string {
	// get last name list
	list := getData("Address", "CityPrefix")
	addr.City.Prefix = sample(list)
	return addr.City.Prefix
}

// CitySuffixʹ generate suffix for City.
func (addr *Address) CitySuffixʹ() string {
	// get last name list
	list := getData("Address", "CitySuffix")
	addr.City.Suffix = sample(list)
	return addr.City.Suffix
}

// StreetNameʹ generate random street name for Address,
// based on first name and last name from Name type.
func (addr *Address) StreetNameʹ() string {
	suffix := addr.StreetSuffixʹ()

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

// BuildingNumberʹ generate building number based on template in locale data.
func (addr *Address) BuildingNumberʹ() string {
	format := getData("Address", "BuildingNumber")
	return replaceSymbolWithNumber(sample(format), '#')
}

// StreetAddressʹ generate a street address from locale data.
func (addr *Address) StreetAddressʹ(params ...interface{}) string {
	var address string
	useFullAddress := false

	kinds := kindOf(params...)
	if len(kinds) >= 1 && kinds[0] == reflect.Bool {
		useFullAddress = params[0].(bool)
	}

	r := rand.Intn(3)
	symbol := strings.Repeat("#", 5-r)
	address = symbol + " " + addr.StreetNameʹ()

	if useFullAddress {
		address = address + " " + addr.SecondaryAddressʹ()
	}

	addr.Street.Address = address
	return addr.Street.Address
}

// StreetPrefixʹ generate a street prefix from locale data.
func (addr *Address) StreetPrefixʹ() string {
	prefix := getData("Address", "StreetPrefix")
	addr.Street.Prefix = sample(prefix)
	return addr.Street.Prefix
}

// StreetSuffixʹ generate a street prefix from locale data.
func (addr *Address) StreetSuffixʹ() string {
	suffix := getData("Address", "StreetSuffix")
	addr.Street.Suffix = sample(suffix)
	return addr.Street.Suffix
}

// SecondaryAddressʹ generate a secondary address template from locale data.
func (addr *Address) SecondaryAddressʹ() string {
	formats := getData("Address", "SecondaryAddress")
	secondaryAddress := sample(formats)
	addr.Secondary = replaceSymbolWithNumber(secondaryAddress, '#')
	return addr.Secondary
}

// Countyʹ generate county from locale date.
func (addr *Address) Countyʹ() string {
	county := getData("Address", "County")
	addr.County = sample(county)
	return addr.County
}

// Countryʹ generate country from locale data.
func (addr *Address) Countryʹ() string {
	country := getData("Address", "Country")
	addr.Country = sample(country)
	return addr.Country
}

// CountryCodeʹ generate country code from locale data.
func (addr *Address) CountryCodeʹ() string {
	countryCode := getData("Address", "CountryCode")
	addr.CountryCode = sample(countryCode)
	return addr.CountryCode
}

// Stateʹ generate random state from locale data.
func (addr *Address) Stateʹ() string {
	state := getData("Address", "State")
	addr.State = sample(state)
	return addr.State
}

// StateAbbrʹ generate random abbr of state from locale data.
func (addr *Address) StateAbbrʹ() string {
	stateAbbr := getData("Address", "StateAbbr")
	addr.StateAbbr = sample(stateAbbr)
	return addr.StateAbbr
}

// Latitudeʹ generate random latitude position to Geo.
func (addr *Address) Latitudeʹ() string {
	rnd := rand.Intn(180 * 10000)
	latitude := (float64(rnd) / 10000.0) - 90.0
	addr.Geo.Latitude = strconv.FormatFloat(latitude, 'f', 4, 64)
	return addr.Geo.Latitude

}

// Longitudeʹ generate random longitude position to Geo.
func (addr *Address) Longitudeʹ() string {
	// (faker.random.number(360 * 10000) / 10000.0 - 180.0).toFixed(4);
	rnd := rand.Intn(360 * 10000)
	longitude := (float64(rnd) / 10000.0) - 180.0
	addr.Geo.Longitude = strconv.FormatFloat(longitude, 'f', 4, 64)
	return addr.Geo.Longitude
}

// ToJSON encodes Address to JSON format.
func (addr *Address) ToJSON() (s string, err error) {
	result, err := json.Marshal(addr)
	if err != nil {
		return
	}
	return string(result), err
}

// ToXML encodes Address to XML format.
func (addr *Address) ToXML() (s string, err error) {
	result, err := xml.Marshal(addr)

	if err != nil {
		return
	}
	return string(result), err
}
