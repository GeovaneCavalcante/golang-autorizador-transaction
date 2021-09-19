package entity

import "errors"

//ErrUpdateAccount invalid update
var ErrUpdateAccount = errors.New("Error updating account")

//ErrGetAccount invalid get
var ErrGetAccount = errors.New("Error getting account")

//ErrCreateAccount invalid create
var ErrCreateAccount = errors.New("Error create account")

//ErrListTransactions invalid list
var ErrListTransactions = errors.New("Error getting transactions")

//ErrCreateTransaction invalid create
var ErrCreateTransaction = errors.New("Error create transaction")

//ErrValidateTransaction invalid transaction
var ErrValidateTransaction = errors.New("Error create transaction")
