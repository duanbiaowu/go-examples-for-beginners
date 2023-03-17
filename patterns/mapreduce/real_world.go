package mapreduce

import "reflect"

// Map/Reduce/Filter 只是一种控制逻辑
// 真正的业务逻辑是在传给他们的数据和那个函数来定义的
// 这是一个很经典的“业务逻辑”和“控制逻辑”分离解耦的编程模式
// 业务逻辑多变，控制逻辑抽象，保持两者分离，然后组合使用，确实经典

// Map/Reduce/Filter 基础版本 ------------------------------------

type Employee struct {
	Name     string
	Age      int
	Vacation int
	Salary   int
}

func EmployeeCountIf(employees []Employee, fn func(e *Employee) bool) int {
	counts := 0
	for i := range employees {
		if fn(&employees[i]) {
			counts += 1
		}
	}
	return counts
}

func EmployeeFilterIn(employees []Employee, fn func(e *Employee) bool) []Employee {
	var newList []Employee
	for i := range employees {
		if fn(&employees[i]) {
			newList = append(newList, employees[i])
		}
	}
	return newList
}

func EmployeeSumIf(employees []Employee, fn func(e *Employee) int) int {
	sum := 0
	for i := range employees {
		sum += fn(&employees[i])
	}
	return sum
}

// Map 反射版本 ------------------------------------

func Transform(slice, fn interface{}) interface{} {
	return transform(slice, fn, false)
}

func TransformInPlace(slice, fn interface{}) interface{} {
	return transform(slice, fn, true)
}

func transform(slice, function interface{}, inPlace bool) interface{} {
	// check slice type is Slice
	sliceInType := reflect.ValueOf(slice)
	if sliceInType.Kind() != reflect.Slice {
		panic("transform: not slice")
	}

	// check the fn signature
	fn := reflect.ValueOf(function)
	elemType := sliceInType.Type().Elem()
	if !verifyFuncSignature(fn, elemType, nil) {
		panic("transform: function must be of type func(" + sliceInType.Type().Elem().String() + ") outputElemType")
	}

	sliceOutType := sliceInType
	if !inPlace {
		sliceOutType = reflect.MakeSlice(reflect.SliceOf(fn.Type().Out(0)), sliceInType.Len(), sliceInType.Len())
	}
	for i := 0; i < sliceInType.Len(); i++ {
		sliceOutType.Index(i).Set(fn.Call([]reflect.Value{sliceOutType.Index(i)})[0])
	}
	return sliceOutType.Interface()
}

// Reduce 反射版本 ------------------------------------

func ReflectReduce(slice, pairFunc, zero interface{}) interface{} {
	sliceInType := reflect.ValueOf(slice)
	if sliceInType.Kind() != reflect.Slice {
		panic("reduce: wrong type, not slice")
	}

	length := sliceInType.Len()
	if length == 0 {
		return zero
	} else if length == 1 {
		return sliceInType.Index(0)
	}

	elemType := sliceInType.Type().Elem()
	fn := reflect.ValueOf(pairFunc)
	if !verifyFuncSignature(fn, elemType, elemType, elemType) {
		t := elemType.String()
		panic("reduce: function must be of type func(" + t + ", " + t + ") " + t)
	}

	// 除了数组，也可以使用两个变量实现: prev, cur
	var ins [2]reflect.Value
	ins[0] = sliceInType.Index(0)
	ins[1] = sliceInType.Index(1)
	out := fn.Call(ins[:])[0]

	for i := 2; i < length; i++ {
		ins[0] = out
		ins[1] = sliceInType.Index(i)
		out = fn.Call(ins[:])[0]
	}
	return out.Interface()
}

// Filter 反射版本 ------------------------------------

func ReflectFilter(slice, fn interface{}) interface{} {
	result, _ := filter(slice, fn, false)
	return result
}

func ReflectFilterInPlace(slicePtr, fn interface{}) {
	in := reflect.ValueOf(slicePtr)
	if in.Kind() != reflect.Ptr {
		panic("FilterInPlace: wrong type, " +
			"not a pointer to slice")
	}
	_, n := filter(in.Elem().Interface(), fn, true)
	in.Elem().SetLen(n)
}

func filter(slice, function interface{}, isPlace bool) (interface{}, int) {
	sliceInType := reflect.ValueOf(slice)
	if sliceInType.Kind() != reflect.Slice {
		panic("filter: wrong type, not a slice")
	}

	fn := reflect.ValueOf(function)
	elemType := sliceInType.Type().Elem()
	if !verifyFuncSignature(fn, elemType, reflect.ValueOf(true).Type()) {
		panic("filter: function must be of type func(" + elemType.String() + ") bool")
	}

	var which []int
	for i := 0; i < sliceInType.Len(); i++ {
		if fn.Call([]reflect.Value{sliceInType.Index(i)})[0].Bool() {
			which = append(which, i)
		}
	}

	out := sliceInType
	if !isPlace {
		out = reflect.MakeSlice(sliceInType.Type(), len(which), len(which))
	}
	for i := range which {
		out.Index(i).Set(sliceInType.Index(which[i]))
	}
	return out.Interface(), len(which)
}

// 验证函数, 参数合法性
func verifyFuncSignature(fn reflect.Value, types ...reflect.Type) bool {
	// check it is a function
	if fn.Kind() != reflect.Func {
		return false
	}

	// NumIn() - returns a function type's input parameter count.
	// NumOut() - returns a function type's output parameter count.
	if (fn.Type().NumIn() != len(types)-1) || (fn.Type().NumOut() != 1) {
		return false
	}

	// In() - returns the type of function type's input parameter
	for i := 0; i < len(types)-1; i++ {
		if fn.Type().In(i) != types[i] {
			return false
		}
	}

	// Out() - returns the type of function type's output parameter.
	outType := types[len(types)-1]
	if outType != nil && fn.Type().Out(0) != outType {
		return false
	}
	return true
}
