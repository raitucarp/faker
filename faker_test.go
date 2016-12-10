package faker_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/olekukonko/tablewriter"
	"github.com/raitucarp/faker"
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

func TestStaticSeed(t *testing.T) {
	f := faker.New()

	faker.StaticSeed(1234)
	name := f.Name.FirstNameʹ()
	name2 := f.Name.FirstNameʹ()

	if name != name2 {
		t.Fatal("name is not equals to name2, static seed failed",
			"name = ", name,
			"and name2 = ", name2)
	}

	faker.ResetStaticSeed()
	currentSeed := faker.GetStaticSeed()
	if currentSeed != 0 {
		t.Fatal("seed not reset")
	}

	name3 := f.Name.FirstNameʹ()
	name4 := f.Name.FirstNameʹ()
	if name3 == name4 {
		t.Fatal("seed is resetted but still same")
	}
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
	tbl.AddData("Color", f.Commerce.Colorʹ())
	tbl.AddData("Department", f.Commerce.Departmentʹ())
	tbl.AddData("ProductName", f.Commerce.ProductNameʹ())
	tbl.AddData("Price", f.Commerce.Priceʹ(20000, 3, "$"))
	tbl.AddData("ProductAdjective", f.Commerce.ProductAdjectiveʹ())
	tbl.AddData("ProductMaterial", f.Commerce.ProductMaterialʹ())
	tbl.AddData("Product", f.Commerce.Productʹ())

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
	tbl.AddData("Suffixes", strings.Join(f.Company.Suffixesʹ(), ", "))
	tbl.AddData("CompanyName", f.Company.CompanyNameʹ())
	tbl.AddData("CompanySuffix", f.Company.CompanySuffixʹ())
	tbl.AddData("CatchPhrase", f.Company.CatchPhraseʹ())
	tbl.AddData("CatchPhraseAdjective", f.Company.CatchPhraseAdjectiveʹ())
	tbl.AddData("CatchPhraseDescriptor", f.Company.CatchPhraseAdjectiveʹ())
	tbl.AddData("CatchPhraseNoun", f.Company.CatchPhraseNounʹ())
	tbl.AddData("Bs", f.Company.BSʹ())
	tbl.AddData("BsAdjective", f.Company.BsAdjectiveʹ())
	tbl.AddData("BsBuzz", f.Company.BsBuzzʹ())
	tbl.AddData("BsNoun", f.Company.BsNounʹ())
	tbl.Render()

	// logging
	t.Log("Fake Company\n", tbl.String())

}

func TestDate(t *testing.T) {
	d := faker.Date{}
	tbl := NewTable()

	tbl.SetHeader("Field Name", "Data")
	tbl.AddData("Past", d.Pastʹ().String())
	tbl.AddData("Future", d.Futureʹ().String())
	tbl.AddData("Between", d.Betweenʹ("02/21/2016", "02/22/2016", "01/2/20006").String())
	tbl.AddData("Recent", d.Recentʹ(1).String())
	tbl.AddData("Month", d.Monthʹ())
	tbl.AddData("Weekday", d.Weekdayʹ())

	tbl.Render()
	// logging
	t.Log("Fake Date\n", tbl.String())
}

func TestFinance(t *testing.T) {
	f := faker.Finance{}

	tbl := NewTable()

	tbl.SetHeader("Field Name", "Data")
	tbl.AddData("Account", strconv.Itoa(f.Accountʹ()))
	tbl.AddData("Account Name", f.AccountNameʹ())
	tbl.AddData("Amount", f.Amountʹ())
	tbl.AddData("Mask", f.Maskʹ())
	tbl.AddData("TransactionType", f.TransactionTypeʹ())
	tbl.AddData("CurrencySymbol", f.CurrencySymbolʹ())
	tbl.AddData("CurrencyCode", f.CurrencyCodeʹ())
	tbl.AddData("CurrencyName", f.CurrencyNameʹ())

	tbl.Render()
	// logging
	t.Log("Fake Finance\n", tbl.String())
}

