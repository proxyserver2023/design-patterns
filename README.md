# Design Patterns in Golang

## Creational Patterns
### Abstract Factory
* Intent
  - Provide an interface for creating families of related or dependent objects without specifying their concrete classes.
  - A hierarchy that encapsulates: many possible “platforms”, and the construction of a suite of “products”.
  - The new operator considered harmful.
* Discussion
  - The “factory” object has the responsibility for providing creation services for the entire platform family. Clients never create platform objects directly, they ask the factory to do that for them.

``` go
type IBook interface {
	getAuthor() string
	getTitle() string
}

type IBookFactory interface {
	makeGoBook() IBook
	makeKubeBook() IBook
}

func main() {
	
}

func testConcreteFactory(bookFactory IBookFactory){
	goBook := bookFactory.makeGoBook()
	fmt.Println(goBook)
	
	goBook2 := bookFactory.makeGoBook()
	fmt.Println(goBook2)
	
	kubeBook := bookFactory.makeKubeBook()
	fmt.Println(kubeBook)
}
```

