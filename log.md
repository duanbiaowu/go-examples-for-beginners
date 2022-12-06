# 概述

调用 `log` 包即可，包里面的方法输出日志时会自动加上日期时间前缀字符。

# 例子

## 输出到终端

```go
package main

import (
	"log"
	"os"
)

func main() {
	log.SetOutput(os.Stdout)

	log.Println("[4.426ms] [rows:1] SELECT * FROM `users` WHERE `id` = 1024")
	log.Printf("[GET] %d %s %s", 200, "OK", "/api/v1/users")
}

// $ go run main.go
// 输出如下，你的输出可能和这里的不一样
/**
  2021/01/03 15:18:55 [4.426ms] [rows:1] SELECT * FROM `users` WHERE `id` = 1024
  2021/01/03 15:18:55 [GET] 200 OK /api/v1/users
*/
```

## 输出到文件

建议先阅读 [创建, 删除文件](file_create_delete.md)。

```go
package main

import (
	"log"
	"os"
)

func main() {
	logFile := "/tmp/test_main_go_server.log"
	file, err := os.Create(logFile)
	if err != nil {
		panic(err)
	}

	defer func() {
		err = file.Close()
		if err != nil {
			panic(err)
		}
	}()

	log.SetOutput(file)

	log.Println("[4.426ms] [rows:1] SELECT * FROM `users` WHERE `id` = 1024")
	log.Printf("[GET] %d %s %s", 200, "OK", "/api/v1/users")
}

// $ go run main.go
// $ cat /tmp/test_main_go_server.log
// 输出如下，你的输出可能和这里的不一样
/**
  2021/01/03 15:25:23 [4.426ms] [rows:1] SELECT * FROM `users` WHERE `id` = 1024
  2021/01/03 15:25:23 [GET] 200 OK /api/v1/users
*/
```
