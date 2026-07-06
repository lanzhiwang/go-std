
--------------------------------------------------------------------------------------

```bash
$ docker run -ti --rm --name golang golang:1.26.4 bash
root@db3a6980e246:/go#
root@db3a6980e246:/go# go version
go version go1.26.4 linux/amd64
root@db3a6980e246:/go#
root@db3a6980e246:/go# go help generate
usage: go generate [-run regexp] [-n] [-v] [-x] [build flags] [file.go... | packages]

Generate runs commands described by directives within existing
files. Those commands can run any process but the intent is to
create or update Go source files.

Go generate is never run automatically by go build, go test,
and so on. It must be run explicitly.

Go generate scans the file for directives, which are lines of
the form,

	//go:generate command argument...

(note: no leading spaces and no space in "//go") where command
is the generator to be run, corresponding to an executable file
that can be run locally. It must either be in the shell path
(gofmt), a fully qualified path (/usr/you/bin/mytool), or a
command alias, described below.

Note that go generate does not parse the file, so lines that look
like directives in comments or multiline strings will be treated
as directives.

The arguments to the directive are space-separated tokens or
double-quoted strings passed to the generator as individual
arguments when it is run.

Quoted strings use Go syntax and are evaluated before execution; a
quoted string appears as a single argument to the generator.

To convey to humans and machine tools that code is generated,
generated source should have a line that matches the following
regular expression (in Go syntax):

	^// Code generated .* DO NOT EDIT\.$

This line must appear before the first non-comment, non-blank
text in the file.

Go generate sets several variables when it runs the generator:

	$GOARCH
		The execution architecture (arm, amd64, etc.)
	$GOOS
		The execution operating system (linux, windows, etc.)
	$GOFILE
		The base name of the file.
	$GOLINE
		The line number of the directive in the source file.
	$GOPACKAGE
		The name of the package of the file containing the directive.
	$GOROOT
		The GOROOT directory for the 'go' command that invoked the
		generator, containing the Go toolchain and standard library.
	$DOLLAR
		A dollar sign.
	$PATH
		The $PATH of the parent process, with $GOROOT/bin
		placed at the beginning. This causes generators
		that execute 'go' commands to use the same 'go'
		as the parent 'go generate' command.

Other than variable substitution and quoted-string evaluation, no
special processing such as "globbing" is performed on the command
line.

As a last step before running the command, any invocations of any
environment variables with alphanumeric names, such as $GOFILE or
$HOME, are expanded throughout the command line. The syntax for
variable expansion is $NAME on all operating systems. Due to the
order of evaluation, variables are expanded even inside quoted
strings. If the variable NAME is not set, $NAME expands to the
empty string.

A directive of the form,

	//go:generate -command xxx args...

specifies, for the remainder of this source file only, that the
string xxx represents the command identified by the arguments. This
can be used to create aliases or to handle multiword generators.
For example,

	//go:generate -command foo go tool foo

specifies that the command "foo" represents the generator
"go tool foo".

Generate processes packages in the order given on the command line,
one at a time. If the command line lists .go files from a single directory,
they are treated as a single package. Within a package, generate processes the
source files in a package in file name order, one at a time. Within
a source file, generate runs generators in the order they appear
in the file, one at a time. The go generate tool also sets the build
tag "generate" so that files may be examined by go generate but ignored
during build.

For packages with invalid code, generate processes only source files with a
valid package clause.

If any generator returns an error exit status, "go generate" skips
all further processing for that package.

The generator is run in the package's source directory.

Go generate accepts two specific flags:

	-run=""
		if non-empty, specifies a regular expression to select
		directives whose full original source text (excluding
		any trailing spaces and final newline) matches the
		expression.

	-skip=""
		if non-empty, specifies a regular expression to suppress
		directives whose full original source text (excluding
		any trailing spaces and final newline) matches the
		expression. If a directive matches both the -run and
		the -skip arguments, it is skipped.

It also accepts the standard build flags including -v, -n, and -x.
The -v flag prints the names of packages and files as they are
processed.
The -n flag prints commands that would be executed.
The -x flag prints commands as they are executed.

For more about build flags, see 'go help build'.

For more about specifying packages, see 'go help packages'.
root@db3a6980e246:/go#
```

--------------------------------------------------------------------------------------

