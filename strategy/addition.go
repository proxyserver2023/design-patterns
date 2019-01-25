package strategy

// Addition is the struct for +
type Addition struct{}

// Apply - implements the + operator
func (Addition) Apply(lval, rval int) int {
	return lval + rval
}
