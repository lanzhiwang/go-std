```bash
$ go help buildconstraint
```

A **build constraint**, also known as a **build tag**, is a condition under which a
file should be included in the package. Build constraints are given by a
line comment that begins
构建约束，也称为构建标签，是文件应包含在包中的条件。构建约束由开始的行注释给出

	//go:build

Constraints may appear in any kind of source file (not just Go), but
they must appear near the top of the file, preceded
only by blank lines and other comments. These rules mean that in Go
files a build constraint must appear before the package clause.
约束可以出现在任何类型的源文件中（不仅仅是 Go），但它们必须出现在文件顶部附近，前面只能有空行和其他注释。
这些规则意味着在 Go 文件中，构建约束必须出现在 package 子句之前。

To distinguish build constraints from package documentation,
a build constraint should be followed by a blank line.
为了区分构建约束和包文档，构建约束后面应该跟一个空行。

A build constraint comment is evaluated as an expression containing
build tags combined by `||`, `&&`, and `!` operators and `parentheses`.
Operators have the same meaning as in Go.
构建约束注释被评估为包含由 ||、&& 和 ! 组合的构建标签的表达式。运算符和括号。运算符的含义与 Go 中相同。

For example, the following build constraint constrains a file to
build when the "linux" and "386" constraints are satisfied, or when
"darwin" is satisfied and "cgo" is not:
例如，以下构建约束会在满足“linux”和“386”约束时，或者在满足“darwin”但不满足“cgo”时约束要构建的文件：

	//go:build (linux && 386) || (darwin && !cgo)

It is an error for a file to have more than one //go:build line.
一个文件包含多个 //go:build 行是错误的。

During a particular build, the following build tags are satisfied:
在特定构建期间，满足以下构建标签：

- the target operating system, as spelled by `runtime.GOOS`, set with the `GOOS` environment variable.

- the target architecture, as spelled by `runtime.GOARCH`, set with the `GOARCH` environment variable.

- any architecture features, in the form `GOARCH.feature` (for example, "amd64.v2"), as detailed below.

- "unix", if GOOS is a Unix or Unix-like system.

- the compiler being used, either "gc" or "gccgo"

- "cgo", if the cgo command is supported (see CGO_ENABLED in 'go help environment').

- a term for each Go major release, through the current version:
  "go1.1" from Go version 1.1 onward, "go1.12" from Go 1.12, and so on.

- any additional tags given by the -tags flag (see 'go help build').

There are no separate build tags for beta or minor releases.
测试版或次要版本没有单独的构建标签。

If a file's name, after stripping the extension and a possible _test suffix,
matches any of the following patterns:
	*_GOOS
	*_GOARCH
	*_GOOS_GOARCH
(example: source_windows_amd64.go) where GOOS and GOARCH represent
any known operating system and architecture values respectively, then
the file is considered to have an implicit build constraint requiring
those terms (in addition to any explicit constraints in the file).
如果文件名在去除扩展名和可能的 _test 后缀后，与以下任何模式匹配： *_GOOS *_GOARCH *_GOOS_GOARCH（示例：source_windows_amd64.go），
其中 GOOS 和 GOARCH 分别表示任何已知的操作系统和体系结构值，则该文件被认为具有需要这些术语的隐式构建约束（除了文件中的任何显式约束之外）。

Using GOOS=android matches build tags and files as for GOOS=linux
in addition to android tags and files.
使用 GOOS=android 除了 android 标签和文件之外，还与 GOOS=linux 匹配构建标签和文件。

Using GOOS=illumos matches build tags and files as for GOOS=solaris
in addition to illumos tags and files.
除了 illumos 标签和文件之外，使用 GOOS=illumos 还可以与 GOOS=solaris 一样匹配构建标签和文件。

Using GOOS=ios matches build tags and files as for GOOS=darwin
in addition to ios tags and files.
除了 ios 标签和文件之外，使用 GOOS=ios 还可以与 GOOS=darwin 一样匹配构建标签和文件。

The defined architecture feature build tags are:
定义的架构功能构建标签是：

	- For GOARCH=386, GO386=387 and GO386=sse2
	  set the 386.387 and 386.sse2 build tags, respectively.

	- For GOARCH=amd64, GOAMD64=v1, v2, and v3
	  correspond to the amd64.v1, amd64.v2, and amd64.v3 feature build tags.

	- For GOARCH=arm, GOARM=5, 6, and 7
	  correspond to the arm.5, arm.6, and arm.7 feature build tags.

	- For GOARCH=mips or mipsle,
	  GOMIPS=hardfloat and softfloat
	  correspond to the mips.hardfloat and mips.softfloat
	  (or mipsle.hardfloat and mipsle.softfloat) feature build tags.

	- For GOARCH=mips64 or mips64le,
	  GOMIPS64=hardfloat and softfloat
	  correspond to the mips64.hardfloat and mips64.softfloat
	  (or mips64le.hardfloat and mips64le.softfloat) feature build tags.

	- For GOARCH=ppc64 or ppc64le,
	  GOPPC64=power8, power9, and power10 correspond to the
	  ppc64.power8, ppc64.power9, and ppc64.power10
	  (or ppc64le.power8, ppc64le.power9, and ppc64le.power10)
	  feature build tags.

	- For GOARCH=wasm, GOWASM=satconv and signext
	  correspond to the wasm.satconv and wasm.signext feature build tags.

For GOARCH=amd64, arm, ppc64, and ppc64le, a particular feature level
sets the feature build tags for all previous levels as well.
For example, GOAMD64=v2 sets the amd64.v1 and amd64.v2 feature flags.
This ensures that code making use of v2 features continues to compile
when, say, GOAMD64=v4 is introduced.
Code handling the absence of a particular feature level
should use a negation:
对于 GOARCH=amd64、arm、ppc64 和 ppc64le，特定功能级别也会为所有先前级别设置功能构建标签。
例如，GOAMD64=v2 设置 amd64.v1 和 amd64.v2 功能标志。
这可以确保在引入 GOAMD64=v4 时，使用 v2 功能的代码可以继续编译。
处理缺少特定功能级别的代码应使用否定：

	//go:build !amd64.v2

To keep a file from being considered for any build:
要防止某个文件被考虑用于任何构建：

	//go:build ignore

(Any other unsatisfied word will work as well, but "ignore" is conventional.)
（任何其他不满意的词也可以，但“忽略”是约定俗成的。）

To build a file only when using cgo, and only on Linux and OS X:
仅在使用 cgo 时且仅在 Linux 和 OS X 上构建文件：

	//go:build cgo && (linux || darwin)

Such a file is usually paired with another file implementing the
default functionality for other systems, which in this case would
carry the constraint:
这样的文件通常与实现其他系统默认功能的另一个文件配对，在这种情况下，该文件将带有约束：

	//go:build !(cgo && (linux || darwin))

Naming a file dns_windows.go will cause it to be included only when
building the package for Windows; similarly, math_386.s will be included
only when building the package for 32-bit x86.
将文件命名为 dns_windows.go 将导致仅在为 Windows 构建包时才包含该文件；同样，仅在构建 32 位 x86 包时才会包含 math_386.s。

Go versions 1.16 and earlier used a different syntax for build constraints,
with a "// +build" prefix. The gofmt command will add an equivalent //go:build
constraint when encountering the older syntax.
Go 版本 1.16 及更早版本使用不同的语法来构建约束，并带有“// +build”前缀。当遇到旧语法时，gofmt 命令将添加等效的 //go:build 约束。
