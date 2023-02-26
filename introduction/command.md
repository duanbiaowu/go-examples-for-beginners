# 概述

调用 `os/exec` 包即可。

# 例子

```go
package main

import (
	"fmt"
	"os/exec"
)

func main() {
	out, err := exec.Command("date").Output()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", out)

	out, err = exec.Command("git", "--version").Output()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", out)
}

// $ go run main.go
// 输出如下，你的输出可能和这里的不一样
/**
  Thu Nov  3 08:14:57 CST 2022

  git version 2.30.1 (Apple Git-130)
*/
```