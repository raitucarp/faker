package faker

import "reflect"

type CatchPhrase struct {
	Adjective  string
	Descriptor string
	Noun       string
}

type BusinessSlogan struct {
	Adjective string
	Buzz      string
	Noun      string
}

type Company struct {
	Name           string
	Suffix         string
	CatchPhrase    CatchPhrase
	BusinessSlogan BusinessSlogan
}

func (company *Company) Suffixesʹ() []string {
	suffixes := getData("Company", "Suffix")
	return suffixes
}

func (company *Company) CompanyNameʹ(params ...interface{}) string {
	formats := getData("Company", "Name")
	format := sample(formats)

	kinds := kindOf(params...)
	if len(kinds) >= 1 && kinds[0] == reflect.String {
		format = params[0].(string)
	}

	name, err := Parse(format)
	if err != nil {
		panic(err)
	}
	company.Name = name
	return company.Name
}

func (company *Company) CompanySuffixʹ() string {
	suffixes := company.Suffixesʹ()
	company.Suffix = sample(suffixes)
	return company.Suffix
}

func (company *Company) CatchPhraseʹ() string {
	template := getData("Company", "CatchPhrase", "Template")
	catchPhrase := sample(template)
	result, err := Parse(catchPhrase)
	if err != nil {
		panic(err)
	}

	return result
}

func (company *Company) BSʹ() string {
	template := getData("Company", "BusinessSlogan", "Template")
	bs := sample(template)
	result, err := Parse(bs)
	if err != nil {
		panic(err)
	}

	return result
}

func (company *Company) CatchPhraseAdjectiveʹ() string {
	adjective := getData("Company", "CatchPhrase", "Adjective")
	company.CatchPhrase.Adjective = sample(adjective)
	return company.CatchPhrase.Adjective
}

func (company *Company) CatchPhraseDescriptorʹ() string {
	descriptor := getData("Company", "CatchPhrase", "Descriptor")
	company.CatchPhrase.Descriptor = sample(descriptor)
	return company.CatchPhrase.Descriptor
}

func (company *Company) CatchPhraseNounʹ() string {
	noun := getData("Company", "CatchPhrase", "Noun")
	company.CatchPhrase.Noun = sample(noun)
	return company.CatchPhrase.Noun
}

func (company *Company) BsAdjectiveʹ() string {
	adjective := getData("Company", "BusinessSlogan", "Adjective")
	company.BusinessSlogan.Adjective = sample(adjective)
	return company.BusinessSlogan.Adjective
}

func (company *Company) BsBuzzʹ() string {
	verb := getData("Company", "BusinessSlogan", "Buzz")
	company.BusinessSlogan.Buzz = sample(verb)
	return company.BusinessSlogan.Buzz
}

func (company *Company) BsNounʹ() string {
	noun := getData("Company", "BusinessSlogan", "Noun")
	company.BusinessSlogan.Noun = sample(noun)
	return company.BusinessSlogan.Noun
}

// Fake Generate random data for all field
func (company *Company) Fake() {
	company.Suffixesʹ()
	company.CompanyNameʹ()
	company.CompanySuffixʹ()
	company.CatchPhraseʹ()
	company.BSʹ()
	company.CatchPhraseAdjectiveʹ()
	company.CatchPhraseDescriptorʹ()
	company.CatchPhraseNounʹ()
	company.BsAdjectiveʹ()
	company.BsBuzzʹ()
	company.BsNounʹ()
}

// ToJSON Encode name and its fields to JSON.
func (company *Company) ToJSON() (string, error) {
	return ToJSON(company)
}

// ToXML Encode name and its fields to XML.
func (company *Company) ToXML() (string, error) {
	return ToXML(company)
}
