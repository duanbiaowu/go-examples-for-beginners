//go:build wireinject
// +build wireinject

package fields

import "github.com/google/wire"

func InjectMessage() string {
	wire.Build(provideFoo, Gets)
	return ""
}

// You can instead use wire.FieldsOf to use those fields directly without writing getS:
func InjectMessage2() string {
	wire.
		wire.Build(provideFoo, wire.FieldsOf(new(Foo), "S"))
	return ""
}
