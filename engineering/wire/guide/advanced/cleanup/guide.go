package cleanup

import (
	"errors"
	"fmt"
	"os"
	"time"
)

type Foo int
type Bar int

func ProvideFoo() Foo {
	return 0
}

func ProvideBar() (Bar, func(), error) {
	cleanup := func() {
		println("cleanup!")
	}
	if time.Now().Unix()%2 == 0 {
		return 0, nil, errors.New("I made error")
	}
	return 1, cleanup, nil
}

type FooBar struct {
	MyFoo Foo
	MyBar Bar
}

func provideFile(path string) (*os.File, func(), error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, nil, err
	}
	cleanup := func() {
		if err = f.Close(); err != nil {
			fmt.Println(err)
		}
	}
	return f, cleanup, nil
}