Q: 假设你是一位精通 golang 语言的高级开发专家, 我使用命令 `go help generate` 查看了 go generate 的帮助文档, 内容如下
```
root@db3a6980e246:/go# go version
go version go1.26.4 linux/amd64
root@db3a6980e246:/go#
root@db3a6980e246:/go# go help generate
usage: go generate [-run regexp] [-n] [-v] [-x] [build flags] [file.go... | packages]

Generate runs commands described by directives within existing
files. Those commands can run any process but the intent is to
create or update Go source files.

Go generate is never run automatically by go build, go test,
and so on. It must be run explicitly.

Go generate scans the file for directives, which are lines of
the form,

	//go:generate command argument...

(note: no leading spaces and no space in "//go") where command
is the generator to be run, corresponding to an executable file
that can be run locally. It must either be in the shell path
(gofmt), a fully qualified path (/usr/you/bin/mytool), or a
command alias, described below.

Note that go generate does not parse the file, so lines that look
like directives in comments or multiline strings will be treated
as directives.

The arguments to the directive are space-separated tokens or
double-quoted strings passed to the generator as individual
arguments when it is run.

Quoted strings use Go syntax and are evaluated before execution; a
quoted string appears as a single argument to the generator.

To convey to humans and machine tools that code is generated,
generated source should have a line that matches the following
regular expression (in Go syntax):

	^// Code generated .* DO NOT EDIT\.$

This line must appear before the first non-comment, non-blank
text in the file.

Go generate sets several variables when it runs the generator:

	$GOARCH
		The execution architecture (arm, amd64, etc.)
	$GOOS
		The execution operating system (linux, windows, etc.)
	$GOFILE
		The base name of the file.
	$GOLINE
		The line number of the directive in the source file.
	$GOPACKAGE
		The name of the package of the file containing the directive.
	$GOROOT
		The GOROOT directory for the 'go' command that invoked the
		generator, containing the Go toolchain and standard library.
	$DOLLAR
		A dollar sign.
	$PATH
		The $PATH of the parent process, with $GOROOT/bin
		placed at the beginning. This causes generators
		that execute 'go' commands to use the same 'go'
		as the parent 'go generate' command.

Other than variable substitution and quoted-string evaluation, no
special processing such as "globbing" is performed on the command
line.

As a last step before running the command, any invocations of any
environment variables with alphanumeric names, such as $GOFILE or
$HOME, are expanded throughout the command line. The syntax for
variable expansion is $NAME on all operating systems. Due to the
order of evaluation, variables are expanded even inside quoted
strings. If the variable NAME is not set, $NAME expands to the
empty string.

A directive of the form,

	//go:generate -command xxx args...

specifies, for the remainder of this source file only, that the
string xxx represents the command identified by the arguments. This
can be used to create aliases or to handle multiword generators.
For example,

	//go:generate -command foo go tool foo

specifies that the command "foo" represents the generator
"go tool foo".

Generate processes packages in the order given on the command line,
one at a time. If the command line lists .go files from a single directory,
they are treated as a single package. Within a package, generate processes the
source files in a package in file name order, one at a time. Within
a source file, generate runs generators in the order they appear
in the file, one at a time. The go generate tool also sets the build
tag "generate" so that files may be examined by go generate but ignored
during build.

For packages with invalid code, generate processes only source files with a
valid package clause.

If any generator returns an error exit status, "go generate" skips
all further processing for that package.

The generator is run in the package's source directory.

Go generate accepts two specific flags:

	-run=""
		if non-empty, specifies a regular expression to select
		directives whose full original source text (excluding
		any trailing spaces and final newline) matches the
		expression.

	-skip=""
		if non-empty, specifies a regular expression to suppress
		directives whose full original source text (excluding
		any trailing spaces and final newline) matches the
		expression. If a directive matches both the -run and
		the -skip arguments, it is skipped.

It also accepts the standard build flags including -v, -n, and -x.
The -v flag prints the names of packages and files as they are
processed.
The -n flag prints commands that would be executed.
The -x flag prints commands as they are executed.

For more about build flags, see 'go help build'.

For more about specifying packages, see 'go help packages'.
root@db3a6980e246:/go#
```
请帮我详细总结上述帮助文档的所有内容, 并且根据不同的场景举例说明 go generate 的用法

这份 `go help generate` 的官方文档详细阐述了 `go generate` 命令的设计定位、指令格式、执行规则、环境变量、别名定义以及控制参数.

