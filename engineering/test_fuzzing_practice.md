# 概述

这篇文章将介绍 `模糊测试` 的基础知识。通过模糊测试，随机数据会针对测试运行并试图找到漏洞或导致程序异常退出的输入数据。
可以通过 `模糊测试` 发现的漏洞类型包括 `SQL 注入`, `缓冲区溢出攻击`, `DOS` 和 `CSRF`。  

我们通过一个小例子来学习，先为一个简单的函数编写模糊测试，然后运行、调试和修复代码中存在的问题。文章中涉及到 `模糊测试` 的名词和前置条件，
在 [模糊测试-理论](test_fuzzing_theory.md) 一文中已经讲过，这里就不再赘述了。

通过示例程序学习分为以下几个步骤:

1. 创建一个目录用于保存代码
2. 编写代码并进行测试
3. 添加单元测试
4. 添加模糊测试
5. 修复两个 `Bug`
6. 学习更多资源

# 创建一个目录用于保存代码

**Linux/Mac:** 

```shell
$ mkdir fuzz
$ cd fuzz  
```

**Windows:**

```shell
C:\> cd %HOMEPATH%
mkdir fuzz
cd fuzz
```

创建目录完成后，创建 `module`:

```shell
$ go mod init example/fuzz
go: creating new go.mod: module example/fuzz
```

# 编写代码并进行测试

编写一个函数，实现功能: `反转字符串`。

将如下代码写入 `main.go` 文件:

```go
package main

import "fmt"

func Reverse(s string) string {
	b := []byte(s)
	for i, j := 0, len(b)-1; i < len(b)/2; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}

func main() {
	input := "The quick brown fox jumped over the lazy dog"
	rev := Reverse(input)
	doubleRev := Reverse(rev)
	fmt.Printf("original: %q\n", input)
	fmt.Printf("reversed: %q\n", rev)
	fmt.Printf("reversed again: %q\n", doubleRev)
}
```

写入完成后，运行代码:

```shell
go run .

# 输出如下
original: "The quick brown fox jumped over the lazy dog"
reversed: "god yzal eht revo depmuj xof nworb kciuq ehT"
reversed again: "The quick brown fox jumped over the lazy dog"
```

# 添加单元测试

将如下代码写入 `reverse_test.go` 文件:

```go
package main

import (
    "testing"
)

func TestReverse(t *testing.T) {
    testcases := []struct {
        in, want string
    }{
        {"Hello, world", "dlrow ,olleH"},
        {" ", " "},
        {"!12345", "54321!"},
    }
    for _, tc := range testcases {
        rev := Reverse(tc.in)
        if rev != tc.want {
                t.Errorf("Reverse: %q, want %q", rev, tc.want)
        }
    }
}
```

写入完成后，运行测试:

```shell
$ go test

# 输出如下
PASS
ok      example/fuzz  0.013s

```

# 添加模糊测试

**`单元测试` 的局限性在于：每个测试用例都必须由开发者手动添加**。`模糊测试` 可以通过自动化添加测试用例，并且覆盖开发者可能没有考虑到的 `边缘场景`。

与 `单元测试` 不同，`模糊测试` 因为无法手动控制测试用例输入，所以自然也就无法指定预期的结果输出。也就是说，需要开发者转变 `测试用例` 的观念，
不再一个一个指定测试用例和期望结果，而是 **告诉 `模糊测试` 需要验证的逻辑规则属性**。

比如，这个例子中需要验证的规则有 2 个:

1. 反转字符串两次之后，返回值和原始值一样
2. 反转字符串之后，字符编码格式为 `UTF8`

`单元测试` 和 `模糊测试` 的语法差异:

1. `模糊测试` 函数以 `FuzzXxx` 而不是 `TestXxx` 开头，参数为 `*testing.F` 而不是 `*testing.T`
2. `单元测试` 代码中的 `t.Run`，`模糊测试` 应该替换为 `f.Fuzz`，并且使用 `f.Add` 作为种子语料库自动化提供 `测试用例`

将如下代码写入 `reverse_test.go` 文件中:

```go
package main

import (
	"testing"
	"unicode/utf8"
)

func TestReverse(t *testing.T) {
	testcases := []struct {
		in, want string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{" ", " "},
		{"!12345", "54321!"},
	}
	for _, tc := range testcases {
		rev := Reverse(tc.in)
		if rev != tc.want {
			t.Errorf("Reverse: %q, want %q", rev, tc.want)
		}
	}
}

func FuzzReverse(f *testing.F) {
    testcases := []string{"Hello, world", " ", "!12345"}
    for _, tc := range testcases {
        f.Add(tc)  // Use f.Add to provide a seed corpus
    }
    f.Fuzz(func(t *testing.T, orig string) {
        rev := Reverse(orig)
        doubleRev := Reverse(rev)
        if orig != doubleRev {
            t.Errorf("Before: %q, after: %q", orig, doubleRev)
        }
        if utf8.ValidString(orig) && !utf8.ValidString(rev) {
            t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
        }
    })
}
```

