# 自定义类型
关键字 `type`, 主要用来对同一种类型进行抽象。

# 语法规则
```shell
type 自定义类型名称 具体类型

# 例子
type Number       int
```

## 同时定义多个自定义类型
```go
type (
    Number int
    Name string
    Has bool
)
```

## 嵌套定义
可以基于已有的自定义类型，定义一个新的自定义类型。
```go
type Number2 Number
```

# 使用规则
和变量使用规则一样。
```go
var x Number = 1024
var n Name = "abc"
var h Has = true
```