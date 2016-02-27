package faker_test

import (
	"bytes"
	"github.com/olekukonko/tablewriter"
	"github.com/raitucarp/faker"
	"testing"
)

type Table struct {
	data   [][]string
	header []string
	buf    *bytes.Buffer
	table  *tablewriter.Table
}

func (t *Table) SetHeader(header ...string) {
	t.table.SetHeader(header)
}

func (t *Table) AddData(data ...string) {
	t.table.Append(data)
}

func (t *Table) Render() {
	t.table.Render()
}

func (t *Table) String() string {
	return t.buf.String()
}

func NewTable() Table {
	t := Table{}
	t.buf = new(bytes.Buffer)
	t.table = tablewriter.NewWriter(t.buf)

	return t
}

// Testing begin here
func TestAddress(t *testing.T) {
	f := faker.New()
	tbl := NewTable()
	tbl.SetHeader("Field Name", "Data")
	tbl.AddData("ZipCode", f.Address.ZipCode_())
	tbl.AddData("CityPrefix", f.Address.CityPrefix_())
	tbl.AddData("City", f.Address.City_())
	tbl.AddData("CitySuffix", f.Address.CitySuffix_())
	tbl.AddData("StreetName", f.Address.StreetName_())
	tbl.AddData("StreetAddress", f.Address.StreetAddress_())
	tbl.AddData("StreetSuffix", f.Address.StreetSuffix_())
	tbl.AddData("StreetPrefix", f.Address.StreetPrefix_())
	tbl.AddData("SecondaryAddress", f.Address.SecondaryAddress_())
	tbl.AddData("County", f.Address.County_())
	tbl.AddData("Country", f.Address.Country_())
	tbl.AddData("CountryCode", f.Address.CountryCode_())
	tbl.AddData("State", f.Address.State_())
	tbl.AddData("StateAbbr", f.Address.StateAbbr_())
	tbl.AddData("Latitude", f.Address.Latitude_())
	tbl.AddData("Longitude", f.Address.Longitude_())

	JSON, err := f.Address.ToJSON()
	checkError(t, err)
	tbl.AddData("JSON", JSON)

	XML, err := f.Address.ToXML()
	checkError(t, err)
	tbl.AddData("XML", XML)

	tbl.Render()

	// logging
	t.Log("Fake Address\n", tbl.String())
}

func TestCommerce(t *testing.T) {
	f := faker.New()
	tbl := NewTable()
	tbl.SetHeader("Field Name", "Data")
	tbl.AddData("Color", f.Commerce.Color_())
	tbl.AddData("Department", f.Commerce.Department_())
	tbl.AddData("ProductName", f.Commerce.ProductName_())
	tbl.AddData("Price", f.Commerce.Price_(20000, 3, "$"))
	tbl.AddData("ProductAdjective", f.Commerce.ProductAdjective_())
	tbl.AddData("ProductMaterial", f.Commerce.ProductMaterial_())
	tbl.AddData("Product", f.Commerce.Product_())

	JSON, err := f.Commerce.ToJSON()
	checkError(t, err)
	tbl.AddData("JSON", JSON)

	XML, err := f.Commerce.ToXML()
	checkError(t, err)
	tbl.AddData("XML", XML)

	tbl.Render()

	// logging
	t.Log("Fake Commerce\n", tbl.String())
}

func TestName(t *testing.T) {
	f := faker.New()

	tbl := NewTable()
	tbl.SetHeader("Field Name", "Data")
	tbl.AddData("FirstName", f.Name.FirstName_())
	tbl.AddData("LastName", f.Name.LastName_())
	tbl.AddData("FindName", f.Name.FindName_())
	tbl.AddData("JobTitle", f.Name.JobTitle_())
	tbl.AddData("Prefix", f.Name.Prefix_())
	tbl.AddData("Suffix", f.Name.Suffix_())
	tbl.AddData("Title", f.Name.Title_())
	tbl.AddData("JobDescriptor", f.Name.JobDescriptor_())
	tbl.AddData("JobArea", f.Name.JobArea_())
	tbl.AddData("JobType", f.Name.JobType_())

	JSON, err := f.Name.ToJSON()
	checkError(t, err)
	tbl.AddData("JSON", JSON)

	XML, err := f.Name.ToXML()
	checkError(t, err)
	tbl.AddData("XML", XML)

	tbl.Render()

	// logging
	t.Log("Fake Name\n", tbl.String())
}

func checkError(t *testing.T, err error) {
	t.Error(err)
}
