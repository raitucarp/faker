package faker

import (
	"math"
	"math/rand"
	"reflect"
	"strconv"
	"strings"
)

type ProductName struct {
	Product   string `json:"product"`
	Adjective string `json:"adjective"`
	Material  string `json:"material"`
}

type Commerce struct {
	Color       string      `json:"color"`
	Department  string      `json:"color"`
	Price       string      `json:"color"`
	ProductName ProductName `json:"product_name"`
}

func (commerce *Commerce) Colorʹ() string {
	// get color
	color := getData("Commerce", "Color")

	// assign to commerce
	commerce.Color = sample(color)
	return commerce.Color
}

func (commerce *Commerce) Departmentʹ() string {
	department := getData("Commerce", "Department")
	commerce.Department = sample(department)
	return commerce.Department
}

func (commerce *Commerce) Priceʹ(params ...interface{}) string {
	// max int, dec float64, symbol string
	min := 0
	max := 1000
	dec := float64(2)
	var symbol string

	kinds := kindOf(params...)
	if len(kinds) >= 1 && kinds[0] == reflect.Int {
		min = params[0].(int)
	}

	if len(kinds) >= 2 && kinds[1] == reflect.Int {
		max = params[1].(int)
	}

	if len(kinds) >= 3 && kinds[2] == reflect.Float64 {
		dec = params[2].(float64)
	}

	if len(kinds) >= 4 && kinds[3] == reflect.String {
		symbol = params[3].(string)
	}

	if min < 0 || max < 0 {
		return symbol + "0.00"
	}

	rnd := float64(rand.Intn(max))
	pow := math.Pow(10, dec)
	val := float64(round(rnd*pow)) / pow

	result := strconv.FormatFloat(val, 'f', int(dec), 32)
	return symbol + result
}

func (commerce *Commerce) ProductNameʹ() string {
	separator := " "
	productName := []string{
		commerce.ProductAdjectiveʹ(),
		commerce.ProductMaterialʹ(),
		commerce.Productʹ(),
	}
	return strings.Join(productName, separator)
}

func (commerce *Commerce) ProductAdjectiveʹ() string {
	adjective := getData("Commerce", "ProductName", "Adjective")
	commerce.ProductName.Adjective = sample(adjective)
	return commerce.ProductName.Adjective
}

func (commerce *Commerce) ProductMaterialʹ() string {
	material := getData("Commerce", "ProductName", "Material")
	commerce.ProductName.Material = sample(material)
	return commerce.ProductName.Material
}

func (commerce *Commerce) Productʹ() string {
	product := getData("Commerce", "ProductName", "Product")
	commerce.ProductName.Product = sample(product)
	return commerce.ProductName.Product

}

// Fake Generate random data for all field
func (commerce *Commerce) Fake() {
	commerce.Colorʹ()
	commerce.Departmentʹ()
	commerce.Priceʹ()
	commerce.ProductNameʹ()
	commerce.ProductAdjectiveʹ()
	commerce.ProductMaterialʹ()
	commerce.Productʹ()
}

// ToJSON Encode name and its fields to JSON.
func (commerce *Commerce) ToJSON() (string, error) {
	return ToJSON(commerce)
}

// ToXML Encode name and its fields to XML.
func (commerce *Commerce) ToXML() (string, error) {
	return ToJSON(commerce)
}