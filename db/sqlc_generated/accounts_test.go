package db

import "testing"

// Unit test for func CreateAccount - First the func needs to connect to the database
//		CreateAccount - input: para{owner, balance, currency} + context
//						output: Account {...} + error
func TestCreateAccount(t *testing.T) {
	arg := Account{
		Owner:    "JackLong",
		Balance:  152,
		Currency: "USD",
	}

}
