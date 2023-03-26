//go:build wireinject
// +build wireinject

package interfaces

import "github.com/google/wire"

var Set = wire.NewSet(providerMyFoo, wire.Bind(new(Foo), new(*MyFoo)), providerBar)

func InitializeFooBar() string {
	wire.Build(Set)
	return ""
}
