package company

type catchPhrase struct {
	Template   []string
	Adjective  []string
	Descriptor []string
	Noun       []string
}

var catchPhraseTemplate = []string{
	"#{.Company.CatchPhrase.Adjective} #{.Company.CatchPhrase.Descriptor} #{.Company.CatchPhrase.Noun}",
}
var CatchPhrase = catchPhrase{
	Template:   catchPhraseTemplate,
	Adjective:  Adjective,
	Descriptor: Descriptor,
	Noun:       Noun,
}
