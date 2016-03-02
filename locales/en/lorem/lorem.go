package lorem

type Definitions struct {
	Supplemental []string
	Words        []string
}

func Export() Definitions {
	def := Definitions{
		Supplemental: Supplemental,
		Words:        Words,
	}
	return def
}
