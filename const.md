# 常量

关键字 `const`, 和其他编程语言中常量的语义一样，定义后无法修改。

# 语法规则

```shell
const 常量名称 [常量类型] =  常量值

# 例子
const Pi      float     =  3.14159
```

其中，常量类型为可选，因为编译器可以根据值来推断其类型 (建议指定类型，可以增强语义性)。

## 同时定义多个常量

```shell
const (
  常量名称 [常量类型] =  常量值
  常量名称 [常量类型] =  常量值
  常量名称 [常量类型] =  常量值
  ...
)
````
  
### 例子

```go
package main

const (
   Sunday    = 0
   Monday    = 1
   Tuesday   = 2
   Wednesday = 3
   Thursday  = 4
   Friday    = 5
   Saturday  = 6
)

func main() {
   println(Sunday)
   println(Monday)
   println(Tuesday)
   println(Wednesday)
   println(Thursday)
   println(Friday)
   println(Saturday)
}

// $ go run main.go
// 输出如下 
/**
    0
    1
    2
    3
    4
    5
    6
*/
```

# 显式/隐式

1. 浮点类型
   * 显式类型定义： `const Pi float64 = 3.14159`
   * 隐式类型定义： `const Pi = 3.14159`
2. 整型
   * 显式类型定义： `const Page int = 1`
   * 隐式类型定义： `const Page = 1`
3. 字符串
    * 显式类型定义： `const Name string = "abc"`
    * 隐式类型定义： `const Name = "abc"`
4. 其他类型以此类推

# 赋值规则

常量的值必须在编译时就能确定。
* 正确的：`const N = 10/2`
* 错误的：`const N = getNumber()  // 引发构建错误: getNumber() used as value`