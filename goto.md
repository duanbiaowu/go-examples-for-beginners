# 概述
关键字 `goto` 可以指定程序跳转到指定的位置执行，那么这个位置如何表示呢？使用 `标签` 来表示 (可以理解为标签就是一个变量)。

# 语法规则
标签的名称大小写敏感，可以搭配 `for`, `switch` 语句使用。

```shell
# 配合 for 使用
标签名称:
    for 初始表达式; 条件表达式; 迭代表达式 {
        // do something
        // [goto|continue|break] 标签名称
    }
```

# 例子
**注意: 该示例会无限输出，需要请按 `Ctrl + C`。**

```go
package main

import "fmt"

func main() {
LABEL1:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if j == 2 {
				goto LABEL1     // 感兴趣的读者可以将这里的 goto 改为 break 或 continue, 体验下不同的用法
			}
			fmt.Printf("i = %d, j = %d\n", i, j)
		}
	}
}
// $ go run main.go
// 无限输出 
/**
    i = 0, j = 0
    i = 0, j = 1
    ...
 */
```