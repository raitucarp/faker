package finance

import "math/rand"

type Definitions struct {
	AccountType     []string
	Currency        []currency
	TransactionType []string
}

func (def Definitions) RandomCurrency() map[string]string {
	rnd := rand.Intn(len(Currency))
	c := Currency[rnd]
	result := map[string]string{
		"Symbol": c.Symbol,
		"Name":   c.Name,
		"Code":   c.Code,
	}
	return result
}

func Export() Definitions {
	def := Definitions{
		AccountType:     AccountType,
		Currency:        Currency,
		TransactionType: TransactionType,
	}
	return def
}