下面为您将这份文档的核心要点进行结构化整理, 并在后文中结合实际开发场景给出具体的应用示例.

---

### 第一部分: `go help generate` 文档要点总结

#### 1. 核心定位与设计原则

* 用途: 通过识别源文件中的特定注释指令来运行外部工具, 主要用于创建或更新 Go 源代码(如自动生成代码、序列化/反序列化方法等).
* 非自动触发: `go generate` 永远不会被 `go build`、`go test` 等命令自动调用, 必须由开发者手动、显式地运行.
* 非语法解析: `go generate` 不会对 Go 代码进行真正的 AST(抽象语法树)解析. 这意味着, 即使指令写在普通的注释块或多行字符串中, 只要符合格式, 也会被当作指令执行.

#### 2. 指令格式要求

* 基本格式:

    ```go
    //go:generate command argument...
    ```

    * 限制: `//` 前不能有空格, `//go` 之间也不能有空格.
    * `command` 必须是本地可执行文件. 它可以存在于系统的 `PATH` 路径中、或使用绝对路径、或者是定义的别名(Alias).

* 参数规则: 参数使用空格分隔. 如果参数包含空格, 可以使用 Go 语法中的双引号(`" "`)括起来. 在执行前, 双引号内部的变量会被评估解析.

#### 3. 规范要求: 生成代码的标识

为了让其他开发者和静态代码分析工具识别出该文件是自动生成的, 生成的 Go 源文件开头(在第一个非注释、非空白文本之前)应该包含匹配以下正则表达式的注释:

```go
^// Code generated .* DO NOT EDIT\.$
```
例如: `// Code generated by templ; DO NOT EDIT.`

#### 4. 环境变量

`go generate` 在运行生成器工具时, 会向子进程注入以下环境变量, 以便工具能够获取当前上下文:

* `$GOARCH`: 目标架构(如 `amd64`, `arm`).
* `$GOOS`: 目标操作系统(如 `linux`, `windows`).
* `$GOFILE`: 当前包含指令的源文件名(仅文件名, 不含路径).
* `$GOLINE`: 指令在源文件中的行号.
* `$GOPACKAGE`: 当前文件所属的包名.
* `$GOROOT`: 调用当前生成器的 Go SDK 根目录.
* `$DOLLAR`: 美元符号字面量 `$`.
* `$PATH`: 系统的环境变量 `$PATH`, 其中 `$GOROOT/bin` 会被前置到最前面, 确保生成器内部调用 `go` 命令时使用的是同一个 Go 版本.

变量替换机制: 任何 `$NAME` 格式的系统环境变量(包括上述变量及外部环境的 `$HOME` 等)在命令执行前都会被替换. 这种替换在双引号内同样生效. 如果变量未定义, 则替换为空字符串.

#### 5. 命令别名定义(`-command`)

可以使用 `-command` 关键字定义当前源文件内有效的别名, 以便简化后续指令或处理多参数命令:

```go
//go:generate -command xxx args...
```

范围: 仅在当前定义该别名的源文件内有效.

#### 6. 执行顺序与工作目录

* 工作目录: 生成器工具是在该包的源文件所在目录下运行的.
* 包顺序: 按命令行中指定的包顺序, 依次处理. 如果是直接指定一组 `.go` 文件, 它们会被视为一个包.
* 文件顺序: 在同一个包内, 按照文件名的字母顺序(alphabetical order)依次处理.
* 指令顺序: 在同一个文件内, 按照指令从上到下的出现顺序依次执行.
* 中断机制: 如果任何一个指令执行时返回了非 0 的错误状态码, `go generate` 将会终止当前包后续的所有处理.

#### 7. 构建标签(Build Tag)

`go generate` 在运行时会默认隐式设置构建标签 `generate`. 这允许我们在项目中包含一些仅用于提供生成上下文、但在实际编译/构建(`go build`)时需要被忽略的辅助文件.

#### 8. 命令行参数(Flags)

* `-run="regexp"`: 仅执行匹配该正则表达式的指令.

* `-skip="regexp"`: 跳过匹配该正则表达式的指令. 如果一条指令同时匹配了 `-run` 和 `-skip`, 则会被跳过.

