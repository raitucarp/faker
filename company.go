package faker

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

func (comp *Company) Suffixes_() []string {
	suffixes := getData("Company", "Suffix")
	return suffixes
}

func (comp *Company) CompanyName_() string {
	formats := getData("Company", "Name")
	format := sample(formats)

	name, err := Fake(format)
	if err != nil {
		panic(err)
	}
	comp.Name = name
	return comp.Name
}

func (comp *Company) CompanySuffix_() string {
	suffixes := comp.Suffixes_()
	comp.Suffix = sample(suffixes)
	return comp.Suffix
}

func (comp *Company) CatchPhrase_() string {
	template := getData("Company", "CatchPhrase", "Template")
	catchPhrase := sample(template)
	result, err := Fake(catchPhrase)
	if err != nil {
		panic(err)
	}

	return result
}

func (comp *Company) BS_() string {
	template := getData("Company", "BusinessSlogan", "Template")
	bs := sample(template)
	result, err := Fake(bs)
	if err != nil {
		panic(err)
	}

	return result
}

func (comp *Company) CatchPhraseAdjective_() string {
	adjective := getData("Company", "CatchPhrase", "Adjective")
	comp.CatchPhrase.Adjective = sample(adjective)
	return comp.CatchPhrase.Adjective
}

func (comp *Company) CatchPhraseDescriptor_() string {
	descriptor := getData("Company", "CatchPhrase", "Descriptor")
	comp.CatchPhrase.Descriptor = sample(descriptor)
	return comp.CatchPhrase.Descriptor
}

func (comp *Company) CatchPhraseNoun_() string {
	noun := getData("Company", "CatchPhrase", "Noun")
	comp.CatchPhrase.Noun = sample(noun)
	return comp.CatchPhrase.Noun
}

func (comp *Company) BsAdjective_() string {
	adjective := getData("Company", "BusinessSlogan", "Adjective")
	comp.BusinessSlogan.Adjective = sample(adjective)
	return comp.BusinessSlogan.Adjective
}

func (comp *Company) BsBuzz_() string {
	verb := getData("Company", "BusinessSlogan", "Buzz")
	comp.BusinessSlogan.Buzz = sample(verb)
	return comp.BusinessSlogan.Buzz
}

func (comp *Company) BsNoun_() string {
	noun := getData("Company", "BusinessSlogan", "Noun")
	comp.BusinessSlogan.Noun = sample(noun)
	return comp.BusinessSlogan.Noun
}
