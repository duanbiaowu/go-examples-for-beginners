//go:build wireinject
// +build wireinject

package base

import (
	"context"
	"github.com/google/wire"
)

// Any non-injector declarations found in a file with injectors will be copied into the generated file.
var SuperSet = wire.NewSet(ProvideFoo, ProvideBar, ProvideBaz)

// You can also add other provider sets into a provider set.
// var MegaSet = wire.NewSet(SuperSet, pkg.OtherSuperSet)

func initializeBaz(ctx context.Context) (Baz, error) {
	wire.Build(SuperSet)
	return Baz{}, nil
}