* 标准构建参数:
    * `-n`: 打印将要执行的命令, 但不实际运行(干跑/Dry-run).
    * `-x`: 在执行时, 打印出当前正在执行的命令.
    * `-v`: 打印正在处理的包和文件名.

---

### 第二部分: 场景化使用示例

#### 场景一: 基础用法(生成枚举值的 `String()` 方法)

在 Go 中, 我们经常使用 `iota` 定义枚举, 并希望其能直接输出可读的字符串. Go 官方提供了 `stringer` 工具来自动生成此方法.

假设在一个名为 `painkiller.go` 的文件中:

```go
package painkiller

// 1. 定义生成指令, 调用 stringer 工具, 针对 Pill 类型生成代码
//go:generate stringer -type=Pill

type Pill int

const (
	Placebo Pill = iota
	Aspirin
	Ibuprofen
)
```

使用步骤:
1. 安装工具(如未安装): `go install golang.org/x/tools/cmd/stringer@latest`
2. 在项目根目录运行: `go generate ./...`
3. 结果: 在同目录下会自动生成 `pill_string.go` 文件, 其中包含 `Pill` 的 `String()` 方法, 并自动带有 `// Code generated by "stringer ..."; DO NOT EDIT.` 规范标识.

---

#### 场景二: 利用内置环境变量动态生成信息

当我们需要在生成的代码中嵌入"源文件名"、"包名"或"架构信息"时, 可以利用 `go generate` 自动注入的环境变量.

假设我们有一个自定义的生成脚本 `gen_info.sh`, 我们希望将当前文件名和包名传入:

```bash
#!/bin/bash
# gen_info.sh 内容
echo "// Code generated by gen_info.sh; DO NOT EDIT." > info.go
echo "package $1" >> info.go
echo "const SourceFile = \"$2\"" >> info.go
```

在我们的 Go 业务文件 `main.go` 中:

```go
package main

import "fmt"

// 利用内置变量 $GOPACKAGE 和 $GOFILE 作为参数传递
//go:generate ./gen_info.sh $GOPACKAGE $GOFILE

func main() {
	// SourceFile 会由自动生成的 info.go 提供
	fmt.Println("Generated from file:", SourceFile)
}
```

运行 `go generate` 后, 会自动生成 `info.go`:

```go
// Code generated by gen_info.sh; DO NOT EDIT.
package main
const SourceFile = "main.go"
```

---

#### 场景三: 定义和使用命令别名(`-command`)

如果需要调用的生成命令比较长, 或者想用特定的 Go 工具链(比如本地未安装, 通过 `go run` 临时运行某个 package), 可以使用 `-command` 来定义别名.

```go
package db

// 定义别名 'sqlgen', 代表执行 'go run github.com/example/tools/sqlgen'
//go:generate -command sqlgen go run github.com/example/tools/sqlgen

// 后续指令可以直接使用 'sqlgen' 别名
//go:generate sqlgen --table=users --out=users_gen.go
//go:generate sqlgen --table=orders --out=orders_gen.go
```

这样的好处是:
1. 简化了后续每条生成命令的长度.
2. 约束了别名作用域仅在当前文件, 不会污染其他文件的生成逻辑.

---

#### 场景四: 利用构建标签(Build Tag)隔离生成器逻辑

有时我们需要写一些专门用来做"代码生成"或"准备数据"的 Go 辅助程序, 但又不希望这些程序被打包进最终的业务二进制文件中.

我们可以创建一个辅助文件 `generator_tool.go`:

```go
//go:build generate
// +build generate

package main

import (
	"log"
	"os"
)

func main() {
	// 复杂的生成逻辑, 比如读取本地数据库生成 schema.go
	err := os.WriteFile("schema_gen.go", []byte("// Code generated. DO NOT EDIT.\npackage main\n"), 0644)
	if err != nil {
		log.Fatal(err)
	}
}
```

在主业务文件 `main.go` 中, 调用这个辅助程序:

```go
package main

//go:generate go run generator_tool.go

func main() {
	// 业务逻辑
}
```

* 在 `go build` 时: 因为 `generator_tool.go` 头部带有 `//go:build generate` 标签, 而编译时默认没有开启该标签, 因此它不会被编译进最终的程序.
* 在 `go generate` 时: 工具会自动激活 `generate` 标签, 从而可以顺利读取和执行 `generator_tool.go`.

---

#### 场景五: 通过命令行 Flag 精确控制执行

