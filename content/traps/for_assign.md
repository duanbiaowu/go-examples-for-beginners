# 循环赋值错误

## 错误的做法

```go
package main

import "fmt"

func main() {
	number := make([]int, 5)

	for i, v := range number {
		v = i
		fmt.Printf(" %v", v)
	}

	fmt.Println("\n", number)
}
// $ go run main.go
// 输出如下 
/**
     0 1 2 3 4
    [0 0 0 0 0]
*/
```

**错误的原因在于**: 循环时的 `v` 变量是从当前元素复制出来的一个临时变量，修改它的值并不会影响到当前元素的值。

## 正确的做法

```go
package main

import "fmt"

func main() {
	number := make([]int, 5)

	for i, _ := range number {
		number[i] = i
		fmt.Printf(" %v", number[i])
	}

	fmt.Println("\n", number)
}
// $ go run main.go
// 输出如下 
/**
    0 1 2 3 4
    [0 1 2 3 4]
*/
```

# 循环引用错误

## 错误的做法

```go
package main

import (
	"fmt"
)

func main() {
	numbers := make([]*int, 5)
	for i, _ := range numbers {
		numbers[i] = &i
		fmt.Printf("%v ", *(numbers[i]))
	}
	fmt.Println()

	for i, _ := range numbers {
		fmt.Printf("%v ", *(numbers[i]))
	}
	fmt.Println()
}
// $ go run main.go
// 输出如下 
/**
    0 1 2 3 4
    4 4 4 4 4
*/
```

**错误的原因在于**: 每次循环时的 `i` 都是同一个变量，将它的地址赋值给当前元素，循环结束后，所有元素指向的都是 `i` 的地址，
也就是说，所有的元素指向的地址对应的值都是 `i` 最后一次修改后的值，也就是 `4` 。

## 正确的做法

```go
package main

import (
	"fmt"
)

func main() {
	numbers := make([]*int, 5)
	for i, _ := range numbers {
		cur := i          // 复制当前值，保证每次赋值给当前元素的变量不一样
		numbers[i] = &cur // 获取复制后的值的地址
		fmt.Printf("%v ", *(numbers[i]))
	}
	fmt.Println()

	for i, _ := range numbers {
		fmt.Printf("%v ", *(numbers[i]))
	}
	fmt.Println()
}
// $ go run main.go
// 输出如下 
/**
    0 1 2 3 4
    0 1 2 3 4
*/
```


