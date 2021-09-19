package entity

import (
	"time"
)

// Transaction data
type Transaction struct {
	ID         ID
	Merchant   string
	Amount     int
	Violations []string
	Time       time.Time
}

//NewTransaction create a new transaction
func NewTransaction(merchant string, amount int, timeData string, violations []string) (*Transaction, error) {
	tm, err := time.Parse(time.RFC3339, timeData)
	if err != nil {
		return nil, err
	}
	t := &Transaction{
		ID:         NewID(),
		Merchant:   merchant,
		Amount:     amount,
		Time:       tm,
		Violations: violations,
	}
	return t, nil
}
