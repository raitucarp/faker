package address

var City = []string{
	"#{.Address.CityPrefix} #{.Name.FirstName} #{.Address.CitySuffix}",
	"#{.Address.CityPrefix} #{.Name.FirstName}",
	"#{.Name.FirstName} #{.Address.CitySuffix}",
	"#{.Name.LastName} #{.Address.CitySuffix}",
}
