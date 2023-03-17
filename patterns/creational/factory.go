package creational

type Gun interface {
	Name() string
	Power() float64
}

type AK47 struct {
}

type A92F struct {
}

type Fist struct {
}

type GunType uint

const (
	FistType GunType = 1 << iota
	RifleType
	PistolType
)

func NewGun(t GunType) Gun {
	switch t {
	case RifleType:
		return newAK47()
	case PistolType:
		return newA92F()
	default:
		return newFist()
	}
}

func (a *AK47) Name() string {
	return "AK47"
}

func (a *AK47) Power() float64 {
	return 30.0
}

func (a *A92F) Name() string {
	return "A92F"
}

func (a *A92F) Power() float64 {
	return 10.0
}

func (f *Fist) Name() string {
	return "Fist"
}

func (f *Fist) Power() float64 {
	return 0.01
}

func newAK47() *AK47 {
	return &AK47{}
}

func newA92F() *A92F {
	return &A92F{}
}

func newFist() *Fist {
	return &Fist{}
}