当项目中有非常多的 `go:generate` 指令时, 我们有时只想调试或运行其中的某一部分, 此时可以使用 `-run`、`-skip`、`-n`、`-x`.

1. Dry-run 预览(不实际执行):

   ```bash
   go generate -n ./...
   ```

   这会把所有即将执行的命令打印到控制台, 以便检查变量替换是否正确, 而不会产生任何文件写入.

2. 打印执行详情:

   ```bash
   go generate -x ./...
   ```

   在真正运行命令的同时, 将其打印出来, 便于观察执行进度.

3. 仅执行匹配特定模式的指令:
   假设项目中既有 `stringer` 又有 `mockgen`, 我们只想生成 mock 代码:

   ```bash
   go generate -run "mockgen" ./...
   ```

4. 跳过特定指令:
   如果某些生成任务非常耗时, 我们想在本地开发时临时跳过它们:

   ```bash
   go generate -skip "slow_tool" ./...
   ```

--------------------------------------------------------------------------------------

Q: 假设你是一位精通 golang 语言的高级开发专家, 我使用命令 `go help generate` 查看了 go generate 的帮助文档, 内容如下
```
root@db3a6980e246:/go# go help generate
usage: go generate [-run regexp] [-n] [-v] [-x] [build flags] [file.go... | packages]

Generate runs commands described by directives within existing
files. Those commands can run any process but the intent is to
create or update Go source files.

Go generate is never run automatically by go build, go test,
and so on. It must be run explicitly.

Go generate scans the file for directives, which are lines of
the form,

	//go:generate command argument...

(note: no leading spaces and no space in "//go") where command
is the generator to be run, corresponding to an executable file
that can be run locally. It must either be in the shell path
(gofmt), a fully qualified path (/usr/you/bin/mytool), or a
command alias, described below.

Note that go generate does not parse the file, so lines that look
like directives in comments or multiline strings will be treated
as directives.

The arguments to the directive are space-separated tokens or
double-quoted strings passed to the generator as individual
arguments when it is run.

Quoted strings use Go syntax and are evaluated before execution; a
quoted string appears as a single argument to the generator.

To convey to humans and machine tools that code is generated,
generated source should have a line that matches the following
regular expression (in Go syntax):

	^// Code generated .* DO NOT EDIT\.$

This line must appear before the first non-comment, non-blank
text in the file.

Go generate sets several variables when it runs the generator:

	$GOARCH
		The execution architecture (arm, amd64, etc.)
	$GOOS
		The execution operating system (linux, windows, etc.)
	$GOFILE
		The base name of the file.
	$GOLINE
		The line number of the directive in the source file.
	$GOPACKAGE
		The name of the package of the file containing the directive.
	$GOROOT
		The GOROOT directory for the 'go' command that invoked the
		generator, containing the Go toolchain and standard library.
	$DOLLAR
		A dollar sign.
	$PATH
		The $PATH of the parent process, with $GOROOT/bin
		placed at the beginning. This causes generators
		that execute 'go' commands to use the same 'go'
		as the parent 'go generate' command.

Other than variable substitution and quoted-string evaluation, no
special processing such as "globbing" is performed on the command
line.

As a last step before running the command, any invocations of any
environment variables with alphanumeric names, such as $GOFILE or
$HOME, are expanded throughout the command line. The syntax for
variable expansion is $NAME on all operating systems. Due to the
order of evaluation, variables are expanded even inside quoted
strings. If the variable NAME is not set, $NAME expands to the
empty string.

A directive of the form,

	//go:generate -command xxx args...

specifies, for the remainder of this source file only, that the
string xxx represents the command identified by the arguments. This
can be used to create aliases or to handle multiword generators.
For example,

	//go:generate -command foo go tool foo

specifies that the command "foo" represents the generator
"go tool foo".

Generate processes packages in the order given on the command line,
one at a time. If the command line lists .go files from a single directory,
they are treated as a single package. Within a package, generate processes the
source files in a package in file name order, one at a time. Within
a source file, generate runs generators in the order they appear
in the file, one at a time. The go generate tool also sets the build
tag "generate" so that files may be examined by go generate but ignored
during build.

For packages with invalid code, generate processes only source files with a
valid package clause.

If any generator returns an error exit status, "go generate" skips
all further processing for that package.

The generator is run in the package's source directory.

Go generate accepts two specific flags:

	-run=""
		if non-empty, specifies a regular expression to select
		directives whose full original source text (excluding
		any trailing spaces and final newline) matches the
		expression.

	-skip=""
		if non-empty, specifies a regular expression to suppress
		directives whose full original source text (excluding
		any trailing spaces and final newline) matches the
		expression. If a directive matches both the -run and
		the -skip arguments, it is skipped.

It also accepts the standard build flags including -v, -n, and -x.
The -v flag prints the names of packages and files as they are
processed.
The -n flag prints commands that would be executed.
The -x flag prints commands as they are executed.

For more about build flags, see 'go help build'.

For more about specifying packages, see 'go help packages'.
root@db3a6980e246:/go#
```
请帮我将上述帮助文档翻译为中文

