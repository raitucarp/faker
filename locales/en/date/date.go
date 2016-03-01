package date

type Definitions struct {
	Month   month
	Weekday weekday
}

func Export() Definitions {
	def := Definitions{
		Month:   Month,
		Weekday: Weekday,
	}
	return def
}
