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
	ingverb := getData("Hacker", "IngVerb")
	h.IngVerb = sample(ingverb)
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
