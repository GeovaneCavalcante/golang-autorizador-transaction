package presenter

import "github.com/authorizer/entity"

//Account data
type Account struct {
	ActiveCard     bool `json:"active-card"`
	AvailableLimit int  `json:"available-limit"`
}

// AccountWrapper data
type AccountWrapper struct {
	Account    interface{} `json:"account"`
	Violations []string    `json:"violations"`
}

// NewAccountWrapper account wrapper creator
func NewAccountWrapper(c *entity.Account, violations []string) *AccountWrapper {
	if c == nil {
		return &AccountWrapper{
			Account:    make(map[string]interface{}),
			Violations: violations,
		}
	}
	return &AccountWrapper{
		Account: Account{
			ActiveCard:     c.ActiveCard,
			AvailableLimit: c.AvailableLimit,
		},
		Violations: violations,
	}
}
