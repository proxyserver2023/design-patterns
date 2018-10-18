## Strategy Pattern


1. enables an algorithm's behavior to be selected at runtime.
2. defines algorithm, encapsulates them and uses them interchangably.


## Implementation


``` go
type Operator interface {
    Apply(int, int) int
}


type Operation struct {
    Operator Operator
}


func (o *Operation) Operate(leftValue, rightValue int) int {
    return o.Operator.Apply(leftValue, rightValue)
}
```


## Usage


### Addition Operator

``` go
type Addition struct {
    
}

func (Addition) Apply(lval, rval int) int {
    return lval + rval
}

add := Operation{Addition{}}
add.Operate(3,5) 
```

### Multiplication Operator

``` go
type Multiplication struct{}

func (Multiplication) Apply(lval, rval int) int {
	return lval * rval
}

multi := Operation{Multiplication{}}
multi.Operate(3,5) // 15
```


## Rules of Thumb
Strategy pattern is similar to Template pattern except in its granularity.
Strategy pattern lets you change the guts of an object. Decorator pattern lets you change the skin.