## 运行代码

1. 运行测试，但是不指定运行模糊测试，确保 `种子语料库` 通过

```shell
$ go test

# 输出如下
PASS
ok      example/fuzz  0.013s
```

2. 指定运行模糊测试，使用标志 `-fuzz` 

```shell
$ go test -fuzz=Fuzz
fuzz: elapsed: 0s, gathering baseline coverage: 0/3 completed
fuzz: elapsed: 0s, gathering baseline coverage: 3/3 completed, now fuzzing with 8 workers
fuzz: minimizing 38-byte failing input file...
--- FAIL: FuzzReverse (0.01s)
    --- FAIL: FuzzReverse (0.00s)
        reverse_test.go:20: Reverse produced invalid UTF-8 string "\x9c\xdd"

    Failing input written to testdata/fuzz/FuzzReverse/af69258a12129d6cbba438df5d5f25ba0ec050461c116f777e77ea7c9a0d217a
    To re-run:
    go test -run=FuzzReverse/af69258a12129d6cbba438df5d5f25ba0ec050461c116f777e77ea7c9a0d217a
FAIL
exit status 1
FAIL    example/fuzz  0.030s
```

`模糊测试` 发生报错，测试失败，导致报错的用例被写入 `种子语料库`，该文件将会在下次调用 `go test` 时候运行，即使不指定 `-fuzz` 标志。
要查看导致失败的测试用例，请在文本编辑器中打开写入 `testdata/fuzz/FuzzReverse` 目录的语料库文件。你的种子语料库文件可能包含不同的字符串，但格式是相同的。

```shell
# 示例文件内容如下
# 文件名称: fuzz/testdata/fuzz/FuzzReverse/af69258a12129d6cbba438df5d5f25ba0ec050461c116f777e77ea7c9a0d217a
go test fuzz v1
string("泃")
```

3. 在不指定 `-fuzz` 标志的情况下，再次运行 `go test`, 将使用新的失败种子语料库

```shell
$ go test
--- FAIL: FuzzReverse (0.00s)
    --- FAIL: FuzzReverse/af69258a12129d6cbba438df5d5f25ba0ec050461c116f777e77ea7c9a0d217a (0.00s)
        reverse_test.go:20: Reverse produced invalid string
FAIL
exit status 1
FAIL    example/fuzz  0.016s
```

# 修复两个 `Bug`

现在，我们来修复上述代码中的 `Bug`, 如果你有时间的话，可以先尝试一下自己解决问题。

## 诊断错误

首先，看一下 `utf8.ValidString` 的文档:

> ValidString reports whether s consists entirely of valid UTF-8-encoded runes.

目前，我们实现的 `Reverse` 函数逐字节反转字符串，显然这是问题所在 (**因为中文需要 3 个字节表示一个字符，反转后就和原始字符的不一样了**)，
所以为了保留原始字符串的 `UTF-8` 编码，必须逐个 `字符` 反转字符串。

将 `reverse_test.go` 文件中的 `FuzzReverse` 函数替换为如下内容:

```go
func FuzzReverse(f *testing.F) {
    testcases := []string{"Hello, world", " ", "!12345"}
    for _, tc := range testcases {
        f.Add(tc)  // Use f.Add to provide a seed corpus
    }
    f.Fuzz(func(t *testing.T, orig string) {
        rev := Reverse(orig)
        doubleRev := Reverse(rev)
        t.Logf("Number of runes: orig=%d, rev=%d, doubleRev=%d", utf8.RuneCountInString(orig), utf8.RuneCountInString(rev), utf8.RuneCountInString(doubleRev))
        if orig != doubleRev {
            t.Errorf("Before: %q, after: %q", orig, doubleRev)
        }
        if utf8.ValidString(orig) && !utf8.ValidString(rev) {
            t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
        }
  })
}
```

主要添加了打印相关代码，这样在测试失败时，可以打印出相关字符串，辅助我们 `Debug`。

## 运行代码

```shell
$ go test
--- FAIL: FuzzReverse (0.00s)
    --- FAIL: FuzzReverse/28f36ef487f23e6c7a81ebdaa9feffe2f2b02b4cddaa6252e87f69863046a5e0 (0.00s)
        reverse_test.go:16: Number of runes: orig=1, rev=3, doubleRev=1
        reverse_test.go:21: Reverse produced invalid UTF-8 string "\x83\xb3\xe6"
FAIL
exit status 1
FAIL    example/fuzz    0.598s
```

