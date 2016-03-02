package phone

type Definitions struct {
	PhoneFormat []string
}

func Export() Definitions {
	def := Definitions{
		PhoneFormat: PhoneFormat,
	}
	return def
}
