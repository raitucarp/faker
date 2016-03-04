package faker

import (
	"strings"
)

//import "math/rand"

type Transaction struct {
	Amount   string
	Date     string
	Business string
	Name     string
	Type     string
	Account  int
}

type Post struct {
	Words     []string
	Sentence  string
	Sentences string
	Paragraph string
}

type ContextualCard struct {
	Name        string
	Username    string
	Avatar      string
	Email       string
	DateOfBirth string
	Phone       string
	Address     struct {
		Street  string
		Suite   string
		City    string
		ZipCode string
		Geo     Geo
	}
	Website string
	Company CompanyCard
}

type Card struct {
	Name     string
	Username string
	Email    string
	Address  struct {
		StreetA string
		StreetB string
		StreetC string
		StreetD string
		City    string
		State   string
		Country string
		ZipCode string
		Geo     Geo
	}
	Phone          string
	Website        string
	Company        CompanyCard
	Posts          []Post
	AccountHistory []Transaction
}

type CompanyCard struct {
	Name        string
	CatchPhrase string
	Bs          string
}

type UserCard struct {
	Name     string
	Username string
	Email    string
	Address  struct {
		Street  string
		Suite   string
		City    string
		ZipCode string
		Geo     Geo
	}
	Phone   string
	Website string
	Company CompanyCard
}

func (c *Card) ToJSON() (s string, err error) {
	return ToJSON(c)
}

func (c *ContextualCard) ToJSON() (s string, err error) {
	return ToJSON(c)
}

func (c *UserCard) ToJSON() (s string, err error) {
	return ToJSON(c)
}

func (c *Transaction) ToJSON() (s string, err error) {
	return ToJSON(c)
}

func CreateCard() (card Card) {
	f := New()
	f.Fake()

	card.Name = f.Name.FindName_()
	card.Username = f.Internet.Username_()
	card.Email = f.Internet.Email_()
	card.Address.StreetA = f.Address.StreetNameʹ()
	card.Address.StreetB = f.Address.StreetAddressʹ()
	card.Address.StreetC = f.Address.StreetAddressʹ(true)
	card.Address.StreetD = f.Address.SecondaryAddressʹ()
	card.Address.City = f.Address.Cityʹ()
	card.Address.State = f.Address.Stateʹ()
	card.Address.Country = f.Address.Countryʹ()
	card.Address.ZipCode = f.Address.ZipCodeʹ()
	card.Address.Geo = Geo{
		Longitude: f.Address.Longitudeʹ(),
		Latitude:  f.Address.Latitudeʹ(),
	}
	card.Phone = f.Phone.Number_()
	card.Website = f.Internet.URL_()
	card.Company.Name = f.Company.CompanyName_()
	card.Company.CatchPhrase = f.Company.CatchPhrase_()
	card.Company.Bs = f.Company.BS_()

	rnd := random(7, 15)
	rndTransaction := random(4, 10)

	var posts []Post
	for i := 0; i < rnd; i++ {
		post := CreatePost()
		posts = append(posts, post)
	}

	var transactions []Transaction
	for i := 0; i < rndTransaction; i++ {
		transaction := CreateTransaction()
		transactions = append(transactions, transaction)
	}

	card.Posts = posts
	card.AccountHistory = transactions
	return
}

func CreatePost() (post Post) {
	l := new(Lorem)
	l.Fake()

	post.Words = l.Words
	post.Sentence = l.Sentence
	post.Sentences = l.Sentences
	post.Paragraph = l.Paragraph

	return post
}

func CreateContextualCard() (card ContextualCard) {
	name := new(Name)
	internet := new(Internet)
	phone := new(Phone)
	company := new(Company)
	date := new(Date)
	address := new(Address)

	past := random(10, 40)

	card.Name = name.FirstName_()
	card.Username = internet.Username_()
	card.Avatar = internet.Avatar_()
	card.Email = internet.Email_()
	card.DateOfBirth = date.Past_(past, "02-01-2016").UTC().String()
	card.Phone = phone.Number_()
	card.Address.Street = address.StreetNameʹ()
	card.Address.Suite = address.SecondaryAddressʹ()
	card.Address.City = address.Cityʹ()
	card.Address.ZipCode = address.ZipCodeʹ()
	card.Address.Geo = Geo{
		Latitude:  address.Latitudeʹ(),
		Longitude: address.Longitudeʹ(),
	}
	card.Phone = phone.Number_()
	card.Website = internet.DomainName_()
	card.Company.Name = company.CompanyName_()
	card.Company.CatchPhrase = company.CatchPhrase_()
	card.Company.Bs = company.BS_()
	return
}

func CreateUserCard() (card UserCard) {
	name := new(Name)
	internet := new(Internet)
	address := new(Address)
	phone := new(Phone)
	company := new(Company)

	card.Name = name.FindName_()
	card.Username = internet.Username_()
	card.Email = internet.Email_()
	card.Address.Street = address.StreetNameʹ()
	card.Address.Suite = address.SecondaryAddressʹ()
	card.Address.City = address.Cityʹ()
	card.Address.ZipCode = address.ZipCodeʹ()
	card.Address.Geo = Geo{
		Latitude:  address.Latitudeʹ(),
		Longitude: address.Longitudeʹ(),
	}
	card.Phone = phone.Number_()
	card.Website = internet.DomainName_()
	card.Company.Name = company.CompanyName_()
	card.Company.CatchPhrase = company.CatchPhrase_()
	card.Company.Bs = company.BS_()

	return
}

func CreateTransaction() (transaction Transaction) {
	finance := new(Finance)
	finance.Fake()

	date := new(Date)

	company := new(Company)
	company.Fake()

	transaction.Amount = finance.Amount
	transaction.Date = date.Past_().String()
	transaction.Business = company.Name
	transaction.Type = finance.TransactionType
	transaction.Account = finance.Account.Number

	name := strings.Join([]string{finance.AccountName_(), finance.Mask_()}, " ")
	transaction.Name = name
	return
}
