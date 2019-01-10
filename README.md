# Design Patterns in Golang

## Creational Patterns
### Abstract Factory
* Intent
  - Provide an interface for creating families of related or dependent objects without specifying their concrete classes.
  - A hierarchy that encapsulates: many possible “platforms”, and the construction of a suite of “products”.
  - The new operator considered harmful.
* Discussion
  - The “factory” object has the responsibility for providing creation services for the entire platform family. Clients never create platform objects directly, they ask the factory to do that for them.

[Golang Abstract Factory Example](https://play.golang.org/p/STN2-LezKix)

``` go
package main

import "fmt"

type IBook interface {
	getAuthor() string
	getTitle() string
}

type GoBook struct{title, author string}
func (g *GoBook) getAuthor() string { return fmt.Sprintf("%s", g.author) }
func (g *GoBook) getTitle() string { return fmt.Sprintf("%s", g.title) }
func (g *GoBook) String() string { return fmt.Sprintf("%v - %v", g.getTitle(), g.getAuthor())}

type KubeBook struct{title, author string}
func (k *KubeBook) getAuthor() string { return fmt.Sprintf("%s", k.author) }
func (k *KubeBook) getTitle() string { return fmt.Sprintf("%s", k.title) }
func (k *KubeBook) String() string { return fmt.Sprintf("%v - %v", k.getTitle(), k.getAuthor())}

type IBookFactory interface {
	makeGoBook() IBook
	makeKubeBook() IBook
}

type OReillyBookFactory struct{}

func (o OReillyBookFactory) makeGoBook() IBook {
	return &GoBook{title: "Title of Golang Book", author: "J.K. Rowling"}
}

func (o OReillyBookFactory) makeKubeBook() IBook {
	return &KubeBook{title: "Title of Kube Book", author: "Ryan Reynolds"}
}

func main() {
	fmt.Println("Begin Testing Abstract Factory Pattern")
	fmt.Println("")

	fmt.Println("Testing O'Reilly Book Factory")
	bookFactory := OReillyBookFactory{}
	testConcreteFactory(bookFactory)
	fmt.Println("End Testing O'Reilly Book Factory")

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
```

