package faker

import (
	"math"
	"reflect"
	"strconv"
	"strings"
)

type FinanceAccount struct {
	Number string
	Name   string
}

func (fa *FinanceAccount) Number_(args ...interface{}) string {
	types := typeof(args)

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
	n := replaceSymbolsWithNumber(template, '#')

	fa.Number = n
	return fa.Number
}

func (fa *FinanceAccount) Name_() string {
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

func (fc *FinanceCurrency) Code_() string {
	fc.pick()
	return fc.Code
}

func (fc *FinanceCurrency) Name_() string {
	fc.pick()
	return fc.Name
}

func (fc *FinanceCurrency) Symbol_() string {
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

func (f *Finance) Account_(args ...interface{}) string {
	f.Account.Number_(args)
	return f.Account.Number
}

func (f *Finance) AccountName_() string {
	f.Account.Name_()
	return f.Account.Name
}

func (f *Finance) Mask_(params ...interface{}) string {
	length := 4
	parenthesis := true
	ellipsis := true

	types := typeof(params...)
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

	return replaceSymbolsWithNumber(template, '#')
}

func (f *Finance) Amount_(params ...interface{}) string {
	min := 0
	max := 1000
	dec := float64(2)
	symbol := ""

	types := typeof(params)

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

func (f *Finance) TransactionType_() string {
	transactionTypes := getData("Finance", "TransactionType")
	f.TransactionType = sample(transactionTypes)
	return f.TransactionType
}

func (f *Finance) CurrencyCode_() string {
	f.Currency.Code_()
	return f.Currency.Code
}

func (f *Finance) CurrencyName_() string {
	f.Currency.Name_()
	return f.Currency.Name
}

func (f *Finance) CurrencySymbol_() string {
	f.Currency.Symbol_()
	return f.Currency.Symbol
}
