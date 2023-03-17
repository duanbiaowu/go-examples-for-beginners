package creational

type Builder interface {
	WithColor(Color) Builder
	WithWheels(Wheels) Builder
	TopSpeed(Speed) Builder
	Build() Car
}

type Car interface {
	Drive() error
	Stop() error
	Color() Color
	Wheels() Wheels
	Speed() Speed
}

type Assembly struct {
	color  Color
	wheels Wheels
	speed  Speed
}

type BMW struct {
	color  Color
	wheels Wheels
	speed  Speed
}

type Speed float64

type Color string

type Wheels string

const (
	MPH Speed = 1
	KPH Speed = 1.60934
)

const (
	BlueColor  Color = "blue"
	GreenColor Color = "green"
	RedColor   Color = "red"
)

const (
	SportsWheels Wheels = "sports"
	SteelWheels  Wheels = "steel"
)

func (a *Assembly) WithColor(c Color) Builder {
	a.color = c
	return a
}

func (a *Assembly) WithWheels(w Wheels) Builder {
	a.wheels = w
	return a
}

func (a *Assembly) TopSpeed(s Speed) Builder {
	a.speed = s
	return a
}

func (a *Assembly) Build() Car {
	return &BMW{
		color:  a.color,
		wheels: a.wheels,
		speed:  a.speed,
	}
}

func (b *BMW) Drive() error {
	return nil
}

func (b *BMW) Stop() error {
	return nil
}

func (b *BMW) Color() Color {
	return b.color
}

func (b *BMW) Wheels() Wheels {
	return b.wheels
}

func (b *BMW) Speed() Speed {
	return b.speed
}

func NewBuilder() Builder {
	return &Assembly{}
}
