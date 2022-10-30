# 导入包
关键字 `import`

# 语法规则
1. 单个导入 
    ```shell
    import "包名"
    ``` 
2. 多个导入
    ```shell
    import (
        "包名1"
        "包名2"
        "包名3"
        ...
    )
    ```
   
## 例子
1. 导入 `打印包`
    ```go
    package main

    import "fmt"
    
    func main() {
        fmt.Println("hello world")
    }
    ```
2. 导入 `打印包` 和 `字符串包`
    ```go
    package main
    
    import (
        "fmt"
        "strings"
    )
    
    func main() {
        fmt.Println("hello world")
        fmt.Println(strings.Contains("hello world", "hello"))
    }
    
    // $ go run main.go
    // 输出如下
    /**
        hello world
        true
    */
    ```