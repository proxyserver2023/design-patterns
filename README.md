# Design Patterns in Golang

## Creational Patterns
### Abstract Factory
* Intent
  - Provide an interface for creating families of related or dependent objects without specifying their concrete classes.
  - A hierarchy that encapsulates: many possible “platforms”, and the construction of a suite of “products”.
  - The new operator considered harmful.
* Discussion
  - The “factory” object has the responsibility for providing creation services for the entire platform family. Clients never create platform objects directly, they ask the factory to do that for them.

[Golang Abstract Factory Example](https://play.golang.org/p/85I07-GqVKx)

``` go
package main

import "fmt"

type IBook interface {
	getAuthor() string
	getTitle() string
}

type GoBook struct{ title, author string }

func (g *GoBook) getAuthor() string { return fmt.Sprintf("%s", g.author) }
func (g *GoBook) getTitle() string  { return fmt.Sprintf("%s", g.title) }
func (g *GoBook) String() string    { return fmt.Sprintf("%v - %v", g.getTitle(), g.getAuthor()) }

type KubeBook struct{ title, author string }

func (k *KubeBook) getAuthor() string { return fmt.Sprintf("%s", k.author) }
func (k *KubeBook) getTitle() string  { return fmt.Sprintf("%s", k.title) }
func (k *KubeBook) String() string    { return fmt.Sprintf("%v - %v", k.getTitle(), k.getAuthor()) }

type IBookFactory interface {
	makeGoBook() IBook
	makeKubeBook() IBook
}

type OReillyBookFactory struct{}

func (o OReillyBookFactory) makeGoBook() IBook {
	return &GoBook{title: "O'Reilly Title of Golang Book", author: "O'Reilly J.K. Rowling"}
}
func (o OReillyBookFactory) makeKubeBook() IBook {
	return &KubeBook{title: "O'Reilly Title of Kube Book", author: "Ryan Reynolds O'REilly"}
}

type SamsBookFactory struct{}

func (s SamsBookFactory) makeGoBook() IBook {
	return &GoBook{title: "Sams Title of Golang Book", author: "Sams J.K. Rowling"}
}
func (s SamsBookFactory) makeKubeBook() IBook {
	return &KubeBook{title: "Sams Title of Kube Book", author: "Ryan Sams"}
}

func main() {
	fmt.Println("Begin Testing Abstract Factory Pattern")
	fmt.Println("")

	fmt.Println("Testing O'Reilly Book Factory")
	bookFactory := OReillyBookFactory{}
	testConcreteFactory(bookFactory)
	fmt.Println("End Testing O'Reilly Book Factory")
	fmt.Println("")

	fmt.Println("Testing Sams Book Factory")
	bookFactory2 := SamsBookFactory{}
	testConcreteFactory(bookFactory2)
	fmt.Println("End Testing Sams Book Factory")
	fmt.Println("")

	fmt.Println("End Testing Abstract Factory Pattern")
}

func testConcreteFactory(bookFactory IBookFactory) {
	goBook := bookFactory.makeGoBook()
	fmt.Println(goBook)

	goBook2 := bookFactory.makeGoBook()
	fmt.Println(goBook2)

	kubeBook := bookFactory.makeKubeBook()
	fmt.Println(kubeBook)
}


/*
Outputs
--------
Begin Testing Abstract Factory Pattern

Testing O'Reilly Book Factory
O'Reilly Title of Golang Book - O'Reilly J.K. Rowling
O'Reilly Title of Golang Book - O'Reilly J.K. Rowling
O'Reilly Title of Kube Book - Ryan Reynolds O'REilly
End Testing O'Reilly Book Factory

Testing Sams Book Factory
Sams Title of Golang Book - Sams J.K. Rowling
Sams Title of Golang Book - Sams J.K. Rowling
Sams Title of Kube Book - Ryan Sams
End Testing Sams Book Factory

End Testing Abstract Factory Pattern

*/
```

## Builder Design Pattern
Builder pattern separates the construction of a complex object from its representation so that the same construction process can create different representations.

* Example
  - This pattern is used by fast food restaurants to construct children’s meals.
  - Children’s meals typically consist of a main item, a side item, a drink, and a toy (e.g., a hamburger, fries, Coke, and toy dinosaur).
  - Note that there can be variation in the content of the children’s meal, but the construction process is the same.
  - Whether a customer orders a hamburger, cheeseburger, or chicken, the process is the same.
  - The employee at the counter directs the crew to assemble a main item, side item, and toy.
  - These items are then placed in a bag. The drink is placed in a cup and remains outside of the bag

* Checklist
  - common input and many possible representations
  - common input parsing in a reader class
  - design a standard protocol for creating all possible output representations.
  - capture all steps in a builder interface
  - builder derived class for each target representation.
  - client creates a Reader Object and a Builder object and registers the latter with the former.
  - client asks the reader to construct
  - client asks the builder to return the result.
  [Contd.]

[Example Code](https://play.golang.org/p/l6iliNvQGgc)
``` go
package main

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

func main() {
	newAccount := NewAccountBuilder("12345ABCDEF0123").
		withOwner("Alamin Mahamud").
		atBranch("Manchester").
		withOpeningBalance(1000000).
		atRate(2.5).
		build()
	fmt.Println(newAccount)
}

/* Outputs
------------
Account Owner -> Alamin Mahamud
Opening balance -> 1000000.00
*/

```
