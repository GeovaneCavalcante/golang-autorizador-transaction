package entity

//Account data
type Account struct {
	ID             ID
	ActiveCard     bool
	AvailableLimit int
	Violations     []string
}

//NewAccount create a new account
func NewAccount(activeCard bool, availableLimit int, violations []string) *Account {
	c := &Account{
		ID:             NewID(),
		ActiveCard:     activeCard,
		AvailableLimit: availableLimit,
		Violations:     violations,
	}
	return c
}
