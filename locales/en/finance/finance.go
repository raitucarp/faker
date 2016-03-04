package finance

import "math/rand"

type Definitions struct {
	AccountType     []string
	Currency        []currency
	TransactionType []string
}

func (def Definitions) RandomCurrency() map[string]string {
	currencies := []currency{}
	for _, v := range Currency {
		if v.Symbol != "" {
			currencies = append(currencies, v)
		}
	}
	rnd := rand.Intn(len(currencies))
	c := currencies[rnd]
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
