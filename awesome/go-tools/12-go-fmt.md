
--------------------------------------------------------------------------------------

```bash
$ docker run -ti --rm --name golang golang:1.26.4 bash
root@1790643b58ae:/go#
root@1790643b58ae:/go# go version
go version go1.26.4 linux/amd64
root@1790643b58ae:/go#
root@1790643b58ae:/go# go help fmt
usage: go fmt [-n] [-x] [packages]

Fmt runs the command 'gofmt -l -w' on the packages named
by the import paths. It prints the names of the files that are modified.

For more about gofmt, see 'go doc cmd/gofmt'.
For more about specifying packages, see 'go help packages'.

The -n flag prints commands that would be executed.
The -x flag prints commands as they are executed.

The -mod flag's value sets which module download mode
to use: readonly or vendor. See 'go help modules' for more.

To run gofmt with specific options, run gofmt itself.

See also: go fix, go vet.
root@1790643b58ae:/go# go doc cmd/gofmt
Gofmt formats Go programs. It uses tabs for indentation and blanks for
alignment. Alignment assumes that an editor is using a fixed-width font.

Without an explicit path, it processes the standard input. Given a file, it
operates on that file; given a directory, it operates on all .go files in that
directory, recursively. (Files starting with a period are ignored.) By default,
gofmt prints the reformatted sources to standard output.

Usage:

    gofmt [flags] [path ...]

The flags are:

    -d
    	Do not print reformatted sources to standard output.
    	If a file's formatting is different than gofmt's, print diffs
    	to standard output.
    -e
    	Print all (including spurious) errors.
    -l
    	Do not print reformatted sources to standard output.
    	If a file's formatting is different from gofmt's, print its name
    	to standard output.
    -r rule
    	Apply the rewrite rule to the source before reformatting.
    -s
    	Try to simplify code (after applying the rewrite rule, if any).
    -w
    	Do not print reformatted sources to standard output.
    	If a file's formatting is different from gofmt's, overwrite it
    	with gofmt's version. If an error occurred during overwriting,
    	the original file is restored from an automatic backup.

Debugging support:

    -cpuprofile filename
    	Write cpu profile to the specified file.

The rewrite rule specified with the -r flag must be a string of the form:

    pattern -> replacement

Both pattern and replacement must be valid Go expressions. In the pattern,
single-character lowercase identifiers serve as wildcards matching arbitrary
sub-expressions; those expressions will be substituted for the same identifiers
in the replacement.

When gofmt reads from standard input, it accepts either a full Go program or a
program fragment. A program fragment must be a syntactically valid declaration
list, statement list, or expression. When formatting such a fragment,
gofmt preserves leading indentation as well as leading and trailing spaces,
so that individual sections of a Go program can be formatted by piping them
through gofmt.

# Examples

To check files for unnecessary parentheses:

    gofmt -r '(a) -> a' -l *.go

To remove the parentheses:

    gofmt -r '(a) -> a' -w *.go

To convert the package tree from explicit slice upper bounds to implicit ones:

    gofmt -r 'α[β:len(α)] -> α[β:]' -w $GOROOT/src

# The simplify command

When invoked with -s gofmt will make the following source transformations where
possible.

    An array, slice, or map composite literal of the form:
    	[]T{T{}, T{}}
    will be simplified to:
    	[]T{{}, {}}

    A slice expression of the form:
    	s[a:len(s)]
    will be simplified to:
    	s[a:]

    A range of the form:
    	for x, _ = range v {...}
    will be simplified to:
    	for x = range v {...}

    A range of the form:
    	for _ = range v {...}
    will be simplified to:
    	for range v {...}

This may result in changes that are incompatible with earlier versions of Go.
root@1790643b58ae:/go#
```

--------------------------------------------------------------------------------------

Q: 假设你是一位精通 golang 语言的高级开发专家, 我使用命令 `go help fmt` 查看了 go fmt 的帮助文档, 内容如下
```
```
请帮我详细总结上述帮助文档的所有内容, 并且根据不同的场景举例说明 go fmt 的用法

你好! 这份帮助文档详细介绍了 Go 语言中代码格式化工具 `go fmt` 和 `gofmt` 的工作原理、参数选项以及高级重构功能.

作为 Go 开发者, 理解这两个工具的区别和配合使用非常重要:
* `go fmt` 是一个高级包装命令, 主要面向 Go 包(packages)进行操作, 它在后台默认调用了 `gofmt -l -w`.
* `gofmt` 是底层的格式化引擎, 直接面向文件或目录, 并提供了更丰富的底层选项(如代码简化 `-s`、代码重写 `-r` 等).

以下是对上述文档内容的详细梳理以及在实际开发场景中的应用示例.

---

### 一、帮助文档内容详细总结

#### 1. `go fmt` 命令总结

* 语法: `go fmt [-n] [-x] [packages]`
* 核心行为: 对指定的包执行 `gofmt -l -w`. 它会直接修改(覆盖)格式不规范的文件, 并打印出被修改的文件名.
* 常用参数:
  * `-n`: 模拟执行(Dry-run). 只打印出将要执行的命令, 而不实际执行它们.
  * `-x`: 打印并执行. 在执行格式化命令时, 将具体的底层命令打印出来.
  * `-mod`: 设置模块下载模式(值可为 `readonly` 或 `vendor`).

#### 2. `gofmt` 命令总结

