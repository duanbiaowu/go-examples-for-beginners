# 概述

Go 编译后是一个 `二进制可执行文件`，部署非常简单。 在 `1.16` 版本之后， Go 提供了 `embed` 包支持编译时嵌入静态文件，
这样就可以直接在程序中访问静态文件的内容了。 结合这两个特性，可以将应用整体打包进一个二进制可执行文件。

# 常见应用场景

- 将音频、视频文件嵌入小工具内，比如笔者曾经做过一个后台下载小工具，在下载成功/失败时会有对应的提示语音
- 网站的静态资源，比如网站只有少量的静态资源 (icon, 图片，图表等) 文件，可以将这些文件打包进可执行文件 
- 应用配置文件直接打包进可执行文件
- 自定义的静态文件服务

# 例子

## 测试文件

- 新建两个文件并写入一些字符串用作演示
- 新建一个 `server.log` 文件，写入如下内容

```shell
[Server]
[Server]
[Server]
```
 
- 新建一个 `client.log` 文件，写入如下内容

```shell
[Client]
[Client]
[Client]
```

## 单个文件的内容嵌入到字符串

注意 `embed` 注解的写法。

```go
package main

import (
	_ "embed"
	"fmt"
)

//go:embed server.log
var log string

func main() {
	fmt.Printf("%s\n", log)
}

// $ go run main.go
// 输出如下 
/**
  [Server]
  [Server]
  [Server]
*/
```

## 嵌入文件系统 FS

**FS: 表示只读的文件句柄集合**

注意 `embed` 注解的写法。

### 嵌入单个文件

```go
package main

import (
	"embed"
	"fmt"
)

//go:embed server.log
var license embed.FS

func main() {
	content, err := license.ReadFile("server.log")
	if err != nil {
		panic(err)
	} else {
		fmt.Printf("%s\n", content)
	}
}

// $ go run main.go
// 输出如下 
/**
  [Server]
  [Server]
  [Server]
*/
```

### 嵌入多个文件

注意 `embed` 注解的写法。

```go
package main

import (
	"embed"
	"fmt"
)

//go:embed server.log client.log
var license embed.FS

func main() {
	content, err := license.ReadFile("server.log")
	if err != nil {
		panic(err)
	} else {
		fmt.Printf("%s\n", content)
	}

	content, err = license.ReadFile("client.log")
	if err != nil {
		panic(err)
	} else {
		fmt.Printf("%s\n", content)
	}
}

// $ go run main.go
// 输出如下 
/**
  [Server]
  [Server]
  [Server]
  [Client]
  [Client]
  [Client]
*/
```

# 小结

怎么样？`embed` 是不是功能强大，使用简单 :)