# 常量生成器

关键字 `iota`, 创建一系列相关的值，省略逐个定义。

# 语法规则

```shell
const (
    常量1 [常量类型] = iota
    常量2
    常量3
    常量4
    常量5
    ...
)
```

## 例子

```go
package main

const (
    Sunday int = iota
    Monday      // 1
    Tuesday     // 2
    Wednesday   // 3
    Thursday    // 4
    Friday      // 5
    Saturday    // 6
)

println(Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday)

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

在上面的声明中，Sunday 的值为 0, Monday 的值为 1, 以此类推。
 