整个 `种子语料库` 使用字符串，其中每个字符都是一个字节。但是诸如 `泃` 之类的中文字符可能需要几个字节。因此，中文字符串导致测试失败。

## 修正错误

将 `main.go` 文件中的 `Reverse` 函数替换为如下内容:

```go
func Reverse(s string) string {
    r := []rune(s)
    for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
        r[i], r[j] = r[j], r[i]
    }
    return string(r)
}
```

## 运行代码

1. 运行测试，但是不指定运行模糊测试

```shell
$ go test

# 输出如下
PASS
ok      example/fuzz  0.016s
```

2. 指定  `-fuzz` 标识进行模糊测试，查看是否有新的错误

```shell
$ go test -fuzz=Fuzz
fuzz: elapsed: 0s, gathering baseline coverage: 0/37 completed
fuzz: minimizing 506-byte failing input file...
fuzz: elapsed: 0s, gathering baseline coverage: 5/37 completed
--- FAIL: FuzzReverse (0.02s)
    --- FAIL: FuzzReverse (0.00s)
        reverse_test.go:33: Before: "\x91", after: "�"

    Failing input written to testdata/fuzz/FuzzReverse/1ffc28f7538e29d79fce69fef20ce5ea72648529a9ca10bea392bcff28cd015c
    To re-run:
    go test -run=FuzzReverse/1ffc28f7538e29d79fce69fef20ce5ea72648529a9ca10bea392bcff28cd015c
FAIL
exit status 1
FAIL    example/fuzz  0.032s
```

我们可以看到这个字符串经过两次反转后和原字符串不一样了。这次输入本身是无效的 `unicode`。继续 `Debug` ...

## 诊断错误

现在，我们来修复刚才新产生的 `Bug`, 如果你有时间的话，可以先尝试一下自己解决问题。

Go 中的 `字符串` 是只读的 `字节` 切片，可以包含无效的 `UTF-8` 字节。刚才的测试用例中，原始字符串是一个字节切片，包含一个字节 `\x91`。
当 `Reverse` 函数内部将输入字符串设置为 `[]rune` 时，Go 将字节切片编码为 `UTF-8`，并将字节替换为 `UTF-8` 字符 `�`。
将替换的 `UTF-8` 字符与输入字节片进行比较时，它们显然不相等，于是测试就失败了。

## 修正错误

将 `main.go` 文件中的 `Reverse` 函数替换为如下内容:

```go
func Reverse(s string) string {
    fmt.Printf("input: %q\n", s)
    r := []rune(s)
    fmt.Printf("runes: %q\n", r)
    for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
        r[i], r[j] = r[j], r[i]
    }
    return string(r)
}
```

主要添加了打印相关代码，这样在测试失败时，可以打印出相关字符串，辅助我们 `Debug`。

## 运行代码

这一次，我们只运行失败的测试以检查日志，使用 `go test -run`

```shell
$ go test -run=FuzzReverse/28f36ef487f23e6c7a81ebdaa9feffe2f2b02b4cddaa6252e87f69863046a5e0
input: "\x91"
runes: ['�']
input: "�"
runes: ['�']
--- FAIL: FuzzReverse (0.00s)
    --- FAIL: FuzzReverse/28f36ef487f23e6c7a81ebdaa9feffe2f2b02b4cddaa6252e87f69863046a5e0 (0.00s)
        reverse_test.go:16: Number of runes: orig=1, rev=1, doubleRev=1
        reverse_test.go:18: Before: "\x91", after: "�"
FAIL
exit status 1
FAIL    example/fuzz    0.145s
```

要运行 `FuzzXxx/testdata` 中的 `特定语料库条目`，可以通过指定 `{FuzzTestName}/{filename}` 来运行，这在很有用的调试技巧。

## 修正错误

如果 `Reverse` 的输入不是有效的 `UTF-8`，直接返回一个错误。

1. 将如下代码写入 `main.go` 文件

```go
package main

import (
    "errors"
    "fmt"
    "unicode/utf8"
)

func main() {
    input := "The quick brown fox jumped over the lazy dog"
    rev, revErr := Reverse(input)
    doubleRev, doubleRevErr := Reverse(rev)
    fmt.Printf("original: %q\n", input)
    fmt.Printf("reversed: %q, err: %v\n", rev, revErr)
    fmt.Printf("reversed again: %q, err: %v\n", doubleRev, doubleRevErr)
}

func Reverse(s string) (string, error) {
    if !utf8.ValidString(s) {
        return s, errors.New("input is not valid UTF-8")
    }
    r := []rune(s)
    for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
        r[i], r[j] = r[j], r[i]
    }
    return string(r), nil
}
```

