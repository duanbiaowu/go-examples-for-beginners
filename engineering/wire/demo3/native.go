package demo3

import (
	"errors"
	"fmt"
	"time"
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

func NewPlayer(name PlayerParam) (Player, error) {
	if time.Now().Unix()%2 == 0 {
		return Player{}, errors.New("player dead")
	}
	return Player{Name: string(name)}, nil
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

func InitMissionNative(name string) Mission {
	return Mission{Player{"dj"}, Monster{"kitty"}}
}
