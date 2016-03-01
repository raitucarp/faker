package address

type Definitions struct {
	PostCode, CityPrefix, CitySuffix,
	City, StreetSuffix, BuildingNumber,
	StreetAddress, SecondaryAddress, County,
	Country, CountryCode, State,
	StateAbbr []string
}

func Export() Definitions {
	def := Definitions{
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
