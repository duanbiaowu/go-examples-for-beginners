package k8s

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

// 1. Simple Visitor Example ----------------------------------------

type Visitor func(shape Shape) ([]byte, error)

type Shape interface {
	accept(Visitor) ([]byte, error)
}

type Circle struct {
	Radius int
}

type Rectangle struct {
	Width  int
	Height int
}

func (c *Circle) accept(v Visitor) ([]byte, error) {
	return v(c)
}

func (r *Rectangle) accept(v Visitor) ([]byte, error) {
	return v(r)
}

func JsonVisitor(shape Shape) ([]byte, error) {
	bytes, err := json.Marshal(shape)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func XmlVisitor(shape Shape) ([]byte, error) {
	bytes, err := xml.Marshal(shape)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

// 2. Simplified Kubectl Example ----------------------------------------

type VisitorFunc func(*Info, error) error

type VisitorCtl interface {
	Visit(VisitorFunc) error
}

type Info struct {
	Namespace   string
	Name        string
	OtherThings string
}

func (info *Info) Visit(fn VisitorFunc) error {
	return fn(info, nil)
}

type NameVisitor struct {
	visitor VisitorCtl
}

func (v NameVisitor) Visit(fn VisitorFunc) error {
	return v.visitor.Visit(func(info *Info, err error) error {
		fmt.Println("NameVisitor() before call function")
		err = fn(info, err)
		if err == nil {
			fmt.Printf("==> Name=%s, NameSpace=%s\n", info.Name, info.Namespace)
		}
		fmt.Println("NameVisitor() after call function")
		return err
	})
}

type OtherThingsVisitor struct {
	visitor VisitorCtl
}

func (v OtherThingsVisitor) Visit(fn VisitorFunc) error {
	return v.visitor.Visit(func(info *Info, err error) error {
		fmt.Println("OtherThingsVisitor() before call function")
		err = fn(info, err)
		if err == nil {
			fmt.Printf("==> OtherThings=%s\n", info.OtherThings)
		}
		fmt.Println("OtherThingsVisitor() after call function")
		return err
	})
}

type LogVisitor struct {
	visitor VisitorCtl
}

func (v LogVisitor) Visit(fn VisitorFunc) error {
	return v.visitor.Visit(func(info *Info, err error) error {
		fmt.Println("LogVisitor() before call function")
		err = fn(info, err)
		fmt.Println("LogVisitor() after call function")
		return err
	})
}

// 3. Refactor with DecoratedVisitor ----------------------------------------

type DecoratedVisitor struct {
	visitor    VisitorCtl
	decorators []VisitorFunc
}

func NewDecoratedVisitorCtl(v VisitorCtl, fn ...VisitorFunc) VisitorCtl {
	if len(fn) == 0 {
		return v
	}
	return DecoratedVisitor{v, fn}
}

func (v DecoratedVisitor) Visit(fn VisitorFunc) error {
	return v.visitor.Visit(func(info *Info, err error) error {
		if err != nil {
			return nil
		}
		if err = fn(info, nil); err != nil {
			return err
		}
		for i := range v.decorators {
			if err = v.decorators[i](info, nil); err != nil {
				return err
			}
		}
		return nil
	})
}

func NameVisitorFun(info *Info, err error) error {
	if err == nil {
		fmt.Printf("==> Name=%s, NameSpace=%s\n", info.Name, info.Namespace)
	}
	return err
}

func OtherThingsVisitorFun(info *Info, err error) error {
	if err == nil {
		fmt.Printf("==> OtherThings=%s\n", info.OtherThings)
	}
	return err
}
