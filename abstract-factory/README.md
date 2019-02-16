# Abstract Factory

- Intent
  - Provide an interface for creating families of related or dependent objects without specifying their concrete classes.
  - A hierarchy that encapsulates: many possible “platforms”, and the construction of a suite of “products”.
  - The new operator considered harmful.
- Discussion
  - The “factory” object has the responsibility for providing creation services for the entire platform family. Clients never create platform objects directly, they ask the factory to do that for them.

## Class Diagram

![Abstract Factory](https://sourcemaking.com/files/v2/content/patterns/Abstract_Factory.png)

## Example

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
