package faker_test

import (
	"bytes"
	"github.com/olekukonko/tablewriter"
	"github.com/raitucarp/faker"
	"strings"
	"testing"
	//"time"
	"strconv"
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
	tbl.AddData("ZipCode", f.Address.ZipCodeʹ())
	tbl.AddData("CityPrefix", f.Address.CityPrefixʹ())
	tbl.AddData("City", f.Address.Cityʹ())
	tbl.AddData("CitySuffix", f.Address.CitySuffixʹ())
	tbl.AddData("StreetName", f.Address.StreetNameʹ())
	tbl.AddData("StreetAddress", f.Address.StreetAddressʹ())
	tbl.AddData("StreetSuffix", f.Address.StreetSuffixʹ())
	tbl.AddData("StreetPrefix", f.Address.StreetPrefixʹ())
	tbl.AddData("SecondaryAddress", f.Address.SecondaryAddressʹ())
	tbl.AddData("County", f.Address.Countyʹ())
	tbl.AddData("Country", f.Address.Countryʹ())
	tbl.AddData("CountryCode", f.Address.CountryCodeʹ())
	tbl.AddData("State", f.Address.Stateʹ())
	tbl.AddData("StateAbbr", f.Address.StateAbbrʹ())
	tbl.AddData("Latitude", f.Address.Latitudeʹ())
	tbl.AddData("Longitude", f.Address.Longitudeʹ())

	/*JSON, err := f.Address.ToJSON()
	checkError(t, err)
	tbl.AddData("JSON", JSON)

	XML, err := f.Address.ToXML()
	checkError(t, err)
	tbl.AddData("XML", XML)*/

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

	/*JSON, err := f.Commerce.ToJSON()
	checkError(t, err)
	tbl.AddData("JSON", JSON)

	XML, err := f.Commerce.ToXML()
	checkError(t, err)
	tbl.AddData("XML", XML)*/

	tbl.Render()

	// logging
	t.Log("Fake Commerce\n", tbl.String())
}

func TestCompany(t *testing.T) {
	f := faker.New()
	tbl := NewTable()
	tbl.SetHeader("Field Name", "Data")
	tbl.AddData("Suffixes", strings.Join(f.Company.Suffixes_(), ", "))
	tbl.AddData("CompanyName", f.Company.CompanyName_())
	tbl.AddData("CompanySuffix", f.Company.CompanySuffix_())
	tbl.AddData("CatchPhrase", f.Company.CatchPhrase_())
	tbl.AddData("CatchPhraseAdjective", f.Company.CatchPhraseAdjective_())
	tbl.AddData("CatchPhraseDescriptor", f.Company.CatchPhraseAdjective_())
	tbl.AddData("CatchPhraseNoun", f.Company.CatchPhraseNoun_())
	tbl.AddData("Bs", f.Company.BS_())
	tbl.AddData("BsAdjective", f.Company.BsAdjective_())
	tbl.AddData("BsBuzz", f.Company.BsBuzz_())
	tbl.AddData("BsNoun", f.Company.BsNoun_())
	tbl.Render()

	// logging
	t.Log("Fake Company\n", tbl.String())

}

func TestDate(t *testing.T) {
	d := faker.Date{}
	tbl := NewTable()

	tbl.SetHeader("Field Name", "Data")
	tbl.AddData("Past", d.Past_().String())
	tbl.AddData("Future", d.Future_().String())
	tbl.AddData("Between", d.Between_("02/21/2016", "02/22/2016", "01/2/20006").String())
	tbl.AddData("Recent", d.Recent_(1).String())
	tbl.AddData("Month", d.Month_())
	tbl.AddData("Weekday", d.Weekday_())

	tbl.Render()
	// logging
	t.Log("Fake Date\n", tbl.String())
}

func TestFinance(t *testing.T) {
	f := faker.Finance{}

	tbl := NewTable()

	tbl.SetHeader("Field Name", "Data")
	tbl.AddData("Account", strconv.Itoa(f.Account_()))
	tbl.AddData("Account Name", f.AccountName_())
	tbl.AddData("Amount", f.Amount_())
	tbl.AddData("Mask", f.Mask_())
	tbl.AddData("TransactionType", f.TransactionType_())
	tbl.AddData("CurrencySymbol", f.CurrencySymbol_())
	tbl.AddData("CurrencyCode", f.CurrencyCode_())
	tbl.AddData("CurrencyName", f.CurrencyName_())

	tbl.Render()
	// logging
	t.Log("Fake Finance\n", tbl.String())
}

