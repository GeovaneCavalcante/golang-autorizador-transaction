package account

import "github.com/authorizer/entity"

func NewFixtureAccount() entity.Account {
	return entity.Account{
		ID:             entity.NewID(),
		ActiveCard:     true,
		AvailableLimit: 1000,
		Violations:     make([]string, 0, 0),
	}
}
