// reference : https://coolshell.cn/articles/21140.html
// see more:
// 	https://github.com/pkg/errors
//	https://pkg.go.dev/errors
// 	一些不错的实践：https://lailin.xyz/post/go-training-03.html

package idiom

import (
	"encoding/binary"
	"io"
)

type Point struct {
	PI   float64
	Uate uint8
	Mine [3]byte
	Too  uint16
}

type Person struct {
	Name   [10]byte
	Age    uint8
	Weight uint8
	err    error
}

// 1. Error Check Hell
func parse(r io.Reader) (*Point, error) {
	var p Point

	if err := binary.Read(r, binary.LittleEndian, &p.PI); err != nil {
		return nil, err
	}
	if err := binary.Read(r, binary.LittleEndian, &p.Uate); err != nil {
		return nil, err
	}
	if err := binary.Read(r, binary.LittleEndian, &p.Mine); err != nil {
		return nil, err
	}
	if err := binary.Read(r, binary.LittleEndian, &p.Too); err != nil {
		return nil, err
	}

	return &p, nil
}

// 2. 使用函数式编程方式
func parse2(r io.Reader) (*Point, error) {
	var p Point
	var err error
	read := func(data interface{}) {
		if err != nil {
			return
		}
		err = binary.Read(r, binary.LittleEndian, data)
	}

	read(&p.PI)
	read(&p.Uate)
	read(&p.Mine)
	read(&p.Too)
	if err != nil {
		return &p, err
	}
	return &p, nil
}

// 3. 清除内部函数

type Reader struct {
	r   io.Reader
	err error
}

func (r *Reader) read(data interface{}) {
	if r.err == nil {
		r.err = binary.Read(r.r, binary.LittleEndian, data)
	}
}

func parse3(input io.Reader) (*Point, error) {
	var p Point
	r := Reader{r: input}

	r.read(&p.PI)
	r.read(&p.Uate)
	r.read(&p.Mine)
	r.read(&p.Too)
	if r.err != nil {
		return nil, r.err
	}
	return &p, nil
}

// 4. 流式接口 Fluent Interface 长度不够，少一个 Weight 字段

func (p *Person) read(input io.Reader, data interface{}) {
	if p.err == nil {
		p.err = binary.Read(input, binary.BigEndian, data)
	}
}

func (p *Person) ReadName(input io.Reader) *Person {
	p.read(input, &p.Name)
	return p
}

func (p *Person) ReadAge(input io.Reader) *Person {
	p.read(input, &p.Age)
	return p
}

func (p *Person) ReadWeight(input io.Reader) *Person {
	p.read(input, &p.Weight)
	return p
}

func (p *Person) Print() *Person {
	return p
}
