# 概述

`bufio` 包实现了 `缓冲 IO`，它通过在内部封装一个 `io.Reader` 或 `io.Writer` 来实现具体的读写操作。
通过 `缓冲 IO` 可以大大提升 IO 操作的性能，但是有时候，缓冲区也会带来一些违反直觉的问题。

下面的这个小案例是笔者在真实项目中遇到的，整理一下，分享给大家。

## 错误的做法

```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	_, err := fmt.Fprintf(os.Stdout, "%s\n", "[unbuffered] hello world")
	if err != nil {
		panic(err)
	}

	buf := bufio.NewWriter(os.Stdout)
	_, err = fmt.Fprintf(buf, "%s\n", "[buffered] hello world") // 不会输出
	if err != nil {
		panic(err)
	}
}

// $ go run main.go
// 输出如下 
/**
  [unbuffered] hello world
*/
```

**通过输出的结果可以看到，缓冲区数据并没有刷出**。

### Debug

追踪一下源代码，看看问题出在哪里。

**文件名称: `$GOROOT/go/src/bufio/bufio.go`**

`bufio.NewWriter` 返回的对象为 `bufio.Writer`, 实现了标准库的 `Write` 接口:

```go
func (b *Writer) Write(p []byte) (nn int, err error) {
	// 如果写入数据长度大于缓冲区可用数据长度
	for len(p) > b.Available() && b.err == nil {
		var n int
		if b.Buffered() == 0 {
			// 如果缓冲区未写入任何数据，直接写入
			n, b.err = b.wr.Write(p)
		} else {
			// 如果缓冲区存在写入的数据
			// 复制数据，然后调用 Flush
			n = copy(b.buf[b.n:], p)
			b.n += n
			b.Flush()
		}
		nn += n
		p = p[n:]
	}
	// 发生了错误
	if b.err != nil {
		return nn, b.err
	}
	n := copy(b.buf[b.n:], p)
	b.n += n
	nn += n
	return nn, nil
}

// 返回缓冲区可用数据长度
func (b *Writer) Available() int { return len(b.buf) - b.n }

// 返回已写入缓冲区的长度
func (b *Writer) Buffered() int { return b.n }
```

错误的原因在于: 缓冲区的默认长度是 `4096`, 我们写入的数据太少了，远低于剩余缓冲区, 所以没有输出。

## 正确的做法

为了解决这个问题，我们可以模仿 `Write` 方法内部实现的写法，写入数据完成后，主动调用 `Flush` 方法刷出缓冲区。

```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	_, err := fmt.Fprintf(os.Stdout, "%s\n", "[unbuffered] hello world")
	if err != nil {
		panic(err)
	}

	buf := bufio.NewWriter(os.Stdout)
	_, err = fmt.Fprintf(buf, "%s\n", "[buffered] hello world")
	if err != nil {
		panic(err)
	}

	err = buf.Flush()
	if err != nil {
		panic(err)
	}
}

// $ go run main.go
// 输出如下 
/**
  [unbuffered] hello world
  [buffered] hello world
*/
```

**通过输出的结果可以看到，缓冲区数据已经可以正常刷出了**。