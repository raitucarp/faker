package address

type definitions struct {
	PostCode, CityPrefix, CitySuffix,
	City, StreetSuffix, BuildingNumber,
	StreetAddress, SecondaryAddress, County,
	Country, CountryCode, State,
	StateAbbr []string
}

func Export() interface{} {
	def := definitions{
		PostCode:         PostCode,
		CityPrefix:       CityPrefix,
		CitySuffix:       CitySuffix,
		City:             City,
		StreetSuffix:     StreetSuffix,
		BuildingNumber:   BuildingNumber,
		StreetAddress:    StreetAddress,
		SecondaryAddress: SecondaryAddress,
		County:           County,
		Country:          Country,
		CountryCode:      CountryCode,
		State:            State,
		StateAbbr:        StateAbbr,
	}

	return def
}

/*type AddressMethod struct {}

func (m *AddressMethod) PostCode()  {
	return postCode
}

func (m *AddressMethod) CityPrefix()  {
	return cityPrefix
}

func (m *AddressMethod) City()  {
	return city
}

func (m *AddressMethod) CitySuffix()  {
	return citySuffix
}
*/
