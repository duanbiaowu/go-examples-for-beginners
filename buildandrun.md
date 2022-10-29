# 示例代码 
```go
package main    // 包名，必须声明

func main() {
	println("hello world")
}
```

# Go 代码组织方式
Go 代码是使用包来组织的，类似于其他编程语言中的库、模块、命名空间。

# 包
一个包由一个或多个 <code>.go</code> 文件组成，放在一个文件夹中。比如字符串相关处理代码全部放在 <code>string</code> 包中。
每个 <code>.go</code> 文件的开始必须使用 <code>package</code> 声明，比如字符串包声明为 <code>package string</code>。

# main 包
一个特殊的包，用来定义具体的执行程序 (比如说我们的业务程序)。

# main 函数
* 如果当前包是 <code>main 包</code>, 那么 <code>main 函数</code> 就是执行程序的入口。
* 如果当前包不是 <code>main 包</code>, 那么 <code>main 函数</code> 就是一个普通的函数。

# Go 程序的运行方式 
1. 编译并运行 (一步完成)
   * 命令行运行 <code>go run 文件名.go</code>, 比如 <code>go run main.go</code>,
2. 先编译为可执行文件，然后运行 (两步完成)
   * 命令行运行 <code>go run 文件名.go</code>, 比如 <code>go run main.go</code>
   * 生成可执行文件，比如 <code>main.exe</code>
   * 执行可执行文件，<code> .\main.exe</code>

# 备注
**理论性的东西只做了简介，暂时一笔带过，因为笔者认为新学一门编程语言时，能快速地通过编写代码，来了解语法以及程序结构比纯看理论要高效地多。**