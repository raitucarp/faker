package commerce

type definitions struct {
	Color       []string
	Department  []string
	ProductName productName
}

func Export() interface{} {
	def := definitions{
		Color:       Color,
		Department:  Department,
		ProductName: ProductName,
	}
	return def
}
