package interfaces

type Foo interface {
	Foo() string
}

type MyFoo string

func (b *MyFoo) Foo() string {
	return string(*b)
}

func providerMyFoo() *MyFoo {
	b := new(MyFoo)
	*b = "hello world"
	return b
}

type Bar string

func providerBar(f Foo) string {
	return f.Foo()
}
