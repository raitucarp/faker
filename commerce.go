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

func (c *Commerce) Colorʹ() string {
	// get color
	color := getData("Commerce", "Color")

	// assign to commerce
	c.Color = sample(color)
	return c.Color
}

func (c *Commerce) Departmentʹ() string {
	department := getData("Commerce", "Department")
	c.Department = sample(department)
	return c.Department
}

func (c *Commerce) Priceʹ(params ...interface{}) string {
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

func (c *Commerce) ProductNameʹ() string {
	separator := " "
	productName := []string{
		c.ProductAdjectiveʹ(),
		c.ProductMaterialʹ(),
		c.Productʹ(),
	}
	return strings.Join(productName, separator)
}

func (c *Commerce) ProductAdjectiveʹ() string {
	adjective := getData("Commerce", "ProductName", "Adjective")
	c.ProductName.Adjective = sample(adjective)
	return c.ProductName.Adjective
}

func (c *Commerce) ProductMaterialʹ() string {
	material := getData("Commerce", "ProductName", "Material")
	c.ProductName.Material = sample(material)
	return c.ProductName.Material
}

func (c *Commerce) Productʹ() string {
	product := getData("Commerce", "ProductName", "Product")
	c.ProductName.Product = sample(product)
	return c.ProductName.Product

}

// Fake Generate random data for all field
func (this *Commerce) Fake() {
	this.Colorʹ()
	this.Departmentʹ()
	this.Priceʹ()
	this.ProductNameʹ()
	this.ProductAdjectiveʹ()
	this.ProductMaterialʹ()
	this.Productʹ()
}

// ToJSON Encode name and its fields to JSON.
func (c *Commerce) ToJSON() (string, error) {
	return ToJSON(c)
}

// ToJSON Encode name and its fields to JSON.
func (c *Commerce) ToXML() (string, error) {
	return ToJSON(c)
}
