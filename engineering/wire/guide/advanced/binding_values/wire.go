//go:build wireinject
// +build wireinject

package binding_values

import (
	"github.com/google/wire"
	"io"
	"os"
)

func InjectFoo() Foo {
	wire.Build(wire.Value(Foo{X: 1024}))
	return Foo{}
}

func InjectReader() io.Reader {
	wire.Build(wire.InterfaceValue(new(io.Reader), os.Stdin))
	return nil
}
