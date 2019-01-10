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
