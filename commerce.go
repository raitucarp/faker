package faker

import (
	"encoding/json"
	"encoding/xml"
	"math"
	"math/rand"
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

func (c *Commerce) Color_() string {
	// get color
	color := getData("Commerce", "Color")

	// assign to commerce
	c.Color = sample(color)
	return c.Color
}

func (c *Commerce) Department_() string {
	department := getData("Commerce", "Department")
	c.Department = sample(department)
	return c.Department
}

func (c *Commerce) Price_(max int, dec float64, symbol string) string {
	if dec < 2 {
		dec = 2
	}

	if max < 0 {
		return symbol + "0.00"
	}

	rnd := float64(rand.Intn(max))
	pow := math.Pow(10, dec)
	val := float64(round(rnd*pow)) / pow

	result := strconv.FormatFloat(val, 'f', int(dec), 32)
	return symbol + result
}

func (c *Commerce) ProductName_() string {
	separator := " "
	productName := []string{
		c.ProductAdjective_(),
		c.ProductMaterial_(),
		c.Product_(),
	}
	return strings.Join(productName, separator)
}

func (c *Commerce) ProductAdjective_() string {
	adjective := getData("Commerce", "ProductName", "Adjective")
	c.ProductName.Adjective = sample(adjective)
	return c.ProductName.Adjective
}

func (c *Commerce) ProductMaterial_() string {
	material := getData("Commerce", "ProductName", "Material")
	c.ProductName.Material = sample(material)
	return c.ProductName.Material
}

func (c *Commerce) Product_() string {
	product := getData("Commerce", "ProductName", "Product")
	c.ProductName.Product = sample(product)
	return c.ProductName.Product

}

func (c *Commerce) ToJSON() (s string, err error) {
	result, err := json.Marshal(c)

	if err != nil {
		return
	}
	return string(result), err
}

func (c *Commerce) ToXML() (s string, err error) {
	result, err := xml.Marshal(c)

	if err != nil {
		return
	}
	return string(result), err
}
