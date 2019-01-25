package strategy

// Operation kinds of the parent and gives away the abstraction
// and flexibility to strategy
type Operation struct {
	Operator Operator
}

// Operate - implements the strategy pattern
func (o *Operation) Operate(leftValue, rightValue int) int {
	return o.Operator.Apply(leftValue, rightValue)
}
