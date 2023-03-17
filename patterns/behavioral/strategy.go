package behavioral

type Operator interface {
	Apply(int, int) int
}

type Operation struct {
	Operator Operator
}

func (o *Operation) Operate(x, y int) int {
	return o.Operator.Apply(x, y)
}

type Addition struct {
}

type Multiplication struct {
}

func (Addition) Apply(x, y int) int {
	return x + y
}

func (Multiplication) Apply(x, y int) int {
	return x * y
}
