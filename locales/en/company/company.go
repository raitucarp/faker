package company

type Definitions struct {
	Suffix         []string
	Name           []string
	CatchPhrase    catchPhrase
	BusinessSlogan businessSlogan
}

func Export() Definitions {
	def := Definitions{
		Suffix:         Suffix,
		Name:           Name,
		CatchPhrase:    CatchPhrase,
		BusinessSlogan: BusinessSlogan,
	}
	return def
}
