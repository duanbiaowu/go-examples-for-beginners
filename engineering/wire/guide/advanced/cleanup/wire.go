//go:build wireinject
// +build wireinject

package cleanup

import (
	"github.com/google/wire"
	"os"
)

var Set = wire.NewSet(ProvideFoo, ProvideBar, wire.Struct(new(FooBar), "*"))

func InjectFooBar() (FooBar, func(), error) {
	wire.Build(Set)
	return FooBar{}, nil, nil
}

func InjectFile(path string) (*os.File, func(), error) {
	panic(wire.Build(provideFile))
}
