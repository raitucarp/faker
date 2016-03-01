package name

type Definitions struct {
	FirstName []string
	LastName  []string
	Prefix    []string
	Suffix    []string
	Title     title
}

func Export() Definitions {
	def := Definitions{
		FirstName: FirstName,
		LastName:  LastName,
		Prefix:    Prefix,
		Suffix:    Suffix,
		Title:     Title,
	}

	return def
}
