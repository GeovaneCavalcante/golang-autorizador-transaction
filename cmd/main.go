package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/authorizer/cmd/handler"
	"github.com/authorizer/infrastructure/repository"
	"github.com/authorizer/usecase/account"
	"github.com/authorizer/usecase/transaction"
	"log"
	"os"
)

func main() {
	accountRepo := repository.NewAccountInme()
	accountService := account.NewService(accountRepo)

	transactionRepo := repository.NewTransactionInme()
	transactionService := transaction.NewService(transactionRepo, accountService)

	scanner := bufio.NewScanner(os.Stdin)
	var out interface{}
	for scanner.Scan() {
		var input map[string]map[string]interface{}
		err := json.Unmarshal([]byte(scanner.Text()), &input)
		if err != nil {
			log.Println(err.Error())
		}
		if val, ok := input["transaction"]; ok {
			out, err = handler.MakeTransactionHandlers(transactionService, accountService, val)
		}
		if val, ok := input["account"]; ok {
			out, err = handler.MakeAccountHandlers(accountService, val)
		}
		if err != nil {
			log.Println(err)
		}
		b, err := json.Marshal(out)
		if err != nil {
			log.Println(err)
		}
		fmt.Println(string(b))
	}
	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}
