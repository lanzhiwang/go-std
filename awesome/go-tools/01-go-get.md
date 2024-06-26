```bash
$ go help get

# 重点参数
$ go get -x [packages]
```

usage: go get [-t] [-u] [-v] [build flags] [packages]
用法

Get resolves its command-line arguments to packages at specific module versions,
**updates go.mod** to require those versions, and downloads source code into the
module cache.
Get 将其命令行参数解析为特定模块版本的包，更新 go.mod 以要求这些版本，并将源代码下载到模块缓存中。

To add a dependency for a package or upgrade it to its latest version:
要添加包的依赖项或将其升级到最新版本：

	go get example.com/pkg

To upgrade or downgrade a package to a specific version:
要将软件包升级或降级到特定版本：

	go get example.com/pkg@v1.2.3

To remove a dependency on a module and downgrade modules that require it:
要删除对模块的依赖并降级需要它的模块：

	go get example.com/mod@none

To upgrade the minimum required Go version to the latest released Go version:
要将最低所需的 Go 版本升级到最新发布的 Go 版本：

	go get go@latest

To upgrade the Go toolchain to the latest patch release of the current Go toolchain:
将 Go 工具链升级到当前 Go 工具链的最新补丁版本：

	go get toolchain@patch

See https://golang.org/ref/mod#go-get for details.

In earlier versions of Go, 'go get' was used to build and install packages.
Now, 'go get' is dedicated to adjusting dependencies in go.mod. 'go install'
may be used to build and install commands instead. When a version is specified,
'go install' runs in **module-aware mode** and ignores the go.mod file in the
current directory. For example:
在 Go 的早期版本中，“go get”用于构建和安装包。
现在，“go get”专门用于调整 go.mod 中的依赖关系。
“go install”可用于构建和安装命令。
指定版本后，“go install”将以模块感知模式运行并忽略当前目录中的 go.mod 文件。例如：

	go install example.com/pkg@v1.2.3
	go install example.com/pkg@latest

See 'go help install' or https://golang.org/ref/mod#go-install for details.

'go get' accepts the following flags.

The -t flag instructs get to consider modules needed to build tests of
packages specified on the command line.
-t 标志指示 get 考虑构建命令行上指定的包的测试所需的模块。

The -u flag instructs get to update modules providing dependencies
of packages named on the command line to use newer minor or patch
releases when available.
-u 标志指示 get 更新提供命令行上指定的包的依赖项的模块，以使用更新的次要版本或补丁版本（如果可用）。

The -u=patch flag (not -u patch) also instructs get to update dependencies,
but changes the default to select patch releases.
-u=patch 标志（不是 -u patch）还指示 get 更新依赖项，但更改默认值以选择补丁版本。

When the -t and -u flags are used together, get will update
test dependencies as well.
当 -t 和 -u 标志一起使用时， get 也会更新测试依赖项。

The -x flag prints commands as they are executed. This is useful for
debugging version control commands when a module is downloaded directly
from a repository.
-x 标志在执行命令时打印命令。当直接从存储库下载模块时，这对于调试版本控制命令非常有用。

For more about modules, see https://golang.org/ref/mod.

For more about using 'go get' to update the minimum Go version and
suggested Go toolchain, see https://go.dev/doc/toolchain.

For more about specifying packages, see 'go help packages'.

This text describes the behavior of get using modules to manage source
code and dependencies. If instead the go command is running in GOPATH
mode, the details of get's flags and effects change, as does 'go help get'.
See 'go help gopath-get'.
本文描述了 get 使用模块来管理源代码和依赖项的行为。
相反，如果 go 命令在 GOPATH 模式下运行，则 get 的标志和效果的详细信息会发生变化，“go help get”也会发生变化。
请参阅“go help gopath-get”。

See also: go build, go install, go clean, go mod.
