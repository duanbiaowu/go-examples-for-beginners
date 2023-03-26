//go:build wireinject
// +build wireinject

package demo5

import "github.com/google/wire"

// 有些类型是单例，例如配置，数据库对象（sql.DB）
// 可以使用 wire.Value 绑定值，使用wire.InterfaceValue绑定接口
// 怪兽一直是一个Kitty
var kitty = Monster{"Kitty"}

var monsterPlayerSet = wire.NewSet(wire.Value(kitty), NewPlayer)

var endingASet = wire.NewSet(monsterPlayerSet, wire.Struct(new(EndingA), "*"))
var endingBSet = wire.NewSet(monsterPlayerSet, wire.Struct(new(EndingB), "*"))

func InitEndingA(p PlayerParam) EndingA {
	wire.Build(endingASet)
	return EndingA{}
}

func InitEndingB(p PlayerParam) EndingB {
	wire.Build(endingBSet)
	return EndingB{}
}
