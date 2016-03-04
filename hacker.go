package faker

type Hacker struct {
	Abbreviation string
	Adjective    string
	Noun         string
	Verb         string
	IngVerb      string
	Phrase       string
}

func (h *Hacker) Abbreviationʹ() string {
	abbr := getData("Hacker", "Abbreviation")
	h.Abbreviation = sample(abbr)
	return h.Abbreviation
}

func (h *Hacker) Adjectiveʹ() string {
	adjective := getData("Hacker", "Adjective")
	h.Adjective = sample(adjective)
	return h.Adjective
}

func (h *Hacker) Nounʹ() string {
	noun := getData("Hacker", "Noun")
	h.Noun = sample(noun)
	return h.Noun
}

func (h *Hacker) Verbʹ() string {
	verb := getData("Hacker", "Verb")
	h.Verb = sample(verb)
	return h.Verb
}

func (h *Hacker) IngVerbʹ() string {
	ingVerb := getData("Hacker", "IngVerb")
	h.IngVerb = sample(ingVerb)
	return h.IngVerb
}

func (h *Hacker) Phraseʹ() string {
	phrase := getData("Hacker", "Phrase")
	template := sample(phrase)

	result, err := Parse(template)
	if err != nil {
		result = ""
	}

	h.Phrase = result
	return h.Phrase
}

// Fake Generate random data for all field
func (h *Hacker) Fake() {
	h.Abbreviationʹ()
	h.Adjectiveʹ()
	h.Nounʹ()
	h.Verbʹ()
	h.IngVerbʹ()
	h.Phraseʹ()
}

// ToJSON Encode name and its fields to JSON.
func (h *Hacker) ToJSON() (string, error) {
	return ToJSON(h)
}

// ToXML Encode name and its fields to XML.
func (h *Hacker) ToXML() (string, error) {
	return ToXML(h)
}
