package strategy

type Operator interface {
	Apply(int, int) int
}
