package demo4

import (
	"fmt"
)

type Player struct {
	Name string
}

type Monster struct {
	Name string
}

type Mission struct {
	Player  Player
	Monster Monster
}

type PlayerParam string
type MonsterParam string

func NewPlayer(name PlayerParam) Player {
	return Player{Name: string(name)}
}

func NewMonster(name MonsterParam) Monster {
	return Monster{Name: string(name)}
}

func NewMission(p Player, m Monster) Mission {
	return Mission{p, m}
}

func (m Mission) Start() {
	fmt.Printf("%s defeats %s, world peace!\n", m.Player.Name, m.Monster.Name)
}

type EndingA struct {
	Player  Player
	Monster Monster
}

func NewEndingA(p Player, m Monster) EndingA {
	return EndingA{p, m}
}

func (p EndingA) Appear() {
	fmt.Printf("%s defeats %s, world peace!\n", p.Player.Name, p.Monster.Name)
}

type EndingB struct {
	Player  Player
	Monster Monster
}

func NewEndingB(p Player, m Monster) EndingB {
	return EndingB{p, m}
}

func (p EndingB) Appear() {
	fmt.Printf("%s defeats %s, but become monster, world darker!\n", p.Player.Name, p.Monster.Name)
}