这里为您提供该 `go help generate` 帮助文档的中文翻译, 以便您更准确地理解和使用 `go generate`:

```bash
root@db3a6980e246:/go# go help generate
usage: go generate [-run regexp] [-n] [-v] [-x] [build flags] [file.go... | packages]

Generate runs commands described by directives within existing
files. Those commands can run any process but the intent is to
create or update Go source files.
`go generate` 会运行现有文件中由"指令(directives)"所描述的命令.
这些命令可以运行任何进程, 但其设计意图是创建或更新 Go 源文件.

Go generate is never run automatically by go build, go test,
and so on. It must be run explicitly.
`go build`、`go test` 等命令绝不会自动运行 `go generate`. 它必须被显式手动运行.

Go generate scans the file for directives, which are lines of
the form,
`go generate` 会扫描文件以查找指令, 这些指令采用以下格式:

	//go:generate command argument...

(note: no leading spaces and no space in "//go") where command
is the generator to be run, corresponding to an executable file
that can be run locally. It must either be in the shell path
(gofmt), a fully qualified path (/usr/you/bin/mytool), or a
command alias, described below.
(注意: 最前面不能有空格, 并且 "//go" 之间不能有空格).
其中 `command` 是要运行的生成器, 对应于本地可运行的可执行文件.
它必须处于系统的环境变量 PATH 中(例如 `gofmt`),
是完整限定的绝对路径(例如 `/usr/you/bin/mytool`)或者是下文描述的命令别名.

Note that go generate does not parse the file, so lines that look
like directives in comments or multiline strings will be treated
as directives.
请注意, `go generate` 不会解析 Go 文件的语法, 因此在注释或多行字符串中,
只要格式看起来像该指令的行, 都会被视为指令进行处理.

The arguments to the directive are space-separated tokens or
double-quoted strings passed to the generator as individual
arguments when it is run.
指令的参数是以空格分隔的标记或双引号字符串, 在运行生成器时, 它们会作为单独的参数传递给生成器.

Quoted strings use Go syntax and are evaluated before execution; a
quoted string appears as a single argument to the generator.
引号字符串使用 Go 的语法规则, 并在执行之前进行求值; 一个引号字符串在传入生成器时, 会作为单个参数存在.

To convey to humans and machine tools that code is generated,
generated source should have a line that matches the following
regular expression (in Go syntax):
为了向人类和机器工具表明该代码是自动生成的, 生成的源文件中应该包含一行匹配以下正则表达式(Go 语法)的文本:

	^// Code generated .* DO NOT EDIT\.$

This line must appear before the first non-comment, non-blank
text in the file.
该行必须出现在文件中第一个非注释、非空白文本之前.

Go generate sets several variables when it runs the generator:
`go generate` 在运行生成器时会设置以下几个环境变量:

	$GOARCH
		The execution architecture (arm, amd64, etc.)
	$GOOS
		The execution operating system (linux, windows, etc.)
	$GOFILE
		The base name of the file.
	$GOLINE
		The line number of the directive in the source file.
	$GOPACKAGE
		The name of the package of the file containing the directive.
	$GOROOT
		The GOROOT directory for the 'go' command that invoked the
		generator, containing the Go toolchain and standard library.
	$DOLLAR
		A dollar sign.
	$PATH
		The $PATH of the parent process, with $GOROOT/bin
		placed at the beginning. This causes generators
		that execute 'go' commands to use the same 'go'
		as the parent 'go generate' command.

Other than variable substitution and quoted-string evaluation, no
special processing such as "globbing" is performed on the command
line.
除了变量替换和引号字符串求值之外, 命令行上不会进行任何其他特殊处理(例如通配符匹配 "globbing").

As a last step before running the command, any invocations of any
environment variables with alphanumeric names, such as $GOFILE or
$HOME, are expanded throughout the command line. The syntax for
variable expansion is $NAME on all operating systems. Due to the
order of evaluation, variables are expanded even inside quoted
strings. If the variable NAME is not set, $NAME expands to the
empty string.
作为运行命令之前的最后一步, 整个命令行中所有对具有字母数字名称的环境变量(例如 `$GOFILE` 或 `$HOME`)的引用都将被展开.
在所有操作系统上, 变量展开的语法都是 `$NAME`.
由于求值顺序的关系, 即使在双引号字符串内部, 变量也会被展开.
如果未设置变量 `NAME`, 则 `$NAME` 将被展开为空字符串.

A directive of the form,
以下格式的指令:

	//go:generate -command xxx args...

specifies, for the remainder of this source file only, that the
string xxx represents the command identified by the arguments. This
can be used to create aliases or to handle multiword generators.
For example,
仅在此源文件的剩余部分中指定: 字符串 `xxx` 代表由这些参数标识的命令.
这可以用于创建别名或处理由多个单词组成的生成器命令.

	//go:generate -command foo go tool foo

specifies that the command "foo" represents the generator
"go tool foo".
指定命令 `foo` 代表生成器 `go tool foo`.

Generate processes packages in the order given on the command line,
one at a time. If the command line lists .go files from a single directory,
they are treated as a single package. Within a package, generate processes the
source files in a package in file name order, one at a time. Within
a source file, generate runs generators in the order they appear
in the file, one at a time. The go generate tool also sets the build
tag "generate" so that files may be examined by go generate but ignored
during build.
`go generate` 按照命令行中给出的顺序一次处理一个包.
如果命令行中列出了来自单个目录的 `.go` 文件, 它们将被视为同一个包.
在一个包中, `go generate` 按照文件名顺序一次处理一个源文件.
在一个源文件中, `go generate` 按照生成器在文件中出现的顺序一次运行一个.
`go generate` 工具还会设置构建标签(build tag)`"generate"`, 以便文件可以被 `go generate` 检查, 但在日常构建(build)时被忽略.

For packages with invalid code, generate processes only source files with a
valid package clause.
对于代码无效的包, `go generate` 仅处理其中包含有效 `package` 声明子句的源文件.

If any generator returns an error exit status, "go generate" skips
all further processing for that package.
如果任何生成器返回了错误的退出状态码(非零值), `go generate` 将跳过该包的所有后续处理.

The generator is run in the package's source directory.
生成器是在该包的源文件目录下运行的.

Go generate accepts two specific flags:
`go generate` 接受两个特定的标志:

	-run=""
		if non-empty, specifies a regular expression to select
		directives whose full original source text (excluding
		any trailing spaces and final newline) matches the
		expression.
		如果非空, 指定一个正则表达式来选择要执行的指令.
		系统会匹配该指令的完整原始源文本(不包括任何尾随空格和末尾换行符).

	-skip=""
		if non-empty, specifies a regular expression to suppress
		directives whose full original source text (excluding
		any trailing spaces and final newline) matches the
		expression. If a directive matches both the -run and
		the -skip arguments, it is skipped.
		如果非空, 指定一个正则表达式来过滤(跳过)某些指令.
		系统会匹配该指令的完整原始源文本(不包括任何尾随空格和末尾换行符).
		如果一个指令同时匹配了 `-run` 和 `-skip` 参数, 它将被跳过.

It also accepts the standard build flags including -v, -n, and -x.
它还接受标准的构建标志, 包括 `-v`、`-n` 和 `-x`:

The -v flag prints the names of packages and files as they are
processed.
`-v` 标志: 在处理包和文件时打印其名称.

The -n flag prints commands that would be executed.
`-n` 标志: 打印将要执行的命令, 但不实际执行.

The -x flag prints commands as they are executed.
`-x` 标志: 在执行命令时将其打印出来.

For more about build flags, see 'go help build'.

For more about specifying packages, see 'go help packages'.
root@db3a6980e246:/go#

```

--------------------------------------------------------------------------------------


--------------------------------------------------------------------------------------
