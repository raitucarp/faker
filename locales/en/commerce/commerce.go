package commerce

type Definitions struct {
	Color       []string
	Department  []string
	ProductName productName
}

func Export() Definitions {
	def := Definitions{
		Color:       Color,
		Department:  Department,
		ProductName: ProductName,
	}
	return def
}
