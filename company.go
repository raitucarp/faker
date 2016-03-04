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

func (comp *Company) Suffixesʹ() []string {
	suffixes := getData("Company", "Suffix")
	return suffixes
}

func (comp *Company) CompanyNameʹ(params ...interface{}) string {
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
	comp.Name = name
	return comp.Name
}

func (comp *Company) CompanySuffixʹ() string {
	suffixes := comp.Suffixesʹ()
	comp.Suffix = sample(suffixes)
	return comp.Suffix
}

func (comp *Company) CatchPhraseʹ() string {
	template := getData("Company", "CatchPhrase", "Template")
	catchPhrase := sample(template)
	result, err := Parse(catchPhrase)
	if err != nil {
		panic(err)
	}

	return result
}

func (comp *Company) BSʹ() string {
	template := getData("Company", "BusinessSlogan", "Template")
	bs := sample(template)
	result, err := Parse(bs)
	if err != nil {
		panic(err)
	}

	return result
}

func (comp *Company) CatchPhraseAdjectiveʹ() string {
	adjective := getData("Company", "CatchPhrase", "Adjective")
	comp.CatchPhrase.Adjective = sample(adjective)
	return comp.CatchPhrase.Adjective
}

func (comp *Company) CatchPhraseDescriptorʹ() string {
	descriptor := getData("Company", "CatchPhrase", "Descriptor")
	comp.CatchPhrase.Descriptor = sample(descriptor)
	return comp.CatchPhrase.Descriptor
}

func (comp *Company) CatchPhraseNounʹ() string {
	noun := getData("Company", "CatchPhrase", "Noun")
	comp.CatchPhrase.Noun = sample(noun)
	return comp.CatchPhrase.Noun
}

func (comp *Company) BsAdjectiveʹ() string {
	adjective := getData("Company", "BusinessSlogan", "Adjective")
	comp.BusinessSlogan.Adjective = sample(adjective)
	return comp.BusinessSlogan.Adjective
}

func (comp *Company) BsBuzzʹ() string {
	verb := getData("Company", "BusinessSlogan", "Buzz")
	comp.BusinessSlogan.Buzz = sample(verb)
	return comp.BusinessSlogan.Buzz
}

func (comp *Company) BsNounʹ() string {
	noun := getData("Company", "BusinessSlogan", "Noun")
	comp.BusinessSlogan.Noun = sample(noun)
	return comp.BusinessSlogan.Noun
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
	return ToJSON(company)
}
