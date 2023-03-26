package demo

import "fmt"

type Monster struct {
	Name string
}

func NewMonster(name string) Monster {
	return Monster{name}
}

type Player struct {
	Name string
}

func NewPlayer(name string) Player {
	return Player{name}
}

type Mission struct {
	Player  Player
	Monster Monster
}

func NewMission(p Player, m Monster) Mission {
	return Mission{p, m}
}

func (m Mission) Start() {
	fmt.Printf("%s defeats %s, world peace!\n", m.Player.Name, m.Monster.Name)
}

func InitMissionNative(name string) Mission {
	player := NewPlayer(name)
	monster := NewMonster(name)
	return NewMission(player, monster)
}
