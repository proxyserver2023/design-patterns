## Abstract Document Pattern

- flexibility of untyped languages.
- add type safety
- on the fly new property add
- flexible domain organization, using tree like structure
- loosely coupled system.

## UML class Diagram

![abstract-document-pattern-class-diagram](https://upload.wikimedia.org/wikipedia/commons/0/0e/Abstract-document-pattern.svg)

## Pseudocode

```
interface Document
        put(key : String, value : Object) : Object
        get(key : String) : Object
        children(key : String, constructor : Map<String, Object> -> T) : T[]

abstract class BaseDocument : Document
        properties : Map<string, object>

        constructor(properties : Map<string, object>)
                this->properties := properties

        implement put(key : String, value : Object) : Object
                return this->properties->put(key, value)
       
        implement get(key : String) : Object
                return this->properties->get(key)

        implement children(key: string, constructor: Map<string, object> -> T) : T[]
                var result := new T[]
                var children := this->properties->get(key) castTo Map<String, Object>[]
                
                foreach (child in children)
                        result[] := constructor->apply(child)
                
                return result
```

`document.go`

```go
type Document interface {
        Put(key string, value interface{}) interface{}
        Get(key string) interface{}
        Children(key string, func(map[string]interface{}, T)) []T
}
```

`base_document.go`

```go
type BaseDocument struct {
        entries map[string] interface{}
}

func (b *BaseDocument) Put(key string, value interface{}) interface{} {
        b.entries[key] = value
}

func (b *BaseDocument) Get(key string) interface{} {
        if value, ok := b.entries[key]; ok {
                return value
        }
        return nil
}

func (b *BaseDocument) Children(key string, g func(m map[string]interface{}, b interface{})) []interface{} {
        // Unclear about the implementation
}
```

`usage.go`

```go
type Car struct {}

var source map[string]interface{} = ...
var car Car = Car{}
var model string = car.getModel()
var price int = car.getPrice()

// Not Clear about the implementations
var wheels []Wheel = car.Children("wheel", Wheel{}).collect(Collections.toList())
```

## Resources

- [Wiki](https://en.wikipedia.org/wiki/Abstract_Document_Pattern)
- [Martin Fowler](https://martinfowler.com/apsupp/properties.pdf)