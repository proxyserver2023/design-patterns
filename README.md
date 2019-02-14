# Design Patterns in Golang

One of the most resourceful implementations of design patterns in golang

## TOC

- [SOLID](#solid)
  - [Single Responsibility Principle](#single-responsibility-principle)
  - [Open Closed Principle](#open-closed-principle)
  - [Liskov Substitution Principle](#liskov-substitution-principle)
  - [Interface Segregation Principle](#interface-segregation-principle)
  - [Dependency Inversion Principle](#dependency-inversion-principle)
- [Abstract Factory](#abstract-factory)
- [Builder](#builder)
- [Circuit Breaker](#circuit-breaker)
- [Prototype](#prototype)
- [Publish Subscribe](#publish-subscribe)
- [Service Registry](#service-registry)
- [Singleton](#singleton)
- [Strategy](#strategy)

## SOLID

![SOLID Principles](https://www.dckap.com/blog/wp-content/uploads/2018/06/SOLID-PNG.png)

### Single Responsibility Principle

* Generic Interface

```go

```


### Open Closed Principle
### Liskov Substitution Principle
### Interface Segregation Principle
### Dependency Inversion Principle

## Abstract Factory

[Placeholder ...]

## Builder

[Placeholder ...]

## Circuit Breaker

[Placeholder ...]

## Prototype

![Prototype Pattern](http://blog.ralch.com/media/golang/design-patterns/prototype.gif)

The Prototype Pattern creates duplicate objects while keeping performance in mind.

1. It requires implementing a prototype interface which tells to create a clone of the current object.
2. It is used when creation of object directly is costly.

For instance, an object is to be created after a costly database operation. We can cache the object, returns its clone on next request and update the database as and when needed thus reducing the database calls.

**Example - 1**: generate different configuration files depending on our needs

``` go
package configurer

type Config struct {
	workDir string
	user string
}


func NewConfig(user string, workDir string) Config {
	return Config{
		user: user,
		workDir: workDir,
	}
}

func (c Config) WithUser(user string) Config {
	c.user = user
	return c
}

func (c Config) WithWorkDir(workDir string) Config {
	c.workDir = workDir
	return c
}

```

We want to be able to mutate the object without affecting its initial instance. The goal is to be able to generate different configuration files without loosing the flexibility of customizing them without mutation of the initial default configuration.

## Publish Subscribe

[Placeholder ...]

## Service Registry

[Placeholder ...]

## Singleton

1. only one instance
2. global point to access the instance
3. initialization on first use

If app needs one and only one instance of an object.


``` go
type privateStructure struct {
	value string
}

var singleVariable privateStructure

func GetSingletonInstance() privateStructure {
	if singleVariable != nil {
		return singleVariable
	}

	singleVariable = privateStructure{
		value: "some data",
	}

	return singleVariable
}

```

A thread-safe solution might be

``` go
var mu sync.Mutex

func GetInstance() *singleton {
	mu.Lock()
	defer mu.Unlock()
	
	if instance == nil {
		instance = &singleton{}
	}
	
	return instance
	
}
```

`Check-Lock-Check` Pattern

``` go
func GetInstance() *singleton {
	
	if instance == nil {
		mu.Lok()
		defer mu.Unlock()
		
		if instance == nil {
			instance = &singleton{}
		}
	}
	
	return instance
}
```

But using the sync/atomic package, we can atomically load and set a flag that will indicate if we have initialized or not our instance.

``` go
import sync
import sync/atomic

var initialized uint32

func Getinstance() *singleton{
	if atomic.LoadUInt32(&initialized) == 1 {
		return instance
	}
	
	mu.Lock()
	defer mu.Unlock()
	
	if initialized == 0 {
		instance = &singleton{}
		atomic.StoreUint32(&initialized, 1)
	}
}
```

Idiomatic singleton approach in go

``` go
package singleton

import (
	"sync"
)


type singleton struct {}

var instance *singleton
var once sync.Once

func GetInstance() *singleton {
	once.Do(func(){
		instance = &singleton
	})
	return instance
}
```

Example Code
``` go
package main

import (
	"fmt"

	"github.com/alamin-mahamud/go-design-patterns/singleton"
)

func main() {
	s := singleton.GetInstance()
	s.Data = 1
	fmt.Println("1st -> ", s.Data)

	s2 := singleton.GetInstance()
	fmt.Println("2nd -> ", s2.Data)

	s3 := singleton.GetInstance()
	fmt.Println("3rd -> ", s3.Data)

	s2.Data = 20
	fmt.Println("1st -> ", s.Data)
	fmt.Println("2nd -> ", s2.Data)
	fmt.Println("3rd -> ", s3.Data)

	s3.Data = 10
	fmt.Println("1st -> ", s.Data)
	fmt.Println("2nd -> ", s2.Data)
	fmt.Println("3rd -> ", s3.Data)

}
///////////////////////////////////////////
1st ->  1
2nd ->  1
3rd ->  1
1st ->  20
2nd ->  20
3rd ->  20
1st ->  10
2nd ->  10
3rd ->  10
```

## Strategy

[Placeholder ...]
