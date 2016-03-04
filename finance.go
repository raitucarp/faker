package faker

import (
	"math"
	"reflect"
	"strconv"
	"strings"
)

type FinanceAccount struct {
	Number int
	Name   string
}

func (fa *FinanceAccount) Numberʹ(args ...interface{}) int {
	types := kindOf(args)

	// default length
	length := 8

	// set default length
	if len(args) == 1 {
		if types[0] == reflect.Int {
			length = args[0].(int)
		}
	}

	// check length
	if length <= 8 {
		length = 8
	}

	template := strings.Repeat("#", length)
	n := replaceSymbolWithNumber(template, '#')
	num, _ := strconv.Atoi(n)

	fa.Number = num
	return fa.Number
}

func (fa *FinanceAccount) Nameʹ() string {
	name := getData("Finance", "AccountType")
	fa.Name = strings.Join([]string{sample(name), "Account"}, " ")
	return fa.Name
}

// Currency
type FinanceCurrency struct {
	Code   string
	Name   string
	Symbol string
}

func (fc *FinanceCurrency) pick() {
	randomCurrency := call("Finance", "RandomCurrency")
	fc.Code = randomCurrency["Code"]
	fc.Name = randomCurrency["Name"]
	fc.Symbol = randomCurrency["Symbol"]
}

func (fc *FinanceCurrency) Codeʹ() string {
	fc.pick()
	return fc.Code
}

func (fc *FinanceCurrency) Nameʹ() string {
	fc.pick()
	return fc.Name
}

func (fc *FinanceCurrency) Symbolʹ() string {
	for {
		if fc.Symbol == "" {
			fc.pick()
			continue
		} else {
			return fc.Symbol
		}
	}
}

/* Finance */
type Finance struct {
	Account         FinanceAccount
	Amount          string
	TransactionType string
	Currency        FinanceCurrency
}

func (f *Finance) Accountʹ(args ...interface{}) int {
	f.Account.Numberʹ(args...)
	return f.Account.Number
}

func (f *Finance) AccountNameʹ() string {
	f.Account.Nameʹ()
	return f.Account.Name
}

func (f *Finance) Maskʹ(params ...interface{}) string {
	length := 4
	parenthesis := true
	ellipsis := true

	types := kindOf(params...)
	if len(types) >= 1 && types[0] == reflect.Int {
		length = params[0].(int)
	}

	if len(types) >= 2 && types[1] == reflect.Bool {
		parenthesis = params[1].(bool)
	}

	if len(types) >= 3 && types[2] == reflect.Bool {
		ellipsis = params[2].(bool)
	}

	var template string
	template = strings.Repeat("#", length)

	if ellipsis {
		template = strings.Join([]string{"...", template}, "")
	}

	if parenthesis {
		template = strings.Join([]string{"(", template, ")"}, "")
	}

	return replaceSymbolWithNumber(template, '#')
}

func (f *Finance) Amountʹ(params ...interface{}) string {
	min := 0
	max := 1000
	dec := float64(2)
	symbol := ""

	types := kindOf(params)

	if len(types) >= 1 && types[0] == reflect.Int {
		min = params[0].(int)
	}

	if len(types) >= 2 && types[1] == reflect.Int {
		max = params[1].(int)
	}

	if len(types) >= 3 && types[2] == reflect.Int {
		dec = params[2].(float64)
	}

	if len(types) >= 4 && types[3] == reflect.Int {
		symbol = params[3].(string)
	}

	rnd := float64(random(min, max))
	pow := math.Pow(10, dec)
	val := float64(round(rnd*pow)) / pow

	result := strconv.FormatFloat(val, 'f', int(dec), 32)

	f.Amount = symbol + result
	return f.Amount
}

func (f *Finance) TransactionTypeʹ() string {
	transactionTypes := getData("Finance", "TransactionType")
	f.TransactionType = sample(transactionTypes)
	return f.TransactionType
}

func (f *Finance) CurrencyCodeʹ() string {
	f.Currency.Codeʹ()
	return f.Currency.Code
}

func (f *Finance) CurrencyNameʹ() string {
	f.Currency.Nameʹ()
	return f.Currency.Name
}

func (f *Finance) CurrencySymbolʹ() string {
	f.Currency.Symbolʹ()
	return f.Currency.Symbol
}

// Fake Generate random data for all field
func (f *Finance) Fake() {
	f.Accountʹ()
	f.AccountNameʹ()
	f.Maskʹ()
	f.Amountʹ()
	f.TransactionTypeʹ()
	f.CurrencyCodeʹ()
	f.CurrencyNameʹ()
	f.CurrencySymbolʹ()
}

// ToJSON Encode name and its fields to JSON.
func (f *Finance) ToJSON() (string, error) {
	return ToJSON(f)
}

// ToXML Encode name and its fields to XML.
func (f *Finance) ToXML() (string, error) {
	return ToXML(f)
}