func TestHacker(t *testing.T) {
	hacker := faker.Hacker{}

	tbl := NewTable()

	tbl.SetHeader("Field Name", "Data")
	tbl.AddData("Abbreviation", hacker.Abbreviation_())
	tbl.AddData("Adjective", hacker.Adjective_())
	tbl.AddData("Noun", hacker.Noun_())
	tbl.AddData("Verb", hacker.Verb_())
	tbl.AddData("IngVerb", hacker.IngVerb_())
	tbl.AddData("Phrase", hacker.Phrase_())

	tbl.Render()
	// logging
	t.Log("Fake Hacker\n", tbl.String())
}

func TestImage(t *testing.T) {
	image := faker.Image{}

	tbl := NewTable()

	tbl.SetHeader("Field Name", "Data")
	tbl.AddData("Image", image.Image_(1028, 728))
	//tbl.AddData("avatar", image.Image_())
	//tbl.AddData("imageUrl", image.imageURL("abstract"))
	tbl.AddData("abstract", image.Abstract_())
	tbl.AddData("animals", image.Animals_())
	tbl.AddData("business", image.Business_())
	tbl.AddData("cats", image.Cats_())
	tbl.AddData("city", image.City_())
	tbl.AddData("food", image.Food_())
	tbl.AddData("nightlife", image.Nightlife_())
	tbl.AddData("fashion", image.Fashion_())
	tbl.AddData("people", image.People_())
	tbl.AddData("nature", image.Nature_())
	tbl.AddData("sports", image.Sports_())
	tbl.AddData("technics", image.Technics_())
	tbl.AddData("transport", image.Transport_())

	tbl.Render()
	// logging
	t.Log("Fake Image\n", tbl.String())
}

func TestInternet(t *testing.T) {
	internet := faker.Internet{}

	tbl := NewTable()

	tbl.SetHeader("Field Name", "Data")
	tbl.AddData("Avatar", internet.Avatar_())
	tbl.AddData("Email", internet.Email_())
	tbl.AddData("Username", internet.Username_())
	tbl.AddData("Protocol", internet.Protocol_())
	tbl.AddData("URL", internet.URL_())
	tbl.AddData("DomainName", internet.DomainName_())
	tbl.AddData("DomainSuffix", internet.DomainSuffix_())
	tbl.AddData("DomainWord", internet.DomainWord_())
	tbl.AddData("DomainWord", internet.IP_())
	//tbl.AddData("UserAgent", internet.IP_())
	tbl.AddData("Color", internet.Color_())
	tbl.AddData("Mac", internet.Mac_())
	//tbl.AddData("Password", internet.IP_())
	tbl.Render()
	// logging
	t.Log("Fake Internet\n", tbl.String())
}

func TestLorem(t *testing.T) {
	lorem := faker.Lorem{}

	tbl := NewTable()

	tbl.SetHeader("Field Name", "Data")
	tbl.AddData("Word", lorem.Word_())
	tbl.AddData("Words", strings.Join(lorem.Words_(), ", "))
	tbl.AddData("Sentence", lorem.Sentence_())
	tbl.AddData("Sentences", lorem.Sentences_())
	tbl.AddData("Paragraph", lorem.Paragraph_())
	tbl.AddData("Paragraphs", lorem.Paragraphs_())
	tbl.AddData("Text", lorem.Text())
	tbl.Render()

	// logging
	t.Log("Fake Lorem\n", tbl.String())
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

	/*JSON, err := f.Name.ToJSON()
	checkError(t, err)
	tbl.AddData("JSON", JSON)

	XML, err := f.Name.ToXML()
	checkError(t, err)
	tbl.AddData("XML", XML)*/

	tbl.Render()

	// logging
	t.Log("Fake Name\n", tbl.String())
}

func TestRandom(t *testing.T) {

	tbl := NewTable()
	tbl.SetHeader("Field Name", "Data")
	tbl.AddData("UUID", faker.UUID())
	tbl.AddData("Boolean", strconv.FormatBool(faker.Boolean()))

	tbl.Render()

	// logging
	t.Log("Fake Random\n", tbl.String())
}

func TestPhone(t *testing.T) {
	phone := faker.Phone{}
	tbl := NewTable()
	tbl.SetHeader("Field Name", "Data")
	tbl.AddData("Phone number", phone.Number_())
	tbl.AddData("Phone format", phone.Format_())

	tbl.Render()

	// logging
	t.Log("Fake Phone\n", tbl.String())
}

func TestCard(t *testing.T) {

	// logging
	t.Log("Fake Card\n")

	card := faker.CreateCard()
	userCard := faker.CreateUserCard()
	contextualCard := faker.CreateContextualCard()
	transactionCard := faker.CreateTransaction()
	t.Log(card.ToJSON())
	t.Log(userCard.ToJSON())
	t.Log(contextualCard.ToJSON())
	t.Log(transactionCard.ToJSON())

}

func checkError(t *testing.T, err error) {
	t.Error(err)
}