func TestHacker(t *testing.T) {
	hacker := faker.Hacker{}

	tbl := NewTable()

	tbl.SetHeader("Field Name", "Data")
	tbl.AddData("Abbreviation", hacker.Abbreviationʹ())
	tbl.AddData("Adjective", hacker.Adjectiveʹ())
	tbl.AddData("Noun", hacker.Nounʹ())
	tbl.AddData("Verb", hacker.Verbʹ())
	tbl.AddData("IngVerb", hacker.IngVerbʹ())
	tbl.AddData("Phrase", hacker.Phraseʹ())

	tbl.Render()
	// logging
	t.Log("Fake Hacker\n", tbl.String())
}

func TestImage(t *testing.T) {
	image := faker.Image{}

	tbl := NewTable()

	tbl.SetHeader("Field Name", "Data")
	tbl.AddData("Image", image.Imageʹ(1028, 728))
	//tbl.AddData("avatar", image.Image_())
	//tbl.AddData("imageUrl", image.imageURL("abstract"))
	tbl.AddData("abstract", image.Abstractʹ())
	tbl.AddData("animals", image.Animalsʹ())
	tbl.AddData("business", image.Businessʹ())
	tbl.AddData("cats", image.Catsʹ())
	tbl.AddData("city", image.Cityʹ())
	tbl.AddData("food", image.Foodʹ())
	tbl.AddData("nightlife", image.Nightlifeʹ())
	tbl.AddData("fashion", image.Fashionʹ())
	tbl.AddData("people", image.Peopleʹ())
	tbl.AddData("nature", image.Natureʹ())
	tbl.AddData("sports", image.Sportsʹ())
	tbl.AddData("technics", image.Technicsʹ())
	tbl.AddData("transport", image.Transportʹ())

	tbl.Render()
	// logging
	t.Log("Fake Image\n", tbl.String())
}

func TestInternet(t *testing.T) {
	internet := faker.Internet{}

	tbl := NewTable()

	tbl.SetHeader("Field Name", "Data")
	tbl.AddData("Avatar", internet.Avatarʹ())
	tbl.AddData("Email", internet.Emailʹ())
	tbl.AddData("Username", internet.Usernameʹ())
	tbl.AddData("Protocol", internet.Protocolʹ())
	tbl.AddData("URL", internet.URLʹ())
	tbl.AddData("DomainName", internet.DomainNameʹ())
	tbl.AddData("DomainSuffix", internet.DomainSuffixʹ())
	tbl.AddData("DomainWord", internet.DomainWordʹ())
	tbl.AddData("DomainWord", internet.IPʹ())
	//tbl.AddData("UserAgent", internet.IP_())
	tbl.AddData("Color", internet.Colorʹ())
	tbl.AddData("Mac", internet.Macʹ())
	//tbl.AddData("Password", internet.IP_())
	tbl.Render()
	// logging
	t.Log("Fake Internet\n", tbl.String())
}

func TestLorem(t *testing.T) {
	lorem := faker.Lorem{}

	tbl := NewTable()

	tbl.SetHeader("Field Name", "Data")
	tbl.AddData("Word", lorem.Wordʹ())
	tbl.AddData("Words", strings.Join(lorem.Wordsʹ(), ", "))
	tbl.AddData("Sentence", lorem.Sentenceʹ())
	tbl.AddData("Sentences", lorem.Sentencesʹ())
	tbl.AddData("Paragraph", lorem.Paragraphʹ())
	tbl.AddData("Paragraphs", lorem.Paragraphsʹ())
	tbl.AddData("Text", lorem.Text())
	tbl.Render()

	// logging
	t.Log("Fake Lorem\n", tbl.String())
}

func TestName(t *testing.T) {
	f := faker.New()

	tbl := NewTable()
	tbl.SetHeader("Field Name", "Data")
	tbl.AddData("FirstName", f.Name.FirstNameʹ())
	tbl.AddData("LastName", f.Name.LastNameʹ())
	tbl.AddData("FindName", f.Name.FindNameʹ())
	tbl.AddData("JobTitle", f.Name.JobTitleʹ())
	tbl.AddData("Prefix", f.Name.Prefixʹ())
	tbl.AddData("Suffix", f.Name.Suffixʹ())
	tbl.AddData("Title", f.Name.Titleʹ())
	tbl.AddData("JobDescriptor", f.Name.JobDescriptorʹ())
	tbl.AddData("JobArea", f.Name.JobAreaʹ())
	tbl.AddData("JobType", f.Name.JobTypeʹ())

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
	tbl.AddData("Phone number", phone.Numberʹ())
	tbl.AddData("Phone format", phone.Formatʹ())

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