* 语法: `gofmt [flags] [path ...]`
* 核心行为: 默认使用 Tab 键进行缩进, 空格进行对齐. 如果不指定路径, 则处理标准输入(stdin). 如果指定文件则处理该文件, 如果指定目录则递归处理该目录下所有 `.go` 文件(忽略以 `.` 开头的隐藏文件). 默认情况下, 格式化后的内容会输出到标准输出(stdout), 不修改原文件.
* 核心参数:
  * `-w`: 写入文件. 直接用格式化后的内容覆盖原文件, 若覆盖失败会从自动备份中恢复.
  * `-l`: 列出差异文件. 只打印格式不规范的文件路径, 不输出格式化后的内容, 也不修改原文件.
  * `-d`: 展示差异(Diff). 只打印代码格式差异, 不输出完整内容, 也不修改原文件.
  * `-e`: 打印所有语法错误(包括由于前序错误引发的次生错误).
  * `-s`: 代码自动简化. 尝试对代码进行四种特定模式的简化(详见下文).
  * `-r rule`: 结构化代码重写. 使用指定的规则(`pattern -> replacement`)在格式化前对源码进行替换.

#### 3. `-s` 自动简化规则

启用 `-s` 参数时, `gofmt` 会自动执行以下四种代码简化:
1. 复合字面量类型简化: `[]T{T{}, T{}}` 简化为 `[]T{{}, {}}`(省略切片/数组元素中冗余的类型声明).
2. 切片上界简化: `s[a:len(s)]` 简化为 `s[a:]`(省略冗余的 `len` 计算).
3. Range 迭代简化 1: `for x, _ = range v {...}` 简化为 `for x = range v {...}`(省略未使用的匿名占位符).
4. Range 迭代简化 2: `for _ = range v {...}` 简化为 `for range v {...}`(完全省略占位符).

#### 4. `-r` 重写规则语法

* 形式为 `'pattern -> replacement'`, 两者必须是合法的 Go 表达式.
* 在 pattern 中, 单个字符的小写字母作为通配符, 代表任意子表达式. 在 replacement 中, 相同的字母会被替换为匹配到的子表达式.

---

### 二、针对不同场景的实际应用示例

#### 场景 1: 日常开发中快速格式化代码

在日常开发中, 我们通常希望对整个项目或特定包进行格式化, 并直接修改文件.

* 格式化当前目录及所有子目录下的包:

  ```bash
  go fmt ./...
  ```

* 在修改前想确认一下具体会执行哪些底层命令(不真正修改):

  ```bash
  go fmt -n ./...
  # 输出示例:
  # gofmt -l -w /path/to/project/main.go
  ```

#### 场景 2: CI/CD 流程中进行代码格式检查(不修改文件)

在流水线(CI/CD)中, 我们通常只希望检查代码是否符合规范, 如果不符合则构建报错, 但不需要流水线去自动修改代码.

* 列出所有格式不规范的文件名(若输出不为空, 说明有代码未格式化):

  ```bash
  # 如果有未格式化的文件, 会输出文件名
  gofmt -l .
  ```

* 在 CI 中展示不合规代码的具体差异(Diff), 便于开发者修复:

  ```bash
  gofmt -d .
  ```

  CI 脚本示例:

  ```bash
  # 如果 gofmt -l 输出了任何内容, 则脚本退出并报错
  if [ -n "$(gofmt -l .)" ]; then
      echo "以下文件未通过格式化检查, 请运行 'go fmt ./...' 修复: "
      gofmt -d .
      exit 1
  fi
  ```

#### 场景 3: 自动简化项目中的冗余代码

Go 的语法在演进中变得越来越简洁. 我们可以通过 `-s` 参数让工具自动将旧式代码升级为更现代、更精简的写法.

* 对当前目录下的代码进行自动简化并覆盖文件:

  ```bash
  gofmt -s -w .
  ```

* 效果对比:

  * 简化前:

    ```go
    var m = []Point{Point{1, 2}, Point{3, 4}}
    for i, _ := range m {
        _ = m[i:len(m)]
    }
    ```

  * 简化后:

    ```go
    var m = []Point{{1, 2}, {3, 4}}
    for i := range m {
        _ = m[i:]
    }
    ```

#### 场景 4: 批量代码重构(使用 `-r` 重写规则)

这个功能在重构大型项目时非常高效.

* 例子 A: 去除代码中多余的括号
  在开发中可能会不小心写出双重括号或冗余括号.

  ```bash
  # 检查是否有冗余括号
  gofmt -r '(a) -> a' -l *.go

  # 自动去除冗余括号并写入文件
  gofmt -r '(a) -> a' -w *.go
  ```

  效果: `return (x + y)` 会自动变为 `return x + y`.

* 例子 B: 升级废弃的 API 或重构方法调用
  假设你想把项目中所有的 `bytes.Compare(a, b) == 0`(比较切片是否相等)重构为更高效的 `bytes.Equal(a, b)`.

  ```bash
  gofmt -r 'bytes.Compare(a, b) == 0 -> bytes.Equal(a, b)' -w .
  ```

  在这个规则中, `a` 和 `b` 是通配符, 它会自动匹配并替换任何复杂的表达式, 比如:
  * 旧代码: `bytes.Compare(user.Token, getExpectedToken(id)) == 0`
  * 新代码: `bytes.Equal(user.Token, getExpectedToken(id))`

#### 场景 5: 编辑器/IDE 插件集成(处理标准输入)

许多编辑器(如 VS Code、Vim、Emacs)在保存文件时会调用 `gofmt`. 因为编辑器本身已经打开了文件, 它们通常会通过标准输入将临时内容传给 `gofmt`.

* 格式化代码片段(无需保存为文件):

  ```bash
  echo "package main; import \"fmt\"; func main() { fmt.Println( \"hello\" ) }" | gofmt
  ```

  输出:

  ```go
  package main

  import "fmt"

  func main() {
  	fmt.Println("hello")
  }
  ```

--------------------------------------------------------------------------------------

--------------------------------------------------------------------------------------

--------------------------------------------------------------------------------------

--------------------------------------------------------------------------------------

--------------------------------------------------------------------------------------


