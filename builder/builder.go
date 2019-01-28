package builder

import "fmt"

type Account struct {
	accountNumber  string
	interestRate   float64
	openingBranch  string
	openingBalance float64
	ownerName      string
}

func (a *Account) String() string {
	return fmt.Sprintf("Account Owner -> %s\n Opening balance -> %.2f", a.ownerName, a.openingBalance)
}

type Builder interface {
	build() Account
	atRate(rate float64) Builder
	withName(name string) Builder
	atBranch(branch string) Builder
	withOpeningBalance(balance float64) Builder
}

type AccountBuilder struct {
	accountNumber  string
	interestRate   float64
	openingBranch  string
	openingBalance float64
	ownerName      string
}

func (a *AccountBuilder) build() *Account {
	return &Account{
		accountNumber:  a.accountNumber,
		interestRate:   a.interestRate,
		openingBranch:  a.openingBranch,
		openingBalance: a.openingBalance,
		ownerName:      a.ownerName,
	}
}

func (a *AccountBuilder) withOwner(ownerName string) *AccountBuilder {
	a.ownerName = ownerName
	return a
}

func (a *AccountBuilder) atRate(rate float64) *AccountBuilder {
	a.interestRate = rate
	return a
}

func (a *AccountBuilder) atBranch(branchName string) *AccountBuilder {
	a.openingBranch = branchName
	return a
}

func (a *AccountBuilder) withOpeningBalance(balance float64) *AccountBuilder {
	a.openingBalance = balance
	return a
}

func NewAccountBuilder(accountNumber string) *AccountBuilder {
	return &AccountBuilder{
		accountNumber: accountNumber,
	}
}

func Run() {
	newAccount := NewAccountBuilder("12345ABCDEF0123").
		withOwner("Alamin Mahamud").
		atBranch("Manchester").
		withOpeningBalance(1000000).
		atRate(2.5).
		build()
	fmt.Println(newAccount)
}
