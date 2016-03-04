package faker

type Hacker struct {
	Abbreviation string
	Adjective    string
	Noun         string
	Verb         string
	IngVerb      string
	Phrase       string
}

func (h *Hacker) Abbreviation_() string {
	abbr := getData("Hacker", "Abbreviation")
	h.Abbreviation = sample(abbr)
	return h.Abbreviation
}

func (h *Hacker) Adjective_() string {
	adjective := getData("Hacker", "Adjective")
	h.Adjective = sample(adjective)
	return h.Adjective
}

func (h *Hacker) Noun_() string {
	noun := getData("Hacker", "Noun")
	h.Noun = sample(noun)
	return h.Noun
}

func (h *Hacker) Verb_() string {
	verb := getData("Hacker", "Verb")
	h.Verb = sample(verb)
	return h.Verb
}

func (h *Hacker) IngVerb_() string {
	ingVerb := getData("Hacker", "IngVerb")
	h.IngVerb = sample(ingVerb)
	return h.IngVerb
}

func (h *Hacker) Phrase_() string {
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
	h.Abbreviation_()
	h.Adjective_()
	h.Noun_()
	h.Verb_()
	h.IngVerb_()
	h.Phrase_()
}

// ToJSON Encode name and its fields to JSON.
func (h *Hacker) ToJSON() (string, error) {
	return ToJSON(h)
}

// ToJSON Encode name and its fields to JSON.
func (h *Hacker) ToXML() (string, error) {
	return ToJSON(h)
}
