package company

type definitions struct {
	Suffix         []string
	Name           []string
	CatchPhrase    catchPhrase
	BusinessSlogan businessSlogan
}

func Export() interface{} {
	def := definitions{
		Suffix:         Suffix,
		Name:           Name,
		CatchPhrase:    CatchPhrase,
		BusinessSlogan: BusinessSlogan,
	}
	return def
}
