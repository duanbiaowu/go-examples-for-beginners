//go:build wireinject
// +build wireinject

package demo2

import "github.com/google/wire"

func InitMission(p PlayerParam, m MonsterParam) Mission {
	wire.Build(NewPlayer, NewMonster, NewMission)
	return Mission{}
}
