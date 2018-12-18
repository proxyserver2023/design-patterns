package strategy

// Multiplcation -> *
type Multiplication struct{}

// Apply - implementation; so that it can be used for strategy pattern
func (Multiplication) Apply(lval, rval int) int { return lval * rval }