2. 将 `reverse_test.go` 文件中的 `FuzzReverse` 函数替换为如下内容

```go
func FuzzReverse(f *testing.F) {
    testcases := []string {"Hello, world", " ", "!12345"}
    for _, tc := range testcases {
        f.Add(tc)  // Use f.Add to provide a seed corpus
    }
    f.Fuzz(func(t *testing.T, orig string) {
        rev, err1 := Reverse(orig)
        if err1 != nil {
            return
        }
        doubleRev, err2 := Reverse(rev)
        if err2 != nil {
             return
        }
        if orig != doubleRev {
            t.Errorf("Before: %q, after: %q", orig, doubleRev)
        }
        if utf8.ValidString(orig) && !utf8.ValidString(rev) {
            t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
        }
    })
}
```

## 运行代码

运行普通测试:

```shell
$ go test

# 输出如下
PASS
ok      example/fuzz  0.019s
```

除非传递 `-fuzztime` 标志，否则 `模糊测试` 将一直运行，直到它遇到失败的输入。**如果没有发生错误或失败，默认是永远运行**，但是可以使用 `Ctrl-C` 中断进程。
使用 `go test -fuzz=Fuzz` 进行模糊测试，然后在几秒钟后，使用 `Ctrl-C` 停止模糊测试:

```shell
$ go test -fuzz=Fuzz
fuzz: elapsed: 0s, gathering baseline coverage: 0/38 completed
fuzz: elapsed: 0s, gathering baseline coverage: 38/38 completed, now fuzzing with 4 workers
fuzz: elapsed: 3s, execs: 86342 (28778/sec), new interesting: 2 (total: 35)
fuzz: elapsed: 6s, execs: 193490 (35714/sec), new interesting: 4 (total: 37)
fuzz: elapsed: 9s, execs: 304390 (36961/sec), new interesting: 4 (total: 37)
...
fuzz: elapsed: 3m45s, execs: 7246222 (32357/sec), new interesting: 8 (total: 41)
^Cfuzz: elapsed: 3m48s, execs: 7335316 (31648/sec), new interesting: 8 (total: 41)
PASS
ok      example/fuzz  228.000s
```

使用 `go test -fuzz=Fuzz -fuzztime 30s` 进行 `模糊测试`，如果没有发生错误或失败，30 秒后退出 `模糊测试`。

```shell
$ go test -fuzz=Fuzz -fuzztime 30s
fuzz: elapsed: 0s, gathering baseline coverage: 0/5 completed
fuzz: elapsed: 0s, gathering baseline coverage: 5/5 completed, now fuzzing with 4 workers
fuzz: elapsed: 3s, execs: 80290 (26763/sec), new interesting: 12 (total: 12)
fuzz: elapsed: 6s, execs: 210803 (43501/sec), new interesting: 14 (total: 14)
fuzz: elapsed: 9s, execs: 292882 (27360/sec), new interesting: 14 (total: 14)
fuzz: elapsed: 12s, execs: 371872 (26329/sec), new interesting: 14 (total: 14)
fuzz: elapsed: 15s, execs: 517169 (48433/sec), new interesting: 15 (total: 15)
fuzz: elapsed: 18s, execs: 663276 (48699/sec), new interesting: 15 (total: 15)
fuzz: elapsed: 21s, execs: 771698 (36143/sec), new interesting: 15 (total: 15)
fuzz: elapsed: 24s, execs: 924768 (50990/sec), new interesting: 16 (total: 16)
fuzz: elapsed: 27s, execs: 1082025 (52427/sec), new interesting: 17 (total: 17)
fuzz: elapsed: 30s, execs: 1172817 (30281/sec), new interesting: 17 (total: 17)
fuzz: elapsed: 31s, execs: 1172817 (0/sec), new interesting: 17 (total: 17)
PASS
ok      example/fuzz  31.025s
```

`模糊测试` 通过了！

# 总结

恭喜你完成了 Go 中的 `模糊测试` 入门示例。接下来，你可以在项目中选择一个想要模糊测试的函数，然后尝试一下，如果发生了错误或测试失败，正好顺便修复它。

如果您遇到任何问题或对某个功能有想法，[请提出问题](https://github.com/golang/go/issues/new?&labels=fuzz)。

有关该功能的讨论和一般反馈，您还可以参与 `Gophers Slack` 中的 [#fuzzing 频道](https://gophers.slack.com/archives/CH5KV1AKE)。

# Reference

1. [原文](https://go.dev/doc/tutorial/fuzz)

