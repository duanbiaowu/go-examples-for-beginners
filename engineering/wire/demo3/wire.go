//go:build wireinject
// +build wireinject

package demo3

import "github.com/google/wire"

// wire遵循fail-fast的原则，错误必须被处理
// 如果我们的注入器不返回错误，但构造器返回错误，wire工具会报错！
// InitMission Injector
// NewPlayer, NewMonster Provider
// 每个注入器实际上就是一个对象的创建和初始化函数
func InitMission(p PlayerParam, m MonsterParam) (Mission, error) {
	wire.Build(NewPlayer, NewMonster, NewMission)
	return Mission{}, nil
}
