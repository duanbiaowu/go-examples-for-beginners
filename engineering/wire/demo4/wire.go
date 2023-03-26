//go:build wireinject
// +build wireinject

package demo4

import "github.com/google/wire"

var monsterPlayerSet = wire.NewSet(NewMonster, NewPlayer)

//var endingASet = wire.NewSet(monsterPlayerSet, wire.Struct(new(EndingA), "Player", "Monster"))
//var endingBSet = wire.NewSet(monsterPlayerSet, wire.Struct(new(EndingB), "Player", "Monster"))

var endingASet = wire.NewSet(monsterPlayerSet, wire.Struct(new(EndingA), "*"))
var endingBSet = wire.NewSet(monsterPlayerSet, wire.Struct(new(EndingB), "*"))

func InitEndingA(p PlayerParam, m MonsterParam) EndingA {
	wire.Build(endingASet)
	return EndingA{}
}

func InitEndingB(p PlayerParam, m MonsterParam) EndingB {
	wire.Build(endingBSet)
	return EndingB{}
}
