package k8s

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SimpleVisitor(t *testing.T) {
	c := &Circle{10}
	bytes, err := JsonVisitor(c)
	assert.Nil(t, err)
	assert.Equal(t, `{"Radius":10}`, string(bytes))

	bytes, err = XmlVisitor(c)
	assert.Nil(t, err)
	assert.Equal(t, `<Circle><Radius>10</Radius></Circle>`, string(bytes))
}

// 解耦数据结构和算法
// 使用 Strategy 模式也可以实现，而且更优雅
func Test_SimpleVisitorWithStrategyPattern(t *testing.T) {
	c := Circle{10}
	shapes := []Shape{&c}

	for i := range shapes {
		bytes, err := shapes[i].accept(JsonVisitor)
		assert.Nil(t, err)
		assert.Equal(t, `{"Radius":10}`, string(bytes))

		bytes, err = shapes[i].accept(XmlVisitor)
		assert.Equal(t, `<Circle><Radius>10</Radius></Circle>`, string(bytes))
	}
}

// 以装饰器模式来思考执行流程
func Test_VisitorFunc(t *testing.T) {
	info := Info{}
	// &info 实现了 VisitorCtl
	var v VisitorCtl = &info

	// 层层包装后，v 的数据结构类似链表
	// v = LogVisitor
	v = LogVisitor{v}
	fmt.Println(fmt.Sprintf("v = %T\n", v))

	// v.visitor = LogVisitor
	v = NameVisitor{v}
	fmt.Println(fmt.Sprintf("v = %T\n", v.(NameVisitor).visitor))

	// v.visitor.visitor = LogVisitor
	v = OtherThingsVisitor{v}
	fmt.Println(fmt.Sprintf("v = %T\n", v.(OtherThingsVisitor).visitor.(NameVisitor).visitor))

	err := v.Visit(func(info *Info, err error) error {
		info.Name = "Tom"
		info.Namespace = "NONE"
		info.OtherThings = "We are running as remote team."
		return nil
	})
	assert.Nil(t, err)
}

func Test_VisitorFuncWithDecorator(t *testing.T) {
	info := Info{}
	var v VisitorCtl = &info
	v = NewDecoratedVisitorCtl(v, NameVisitorFun, OtherThingsVisitorFun)

	err := v.Visit(func(info *Info, err error) error {
		info.Name = "service"
		info.Namespace = "default"
		info.OtherThings = "arguments..."
		return nil
	})
	assert.Nil(t, err)
}
