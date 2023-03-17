package idiom

import "fmt"

type Country struct {
	Name string
}

type City struct {
	Name string
}

type StringAble interface {
	ToString() string
}

func (c Country) ToString() string {
	return "Country = " + c.Name
}

func (c City) ToString() string {
	return "City = " + c.Name
}

func PrintStr(p StringAble) {
	fmt.Println(p.ToString())
}

type Shape interface {
	Sides() int
	Area() int
}

type Square struct {
	len int
}

func (s *Square) Sides() int {
	return s.len * s.len
}
