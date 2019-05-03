# Design Patterns

One of the most resourceful implementations of design patterns. I tried to use more words to make it more verbose, so that anyone can pick the concepts. But primarily languages used are `Golang`, `Python`

## TOC

- [SOLID Principles](./solid)
- [UML](./uml)
- [Creational](#creational)
  - [Abstract Factory](#abstract-factory)
  - [Builder](#builder)
  - [Factory](#factory)
  - [Object Pool](#object-pool)
  - [Prototype](#prototype)
  - [Singleton](#singleton)
- [Structural](#structural)
  - [Adapter](#adapter)
  - [Bridge](#bridge)
  - [Composite](#composite)
  - [Decorator](#decorator)
  - [Facade](#facade)
  - [Flyweight](#flyweight)
  - [Private Class Data](#private-class-data)
  - [Proxy](#proxy)
- [Behavioural](#behavioural)
  - [Chain of responsibility](#chain-of-responsibility)
  - [Command](#command)
  - [Interpreter](#interpreter)
  - [Iterator](#iterator)
  - [Mediator](#mediator)
  - [Memento](#memento)
  - [Null Object](#null-object)
  - [Observer](#observer)
  - [State](#state)
  - [Strategy](#strategy)
  - [Template Method](#template-method)
  - [Visitor](#visitor)
- [Microservices](#microservices)
  - [Circuit Breaker](#circuit-breaker)
  - [Publish Subscribe](#publish-subscribe)
  - [Service Registry](#service-registry)
- [Others](#others)
  - [Abstract Document](#abstract-document)
  - [Acyclic Visitor](#acyclic-visitor)
- [AntiPatterns](#antipatterns)
- [Resources](#resources)
- [Credits](#credits)

## Creational Patterns

### Abstract Factory

Creates an instance of several families of classes.

### Builder

Separates object construction from its representation.

### Factory

- Create an instance of several derived classes.
- Define an interface for creating an object, but let subclasses decide which class to instantiate.
- Factory Method lets a class defer instantiation to subclasses.

#### Class Diagram

![Sourcemaking-Factory-Pattern-Image](https://sourcemaking.com/files/v2/content/patterns/Factory_Method.png)

![Sourcemaking-Factory-Pattern-Image-Implementation](https://sourcemaking.com/files/v2/content/patterns/Factory_Method_1.png)

#### Example

![factory-pattern](https://sourcemaking.com/files/v2/content/patterns/Factory_Method_example1.png)

#### Notes

- Make all implemented constructors private or protected.

### Object Pool

- Avoid expensive acquisition and release of resources by recycling objects.
- Significant Performance Boost
- Used where
  - cost of initializing a class instance is high,
  - the rate of instantiation of a class is high,
  - the number of instantiation in use at any one time is low.
- Object Caching
- A.K.A resource pool
- the pool will be a growing pool.
- we can restricts the number of objects created.
- It is desirable to keep all Reusable objects that are not currently in use in the same object pool so that they can be managed by one coherent policy. To achieve this, the Reusable Pool class is designed to be a singleton class.
- we don't want a process to have to wait for a particular object to be released, so the Object Pool also instantiates new objects as they are required, but must also implement a facility to clean up unused objects periodically.

#### UML Class Diagram

![connection-pool](https://sourcemaking.com/files/v2/content/patterns/Object_pool1.png)

![connection-pool](https://sourcemaking.com/files/v2/content/patterns/Object_pool_example1.png)

**Credits** - [SourceMaking](https://sourcemaking.com/design_patterns/object_pool)

#### Example

```go
type Reusable struct{}

type ReusablePool struct {
	Objects []Reusable
	MaxPoolSize int
}

func (r *ReusablePool) Acquire() Reusable{

	if r.Objects == nil {
		r.Objects = make([]Reusable, MaxPoolSize)
	}

	r.Objects = r.Objects[1:]
}

func (r *ReusablePool) Release(re *Reusable) {

}


func main() {
	reusablePool := &ReusablePool{MaxPoolSize: 10}
	reusable := reusablePool.Acquire()
	reusablePool.Release(reusable)
}
```

### Prototype

![Prototype Pattern](http://blog.ralch.com/media/golang/design-patterns/prototype.gif)

The Prototype Pattern creates duplicate objects while keeping performance in mind.

- It requires implementing a prototype interface which tells to create a clone of the current object.
- It is used when creation of object directly is costly.

For instance, an object is to be created after a costly database operation. We can cache the object, returns its clone on next request and update the database as and when needed thus reducing the database calls.

**Example - 1**: generate different configuration files depending on our needs

```go
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

### Singleton

- only one instance
- global point to access the instance
- initialization on first use

If app needs one and only one instance of an object.

```go
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

```go
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

```go
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

```go
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

```go
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

```go
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
// 1st ->  1
// 2nd ->  1
// 3rd ->  1
// 1st ->  20
// 2nd ->  20
// 3rd ->  20
// 1st ->  10
// 2nd ->  10
// 3rd ->  10
////////////////////////////////////////////
```

## Structural Patterns

Ease the design by identifying a simple way to realize relationships between entities.

### [Adapter](#adapter)

Match interfaces of different classes.

![type-c-adapter](./images/type-c-adapter.jpg)

We can use `type-c adapter` for using our traditional `usb-2`, `usb-3`, `micro-SD`, `ethernet` devices. They are not same. We used to have different ports for these devices.

But this adapter let's us work with common ground.

#### UML class diagram

![](https://cdncontribute.geeksforgeeks.org/wp-content/uploads/classDiagram.jpg)

#### Example

```go
type RowingBoat interface{
	row()
}


type FishingBoat struct {}
func (f *FishingBoat) sail() {}


type Captain struct {}
func (c *Captain) row(){}


type FishingBoatAdapter struct {
	fishingBoat FishingBoat
}

func (f *FishingBoatAdapter) row() {
	boat.sail()
}
```

```go
type TwoPinCharger interface {
	TwoPinCharge()
}

type TwoPinChargerMobile struct {}
func (t *TwoPinChargerMobile) TwoPinCharge() {}


type ThreePinCharger interface {
	ThreePinCharge()
}
func (t *ThreePinChargerMobile) ThreePinCharge() {}

type TwoPin2ThreePinChargerAdapter struct {
	twoPinCharger TwoPinCharger
}
func (t *TwoPin2ThreePinChargerAdapter) ThreePinCharge() {
	twoPinCharger.TwoPinCharge()
}

// TODO: fix implementation using Class Diagram
// To use the adapter
twoPinCharger := &TwoPinChargerMobile{}
adapter := &TwoPin2ThreePinChargeAdapter{twoPinCharger}
adapter.ThreePinCharge()

```

### [Bridge](#bridge)

Separate an object's interface from it's implementation

### [Composite](#composite)

A tree structure of simple and composite objects.

### [Decorator](#decorator)

Add responsibilities to objects dynamically.

### [Facade](#facade)

A single class that represents an entire subsystem.

### [Flyweight](#flyweight)

A fine-grained instance used for efficient sharing.

### [Private Class Data](#private-class-data)

Restricts accessor/mutator access.

### [Proxy](#proxy)

An object representing another object.

## Behavioural Patterns

### [Chain of responsibility](#chain-of-responsibility)

A way of passing a request to a chain of objects.

### [Command](#command)

Encapsulates a command request as an object

### [Interpreter](#interpreter)

A way to include language elements in a program.

### [Iterator](#iterator)

Sequentially access the elements of a collection.

### [Mediator](#mediator)

Defines simplified communication between classes.

### [Memento](#memento)

Capture and restore an object's internal state.

### [Null Object](#null-object)

Designed to act as a default value of an object.

### [Observer](#observer)

A way of notifying change to a number of classes.

### [State](#state)

Alter an object's behaviour when it's state change.

### [Strategy](#strategy)

Encapsulates an algorithm inside a class.

### [Template Method](#template-method)

Defer the exact steps of an algorithm to a subclass.

### [Visitor](#visitor)

Defines a new operation to a class without change.

## Microservices

### Circuit Breaker

### Publish Subscribe

[Placeholder ...]

### Service Registry

[Placeholder ...]

## Others

### [Abstract Document](#abstract-document)

Add Property on the fly.

### [Acyclic Visitor](#acyclic-visitor)

Allow new functions to be added to existing class hierarchies without affecting those hierarchies, and without creating the troublesome dependency cycles that are inherent to the `GOF VISITOR` Pattern.

#### UML Class Diagram

![acyclic-visitor](https://github.com/iluwatar/java-design-patterns/raw/master/acyclic-visitor/etc/acyclic-visitor.png)

#### Examples

```go
// TODO: implementaion
```

## AntiPatterns

Commonly occuring solutions to a problem, that generates decidedly negative consequences.

### When it happens

When Manager or Developer apply good design patterns into wrong context.

### How to mitigate

Antipatterns provide a detailed plan to reverse the situation.

### Types

- Software Development Antipatterns
- Software Architecture Antipatterns
- Software Product Management Antipatterns

## Resources

- [4 Page Cheatsheet](https://loufranco.com/wp-content/uploads/2012/11/cheatsheet.pdf)
- [Good and Short Explanation](http://edn.embarcadero.com/article/31863)

## Credits

- [Sourcemaking](https://sourcemaking.com/uml)
- [Randy Miller](http://edn.embarcadero.com/article/31863)
- [@iluwatar](https://github.com/iluwatar/java-design-patterns/blob/master/adapter/README.md)
