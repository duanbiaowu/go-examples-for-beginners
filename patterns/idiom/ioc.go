package idiom

import "errors"

type IntSet struct {
	data map[int]bool
}

func NewIntSet() IntSet {
	return IntSet{make(map[int]bool)}
}

func (s *IntSet) Add(x int) {
	s.data[x] = true
}

func (s *IntSet) Delete(x int) {
	delete(s.data, x)
}

func (s *IntSet) Contains(x int) bool {
	return s.data[x]
}

type UndoableIntSet struct {
	IntSet
	functions []func()
}

func NewUndoableIntSet() UndoableIntSet {
	return UndoableIntSet{NewIntSet(), nil}
}

func (s *UndoableIntSet) Add(x int) {
	if !s.Contains(x) {
		s.data[x] = true
		s.functions = append(s.functions, func() {
			s.Delete(x)
		})
	} else {
		s.functions = append(s.functions, nil)
	}
}

func (s *UndoableIntSet) Delete(x int) {
	if !s.Contains(x) {
		s.functions = append(s.functions, nil)
	} else {
		delete(s.data, x)
		s.functions = append(s.functions, func() {
			s.Add(x)
		})
	}
}

// Undo 通过这样的方式来为已有的代码扩展新的功能是一个很好的选择，
// 这样，可以在重用原有代码功能和重新新的功能中达到一个平衡。
// 但是，这种方式最大的问题是，
// Undo 操作其实是一种控制逻辑，并不是业务逻辑，
// 所以，在复用 Undo 这个功能上有问题
// 因为其中加入了大量跟 IntSet 相关的业务逻辑。
func (s *UndoableIntSet) Undo() error {
	if len(s.functions) == 0 {
		return errors.New("no functions to undo")
	}
	index := len(s.functions) - 1
	if function := s.functions[index]; function != nil {
		function()
		s.functions[index] = nil // For garbage collection
	}
	s.functions = s.functions[:index]
	return nil
}

// Undo 反转依赖
type Undo []func()

func (undo *Undo) Add(fn func()) {
	*undo = append(*undo, fn)
}

// Undo 不再依赖于具体的 Set
// 具体的 Set 依赖于 Undo
// Undo 代码达到复用
func (undo *Undo) Undo() error {
	fns := *undo
	if len(fns) == 0 {
		return errors.New("no functions to undo")
	}
	index := len(fns) - 1
	if fn := fns[index]; fn != nil {
		fn()
		fns[index] = nil // For garbage collection
	}
	*undo = fns[:index]
	return nil
}

type FloatSet struct {
	data map[float32]bool
	undo Undo
}

func NewFloatSet() FloatSet {
	return FloatSet{data: make(map[float32]bool)}
}

func (s *FloatSet) Undo() error {
	return s.undo.Undo()
}

func (s *FloatSet) Contains(x float32) bool {
	return s.data[x]
}

func (s *FloatSet) Add(x float32) {
	if !s.Contains(x) {
		s.data[x] = true
		s.undo.Add(func() {
			s.Delete(x)
		})
	} else {
		s.undo.Add(nil)
	}
}

func (s *FloatSet) Delete(x float32) {
	if !s.Contains(x) {
		s.undo.Add(nil)
	} else {
		delete(s.data, x)
		s.undo.Add(func() {
			s.Add(x)
		})
	}
}
