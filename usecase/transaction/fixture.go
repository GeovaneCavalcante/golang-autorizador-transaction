package transaction

import (
	"github.com/authorizer/entity"
	"time"
)

func NewFixtureTransaction() entity.Transaction {
	t, _ := time.Parse(time.RFC3339, "2019-02-13T11:01:31.000Z")
	return entity.Transaction{
		ID:         entity.NewID(),
		Amount:     100,
		Time:       t,
		Violations: make([]string, 0, 0),
	}
}
