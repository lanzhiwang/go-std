```bash
$ go help go.mod
A module version is defined by a tree of source files, with a go.mod
file in its root. When the go command is run, it looks in the current
directory and then successive parent directories to find the go.mod
marking the root of the main (current) module.

The go.mod file format is described in detail at
https://golang.org/ref/mod#go-mod-file.

To create a new go.mod file, use 'go mod init'. For details see
'go help mod init' or https://golang.org/ref/mod#go-mod-init.

To add missing module requirements or remove unneeded requirements,
use 'go mod tidy'. For details, see 'go help mod tidy' or
https://golang.org/ref/mod#go-mod-tidy.

To add, upgrade, downgrade, or remove a specific module requirement, use
'go get'. For details, see 'go help module-get' or
https://golang.org/ref/mod#go-get.

To make other changes or to parse go.mod as JSON for use by other tools,
use 'go mod edit'. See 'go help mod edit' or
https://golang.org/ref/mod#go-mod-edit.

#########################################################################################

$ go help mod init
usage: go mod init [module-path]

Init initializes and writes a new go.mod file in the current directory, in
effect creating a new module rooted at the current directory. The go.mod file
must not already exist.

Init accepts one optional argument, the module path for the new module. If the
module path argument is omitted, init will attempt to infer the module path
using import comments in .go files, vendoring tool configuration files (like
Gopkg.lock), and the current directory (if in GOPATH).

If a configuration file for a vendoring tool is present, init will attempt to
import module requirements from it.

See https://golang.org/ref/mod#go-mod-init for more about 'go mod init'.

[module-path] 见 《go语言学习指南》第九章

#########################################################################################

$ go help mod tidy
usage: go mod tidy [-e] [-v] [-go=version] [-compat=version]

Tidy makes sure go.mod matches the source code in the module.
It adds any missing modules necessary to build the current module\'s
packages and dependencies, and it removes unused modules that
don\'t provide any relevant packages. It also adds any missing entries
to go.sum and removes any unnecessary ones.

The -v flag causes tidy to print information about removed modules
to standard error.
-v标志导致tidy将被移除模块的信息打印到标准错误。

The -e flag causes tidy to attempt to proceed despite errors
encountered while loading packages.
-e标志使tidy尝试继续，尽管在加载包时遇到错误。

The -go flag causes tidy to update the 'go' directive in the go.mod
file to the given version, which may change which module dependencies
are retained as explicit requirements in the go.mod file.
(Go versions 1.17 and higher retain more requirements in order to
support lazy module loading.)
-go标志使tidy在go过程中更新'go'指令。将文件修改为给定的版本，这可能会改变保留
哪些模块依赖关系作为运行中的显式需求。国防部文件。(为了支持延迟模块加载，Go版本
1.17及更高版本保留了更多的需求。)

$ go mod tidy -go=1.18
$ go mod tidy -go=1.17
$ go mod tidy -go=1.16


The -compat flag preserves any additional checksums needed for the
'go' command from the indicated major Go release to successfully load
the module graph, and causes tidy to error out if that version of the
'go' command would load any imported package from a different module
version. By default, tidy acts as if the -compat flag were set to the
version prior to the one indicated by the 'go' directive in the go.mod
file.
-compat标志保留了从指定的主要go版本中'go'命令成功加载模块图所需的任何额外校验
和，并且如果该版本的'go'命令将加载来自不同模块版本的任何导入包，则会导致tidy出
错。默认情况下，tidy的作用就像是-compat标志被设置为go中'go'指令所指示的版本之
前的版本。

See https://golang.org/ref/mod#go-mod-tidy for more about 'go mod tidy'.

#########################################################################################

$ go help module-get
The 'go get' command changes behavior depending on whether the
go command is running in module-aware mode or legacy GOPATH mode.
This help text, accessible as 'go help module-get' even in legacy GOPATH mode,
describes 'go get' as it operates in module-aware mode.

Usage: go get [-t] [-u] [-v] [build flags] [packages]

Get resolves its command-line arguments to packages at specific module versions,
updates go.mod to require those versions, and downloads source code into the
module cache.

To add a dependency for a package or upgrade it to its latest version:

	go get example.com/pkg

To upgrade or downgrade a package to a specific version:

	go get example.com/pkg@v1.2.3

To remove a dependency on a module and downgrade modules that require it:

	go get example.com/mod@none

See https://golang.org/ref/mod#go-get for details.

In earlier versions of Go, 'go get' was used to build and install packages.
Now, 'go get' is dedicated to adjusting dependencies in go.mod. 'go install'
may be used to build and install commands instead. When a version is specified,
'go install' runs in module-aware mode and ignores the go.mod file in the
current directory. For example:

	go install example.com/pkg@v1.2.3
	go install example.com/pkg@latest

See 'go help install' or https://golang.org/ref/mod#go-install for details.

'go get' accepts the following flags.

The -t flag instructs get to consider modules needed to build tests of
packages specified on the command line.
-t标志指示get考虑构建命令行上指定的包的测试所需的模块。

The -u flag instructs get to update modules providing dependencies
of packages named on the command line to use newer minor or patch
releases when available.
-u标志指示get更新提供在命令行中命名的包的依赖项的模块，以便在可用时使用更新的次要或补丁版本。

The -u=patch flag (not -u patch) also instructs get to update dependencies,
but changes the default to select patch releases.
-u=patch标志(不是-u patch)也指示get更新依赖项，但更改默认值以选择补丁版本。

When the -t and -u flags are used together, get will update
test dependencies as well.
当-t和-u标志一起使用时，get也将更新测试依赖项。

The -x flag prints commands as they are executed. This is useful for
debugging version control commands when a module is downloaded directly
from a repository.
-x标志在执行命令时打印命令。当模块直接从存储库下载时，这对于调试版本控制命令非常有用。

For more about modules, see https://golang.org/ref/mod.

For more about specifying packages, see 'go help packages'.

This text describes the behavior of get using modules to manage source
code and dependencies. If instead the go command is running in GOPATH
mode, the details of get\'s flags and effects change, as does 'go help get'.
See 'go help gopath-get'.

See also: go build, go install, go clean, go mod.

#########################################################################################

#########################################################################################

#########################################################################################

#########################################################################################

```

* [Go mod 学习之 replace 篇 解决 go 本地依赖、无法拉取依赖、禁止依赖等问题](https://blog.csdn.net/qq_24433609/article/details/127323097?share_token=20e2c86f-99bd-4cfe-b826-3cd2c57364ac)

* [Go mod 学习之 exclude 篇 禁止引用多个依赖包间接引入的bug版本包](https://blog.csdn.net/qq_24433609/article/details/127325204?share_token=19af3677-9413-4ff0-9390-8da80119adbd)

* v0.0.0-20210930125809-cb0fa318a74b

* [Go 专家编程 go.mod 文件中的indirect准确含义](https://blog.csdn.net/juzipidemimi/article/details/104441398)

* [Go 专家编程 go.mod 文件中 incompatible 包意味着什么](https://my.oschina.net/renhc/blog/3167195)

