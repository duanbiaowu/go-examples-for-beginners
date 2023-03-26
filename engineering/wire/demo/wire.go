//go:build wireinject
// +build wireinject

package demo

import "github.com/google/wire"

// docs: https://zhuanlan.zhihu.com/p/110453784
func InitMission(name string) Mission {
	wire.Build(NewMonster, NewPlayer, NewMission)
	return Mission{}
}
