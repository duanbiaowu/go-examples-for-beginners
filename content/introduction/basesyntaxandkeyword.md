# 文件名
Go 的文件以 `.go` 为后缀，文件名称必须以字母开头 (任何 UTF-8 编码的字符或 _)，后面跟随 0 个或多个字符或者 Unicode 数字。

正确的命名方式：
* filename
* fileName
* file_name
* filename2
* _filename

错误的命名方式：
* 1filename (以数字开头) 
* switch (Go 关键字)
* x+y (运算符)

# 基本代码格式
* 不需要在语句或声明后面使用分号，除非多个语句和声明出现在同一行，比如后面要讲到的 `for 循环`
* `{` 必须和判断语句、循环语句、函数表达式等在同一行，不能独自成行

# 关键字
**Go 一共 25 个关键字**，简洁到了极点。

大部分关键字其他编程语言中也都有，比较特殊的几个是： `chan`, `defer`, `go`, `select`, 不过这里无需记忆，后面章节都会讲到。

|  |     |     |     |     |
|---|---|---|---|---|
| break |   default  |  func   |   interface  |  select   |
| case | defer    | go    | map    | struct    |
| chan | else    | goto    | package    | switch    |
| const  | fallthrough    | if    | range    | type    |
| continue  | for    | import    | return    | var    |


# 预定义标识符
**Go 一共 37 个预定义标识符**

和关键字一样，大部分关键字其他编程语言中也都有，可能名称有所区别，比如 `int64` 为 `long`, `float64` 为 `double`, 这里无需记忆，后面章节都会讲到。

## 常量
* true
* false
* iota
* nil

## 类型

### 整型
* int
* int8
* int16
* int32
* int64

* uint
* uint8
* uint16
* uint32 
* uint64
* uintptr

### 浮点型
* float32
* float64

### 复数型
* complex64
* complex128

### 布尔型
* bool

### 字节
* byte

### 特殊的整型
* rune (其实就是 int32, 主要用来区分字符值和整数值)

### 字符串类型
* string

### 错误类型
* error

## 函数
* make
* len
* cap
* new
* append
* copy
* close
* delete
* complex
* real
* imag
* panic
* recover
