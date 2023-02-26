# 概述

`interface` 类型数据结构内部实现包含 2 个字段， 类型 `Type` 和 值 `Value`。

![interface](images/interface_type.png)

一个接口只有 `Type == nil` 并且 `Value == unset 状态`，该接口才等于 nil 。

**比较规则:**

- 两个 `接口值` 进行比较时，先比较 `Type`，再比较 `Value`
- `接口值` 与 `非接口值` 进行比较时，先将 `非接口值` 转换为 `接口值`，然后再进行比较

## 两个 nil 可能不相等

```go
package main

import (
	"fmt"
)

func main() {
	var p *int = nil
	var v interface{} = p // 赋值完成
	fmt.Println(v == p)   // true
	fmt.Println(p == nil) // true
	fmt.Println(v == nil) // false
}

// $ go run main.go
// 输出如下
/**
  true
  true
  false
*/
```

### 说明

上面的示例代码中，将一个 `指针 nil 值` p 赋值给 `接口 nil 值` v， 赋值完成后，v 的内部字段为 `(Type = *int, V = nil)` 。

1. 第一个比较: v 与 p 比较时，将 p 转换为接口类型后再进行比较
    - p 转换为接口类型后等于 (T = *int, V = nil)
    - v 等于 (T = *int, V = nil)
    - 两者相等 

2. 第二个比较: p 与 nil 比较时，直接比较值，两者相等
 
3. 第三个比较: v 与 nil 比较时，将 nil 转换为接口类型后再进行比较
    - nil 转换为接口类型后等于 (T = nil, V = nil)
    - v 等于 (T = *int, V = nil)
    - 两者不相等

## 接口类型和 nil 可能不相等

非空接口的运行时使用 `runtime.iface` 结构体表示，原型如下:

```go
type iface struct {
    tab  *itab
    data unsafe.Pointer
}
```

其中，**仅当 `tab` 和 `data` 都为 nil 时, 接口类型值才等于 nil**。

```go
package main

import (
	"fmt"
)

func main() {
	var a interface{} = nil         // tab = nil, data = nil
	var b interface{} = (*int)(nil) // tab 包含 *int 类型信息, data = nil

	fmt.Println(a == nil) // true
	fmt.Println(b == nil) // false
}

// $ go run main.go
// 输出如下
/**
  true
  false
*/
```

# 扩展阅读

1. [类型比较](../introduction/type_comparison.md)
2. [make VS new](../introduction/make_with_new.md)