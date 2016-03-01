package hacker

type Definitions struct {
	Abbreviation, Adjective,
	Noun, Verb, IngVerb,
	Phrase []string
}

func Export() Definitions {
	def := Definitions{
		Abbreviation: Abbreviation,
		Adjective:    Adjective,
		Noun:         Noun,
		Verb:         Verb,
		IngVerb:      IngVerb,
		Phrase:       Phrase,
	}
	return def
}
